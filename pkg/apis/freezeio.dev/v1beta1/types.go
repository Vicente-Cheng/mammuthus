package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExportStatus string

const (
	// NFSExportStatusUnapplied means the NFSExport is not applied to the NFS server
	NFSExportStatusUnapplied ExportStatus = "Unapplied"
	// NFSExportStatusApplied means the NFSExport is applied to the NFS server
	NFSExportStatusApplied ExportStatus = "Applied"
	// NFSExportStatusFailed means the NFSExport is failed to apply to the NFS server
	NFSExportStatusFailed ExportStatus = "Failed"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NFSExport is the schema for the nfsexports API
type NFSExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NFSExportSpec   `json:"spec"`
	Status            NFSExportStatus `json:"status,omitempty"`
}

// NFSExportSpec defines the desired state of NFSExport
type NFSExportSpec struct {
	NodeName         string `json:"nodeName"`
	ExportID         int    `json:"exportID"`
	ExportPath       string `json:"exportPath"`
	ExportPseudoPath string `json:"exportPseudoPath"`
	AccessType       string `json:"accessType"`
	Squash           string `json:"squash,omitempty"`
	SecType          string `json:"secType,omitempty"`
	Enabled          bool   `json:"enabled"`
	FSAL             *FSAL  `json:"fsal"`
}

// FSAL defines the File System Abstraction Layer configuration
type FSAL struct {
	FSALType string `json:"fsalType"`
}

// NFSExportStatus defines the observed state of NFSExport
type NFSExportStatus struct {
	ExportID         int          `json:"exportID,omitempty"`
	ExportPath       string       `json:"exportPath,omitempty"`
	ExportPseudoPath string       `json:"exportPseudoPath,omitempty"`
	AccessType       string       `json:"accessType,omitempty"`
	Squash           string       `json:"squash,omitempty"`
	SecType          string       `json:"secType,omitempty"`
	FSAL             *FSAL        `json:"fsal,omitempty"`
	ExportStatus     ExportStatus `json:"exportStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NFSExportList contains a list of NFSExport
type NFSExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NFSExport `json:"items"`
}