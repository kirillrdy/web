{}:

let
  lock = builtins.fromJSON (builtins.readFile ./flake.lock);
  nixpkgs = (fetchTarball
    { url = "https://api.github.com/repos/nixos/nixpkgs/tarball/${lock.nodes.nixpkgs.locked.rev}"; sha256 = lock.nodes.nixpkgs.locked.narHash; });

  pkgs = import nixpkgs { };

in
with pkgs; mkShell {
  shellHook = ''
  '';

  buildInputs = [
    golangci-lint
    minio
    gopls
    go
    (postgresql.withPackages (p: [ p.postgis ]))
  ];
}
