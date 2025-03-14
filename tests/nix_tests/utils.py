import base64
import json
import os
import socket
import subprocess
import sys
import time
from collections import defaultdict
from pathlib import Path

import bech32
from dateutil.parser import isoparse
from dotenv import load_dotenv
from eth_account import Account
from hexbytes import HexBytes
from web3._utils.transactions import fill_nonce, fill_transaction_defaults
from web3.exceptions import TimeExhausted

load_dotenv(Path(__file__).parent.parent.parent / "scripts/.env")
Account.enable_unaudited_hdwallet_features()
ACCOUNTS = {
    "validator": Account.from_mnemonic(os.getenv("VALIDATOR1_MNEMONIC")),
    "community": Account.from_mnemonic(os.getenv("COMMUNITY_MNEMONIC")),
    "signer1": Account.from_mnemonic(os.getenv("SIGNER1_MNEMONIC")),
    "signer2": Account.from_mnemonic(os.getenv("SIGNER2_MNEMONIC")),
}
KEYS = {name: account.key for name, account in ACCOUNTS.items()}
ADDRS = {name: account.address for name, account in ACCOUNTS.items()}
SACAS_ADDRESS_PREFIX = "sacas"
DEFAULT_DENOM = "asacas"
TEST_CONTRACTS = {
    "TestERC20A": "TestERC20A.sol",
    "Greeter": "Greeter.sol",
    "BurnGas": "BurnGas.sol",
    "TestChainID": "ChainID.sol",
    "Mars": "Mars.sol",
    "StateContract": "StateContract.sol",
    "ICS20I": "sacas/ics20/ICS20I.sol",
    "DistributionI": "sacas/distribution/DistributionI.sol",
    "StakingI": "sacas/staking/StakingI.sol",
}


def contract_path(name, filename):
    return (
        Path(__file__).parent
        / "hardhat/artifacts/contracts/"
        / filename
        / (name + ".json")
    )


CONTRACTS = {
    **{
        name: contract_path(name, filename) for name, filename in TEST_CONTRACTS.items()
    },
}


def wait_for_port(port, host="127.0.0.1", timeout=40.0):
    start_time = time.perf_counter()
    while True:
        try:
            with socket.create_connection((host, port), timeout=timeout):
                break
        except OSError as ex:
            time.sleep(0.1)
            if time.perf_counter() - start_time >= timeout:
                raise TimeoutError(
                    "Waited too long for the port {} on host {} to start accepting "
                    "connections.".format(port, host)
                ) from ex


def w3_wait_for_new_blocks(w3, n, sleep=0.5):
    begin_height = w3.eth.block_number
    while True:
        time.sleep(sleep)
        cur_height = w3.eth.block_number
        if cur_height - begin_height >= n:
            break


def wait_for_new_blocks(cli, n, sleep=0.5):
    cur_height = begin_height = int((cli.status())["SyncInfo"]["latest_block_height"])
    while cur_height - begin_height < n:
        time.sleep(sleep)
        cur_height = int((cli.status())["SyncInfo"]["latest_block_height"])
    return cur_height


def wait_for_block(cli, height, timeout=240):
    for _ in range(timeout * 2):
        try:
            status = cli.status()
        except AssertionError as e:
            print(f"get sync status failed: {e}", file=sys.stderr)
        else:
            current_height = int(status["SyncInfo"]["latest_block_height"])
            if current_height >= height:
                break
            print("current block height", current_height)
        time.sleep(0.5)
    else:
        raise TimeoutError(f"wait for block {height} timeout")


def w3_wait_for_block(w3, height, timeout=240):
    for _ in range(timeout * 2):
        try:
            current_height = w3.eth.block_number
        except Exception as e:
            print(f"get json-rpc block number failed: {e}", file=sys.stderr)
        else:
            if current_height >= height:
                break
            print("current block height", current_height)
        time.sleep(0.5)
    else:
        raise TimeoutError(f"wait for block {height} timeout")


def wait_for_block_time(cli, t):
    print("wait for block time", t)
    while True:
        now = isoparse((cli.status())["SyncInfo"]["latest_block_time"])
        print("block time now: ", now)
        if now >= t:
            break
        time.sleep(0.5)


def wait_for_fn(name, fn, *, timeout=240, interval=1):
    for i in range(int(timeout / interval)):
        result = fn()
        print("check", name, result)
        if result:
            break
        time.sleep(interval)
    else:
        raise TimeoutError(f"wait for {name} timeout")


