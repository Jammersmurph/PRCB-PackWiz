<div align="center">

# Project Crowbar Packwiz

**The unofficial, auto-updating Packwiz repository for Project Crowbar.**

[![Minecraft](https://img.shields.io/badge/Minecraft-1.21.1-62b47a?style=flat-square)](https://www.minecraft.net/)
[![NeoForge](https://img.shields.io/badge/NeoForge-21.1.233-f47c20?style=flat-square)](https://neoforged.net/)
[![Update from Modrinth](https://github.com/Jammersmurph/PRCB-PackWiz/actions/workflows/update-modrinth.yml/badge.svg)](https://github.com/Jammersmurph/PRCB-PackWiz/actions/workflows/update-modrinth.yml)
[![Modrinth](https://img.shields.io/badge/Official_pack-Modrinth-1bd96a?style=flat-square)](https://modrinth.com/modpack/project-crowbar-smp)

[Official Modpack](https://modrinth.com/modpack/project-crowbar-smp) · [Updater Release](https://github.com/Jammersmurph/PRCB-PackWiz/releases/latest) · [Project Crowbar](https://projectcrowbar.com/) · [Discord](https://projectcrowbar.com/discord)

</div>

> [!IMPORTANT]
> This is an unofficial Packwiz conversion and update service. It is not the canonical Project Crowbar modpack distribution.

## What This Repository Does

This repository mirrors the official Project Crowbar Modrinth pack as a native [Packwiz](https://packwiz.infra.link/) pack. A scheduled GitHub Actions workflow checks Modrinth every six hours and, when a new compatible release appears, automatically:

1. Downloads the official NeoForge 1.21.1 `.mrpack`.
2. Imports its files and metadata into Packwiz.
3. Refreshes the Packwiz index.
4. Commits the updated pack to `main`.

The result is a stable `pack.toml` URL that launchers and the Project Crowbar updater can synchronize against:

```text
https://raw.githubusercontent.com/Jammersmurph/PRCB-PackWiz/main/pack.toml
```

## Automatic Installation

The optional Project Crowbar Updater can turn an otherwise empty instance into an automatically managed Project Crowbar installation.

1. Create a Minecraft `1.21.1` instance with NeoForge `21.1.233` and Java 21.
2. Download `prcb_updater-1.0.0.jar` from the [latest GitHub release](https://github.com/Jammersmurph/PRCB-PackWiz/releases/latest).
3. Place the JAR in the instance's `mods` directory.
4. Launch the instance with an internet connection.

The updater runs before normal mod scanning, installs the current pack, and checks this repository on later launches. Its source is maintained on the [`prcb-updater`](https://github.com/Jammersmurph/PRCB-PackWiz/tree/prcb-updater) branch.

## Packwiz Usage

Use the manifest URL above with [packwiz-installer](https://packwiz.infra.link/tutorials/installing/packwiz-installer/), or clone the repository and export a Modrinth pack locally:

```sh
packwiz modrinth export
```

## Project Crowbar

Project Crowbar is a community Minecraft server centered on Create, automation, magic, engineering, building, and collaborative play.

- Server: `play.projectcrowbar.com`
- Website: [projectcrowbar.com](https://projectcrowbar.com/)
- Joining requires a Discord account and the server's [account-linking process](https://projectcrowbar.com/wiki/mcdclink).

## Credits

- [Jammersmurph](https://github.com/Jammersmurph) maintains this unofficial Packwiz conversion, updater fork, and update automation.
- [Nikk](https://github.com/nikorsd) created and owns Project Crowbar and maintains its official infrastructure and modpack.
- The Project Crowbar staff, mod authors, artists, musicians, and community created the content represented by this pack.

## Disclaimer

Project Crowbar and all bundled content belong to their respective owners and authors. This repository does not claim ownership of the official pack and does not grant additional rights to redistributed content.
