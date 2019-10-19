package lib

import (
	"strings"
)

var DefaultMime map[string][]string = map[string][]string{
	"Documents":  []string{"doc", "rtf", "pdf", "xls", "xlsx", "csv", "txt", "md"},
	"Music":      []string{"mp3"},
	"Compressed": []string{"zip", "rar", "tar.gz", "tar.xz", "zip", "tar"},
	"Video":      []string{"mp4", "mkv", "3gp"},
	"Pictures":   []string{"jpg", "jpeg", "png"},
}

type File struct {
	FileName  string
	Extension string
	Folder    string
	Path      string
}

func NewFile(fileName string) *File {
	defaultFolder, defaultFormat := "Others", "unknown"
	return &File{
		FileName:  fileName,
		Extension: defaultFormat,
		Folder:    defaultFolder,
		Path:      defaultFolder + "/" + fileName,
	}
}

func (f *File) Investigate() {
	for k, d := range DefaultMime {
		for _, z := range d {
			if strings.HasSuffix(f.FileName, "."+z) {
				f.Extension = z
				f.Folder = k
				f.Path = k + "/" + f.FileName
			}
		}
	}
}
