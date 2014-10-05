package tools

import (
	"archive/zip"
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

//	unzips a file that contains one or more files to a destination
//	directory and creates the directory if it does not exist
func Unzip(src string, destRoot string) error {
	//	 open our archive for reading
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	//	make our dest dir if it does not exist
	if err = os.MkdirAll(destRoot, 755); err != nil {
		return err
	}

	for _, f := range reader.File {
		//	file name
		fPath := filepath.Join(destRoot, f.Name)

		//	output file
		fo, err := os.Create(fPath)
		if err != nil {
			return err
		}
		defer fo.Close()

		// make a write buffer
		w := bufio.NewWriter(fo)

		zippedFile, err := f.Open()
		if err != nil {
			return err
		}

		r := bufio.NewReader(zippedFile)

		// make a buffer to keep chunks that are read
		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := r.Read(buf)
			if err != nil && err != io.EOF {
				return err
			}
			if n == 0 {
				break
			}

			// write a chunk
			if _, err := w.Write(buf[:n]); err != nil {
				return err
			}
		}

		if err = w.Flush(); err != nil {
			return err
		}
	}

	return nil
}
