FROM debian:bullseye-slim

# arm
RUN ln -s /usr/bin/dpkg-split /usr/sbin/dpkg-split
RUN ln -s /usr/bin/dpkg-deb /usr/sbin/dpkg-deb
RUN ln -s /bin/rm /usr/sbin/rm
RUN ln -s /bin/tar /usr/sbin/tar

RUN \
	apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y -q curl apt-transport-https ca-certificates
