#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

echo "Copying config..."
cp /vpn-server/vpn_server.config /usr/vpnserver/vpn_server.config

echo "Starting VPN server..."
vpnserver start

echo "Starting VPN Server code..."
/nalej/vpn-server $@
