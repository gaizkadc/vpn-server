#
# Copyright (C) 2018 Nalej Group - All Rights Reserved
#

FROM siomiz/softethervpn:latest

RUN mkdir /config
ADD dockerenv/vpn_server.config /config/vpn_server.config
ADD dockerenv/startup.sh /config/startup.sh
RUN chmod 0755 /config/startup.sh
EXPOSE 443/tcp 5555/tcp

ENTRYPOINT ["/config/startup.sh"]