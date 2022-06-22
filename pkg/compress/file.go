package compress

import (
	"demo-azure-storage-blob/pkg/utils"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

type File struct {
	fileName  string
	fullPath  string
	folder    string
	byteArray []byte
	file      *os.File
	err       error
}

func NewFile(fileNmae string, byteArray []byte) File {
	return File{fileNmae, "", "", byteArray, nil, nil}
}

func (f *File) CreateFile() *File {
	f.fullPath = utils.IIf(
		f.folder != "",
		fmt.Sprintf("./temp/%s/%s", f.folder, f.fileName),
		fmt.Sprintf("./temp/%s", f.fileName))

	err := ioutil.WriteFile(f.fullPath, f.byteArray, 0644)
	if err != nil {
		f.fullPath = ""
		f.err = err
		return f
	}

	return f
}

func (f *File) CreateFolder() *File {
	f.folder = uuid.NewString()
	err := os.Mkdir(fmt.Sprintf("./temp/%s", f.folder), 0755)
	if err != nil {
		log.Fatal(err)
		f.err = err
		return f
	}

	return f
}

func (f *File) RemoveFolder() error {
	if f.err != nil {
		return f.err
	}

	err := os.RemoveAll(fmt.Sprintf("./temp/%s", f.folder))
	if err != nil {
		return err
	}

	f.file = nil
	f.fullPath = ""

	return nil
}

func (f *File) RemoveFile() error {
	if f.err != nil {
		return f.err
	}

	err := os.Remove(f.fullPath)
	if err != nil {
		return err
	}

	f.file = nil
	f.fullPath = ""

	return nil
}

func (f *File) GetFileInfo() (fs.FileInfo, error) {
	if f.err != nil {
		return nil, f.err
	}

	file, err := os.Open(f.fullPath)
	if err != nil {
		panic(err)
	}

	f.file = file
	info, _ := file.Stat()
	return info, nil
}
