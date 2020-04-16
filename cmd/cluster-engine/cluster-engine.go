package main

import (
	"os"

	clusterengine "github.com/netapp/cake/pkg/cluster-engine"
)

func main() {
	clusterengine.Execute()
	os.Exit(0)
}
