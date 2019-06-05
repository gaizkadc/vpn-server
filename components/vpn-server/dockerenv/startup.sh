#!/bin/bash

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

# Stop old vpn server
vpnserver stop

# Install VPN Server IF it's not yet installed
CONFIG=/vpnserver/vpn_server.config
if [ ! -f "$CONFIG" ]; then
    echo "First time launching VPN server"
    cd /
    cp -r /usr/vpnserver /
    cp /vpn_server.config /vpnserver/vpn_server.config
fi

echo "Starting VPN server..."
/vpnserver/vpnserver start

echo "Starting VPN Server code..."
/nalej/vpn-server $@
