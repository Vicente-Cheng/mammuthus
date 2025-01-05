package ganesha

import (
	cmd "github.com/harvester/go-common/command"
)

type Export struct {
	ExportID   int
	Path       string
	Pseudo     string
	AccessType string
	Squash     string
	SecType    string
	FSALName   string
}

const ExportContentTemplateVFS = `EXPORT
{
    Export_Id = {{.ExportID}};
    Path = {{.Path}};
    Pseudo = {{.Pseudo}};
    Access_Type = {{.AccessType}};
    Squash = {{.Squash}};
    SecType = {{.SecType}};
    FSAL {
        Name = {{.FSALName}};
    }
}`

const (
	GaneshaCMD            = "ganesha.nfsd"
	GaneshaPID            = "/var/run/ganesha.pid"
	GaneshaConfDir        = "/etc/ganesha"
	GaneshaDefaultConfig  = GaneshaConfDir + "/ganesha.conf"
	GaneshaDefaultSquash  = "no_root_squash"
	GaneshaDefaultSecType = "sys"
	LogTarget             = "/proc/1/fd/1" // stdout for container log
)

func RunNFSGanesha() (string, error) {
	executor := cmd.NewExecutor()
	args := []string{"-f", GaneshaDefaultConfig, "-p", GaneshaPID, "-L", LogTarget}
	return executor.Execute(GaneshaCMD, args)
}

func RunDBus() (string, error) {
	executor := cmd.NewExecutor()
	return executor.Execute("dbus-daemon", []string{"--system"})
}
