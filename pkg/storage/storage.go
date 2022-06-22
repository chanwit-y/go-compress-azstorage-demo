package storage

import (
	"context"
	"demo-azure-storage-blob/pkg/env"
	"fmt"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/google/uuid"
)

func GetAccountInfo() (string, string, string, string) {
	azrKey := env.Env().AZURE_ACCOUNT_KEY
	azrBlobAccountName := env.Env().AZURE_ACCOUNT_NAME
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	azrBlobContainer := env.Env().AZURE_BLOB_CONTAINER

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}

func GetBlobName() string {
	t := time.Now()
	uuid := uuid.New()

	// return fmt.Sprintf("%s-%v.jpg", t.Format("20060102"), uuid)
	// return fmt.Sprintf("%s-%v.txt", t.Format("20060102"), uuid)
	return fmt.Sprintf("%s-%v.tar.gzip", t.Format("20060102"), uuid)
}

// The below method assumes you already have the byte array ready to go
func UploadBytesToBlob(b []byte) (string, error) {
	azrKey, accountName, endPoint, container := GetAccountInfo()           // This is our account info method
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", GetBlobName())) // This uses our Blob Name Generator to create individual blob urls
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey) // Finally we create the credentials object required by the uploader
	if errC != nil {
		return "", errC
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Provide any needed options to UploadToBlockBlobOptions (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadToBlockBlobOptions)
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			// ContentType: "image/jpg", //  Add any needed headers here
			// ContentType: "image/jpg", //  Add any needed headers here
			ContentType: "application/zip", //  Add any needed headers here
		},
	}

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	return blockBlobUrl.String(), errU
}
