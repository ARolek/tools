package tools

import (
	"io"
	"net/http"
	"os"
)

//	download a file via HTTP GET to a local destination
func HTTPdownload(src string, dest string) error {
	var err error
	resp, err := http.Get(src)
	if err != nil {
		return err
	}

	// close the body on function complete
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	file, err := os.Create(dest)
	if err != nil {
		return err
	}

	// close file on exit and check for error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
