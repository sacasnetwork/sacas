from pathlib import Path

import pytest
from web3 import Web3

from .network import setup_custom_sacas
from .utils import ADDRS, eth_to_bech32, wait_for_fn


@pytest.fixture(scope="module")
def sacas(request, tmp_path_factory):
    yield from setup_custom_sacas(
        tmp_path_factory.mktemp("zero-fee"),
        26900,
        Path(__file__).parent / "configs/zero-fee.jsonnet",
    )


def test_cosmos_tx(sacas):
    """
    test basic cosmos transaction works with zero fees
    """
    denom = "asacas"
    cli = sacas.cosmos_cli()
    sender = eth_to_bech32(ADDRS["signer1"])
    receiver = eth_to_bech32(ADDRS["signer2"])
    amt = 1000

    old_src_balance = cli.balance(sender, denom)
    old_dst_balance = cli.balance(receiver, denom)

    tx = cli.transfer(
        sender,
        receiver,
        f"{amt}{denom}",
        gas_prices=f"0{denom}",
        generate_only=True,
    )

    tx = cli.sign_tx_json(tx, sender, max_priority_price=0)

    rsp = cli.broadcast_tx_json(tx, broadcast_mode="sync")
    assert rsp["code"] == 0, rsp["raw_log"]

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = cli.balance(receiver, denom)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = cli.balance(sender, denom)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance


def test_eth_tx(sacas):
    """
    test basic Ethereum transaction works with zero fees
    """
    w3: Web3 = sacas.w3

    sender = ADDRS["signer1"]
    receiver = ADDRS["signer2"]
    amt = 1000

    old_src_balance = w3.eth.get_balance(sender)
    old_dst_balance = w3.eth.get_balance(receiver)

    txhash = w3.eth.send_transaction(
        {
            "from": sender,
            "to": receiver,
            "value": amt,
            "gasPrice": 0,
        }
    )
    receipt = w3.eth.wait_for_transaction_receipt(txhash)
    assert receipt.status == 1

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = w3.eth.get_balance(receiver)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = w3.eth.get_balance(sender)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance
