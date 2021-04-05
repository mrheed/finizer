package lib

import (
  "io/ioutil"
  "strings"
)

type Dir struct {
  UpDir               string
  PrevDir             string
  CustomFormatPath    string
  IgnoreDefaultFormat bool
  HighestStairs       int
  MaxStairs           int
  FolderExplored      int
  FileSkipped         int
  FileExplored        int
  TotalItems          int
}

type Moveable struct {
  FileName string
  Source   string
  Dest     string
  Size     int64
  StairsAt int
}

func NewDir(UpDir string, CustomFormatPath string, IgnoreDefaultFormat bool, MaxStairs ...int) *Dir {
  return &Dir{
    UpDir:               UpDir,
    CustomFormatPath:    CustomFormatPath,
    IgnoreDefaultFormat: IgnoreDefaultFormat,
    PrevDir:             "",
    FolderExplored:      1,
    TotalItems:          1,
    MaxStairs:           MaxStairs[0],
  }
}

func (d *Dir) ReadDir(stair int) ([]Moveable, error) {
  var result []Moveable
  var currentDir string = d.PrevDir + "/" + d.UpDir
  if d.PrevDir == "" {
    currentDir = d.UpDir
  }
  files, err := ioutil.ReadDir(currentDir)
  if err != nil {
    return []Moveable{}, err
  }
  for _, file := range files {
    d.TotalItems++
    if d.HighestStairs < stair {
      d.HighestStairs = stair
    }
    if file.IsDir() {
      d.FolderExplored++
      d.UpDir, d.PrevDir = file.Name(), currentDir
      if (d.MaxStairs >= 0 && stair >= d.MaxStairs) || d.mimeDirIsExist(file.Name()) {
        continue
      }
      moveable, err := d.ReadDir(stair + 1)
      if err != nil {
        return []Moveable{}, err
      }
      result = append(result, moveable...)
    } else {
      d.FileExplored++
      fileInf, err := NewFile(file.Name(), file.Size(), d.CustomFormatPath, d.IgnoreDefaultFormat).Investigate()
      if err != nil {
        return []Moveable{}, err
      }
      if fileInf.Skip {
        d.FileSkipped++
        continue
      }
      targetDir := currentDir + "/" + fileInf.Folder + "/" + file.Name()
      source := currentDir + "/" + file.Name()
      result = append(result, Moveable{FileName: file.Name(), Source: trimPrefix(source), Dest: trimPrefix(targetDir), StairsAt: stair, Size: file.Size()})
    }
  }
  return result, nil
}

func (d *Dir) mimeDirIsExist(dirName string) bool {
  if d.CustomFormatPath != "" {
    customRules, err := UnmarshalCustomRules(d.CustomFormatPath)
    if err != nil {
      return false
    }
    for _, d := range customRules {
      if d.Folder == dirName {
        return true
      }
    }
  }
  if d.IgnoreDefaultFormat {
    return false
  }
  if dirName == "Others" {
    return true
  }
  for k, _ := range DefaultMime {
    if k == dirName {
      return true
    }
  }
  return false
}

func trimPrefix(s string) string {
  return strings.TrimPrefix(s, "./")
}
