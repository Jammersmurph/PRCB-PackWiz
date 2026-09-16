package org.projectcrowbar.updater;

import com.mojang.logging.LogUtils;
import net.neoforged.fml.common.Mod;
import org.slf4j.Logger;

@Mod(ProjectCrowbarUpdater.MODID)
public class ProjectCrowbarUpdater {
    public static final String MODID = "prcb_updater";
    public static final Logger LOGGER = LogUtils.getLogger();

    public ProjectCrowbarUpdater() {}
}
