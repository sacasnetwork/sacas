from pathlib import Path

import pytest

from .network import create_snapshots_dir, setup_custom_sacas
from .utils import memiavl_config, wait_for_block


@pytest.fixture(scope="module")
def custom_sacas(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_sacas(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_sacas_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-rocksdb")
    yield from setup_custom_sacas(
        path,
        26810,
        memiavl_config(path, "discard-abci-resp"),
        post_init=create_snapshots_dir,
        chain_binary="sacasd-rocksdb",
    )


@pytest.fixture(scope="module", params=["sacas", "sacas-rocksdb"])
def sacas_cluster(request, custom_sacas, custom_sacas_rocksdb):
    """
    run on sacas and
    sacas built with rocksdb (memIAVL + versionDB)
    """
    provider = request.param
    if provider == "sacas":
        yield custom_sacas

    elif provider == "sacas-rocksdb":
        yield custom_sacas_rocksdb

    else:
        raise NotImplementedError


def test_gas_eth_tx(sacas_cluster):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    wait_for_block(sacas_cluster.cosmos_cli(), 3)
    try:
        sacas_cluster.w3.eth.gas_price  # pylint: disable=pointless-statement
        raise Exception(  # pylint: disable=broad-exception-raised
            "This query should have failed"
        )
    except Exception as error:
        assert "block result not found" in error.args[0]["message"]
