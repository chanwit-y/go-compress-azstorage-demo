package main

import (
	"demo-azure-storage-blob/pkg/compress"
	"fmt"
	"log"
	"os"
)

func main() {
	files := []string{"./upload/000008.JPG", "./upload/data.txt"}

	// Create output file
	out, err := os.Create("output.tar.gz")
	if err != nil {
		log.Fatalln("Error writing archive:", err)
	}
	defer out.Close()
	err = compress.CreateArchive(files, out)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}

	fmt.Println("Archive created successfully")
	// data, _ := os.ReadFile("./output.tar.gz")
	// compress.ExtractTarGz(bytes.NewBuffer(data), "upload")
	// compress.Tar("./upload/data.txt", "./")
	// compress.UnGzip("./output.tar.gz", "./uncompress/")

	// data, err := os.ReadFile("./upload/file.gz")
	// if err != nil {
	// 	panic(err)
	// }
	// UploadBytesToBlob(data)
}
