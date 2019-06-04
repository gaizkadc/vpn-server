#
# Copyright (C) 2018 Nalej Group - All Rights Reserved
#

FROM siomiz/softethervpn:4.29-alpine

RUN mkdir /config
RUN mkdir /vpn-server
ADD dockerenv/vpn_server.config /vpn-server/vpn_server.config
ADD dockerenv/startup.sh /config/startup.sh
RUN chmod 0755 /config/startup.sh
EXPOSE 5555/tcp 5666/tcp

# Now just add the binary\
RUN mkdir /nalej

# Now just add the binary
COPY vpn-server /nalej/

ENTRYPOINT ["/config/startup.sh"]