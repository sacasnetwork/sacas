{ pkgs
, config
, sacas ? (import ../. { inherit pkgs; })
}: rec {
  start-sacas = pkgs.writeShellScriptBin "start-sacas" ''
    # rely on environment to provide sacasd
    export PATH=${pkgs.test-env}/bin:$PATH
    ${../scripts/start-sacas.sh} ${config.sacas-config} ${config.dotenv} $@
  '';
  start-geth = pkgs.writeShellScriptBin "start-geth" ''
    export PATH=${pkgs.test-env}/bin:${pkgs.go-ethereum}/bin:$PATH
    source ${config.dotenv}
    ${../scripts/start-geth.sh} ${config.geth-genesis} $@
  '';
  start-scripts = pkgs.symlinkJoin {
    name = "start-scripts";
    paths = [ start-sacas start-geth ];
  };
}
