package ganesha

import (
	"fmt"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
)

func CreateConfig(configName string, targetExport Export) error {
	logrus.Infof("prepare to update Export: %v", targetExport)
	configFileName := fmt.Sprintf("/etc/ganesha/%s.conf", configName)

	tmpl, err := template.New(configName).Parse(ExportContentTemplateVFS)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Create the output configuration file
	targetFile, err := os.Create(configFileName)
	if err != nil {
		return fmt.Errorf("error creating target file: %v", err)
	}
	defer targetFile.Close()

	// Execute the template with the Export data
	if err := tmpl.Execute(targetFile, targetExport); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Success message
	logrus.Infof("Configuration file %s created successfully.", configFileName)
	return nil
}
