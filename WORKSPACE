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
    build_file_content = _EXPORTS_ALL,
    sha256 = "730f72bf139f017479cc22c17c64f47367212310e08c484ff46c02b3a44efeba",
    strip_prefix = "CloverV2",
    urls = ["https://github.com/CloverHackyColor/CloverBootloader/releases/download/5102/CloverV2-5102.zip"],
)

http_archive(
    name = "com_github_acidanthera_lilu",
    build_file_content = _EXPORTS_ALL,
    sha256 = "f2ea252814449f5297bcfdbee9d771c73329f1c7e7e48659098589ecffd759c9",
    urls = ["https://github.com/acidanthera/Lilu/releases/download/1.4.0/Lilu-1.4.0-RELEASE.zip"],
)

http_archive(
    name = "com_github_acidanthera_virtualsmc",
    build_file_content = _EXPORTS_ALL,
    sha256 = "ba368c4be95cada8491760bb4dfddc6832dc0f7bae6f9a611ffd5862236341f9",
    urls = ["https://github.com/acidanthera/VirtualSMC/releases/download/1.0.9/VirtualSMC-1.0.9-RELEASE.zip"],
)

http_file(
    name = "ovmf_code",
    sha256 = "81ede7af3351a1710a4d4a35655fac10bedb11ae3484ae785820803a5bf69829",
    urls = ["https://cdn.download.clearlinux.org/image/OVMF_CODE.fd"],
)

http_file(
    name = "ovmf_vars",
    sha256 = "5d2ac383371b408398accee7ec27c8c09ea5b74a0de0ceea6513388b15be5d1e",
    urls = ["https://cdn.download.clearlinux.org/image/OVMF_VARS.fd"],
)

new_git_repository(
    name = "com_github_pmj_virtio_net_osx",
    build_file_content = _EXPORTS_ALL,
    commit = "31b4f8e27c077f8c1234f0923a9698330b18b655",
    remote = "https://github.com/pmj/virtio-net-osx.git",
    shallow_since = "1388123323 -0800",
)

http_archive(
    name = "com_github_acidanthera_opencore",
    build_file_content = _EXPORTS_ALL,
    sha256 = "97be3932f2fc294e2f6d3ce9ee3fced48d24f7f2bbf0b7869d89706c72baafa8",
    urls = ["https://github.com/acidanthera/OpenCorePkg/releases/download/0.5.4/OpenCore-0.5.4-DEBUG.zip"],
)

http_archive(
    name = "com_github_acidanthera_applesupport",
    build_file_content = _EXPORTS_ALL,
    sha256 = "80bee4d37e4eadbfa6905aa811b7cb9ac433be58a0c8dd73e7c2c4630dce45ce",
    urls = ["https://github.com/acidanthera/AppleSupportPkg/releases/download/2.1.5/AppleSupport-2.1.5-DEBUG.zip"],
)
