package tools

import (
	"errors"
	"io"
	"net/http"
	"os"
)

//	download a file via HTTP GET to a local destination
//	url: absolute URL where the file is located
//	dest: absolute path including the filename
func HTTPdownload(url, dest string) (err error) {
	//	make sure destination does not already exist
	if FileExist(dest) {
		return errors.New("dest file already exists")
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	// close the body on function complete
	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	// open output file
	file, err := os.Create(dest)
	if err != nil {
		return
	}

	// close file on exit and check for error
	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	return
}
