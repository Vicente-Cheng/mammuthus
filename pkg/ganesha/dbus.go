package ganesha

import (
	"fmt"

	dbus "github.com/godbus/dbus/v5"
	"github.com/sirupsen/logrus"
)

const (
	ganeshaServiceName = "org.ganesha.nfsd"

	ganeshaObjExportPath  = "/org/ganesha/nfsd/ExportMgr"
	exportInterface       = "org.ganesha.nfsd.exportmgr"
	addExportInterface    = exportInterface + ".AddExport"
	removeExportInterface = exportInterface + ".RemoveExport"
	updateExportInterface = exportInterface + ".UpdateExport"
)

func createDBusObject(objPath string) (dbus.BusObject, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	dbusObjPath := dbus.ObjectPath(objPath)
	obj := conn.Object(ganeshaServiceName, dbusObjPath)
	return obj, nil
}

func AddExport(exportName string, exportID int) error {
	targetConfigFile := fmt.Sprintf("%s/%s.conf", GaneshaConfDir, exportName)
	dbusObj, err := createDBusObject(ganeshaObjExportPath)
	if err != nil {
		return err
	}

	var reply string
	exportParam := fmt.Sprintf("EXPORT(export_id=%d)", exportID)
	if err = dbusObj.Call(addExportInterface, 0, targetConfigFile, exportParam).Store(&reply); err != nil {
		return err
	}

	logrus.Infof("Dbus call AddExport: %s", reply)

	return nil
}

func RemoveExport(exportID int) error {
	dbusObj, err := createDBusObject(ganeshaObjExportPath)
	if err != nil {
		return err
	}

	dbusExportID := uint16(exportID)
	if err = dbusObj.Call(removeExportInterface, 0, dbusExportID).Store(); err != nil {
		return err
	}

	return nil
}
