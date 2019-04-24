#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

echo "Copying config..."
cp /config/vpn_server.config /usr/vpnserver/vpn_server.config

echo "Starting VPN server..."
/usr/bin/vpnserver execsvc