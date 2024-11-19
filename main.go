//go:generate go run pkg/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go
//go:generate /bin/bash scripts/generate-manifest

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/rancher/wrangler/v3/pkg/kubeconfig"
	"github.com/rancher/wrangler/v3/pkg/leader"
	"github.com/rancher/wrangler/v3/pkg/signals"
	"github.com/rancher/wrangler/v3/pkg/start"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/kubernetes"

	"github.com/Vicente-Cheng/mammuthus/pkg/controller/nfsexport"
	exportv1 "github.com/Vicente-Cheng/mammuthus/pkg/generated/controllers/freezeio.dev"
	utils "github.com/Vicente-Cheng/mammuthus/pkg/utils"
)

const controllerName = "mammuthus-controller"

func main() {
	var opt utils.Option
	app := cli.NewApp()
	app.Name = controllerName
	app.Version = utils.FriendlyVersion()
	app.Usage = "mammuthus-controller help to manage the NFS export for NFS-Ganesha"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "kubeconfig",
			EnvVars:     []string{"KUBECONFIG"},
			Destination: &opt.KubeConfig,
			Usage:       "Kube config for accessing k8s cluster",
		},
		&cli.IntFlag{
			Name:        "threadiness",
			Value:       2,
			DefaultText: "2",
			Destination: &opt.Threadiness,
		},
		&cli.BoolFlag{
			Name:        "debug",
			EnvVars:     []string{"DEBUG"},
			Usage:       "enable debug logs",
			Destination: &opt.Debug,
		},
		&cli.StringFlag{
			Name:        "namespace",
			Value:       "freezeio-dev",
			DefaultText: "freezeio-dev",
			EnvVars:     []string{"NAMESPACE"},
			Destination: &opt.Namespace,
		},
	}

	app.Action = func(_ *cli.Context) error {
		initLogs(&opt)
		return run(&opt)
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func initLogs(opt *utils.Option) {
	if opt.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("Loglevel set to [%v]", logrus.DebugLevel)
	}
}

func run(opt *utils.Option) error {
	logrus.Infof("Mammuthus Controller %s is starting", utils.FriendlyVersion())
	if opt.Namespace == "" {
		return errors.New("namespace cannot be empty")
	}

	ctx := signals.SetupSignalContext()
	config, err := kubeconfig.GetNonInteractiveClientConfig(opt.KubeConfig).ClientConfig()
	if err != nil {
		return fmt.Errorf("failed to find kubeconfig: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error get client from kubeconfig: %s", err.Error())
	}

	clientNFSExport, err := exportv1.NewFactoryFromConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create NFSExport controller: %v", err)
	}

	nfsExportCtl := clientNFSExport.Freezeio().V1beta1().NFSExport()

	cb := func(ctx context.Context) {
		if err := nfsexport.Register(ctx, nfsExportCtl, opt); err != nil {
			logrus.Errorf("failed to register mammuthus controller: %v", err)
		}

		if err := start.All(ctx, opt.Threadiness, clientNFSExport); err != nil {
			logrus.Errorf("failed to start controller: %v", err)
		}

		<-ctx.Done()
	}

	leader.RunOrDie(ctx, opt.Namespace, controllerName, client, cb)

	logrus.Infof("%s is shutting down", controllerName)
	return nil
}
