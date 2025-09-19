#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
MODULE_NAME="github.com/Vicente-Cheng/mammuthus"

# Generate deepcopy functions
echo "Generating deepcopy functions..."
go run k8s.io/code-generator/cmd/deepcopy-gen \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  --bounding-dirs "${MODULE_NAME}/pkg/apis" \
  --output-file zz_generated.deepcopy.go \
  "${MODULE_NAME}/pkg/apis/freezeio.dev/v1beta1"

# Generate clientset
echo "Generating clientset..."
go run k8s.io/code-generator/cmd/client-gen \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  --clientset-name "versioned" \
  --input-base "${MODULE_NAME}/pkg/apis" \
  --input "freezeio.dev/v1beta1" \
  --output-dir "${SCRIPT_ROOT}/pkg/client" \
  --output-pkg "${MODULE_NAME}/pkg/client/clientset"

# Generate listers
echo "Generating listers..."
go run k8s.io/code-generator/cmd/lister-gen \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  --input-dirs "${MODULE_NAME}/pkg/apis/freezeio.dev/v1beta1" \
  --output-dir "${SCRIPT_ROOT}/pkg/client" \
  --output-pkg "${MODULE_NAME}/pkg/client/listers"

# Generate informers
echo "Generating informers..."
go run k8s.io/code-generator/cmd/informer-gen \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  --input-dirs "${MODULE_NAME}/pkg/apis/freezeio.dev/v1beta1" \
  --versioned-clientset-package "${MODULE_NAME}/pkg/client/clientset/versioned" \
  --listers-package "${MODULE_NAME}/pkg/client/listers" \
  --output-dir "${SCRIPT_ROOT}/pkg/client" \
  --output-pkg "${MODULE_NAME}/pkg/client/informers"

echo "Code generation completed successfully!"