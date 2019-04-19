package backend

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/meltwater/drone-cache/cache"
	"github.com/pkg/errors"
	"io"
)

// azureBlobBackend is an Azure Blob implementation of the Backend
type azureBlobBackend struct {
	account-name   string
	account-key    string
	container-name string
}

func newAzureBlob()

// reading

// writing

