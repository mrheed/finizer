package lib

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Dir struct {
	UpDir          string
	PrevDir        string
	Stairs         int
	HighestStairs  int
	MaxStairs      int
	FolderExplored int
	FileExplored   int
	TotalItems     int
}

func NewDir(UpDir string, MaxStairs ...int) *Dir {
	maxStairs := 0
	if len(MaxStairs) >= 1 {
		maxStairs = MaxStairs[0]
	}
	if UpDir == "" {
		UpDir = "."
	}
	return &Dir{
		UpDir:     UpDir,
		PrevDir:   "",
		MaxStairs: maxStairs,
	}
}

func (d *Dir) ReadDir() error {
	var currentDir string = d.PrevDir + "/" + d.UpDir
	if d.PrevDir == "" {
		currentDir = d.UpDir
	}
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		return err
	}
	if d.PrevDir == "" {
		d.FolderExplored, d.TotalItems = 1, 1
	}
	for _, file := range files {
		sparated := strings.Split(currentDir+"/"+file.Name(), "/")
		d.Stairs = len(sparated) - 1
		if d.HighestStairs < d.Stairs {
			d.HighestStairs = d.Stairs
		}
		if file.IsDir() {
			d.UpDir, d.PrevDir = file.Name(), currentDir
			d.FolderExplored++
			d.TotalItems++
			if d.Stairs == d.MaxStairs {
				continue
			}
			if !mimeDirIsExist(file.Name()) {
				d.ReadDir()
			}
		} else {
			d.TotalItems++
			d.FileExplored++
			fileInf := NewFile(file.Name())
			fileInf.Investigate()
			targetDir := currentDir + "/" + fileInf.Folder + "/" + file.Name()
			source := currentDir + "/" + file.Name()
			fmt.Println("Moving", file.Name(), "to", targetDir)
			if err := MoveFile(file.Size(), source, targetDir); err != nil {
				return err
			}
			fmt.Println(file.Name(), "moved successfully")
		}
	}
	return nil
}

func mimeDirIsExist(dirName string) bool {
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
