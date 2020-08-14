#!/bin/bash
set -e -u -x -o pipefail

log() {
    echo -e "$(date)" ${1}
}

installesdtodmg=${1}
installer_image="${BUILD_WORKSPACE_DIRECTORY=}/${2}"
outfile="${BUILD_WORKSPACE_DIRECTORY=}/${3}"

work="${BUILD_WORKSPACE_DIRECTORY}/work"
installer_mountpoint="${work}/mnt/installer"

function cleanup {
    hdiutil detach "${installer_mountpoint}"
}
trap cleanup EXIT

mkdir -p "${installer_mountpoint}"

log "Attaching Installer image to ${installer_mountpoint}"
hdiutil attach \
    -nobrowse \
    -noautoopen \
    -noverify \
    -owners on \
    -mountpoint "${installer_mountpoint}" \
    "${installer_image}"

touch "template.adtmpl"
sudo ${installesdtodmg} \
    $(id -u) $(id -g) \
    "APFS" \
    "${outfile}.tmp.dmg" \
    "macOS" \
    "32" \
    "template.adtmpl" \
    "${installer_mountpoint}/Applications/Install macOS Catalina.app/Contents/SharedSupport/InstallInfo.plist"

hdiutil convert -format UDRW -o "${outfile}.dmg" "${outfile}.tmp.dmg"
mv "${outfile}.dmg" "${outfile}"
rm -f "${outfile}.tmp.dmg"
