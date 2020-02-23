#!/bin/bash

set -eu -o pipefail

host_device=eth0
tap_device=tap0
host_ip_addr=192.168.254.1
guest_ip_addr=192.168.254.2

ip tuntap add mode tap ${tap_device}
ip address add ${host_ip_addr}/24 dev ${tap_device}
ip link set ${tap_device} up

echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -A POSTROUTING -o ${host_device} -j MASQUERADE
iptables -I FORWARD 1 -i ${tap_device} -j ACCEPT
iptables -I FORWARD 1 -o ${tap_device} -m state --state RELATED,ESTABLISHED -j ACCEPT

# Forward SSH to the guest
iptables -t nat -I PREROUTING -p tcp --dport 2022 -j DNAT --to ${guest_ip_addr}:22
