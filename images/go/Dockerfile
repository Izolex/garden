FROM golang:1.17.1

# arm
RUN ln -s /usr/bin/dpkg-split /usr/sbin/dpkg-split
RUN ln -s /usr/bin/dpkg-deb /usr/sbin/dpkg-deb
RUN ln -s /bin/rm /usr/sbin/rm
RUN ln -s /bin/tar /usr/sbin/tar

RUN apt-get update && \
    apt-get -y install tzdata gcc-arm-linux-gnueabihf

ENV TZ=Europe/Prague
