# macos-virt
The main goal of this project is to easily create macOS virtual machines that
can be booted everywhere for things like development, testing or CI. The idea is
not to provide a desktop like experience, more server-like. As such, GUI is
working, but not the main goal.

macOS has very specific tastes in the hardware it supports. There has been
extensive work by the hackintosh community to ease this to a maximum, but the
real breakthrough came when using QEMU/KVM, since hardware could be emulated by
the virtualization layer (except for the GPU).

This made running macOS dramatically simpler. And this is the approach this
project is taking using Linux/QEMU/KVM has a Type-2 hypervisor for macOS,
thanks to nested virtualization.

Most of the toolchain is sandboxed by Bazel and is expected to run on macOS.

Requirements:
- macOS
- bazelisk: `brew install bazelbuild/tap/bazelisk`
- Docker for Mac (used by `linuxkit`)

## Architecture
The architecture has several layers:

1. macOS
2. OpenCore bootloader
3. Linux/QEMU/KVM
4. Host OS

## macOS
Most of the downloads and temporary files will reside in the `work` directory.

### Installer image
First, create the installer image:
```
$ bazel run //macos:mkinstaller
```
Then choose which version of macOS you want to download and it will create the
installer.

You'll end up with something like `work/Install_macOS_10.15.5-19F2200.sparseimage`.

### macOS booted image
It is possible to install this image directly to an raw disk image without
booting into the installer:
```
$ bazel run //macos:mkimage -- work/Install_macOS_10.15.5-19F2200.sparseimage macOS_10.15.5-19F2200.img
```

This will create `macOS_10.15.5-19F2200.img`, a raw bootable vanilla image.

It is vanilla and doesn't have anything hackintosh specific. It needs a
bootloader.

## OpenCore
The lower level bootloader for the image is OpenCore. It is an advanced
bootloader that can do things like patching the kernel, passing custom arguments
to the kernel, even out CPU specificities and so on.

The image is kept separate from the macOS one in order to iterate faster without
rebuilding a complete macOS image.

To create the OpenCore boot image:
```
$ bazel build //opencore:image
```

The resulting image is bootable both in BIOS and UEFI mode (although UEFI is
recommended).

## Linux
In order to create a minimally bootable Linux image, linuxkit is used.

To create the linuxkit image, make sure Docker for Mac is running and then:
```
$ bazel build //linuxkit:image
```

This can take 90s+ so don't be worried.

This image expects the macOS image to be a block device attached to the virtual
machine on `/dev/sdb` and will try to boot from it.

### booter
The boot sequence is coordinated by `//linuxkit/boot:boot`.
It will:
- load the nescessary kernel modules (`kvm`, `kvm-intel` or `kvm-amd`)
- setup NAT for the VM (`tap` interface and packet forwarding)
- build the QEMU command line dynamically based on available CPU cores, memory
  and so on
