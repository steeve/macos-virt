#!/bin/bash
sudo ./external/com_github_munki_macadmin_scripts/installinstallmacos \
    --raw \
    --workdir ${BUILD_WORKSPACE_DIRECTORY}/work
chown -r $(whoami) ${BUILD_WORKSPACE_DIRECTORY}/work
