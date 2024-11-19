package main

import (
	"os"

	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"

	exportv1 "github.com/Vicente-Cheng/mammuthus/pkg/apis/freezeio.dev/v1beta1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/Vicente-Cheng/mammuthus/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"freezeio.dev": {
				Types: []interface{}{
					exportv1.NFSExport{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})
}
