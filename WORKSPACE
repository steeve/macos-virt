load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

_EXPORTS_ALL = """
filegroup(
    name = "all",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
"""

http_archive(
    name = "com_github_cloverhackycolor_cloverbootloader",
    url = "https://github.com/CloverHackyColor/CloverBootloader/releases/download/5102/CloverV2-5102.zip",
    strip_prefix = "CloverV2",
    build_file_content = _EXPORTS_ALL,
)

http_archive(
    name = "com_github_acidanthera_lilu",
    url = "https://github.com/acidanthera/Lilu/releases/download/1.4.0/Lilu-1.4.0-RELEASE.zip",
    build_file_content = _EXPORTS_ALL,
)

http_archive(
    name = "com_github_acidanthera_virtualsmc",
    url = "https://github.com/acidanthera/VirtualSMC/releases/download/1.0.9/VirtualSMC-1.0.9-RELEASE.zip",
    build_file_content = _EXPORTS_ALL,
)

http_file(
    name = "ovmf_code",
    urls = ["https://cdn.download.clearlinux.org/image/OVMF_CODE.fd"],
)

http_file(
    name = "ovmf_vars",
    urls = ["https://cdn.download.clearlinux.org/image/OVMF_VARS.fd"],
)

new_git_repository(
    name = "com_github_pmj_virtio_net_osx",
    commit = "31b4f8e27c077f8c1234f0923a9698330b18b655",
    remote = "https://github.com/pmj/virtio-net-osx.git",
    shallow_since = "1388123323 -0800",
    build_file_content = _EXPORTS_ALL,
)
