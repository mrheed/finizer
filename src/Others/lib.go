package lib

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"strings"
)

func MakeDirWithFile(file string) error {
	sparated := strings.Split(file, "/")
	if sparated[0] == "." {
		sparated = sparated[1:]
	}
	dir := strings.Join(sparated[:len(sparated)-1], "/")
	_, err := os.Stat(dir)
	if err == nil {
		return errors.New("error: directory " + dir + " is exist")
	}
	return os.MkdirAll(dir, 0777)
}

func MakeFilePath(file, dir string) string {
	if dir == "." {
		return "" + file
	}
	return dir + "/" + file
}

func MoveFile(size int64, source, destination string) error {
	MakeDirWithFile(destination)

	// setup the loading animation
	load := pb.Full.Start64(size)

	// file reader
	input, err := os.Open(source)
	if err != nil {
		return err
	}

	// setup the loading reader
	loadReader := load.NewProxyReader(input)

	// file writer
	output, err := os.Create(destination)
	if err != nil {
		input.Close()
		return err
	}
	defer output.Close()
	_, err = io.Copy(output, loadReader)
	input.Close()
	if err != nil {
		return err
	}

	// finish loading animation
	load.Finish()
	//	err = os.Remove(source)
	//	if err != nil {
	//		return err
	//	}
	return nil
}
