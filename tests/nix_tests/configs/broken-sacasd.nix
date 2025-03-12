{ pkgs ? import ../../../nix { } }:
let sacasd = (pkgs.callPackage ../../../. { });
in
sacasd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-sacasd.patch
  ];
})
