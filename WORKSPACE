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

new_git_repository(
    name = "com_github_munki_macadmin_scripts",
    build_file_content = """\
py_binary(
    name = "installinstallmacos",
    srcs = ["installinstallmacos.py"],
    python_version = "PY2",
    visibility = ["//visibility:public"],
)
    """,
    shallow_since = "1596052888 -0700",
    commit = "e71ade7e92f4aaf5d50cfe488fadb12ddb107d55",
    remote = "https://github.com/munki/macadmin-scripts",
)

new_git_repository(
    name = "com_github_magervalp_autodmg",
    remote = "https://github.com/MagerValp/AutoDMG.git",
    commit = "e5fc641b2f38be742ab083d5d5f1cb952a30af89",
    build_file_content = """\
sh_binary(
    name = "installesdtodmg",
    srcs = ["AutoDMG/installesdtodmg.sh"],
    visibility = ["//visibility:public"],
)
    """,
    shallow_since = "1566401762 +0200",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2697f6bc7c529ee5e6a2d9799870b9ec9eaeb3ee7d70ed50b87a2c2c97e13d9e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()
