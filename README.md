# Project Crowbar Updater

An early-loading auto updater for the Project Crowbar modpack. Forked from the CreateVC auto updater to make the Project Crowbar experience more broad and convenient.

- Reads the Packwiz manifest from the `main` branch of [PRCB-PackWiz](https://github.com/Jammersmurph/PRCB-PackWiz).
- Can be disabled with the JVM argument `-Dprcbupdater.skip=true`.

### Configuration

The update check URL can be configured using JVM arguments:

Use `-Dprcbupdater.url=<YOUR_URL>` to select a different `pack.toml` URL.
