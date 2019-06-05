#!/bin/bash

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

# Install VPN Server IF it's not yet installed
if [ ! -d "$/vpnserver/vpnserver" ]; then
    vpnserver stop
    cd /
    cp -r /usr/vpnserver /
    cp /vpn_server.config /vpnserver/vpn_server.config
fi

echo "Starting VPN server..."
/vpnserver/vpnserver start

echo "Starting VPN Server code..."
/nalej/vpn-server $@
