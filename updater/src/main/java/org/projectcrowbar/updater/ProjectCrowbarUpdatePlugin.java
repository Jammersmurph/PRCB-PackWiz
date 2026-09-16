package org.projectcrowbar.updater;

import cpw.mods.modlauncher.api.IEnvironment;
import cpw.mods.modlauncher.api.IModuleLayerManager;
import cpw.mods.modlauncher.api.ITransformationService;
import cpw.mods.modlauncher.api.ITransformer;
import cpw.mods.modlauncher.api.IncompatibleEnvironmentException;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;

public class ProjectCrowbarUpdatePlugin implements ITransformationService {

    @Override
    public @NotNull String name() {
        return "prcb_updater";
    }

    @Override
    public void initialize(IEnvironment environment) {
        if (Boolean.getBoolean("prcbupdater.skip")) {
            ProjectCrowbarUpdater.LOGGER.info("[ProjectCrowbarUpdater] Skipping update check because 'prcbupdater.skip' is set to true.");
            return;
        }

        ProjectCrowbarUpdater.LOGGER.info("[ProjectCrowbarUpdater] Service initializing. Running Project Crowbar updater...");
        try {
            UpdaterBootstrap.runProjectCrowbarUpdate();
        } catch (Exception e) {
            ProjectCrowbarUpdater.LOGGER.error("[ProjectCrowbarUpdater] Update failed: " + e.getMessage());
            e.printStackTrace();
        }
    }

    @Override
    public void onLoad(IEnvironment environment, Set<String> otherServices) throws IncompatibleEnvironmentException {}

    @Override
    public @NotNull List<ITransformer<?>> transformers() {
        return new ArrayList<>();
    }

    @Override
    public List<Resource> beginScanning(IEnvironment environment) {
        return List.of();
    }

    @Override
    public List<Resource> completeScan(IModuleLayerManager layerManager) {
        return List.of();
    }
}
