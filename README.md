# Project Crowbar Packwiz

Unofficial Packwiz conversion of the [Project Crowbar SMP](https://modrinth.com/server/project-crowbar-smp) modpack.

## Usage

Use [packwiz-installer](https://packwiz.infra.link/tutorials/installing/packwiz-installer/) with the raw URL for `pack.toml`, or export a Modrinth pack locally:

```sh
packwiz modrinth export
```

The server address is `play.projectcrowbar.com`. Project Crowbar requires Discord account linking before joining.

## Updates

The `Update from Modrinth` GitHub Actions workflow checks every six hours for a new official Project Crowbar release. When one is available, it imports the release, refreshes the Packwiz index, and commits the result.

Project Crowbar and the bundled content belong to their respective authors. This repository does not grant additional rights to redistributed content.
