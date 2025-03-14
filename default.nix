{ lib
, buildGoApplication
, buildPackages
, rev ? "dirty"
}:
let
  version = "latest";
  pname = "sacasd";
  tags = [ "netgo" ];
  ldflags = lib.concatStringsSep "\n" ([
    "-X github.com/cosmos/cosmos-sdk/version.Name=sacas"
    "-X github.com/cosmos/cosmos-sdk/version.AppName=${pname}"
    "-X github.com/cosmos/cosmos-sdk/version.Version=${version}"
    "-X github.com/cosmos/cosmos-sdk/version.BuildTags=${lib.concatStringsSep "," tags}"
    "-X github.com/cosmos/cosmos-sdk/version.Commit=${rev}"
  ]);
in
buildGoApplication rec {
  inherit pname version tags ldflags;
  go = buildPackages.go_1_20;
  src = ./.;
  modules = ./gomod2nix.toml;
  doCheck = false;
  pwd = src; # needed to support replace
  subPackages = [ "cmd/sacasd" ];
  CGO_ENABLED = "1";

  meta = with lib; {
    description = "Sacas is a scalable and interoperable blockchain, built on Proof-of-Stake with fast-finality using the Cosmos SDK which runs on top of CometBFT Core consensus engine.";
    homepage = "https://github.com/sacasnetwork/sacas";
    license = licenses.asl20;
    mainProgram = "sacasd";
  };
}