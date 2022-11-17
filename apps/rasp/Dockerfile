# todo - how to extends /docker/go/Dockerfile ? :(
ARG BASE_IMAGE
ARG SRC=/usr/src
ARG SRC_APP=${SRC}/app/src
ARG BIN=/usr/local/bin/main
ARG TZ=Europe/Prague


FROM ${BASE_IMAGE} AS source

ARG APP
ARG SRC
ARG SRC_APP

COPY shared "${SRC}/shared"
COPY "${APP}/src" "${SRC_APP}"
RUN cd "${SRC_APP}" && go mod vendor


FROM ${BASE_IMAGE} AS dev

ARG TZ
ARG SRC
ARG SRC_APP
ARG BIN
ARG DB_PATH

ENV TZ=${TZ}
ENV BIN=${BIN}
ENV SRC=${SRC}
ENV SRC_APP=${SRC_APP}
ENV DB_PATH=${DB_PATH}

RUN apt-get update
RUN apt-get install -y vim iputils-ping curl inotify-tools
RUN go install github.com/go-delve/delve/cmd/dlv@v1.7.3

COPY --from=source "${SRC}" "${SRC}"
COPY docker/go/debug.sh /debug.sh

WORKDIR "${SRC_APP}"

COPY rasp/db.sql .
RUN apt-get install -y sqlite3
RUN dir=$(dirname "${DB_PATH}") && mkdir -p "${dir}"
RUN sqlite3 "${DB_PATH}" < db.sql

EXPOSE 80 40000
CMD sh /debug.sh
