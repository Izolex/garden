FROM docker:20.10.12

# arm
RUN ln -s /usr/bin/dpkg-split /usr/sbin/dpkg-split
RUN ln -s /usr/bin/dpkg-deb /usr/sbin/dpkg-deb
RUN ln -s /bin/rm /usr/sbin/rm
RUN ln -s /bin/tar /usr/sbin/tar

COPY --from=docker/buildx-bin /buildx /usr/libexec/docker/cli-plugins/docker-buildx

ENV DOCKER_HOST=tcp://docker:2376
ENV DOCKER_TLS_CERTDIR="/certs"
ENV DOCKER_CERT_PATH="$DOCKER_TLS_CERTDIR/client"
ENV DOCKER_TLS_VERIFY=1
ENV DOCKER_CLI_EXPERIMENTAL=enabled
