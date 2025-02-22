//   Copyright Steve Sloka 2021
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"

	discoverygrpcv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	log "github.com/sirupsen/logrus"
	"github.com/stevesloka/envoy-xds-server/internal/callbacks"
	"github.com/stevesloka/envoy-xds-server/internal/processor"
	"github.com/stevesloka/envoy-xds-server/internal/server"
	"github.com/stevesloka/envoy-xds-server/internal/watcher"
)

var (
	l log.FieldLogger

	watchDirectoryFileName string
	port                   uint
	basePort               uint
	mode                   string

	nodeID string
)

func init() {
	l = log.New()
	log.SetLevel(log.DebugLevel)

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 9907, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Define the directory to watch for Envoy configuration files
	flag.StringVar(&watchDirectoryFileName, "watchDirectoryFileName", "config/config.yaml", "full path to directory to watch for files")
}

func main() {
	flag.Parse()

	// Context
	ctx := context.Background()

	// Create a cache
	cache := cache.NewSnapshotCache(true, cache.IDHash{}, nil)

	// Create a processor
	proc := processor.NewProcessor(
		ctx, cache, nodeID, log.WithField("context", "processor"))

	// Create a callbacks
	cbs := callbacks.NewCallbacks(ctx, log.WithField("context", "callbacks"), func(request *discoverygrpcv3.DiscoveryRequest) {
		if request.GetErrorDetail() == nil {
			if request.GetVersionInfo() == fmt.Sprintf("%d", proc.GetSnapshotVersion()) {
				l.Info("XDS Update Completed!")
			}
		}
	})

	// Create initial snapshot from file
	proc.ProcessFile(watcher.NotifyMessage{
		Operation: watcher.Create,
		FilePath:  watchDirectoryFileName,
	})

	// Notify channel for file system events
	notifyCh := make(chan watcher.NotifyMessage)

	go func() {
		// Watch for file changes
		watcher.Watch(watchDirectoryFileName, notifyCh)
	}()

	go func() {
		// Run the xDS server
		srv := serverv3.NewServer(ctx, cache, cbs)
		server.RunServer(ctx, srv, port)
	}()

	for {
		select {
		case msg := <-notifyCh:
			proc.ProcessFile(msg)
		}
	}
}
