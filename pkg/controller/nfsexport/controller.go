package nfsexport

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"

	nfsexportv1 "github.com/Vicente-Cheng/mammuthus/pkg/apis/freezeio.dev/v1beta1"
	"github.com/Vicente-Cheng/mammuthus/pkg/ganesha"
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
		logrus.Infof("Skip this round because the nfs export is deleting")
		return nil, nil
	}
	logrus.Infof("Handling NFS Export: %s change event", nfsExport.Name)

	if nfsExport.Spec.NodeName != c.nodeName {
		logrus.Infof("Skip this round because the target node (%s) is not the current node", nfsExport.Spec.NodeName)
		return nil, nil
	}

	// ExportID is 0 means the nfs export nevert applied.
	if nfsExport.Status.ExportID == 0 && !nfsExport.Spec.Enabled {
		logrus.Info("Skip this round because the nfs export is not enabled ever.")
		return nil, nil
	}

	squashVal := ganesha.GaneshaDefaultSquash
	if nfsExport.Spec.Squash != "" {
		squashVal = nfsExport.Spec.Squash
	}
	secTypeVal := ganesha.GaneshaDefaultSecType
	if nfsExport.Spec.SecType != "" {
		secTypeVal = nfsExport.Spec.SecType
	}
	targetExport := ganesha.Export{
		ExportID:   nfsExport.Spec.ExportID,
		Path:       nfsExport.Spec.ExportPath,
		Pseudo:     nfsExport.Spec.ExportPseudoPath,
		AccessType: nfsExport.Spec.AccessType,
		Squash:     squashVal,
		SecType:    secTypeVal,
		FSALName:   nfsExport.Spec.FSAL.FSALType,
	}

	logrus.Infof("prepare to update Export: %v", targetExport)
	if err := ganesha.CreateConfig(nfsExport.Name, targetExport); err != nil {
		logrus.Errorf("failed to update export: %v", err)
		return nil, err
	}

	if err := ganesha.AddExport(nfsExport.Name, nfsExport.Spec.ExportID); err != nil {
		logrus.Errorf("failed to add export: %v", err)
		errString := err.Error()
		if strings.Contains(errString, "already active") {
			if !anyUpdate(nfsExport.Spec, nfsExport.Status) {
				logrus.Infof("return nil because the nfs export did not have any change")
				return nil, nil
			}
			logrus.Infof("We should remove the oldone and readd a new one")
		}
		return nil, err
	}

	nfsExportCpy := nfsExport.DeepCopy()
	// everything is fine, update status
	status := nfsexportv1.NFSExportStatus{
		ExportID:         nfsExport.Spec.ExportID,
		ExportPath:       nfsExport.Spec.ExportPath,
		ExportPseudoPath: nfsExport.Spec.ExportPseudoPath,
		AccessType:       nfsExport.Spec.AccessType,
		FSAL:             nfsExport.Spec.FSAL,
		ExportStatus:     nfsexportv1.NFSExportStatusApplied,
	}
	if nfsExport.Spec.Squash != "" {
		status.Squash = nfsExport.Spec.Squash
	}
	if nfsExport.Spec.SecType != "" {
		status.SecType = nfsExport.Spec.SecType
	}
	nfsExportCpy.Status = status

	return c.NFSExports.UpdateStatus(nfsExportCpy)
}

func (c *Controller) OnNFSExportRemove(_ string, nfsExport *nfsexportv1.NFSExport) (*nfsexportv1.NFSExport, error) {
	if nfsExport == nil {
		logrus.Infof("Skip this round because the nfs export is deleted")
		return nil, nil
	}
	logrus.Infof("Handling NFS Export %s delete event", nfsExport.Name)

	// ExportID is 0 means the nfs export nevert applied.
	if nfsExport.Status.ExportID == 0 {
		logrus.Infof("Skip this round because the nfs export is not applied")
		return nil, nil
	}

	if err := ganesha.RemoveExport(nfsExport.Spec.ExportID); err != nil {
		logrus.Errorf("failed to remove export: %v", err)
		return nil, err
	}

	return nil, nil
}

func anyUpdate(spec nfsexportv1.NFSExportSpec, status nfsexportv1.NFSExportStatus) bool {
	return spec.ExportID != status.ExportID ||
		spec.ExportPath != status.ExportPath ||
		spec.ExportPseudoPath != status.ExportPseudoPath ||
		spec.AccessType != status.AccessType ||
		spec.Squash != status.Squash ||
		spec.SecType != status.SecType ||
		spec.FSAL.FSALType != status.FSAL.FSALType
}
