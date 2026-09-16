# CreateVC Updater 1.0.1 Analysis

Analyzed artifact: `createvc_updater-1.0.1.jar`

SHA-256: `4b0e87464c45657a8ea1e79bb9cc112da4399a270a1a1b0477b9778dcf80faa0`

## Startup Flow

1. NeoForge discovers `org.createvc.updater.UpdatePlugin` through ModLauncher's `ITransformationService` service file.
2. The transformation service runs during ModLauncher initialization, before normal mod scanning.
3. Unless `-Dcreatevcupdater.skip=true` is set, it extracts an embedded Packwiz bootstrap JAR into a temporary directory.
4. It launches that bootstrap in a separate Java process with a Packwiz `pack.toml` URL and the `-g` no-GUI option.
5. The bootstrap checks the latest `Jammersmurph/packwiz-installer` GitHub release, downloads or updates `createvc-updater.jar` in the Minecraft directory, and invokes its `link.infra.packwiz.installer.Main` class.
6. The Packwiz installer compares the remote index with local files, downloads changed files, removes invalidated managed files, and maintains its local cache.
7. The outer mod parses installer output such as `(current/total) Downloaded ...` and reflects it into NeoForge's early loading progress meter.

Update failures are logged and normal mod loading continues. A nonzero child-process exit is also logged without stopping Minecraft.

## Original Endpoints

- Pack manifest: `https://jammersmurph.github.io/CreateVC/pack.toml`
- Development manifest: `https://raw.githubusercontent.com/Jammersmurph/CreateVC/dev/pack.toml`
- Installer release API: `https://api.github.com/repos/Jammersmurph/packwiz-installer/releases/latest`
- Installer release asset: `createvc-updater.jar`

## Project Crowbar Fork

The fork under `updater/` changes the mod identity to `prcb_updater`, uses Project Crowbar branding, and reads this repository's main-branch manifest:

`https://raw.githubusercontent.com/Jammersmurph/PRCB-PackWiz/main/pack.toml`

It can be disabled with `-Dprcbupdater.skip=true` or pointed at another manifest with `-Dprcbupdater.url=<URL>`.

## Trust Model

The updater executes the latest installer JAR published by `Jammersmurph/packwiz-installer` and then applies files referenced by the configured Packwiz manifest. Neither executable is pinned by a hash in the outer mod, so control of either repository is equivalent to control of files installed into the Minecraft instance.
