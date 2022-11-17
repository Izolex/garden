#!/bin/sh

while true; do
    echo "Starting debugging..."

    go mod vendor

    dlv debug --headless --log --listen=:40000 --api-version=2 --accept-multiclient --continue &

    PID=$!

    inotifywait -e modify -e move -e create -e delete -e attrib --exclude vendor --exclude __debug_bin -r "${SRC_APP}" -r "${SRC}/shared"

    echo "Stopping process id: $PID"

    kill -9 $PID
    pkill -f __debug_bin
done