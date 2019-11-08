
FROM siomiz/softethervpn:4.29-alpine

ADD dockerenv/startup.sh /startup.sh
RUN chmod 0755 /startup.sh
EXPOSE 5555/tcp 5666/tcp

# Now just add the binary\
RUN mkdir /nalej \
    && mkdir /vpnserver

# Now just add the binary
COPY vpn-server /nalej/
COPY dockerenv/vpn_server.config /vpn_server.config

ENTRYPOINT ["/startup.sh"]