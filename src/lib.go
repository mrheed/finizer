package lib

import (
  "errors"
  "io"
  "os"
  "strings"

  "github.com/cheggaaa/pb/v3"
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

func Sort(data *[]Moveable, low int, high int) {
  if low < high {
    pi := partition(*data, low, high)
    Sort(data, low, pi-1)
    Sort(data, pi+1, high)
  }
}

func partition(data []Moveable, low int, high int) int {
  pivot := data[high]
  i := low - 1
  for j := low; j < high; j++ {
    if data[j].StairsAt < pivot.StairsAt {
      i++
      temp := data[i]
      data[i] = data[j]
      data[j] = temp
    }
  }
  temp := data[i+1]
  data[i+1] = data[high]
  data[high] = temp
  return i + 1
}

func MoveFile(oldpath, newpath string) error {
  MakeDirWithFile(newpath)
  return os.Rename(oldpath, newpath)
}

func CopyFile(size int64, source, destination string) error {
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
