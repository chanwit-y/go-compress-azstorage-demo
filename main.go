package main

import (
	"bytes"
	"demo-azure-storage-blob/pkg/compress"
	"demo-azure-storage-blob/pkg/storage"
	"os"
)

func main() {

	// fileName := uuid.NewString()

	// // ============ Compress file ============
	// data1, _ := os.ReadFile("./example/000008.JPG")
	// data2, _ := os.ReadFile("./example/data.txt")

	// files := []compress.File{
	// 	compress.NewFile("img.jpg", data1),
	// 	compress.NewFile("data.txt", data2)}

	// compress.TarGzs(files, fmt.Sprintf("./temp/%s.tar.gz", fileName))

	// // ============ Upload to Azure storage ============
	// data, err := os.ReadFile(fmt.Sprintf("./temp/%s.tar.gz", fileName))
	// if err != nil {
	// 	panic(err)
	// }
	// downloadUrl, _ := storage.UploadBytesToBlob(data)
	// sourceURL, _ := url.Parse(downloadUrl)
	// blobName := path.Base(sourceURL.Path)
	// fmt.Println(blobName)

	// ============ Download from Azure storage ============
	storage.DownloadBlob("20220623-92a84e5a-df4c-4bcb-8e62-45e8dc95fd97.tar.gzip", "./download/")

	// ============ Extract file ============
	byteFile, _ := os.ReadFile("./download/20220623-92a84e5a-df4c-4bcb-8e62-45e8dc95fd97.tar.gzip")
	compress.ExtractTarGzs(bytes.NewBuffer(byteFile), "./extract/")
}
