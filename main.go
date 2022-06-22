package main

import (
	"demo-azure-storage-blob/pkg/storage"
)

func main() {

	// // ============ Compress file ============
	// data1, _ := os.ReadFile("./file/000008.JPG")
	// data2, _ := os.ReadFile("./file/data.txt")

	// files := []compress.File{
	// 	compress.NewFile("img.jpg", data1),
	// 	compress.NewFile("data.txt", data2)}

	// // Create output file
	// out, err := os.Create("./temp/output.tar.gz")
	// if err != nil {
	// 	log.Fatalln("Error writing archive:", err)
	// }
	// defer out.Close()

	// compress.TarGzs(files, out)

	// // ============ Upload to Azure storage ============
	// data, err := os.ReadFile("./temp/output.tar.gz")
	// if err != nil {
	// 	panic(err)
	// }
	// storage.UploadBytesToBlob(data)

	// ============ Download from Azure storage ============
	storage.DownloadBlob("20220622-65ce4f66-496c-47ed-a6a9-dda92b58bd08.tar.gzip")

	// // Create output file
	// out, err := os.Create("output.tar.gz")
	// if err != nil {
	// 	log.Fatalln("Error writing archive:", err)
	// }
	// defer out.Close()
	// err = compress.TarGz(files, out)
	// if err != nil {
	// 	log.Fatalln("Error creating archive:", err)
	// }
	// fmt.Println("Archive created successfully")

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
