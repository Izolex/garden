ARG BASE_IMAGE
FROM ${BASE_IMAGE}

RUN echo "deb http://ppa.launchpad.net/ansible/ansible/ubuntu focal main" >> /etc/apt/sources.list

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 93C4A3FD7BB9C367 && \
    apt-get update && \
    apt-get install -y ansible

COPY ./rasp/main /app/main
COPY ./rasp/db.sql /app/db.sql
COPY ./rasp-deploy/src /app/src
COPY ./shared /shared

WORKDIR /app/src
RUN go mod vendor

CMD ["go", "run", "."]
