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
    sha256 = "d67717ef150d9e59068acad564282ef63e228b113274630910deb549562feb7f",
    strip_prefix = "CloverV2",
    urls = ["https://github.com/CloverHackyColor/CloverBootloader/releases/download/5120/CloverV2-5120.zip"],
)

http_archive(
    name = "com_github_acidanthera_lilu",
    build_file_content = _EXPORTS_ALL,
    sha256 = "a7c5463293ea80bd689293169d1929d117ec1095e37ee646018fe1927a9a6eeb",
    urls = ["https://github.com/acidanthera/Lilu/releases/download/1.4.6/Lilu-1.4.6-RELEASE.zip"],
)

http_archive(
    name = "com_github_acidanthera_virtualsmc",
    build_file_content = _EXPORTS_ALL,
    sha256 = "0b01da6e0187091999c63ab0b7de8ecb3ff8143416bc8d3bd07a77749b5f24b1",
    urls = ["https://github.com/acidanthera/VirtualSMC/releases/download/1.1.5/VirtualSMC-1.1.5-RELEASE.zip"],
)

new_git_repository(
    name = "com_github_acidanthera_virtualsmc_src",
    build_file_content = _EXPORTS_ALL,
    remote = "https://github.com/acidanthera/VirtualSMC.git",
    commit = "da7d723ce3949b5f094a77777c40bd31beaa8f38",
    shallow_since = "1596861558 +0300",
)

http_archive(
    name = "com_github_acidanthera_opencore",
    build_file_content = _EXPORTS_ALL,
    sha256 = "ca00d9113f67b55e86f1e39513dadf9b27312598ef20ca401bfb048572ea5def",
    urls = ["https://github.com/acidanthera/OpenCorePkg/releases/download/0.6.0/OpenCore-0.6.0-RELEASE.zip"],
)

http_file(
    name = "com_github_linuxkit_linuxkit",
    urls = ["https://github.com/linuxkit/linuxkit/releases/download/v0.8/linuxkit-darwin-amd64"],
    sha256 = "4dc05ee018f66da9307e996448160166f022ae66b11df6c4c529e8e0f0b1cc34",
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
