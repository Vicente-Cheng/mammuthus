package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestGeneratedTypes ensures the generated code compiles correctly
func TestGeneratedTypes() {
	// Test creating an NFSExport
	export := &NFSExport{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "freezeio.dev/v1beta1",
			Kind:       "NFSExport",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-export",
			Namespace: "default",
		},
		Spec: NFSExportSpec{
			NodeName:         "test-node",
			ExportID:         1,
			ExportPath:       "/test/path",
			ExportPseudoPath: "/test",
			AccessType:       "RW",
			Enabled:          true,
			FSAL: &FSAL{
				FSALType: "VFS",
			},
		},
	}

	// Test deepcopy functionality
	exportCopy := export.DeepCopy()
	_ = exportCopy

	// Test creating an NFSExportList
	exportList := &NFSExportList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "freezeio.dev/v1beta1",
			Kind:       "NFSExportList",
		},
		Items: []NFSExport{*export},
	}

	// Test deepcopy functionality for list
	exportListCopy := exportList.DeepCopy()
	_ = exportListCopy
}