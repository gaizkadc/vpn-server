#!/bin/sh

export PATH=/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin

# Install VPN Server IF it's not yet installed
if [ ! -d "$DIRECTORY" ]; then
    wget https://github.com/SoftEtherVPN/SoftEtherVPN_Stable/releases/download/v4.29-9680-rtm/softether-vpnserver-v4.29-9680-rtm-2019.02.28-linux-x64-64bit.tar.gz
    tar -xvzf softether-vpnserver-v4.29-9680-rtm-2019.02.28-linux-x64-64bit.tar.gz
    rm softether-vpnserver-v4.29-9680-rtm-2019.02.28-linux-x64-64bit.tar.gz
    cd vpnserver
    make
fi

echo "Starting VPN server..."
vpnserver start

echo "Starting VPN Server code..."
/nalej/vpn-server $@
