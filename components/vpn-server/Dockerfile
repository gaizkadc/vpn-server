#
# Copyright (C) 2018 Nalej Group - All Rights Reserved
#

#FROM siomiz/softethervpn:4.29-alpine
FROM debian:stretch

ADD dockerenv/startup.sh /startup.sh
RUN chmod 0755 /startup.sh
EXPOSE 5555/tcp 5666/tcp

# Now just add the binary\
RUN mkdir /nalej

# Now just add the binary
COPY vpn-server /nalej/

ENTRYPOINT ["/startup.sh"]