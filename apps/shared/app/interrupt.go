package app

import (
	"os"
	"os/signal"
)

func OnInterrupt() <-chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	return signalChan
}
