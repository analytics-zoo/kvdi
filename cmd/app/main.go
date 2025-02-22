/*
Copyright 2020,2021 Avi Zimmerman

This file is part of kvdi.

kvdi is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

kvdi is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with kvdi.  If not, see <https://www.gnu.org/licenses/>.
*/

// The main entrypoint to the kVDI App/API server.
package main

import (
	"flag"
	"fmt"
	"os"

	v1 "github.com/kvdi/kvdi/apis/meta/v1"
	"github.com/kvdi/kvdi/pkg/util/common"
	"github.com/kvdi/kvdi/pkg/util/tlsutil"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var applogger = logf.Log.WithName("app")

func main() {
	var vdiCluster string
	var enableCORS bool
	flag.StringVar(&vdiCluster, "vdi-cluster", "", "The VDICluster this application is serving")
	flag.BoolVar(&enableCORS, "enable-cors", false, "Add CORS headers to requests")
	common.ParseFlagsAndSetupLogging()

	common.PrintVersion(applogger)

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		applogger.Error(err, "Failed to load kubernetes configuration")
		os.Exit(1)
	}

	// build the server
	srvr, err := newServer(cfg, vdiCluster, enableCORS)
	if err != nil {
		applogger.Error(err, "Failed to build the server router")
		os.Exit(1)
	}

	// serve
	applogger.Info(fmt.Sprintf("Starting VDI cluster frontend on :%d", v1.WebPort))
	if err := srvr.ListenAndServeTLS(tlsutil.ServerKeypair()); err != nil {
		applogger.Error(err, "Failed to start https server")
		os.Exit(1)
	}
}