def get_precompile_contract(w3, name):
    jsonfile = CONTRACTS[name]
    info = json.loads(jsonfile.read_text())
    if name == "StakingI":
        addr = "0x0000000000000000000000000000000000000800"
    if name == "DistributionI":
        addr = "0x0000000000000000000000000000000000000801"
    if name == "ICS20I":
        addr = "0x0000000000000000000000000000000000000802"
    else:
        raise ValueError(f"invalid precompile contract name: {name}")
    return w3.eth.contract(addr, abi=info["abi"])


def deploy_contract(w3, jsonfile, args=(), key=KEYS["validator"]):
    """
    deploy contract and return the deployed contract instance
    """
    acct = Account.from_key(key)
    info = json.loads(jsonfile.read_text())
    contract = w3.eth.contract(abi=info["abi"], bytecode=info["bytecode"])
    tx = contract.constructor(*args).build_transaction({"from": acct.address})
    txreceipt = send_transaction(w3, tx, key)
    assert txreceipt.status == 1
    address = txreceipt.contractAddress
    return w3.eth.contract(address=address, abi=info["abi"]), txreceipt


def fill_defaults(w3, tx):
    return fill_nonce(w3, fill_transaction_defaults(w3, tx))


def sign_transaction(w3, tx, key=KEYS["validator"]):
    "fill default fields and sign"
    acct = Account.from_key(key)
    tx["from"] = acct.address
    tx = fill_transaction_defaults(w3, tx)
    tx = fill_nonce(w3, tx)
    return acct.sign_transaction(tx)


def send_transaction(w3, tx, key=KEYS["validator"], i=0):
    if i > 3:
        raise TimeExhausted
    signed = sign_transaction(w3, tx, key)
    txhash = w3.eth.send_raw_transaction(signed.rawTransaction)
    try:
        return w3.eth.wait_for_transaction_receipt(txhash, timeout=20)
    except TimeExhausted:
        return send_transaction(w3, tx, key, i + 1)


def send_successful_transaction(w3, i=0):
    if i > 3:
        raise TimeExhausted
    signed = sign_transaction(w3, {"to": ADDRS["community"], "value": 1000})
    txhash = w3.eth.send_raw_transaction(signed.rawTransaction)
    try:
        receipt = w3.eth.wait_for_transaction_receipt(txhash, timeout=20)
        assert receipt.status == 1
    except TimeExhausted:
        return send_successful_transaction(w3, i + 1)
    return txhash


def eth_to_bech32(addr, prefix=SACAS_ADDRESS_PREFIX):
    bz = bech32.convertbits(HexBytes(addr), 8, 5)
    return bech32.bech32_encode(prefix, bz)


def decode_bech32(addr):
    _, bz = bech32.bech32_decode(addr)
    return HexBytes(bytes(bech32.convertbits(bz, 5, 8)))


def supervisorctl(inipath, *args):
    subprocess.run(
        (sys.executable, "-msupervisor.supervisorctl", "-c", inipath, *args),
        check=True,
    )


def parse_events(logs):
    return {
        ev["type"]: {attr["key"]: attr["value"] for attr in ev["attributes"]}
        for ev in logs[0]["events"]
    }


def parse_events_rpc(events):
    result = defaultdict(dict)
    for ev in events:
        for attr in ev["attributes"]:
            if attr["key"] is None:
                continue
            # after sdk v0.47, key and value are strings instead of byte arrays
            if type(attr["key"]) is str:
                result[ev["type"]][attr["key"]] = attr["value"]
            else:
                key = base64.b64decode(attr["key"].encode()).decode()
                if attr["value"] is not None:
                    value = base64.b64decode(attr["value"].encode()).decode()
                else:
                    value = None
                result[ev["type"]][key] = value
    return result


def derive_new_account(n=1):
    # derive a new address
    account_path = f"m/44'/60'/0'/0/{n}"
    mnemonic = os.getenv("COMMUNITY_MNEMONIC")
    return Account.from_mnemonic(mnemonic, account_path=account_path)


def compare_fields(a, b, fields):
    for field in fields:
        assert a[field] == b[field], f"{field} field mismatch"


# get_fees_from_tx_result returns the fees by unpacking them
# from the events contained in the tx_result of a cosmos transaction.
def get_fees_from_tx_result(tx_result, denom=DEFAULT_DENOM):
    for event in tx_result["events"]:
        if event["type"] == "tx":
            for attr in event["attributes"]:
                if attr["key"] == "fee":
                    return int(attr["value"].split(denom)[0])
