package main

import (
	"github.com/balchua/pod-finder/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("pod-finder started")
	cmd.Execute()

}
