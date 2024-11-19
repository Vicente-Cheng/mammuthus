package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=nfsexport;nfsexports,scope=Namespaced
// +kubebuilder:printcolumn:name="ExportID",type="string",JSONPath=`.spec.exportID`
// +kubebuilder:printcolumn:name="ExportPath",type="string",JSONPath=`.spec.exportPath`
// +kubebuilder:printcolumn:name="ExportPseudoPath",type="string",JSONPath=`.spec.exportPseudoPath`
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
	NodeID string `json:"nodeID"`

	// +kubebuilder:validation:Required
	ExportID int `json:"exportID"`

	// +kubebuilder:validation:Required
	ExportPath string `json:"exportPath"`

	// +kubebuilder:validation:Required
	ExportPseudoPath string `json:"exportPseudoPath"`

	// +kubebuilder:validation:Required
	AccessType string `json:"accessType"`

	FSAL *FSAL `json:"fsal,omitempty"`
}

type FSAL struct {
	VFS *VFSFSAL `json:"vfs,omitempty"`
}

type VFSFSAL struct {
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type NFSExportStatus struct {
	ExportID int `json:"exportID"`
}
