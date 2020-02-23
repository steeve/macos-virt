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
    sha256 = "9a909f44c3000caae6a6afcd6487d4041c0887899a25e9ccd958902ec30c504e",
    strip_prefix = "CloverV2",
    urls = ["https://github.com/CloverHackyColor/CloverBootloader/releases/download/5104/CloverV2-5104.zip"],
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
    sha256 = "21d242a5964d824c076095427ee1191414634f1f848a12bb536610b6285993df",
    urls = ["https://github.com/acidanthera/OpenCorePkg/releases/download/0.5.5/OpenCore-0.5.5-DEBUG.zip"],
)

http_archive(
    name = "com_github_acidanthera_applesupport",
    build_file_content = _EXPORTS_ALL,
    sha256 = "80bee4d37e4eadbfa6905aa811b7cb9ac433be58a0c8dd73e7c2c4630dce45ce",
    urls = ["https://github.com/acidanthera/AppleSupportPkg/releases/download/2.1.5/AppleSupport-2.1.5-DEBUG.zip"],
)

http_file(
    name = "com_github_linuxkit_linuxkit",
    urls = ["https://github.com/linuxkit/linuxkit/releases/download/v0.7/linuxkit-darwin-amd64"],
    sha256 = "9ea7fd7c6ba946b06ba398ec342b9f6fb5723f9063dcb94e70f3eac9d8cee179",
    executable = True,
)

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.21.3/rules_go-v0.21.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.21.3/rules_go-v0.21.3.tar.gz",
    ],
    sha256 = "af04c969321e8f428f63ceb73463d6ea817992698974abeff0161e069cd08bd6",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()
