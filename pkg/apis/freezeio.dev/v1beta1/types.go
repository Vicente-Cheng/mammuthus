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
// +kubebuilder:resource:shortName=nfsexport;nfsexports,scope=Namespaced
// +kubebuilder:printcolumn:name="ExportID",type="string",JSONPath=`.spec.exportID`
// +kubebuilder:printcolumn:name="Path",type="string",JSONPath=`.spec.exportPath`
// +kubebuilder:printcolumn:name="PseudoPath",type="string",JSONPath=`.spec.exportPseudoPath`
// +kubebuilder:printcolumn:name="Node",type="string",JSONPath=`.spec.nodeName`
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=`.status.exportStatus`
// +kubebuilder:subresource:status

type NFSExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NFSExportSpec   `json:"spec"`
	Status            NFSExportStatus `json:"status,omitempty"`
}

// NFSExportSpec defines the desired state of NFSExport
type NFSExportSpec struct {
	// +kubebuilder:validation:Required
	NodeName string `json:"nodeName"`

	// +kubebuilder:validation:Required
	ExportID int `json:"exportID"`

	// +kubebuilder:validation:Required
	ExportPath string `json:"exportPath"`

	// +kubebuilder:validation:Required
	ExportPseudoPath string `json:"exportPseudoPath"`

	// +kubebuilder:validation:Required
	AccessType string `json:"accessType"`

	// +kubebuilder:validation:Optional
	Squash string `json:"squash,omitempty"`

	// +kubebuilder:validation:Optional
	SecType string `json:"secType,omitempty"`

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`

	// +kubebuilder:validation:Required
	FSAL *FSAL `json:"fsal"`
}

// only VFS could skip the extra struct, because VFS did not have any extra fields
type FSAL struct {
	// +kubebuilder:validation:Required
	FSALType string `json:"fsalType"`
}

type NFSExportStatus struct {
	// +kubebuilder:validation:Optional
	ExportID int `json:"exportID,omitempty"`

	// +kubebuilder:validation:Optional
	ExportPath string `json:"exportPath,omitempty"`

	// +kubebuilder:validation:Optional
	ExportPseudoPath string `json:"exportPseudoPath,omitempty"`

	// +kubebuilder:validation:Optional
	AccessType string `json:"accessType,omitempty"`

	// +kubebuilder:validation:Optional
	Squash string `json:"squash,omitempty"`

	// +kubebuilder:validation:Optional
	SecType string `json:"secType,omitempty"`

	// +kubebuilder:validation:Optional
	FSAL *FSAL `json:"fsal,omitempty"`

	// +kubebuilder:validation:Required
	ExportStatus ExportStatus `json:"exportStatus"`
}
