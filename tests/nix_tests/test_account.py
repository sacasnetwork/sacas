import pytest
from web3 import Web3

from .network import setup_sacas, setup_sacas_rocksdb
from .utils import (
    ADDRS,
    KEYS,
    derive_new_account,
    send_transaction,
    w3_wait_for_new_blocks,
)


# start a brand new chain for this test
@pytest.fixture(scope="module")
def custom_sacas(tmp_path_factory):
    path = tmp_path_factory.mktemp("account")
    yield from setup_sacas(path, 26700, long_timeout_commit=True)


# ATM rocksdb build is not supported for sdkv0.50
# This is due to cronos dependencies (versionDB, memIAVL)
@pytest.fixture(scope="module")
def custom_sacas_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("account-rocksdb")
    yield from setup_sacas_rocksdb(
        path,
        26777,
    )


@pytest.fixture(scope="module", params=["sacas", "sacas-ws", "sacas-rocksdb", "geth"])
def cluster(request, custom_sacas, custom_sacas_rocksdb, geth):
    """
    run on sacas, sacas websocket,
    sacas built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "sacas":
        yield custom_sacas
    elif provider == "sacas-ws":
        sacas_ws = custom_sacas.copy()
        sacas_ws.use_websocket()
        yield sacas_ws
    # ATM rocksdb build is not supported for sdkv0.50
    # This is due to cronos dependencies (versionDB, memIAVL)
    elif provider == "sacas-rocksdb":
        yield custom_sacas_rocksdb
    elif provider == "geth":
        yield geth
    else:
        raise NotImplementedError


def test_get_transaction_count(cluster):
    w3: Web3 = cluster.w3
    blk = hex(w3.eth.block_number)
    sender = ADDRS["validator"]

    receiver = derive_new_account().address
    n0 = w3.eth.get_transaction_count(receiver, blk)
    # ensure transaction send in new block
    w3_wait_for_new_blocks(w3, 1, sleep=0.1)
    receipt = send_transaction(
        w3,
        {
            "from": sender,
            "to": receiver,
            "value": 1000,
        },
        KEYS["validator"],
    )
    assert receipt.status == 1
    [n1, n2] = [w3.eth.get_transaction_count(receiver, b) for b in [blk, "latest"]]
    assert n0 == n1
    assert n0 == n2


def test_query_future_blk(cluster):
    w3: Web3 = cluster.w3
    acc = derive_new_account(2).address
    current = w3.eth.block_number
    future = current + 1000
    with pytest.raises(ValueError) as exc:
        w3.eth.get_transaction_count(acc, hex(future))
    print(acc, str(exc))
    assert "-32000" in str(exc)
