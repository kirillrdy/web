{ }:

let
  archive = <nixpkgs>;

  pkgs = import archive {};
  pkgsWithOurPostgres = pkgs // {
    postgresql_13 = pkgs.postgresql_13.overrideAttrs(
      old: {patches = pkgs.lib.lists.init old.patches;}
    );
  };

  postgresql = (import "${archive}/pkgs/servers/sql/postgresql/default.nix" pkgsWithOurPostgres).postgresql_13;

in
with pkgs;
mkShell {

  shellHook = ''
  '';

  buildInputs = [
    golangci-lint
    #glibcLocales
    minio
    go
    (postgresql.withPackages ( p : [ p.postgis ]))

  ];
}
