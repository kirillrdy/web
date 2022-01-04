{ }:

let
  archive = <nixpkgs>;

  pkgs = import archive {};
  pkgsWithOurPostgres = pkgs // {
    postgresql_14 = pkgs.postgresql_14.overrideAttrs(
      old: {patches = pkgs.lib.lists.init old.patches;}
    );
  };

  postgresql = (import "${archive}/pkgs/servers/sql/postgresql/default.nix" pkgsWithOurPostgres).postgresql_14;

in
with pkgs;
mkShell {

  shellHook = ''
  '';

  buildInputs = [
    golangci-lint
    #glibcLocales
    minio
    gopls
    go
    (postgresql.withPackages ( p : [ p.postgis ]))

  ];
}
