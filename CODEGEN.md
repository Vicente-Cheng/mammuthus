# Code Generation

This project uses standard Kubernetes code generation tools to generate client code for the NFSExport Custom Resource.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) (for Dapper)
- [Dapper](https://github.com/rancher/dapper) (automatically downloaded by Makefile)

## Generated Code

The following code is automatically generated:

- `pkg/apis/freezeio.dev/v1beta1/zz_generated.deepcopy.go` - DeepCopy methods
- `pkg/client/clientset/` - Typed Kubernetes clientset
- `pkg/client/listers/` - Resource listers
- `pkg/client/informers/` - Kubernetes informers

## Code Generation Commands

### Generate all client code
```bash
make generate
```

### Clean generated code
```bash
make clean-generate
```

### Regenerate everything (clean + generate)
```bash
make regenerate
```

### Using go generate
```bash
go generate ./...
```

## Manual Generation (without Dapper)

If you prefer to run code generation locally:

```bash
./hack/update-codegen.sh
```

## Tools Used

- `deepcopy-gen` - Generates DeepCopy methods
- `client-gen` - Generates typed clientsets
- `lister-gen` - Generates resource listers
- `informer-gen` - Generates Kubernetes informers

All tools are from the standard `k8s.io/code-generator` package (v0.33.1).