import pytest

from .network import setup_sacas, setup_sacas_rocksdb, setup_geth


@pytest.fixture(scope="session")
def sacas(tmp_path_factory):
    path = tmp_path_factory.mktemp("sacas")
    yield from setup_sacas(path, 26650)


@pytest.fixture(scope="session")
def sacas_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("sacas-rocksdb")
    yield from setup_sacas_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["sacas", "sacas-ws"])
def sacas_rpc_ws(request, sacas):
    """
    run on both sacas and sacas websocket
    """
    provider = request.param
    if provider == "sacas":
        yield sacas
    elif provider == "sacas-ws":
        sacas_ws = sacas.copy()
        sacas_ws.use_websocket()
        yield sacas_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["sacas", "sacas-ws", "sacas-rocksdb", "geth"])
def cluster(request, sacas, sacas_rocksdb, geth):
    """
    run on sacas, sacas websocket,
    sacas built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "sacas":
        yield sacas
    elif provider == "sacas-ws":
        sacas_ws = sacas.copy()
        sacas_ws.use_websocket()
        yield sacas_ws
    elif provider == "geth":
        yield geth
    elif provider == "sacas-rocksdb":
        yield sacas_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["sacas", "sacas-rocksdb"])
def sacas_cluster(request, sacas, sacas_rocksdb):
    """
    run on sacas default build &
    sacas with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "sacas":
        yield sacas
    elif provider == "sacas-rocksdb":
        yield sacas_rocksdb
    else:
        raise NotImplementedError
