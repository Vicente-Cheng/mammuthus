package nfsexport

import (
	"context"

	"github.com/sirupsen/logrus"

	nfsexportv1 "github.com/Vicente-Cheng/mammuthus/pkg/apis/freezeio.dev/v1beta1"
	ctlnfsexportv1 "github.com/Vicente-Cheng/mammuthus/pkg/generated/controllers/freezeio.dev/v1beta1"
	"github.com/Vicente-Cheng/mammuthus/pkg/utils"
)

type Controller struct {
	namespace string
	nodeName  string

	NFSExports     ctlnfsexportv1.NFSExportController
	NFSExportCache ctlnfsexportv1.NFSExportCache
}

const (
	nfsExportHandlerName = "mammuthus-nfs-export-handler"
)

func Register(ctx context.Context, nfsexports ctlnfsexportv1.NFSExportController, opt *utils.Option) error {

	c := &Controller{
		namespace:      opt.Namespace,
		nodeName:       opt.NodeName,
		NFSExports:     nfsexports,
		NFSExportCache: nfsexports.Cache(),
	}

	c.NFSExports.OnChange(ctx, nfsExportHandlerName, c.OnNFSExportChange)
	c.NFSExports.OnRemove(ctx, nfsExportHandlerName, c.OnNFSExportRemove)
	return nil
}

func (c *Controller) OnNFSExportChange(_ string, nfsExport *nfsexportv1.NFSExport) (*nfsexportv1.NFSExport, error) {
	if nfsExport == nil || nfsExport.DeletionTimestamp != nil {
		logrus.Infof("Skip this round because the network filesystem is deleting")
		return nil, nil
	}
	logrus.Infof("Handling network filesystem %s change event", nfsExport.Name)

	// To be implemented

	return nil, nil
}

func (c *Controller) OnNFSExportRemove(_ string, nfsExport *nfsexportv1.NFSExport) (*nfsexportv1.NFSExport, error) {
	if nfsExport == nil || nfsExport.DeletionTimestamp != nil {
		logrus.Infof("Skip this round because the network filesystem is deleting")
		return nil, nil
	}
	logrus.Infof("Handling network filesystem %s delete event", nfsExport.Name)

	// To be implemented

	return nil, nil
}
