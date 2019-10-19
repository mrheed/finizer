package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var DefaultMime map[string][]string = map[string][]string{
	"Documents":  []string{"doc", "docx", "rtf", "pdf", "xls", "xlsx", "csv", "txt", "md"},
	"Music":      []string{"mp3"},
	"Compressed": []string{"zip", "rar", "tar.gz", "tar.xz", "tar"},
	"Video":      []string{"mp4", "mkv", "3gp"},
	"Mountable":  []string{"iso"},
	"Pictures":   []string{"jpg", "jpeg", "png", "gif"},
	"Swap":       []string{"swp", "swo", "swn"},
}

type File struct {
	customFormatPath    string
	ignoreDefaultFormat bool
	Skip                bool
	FileName            string
	FileSize            int64
	Extension           string
	Folder              string
	Path                string
}

type CustomRules struct {
	Folder             string
	PermitSizeOverflow bool `json:"permit_size_overflow" bson:"permit_size_overflow"`
	Format             string
	MaxFileSize        int    `json:"max_file_size" bson:"max_file_size"`
	BiggerSizeFolder   string `json:"bigger_size_folder" bson:"bigger_size_folder"`
}

func NewFile(fileName string, fileSize int64, customFormatPath string, ignoreDefaultFormat bool) *File {
	defaultFolder, defaultFormat := "Others", "unknown"
	return &File{
		FileName:            fileName,
		FileSize:            fileSize,
		Extension:           defaultFormat,
		Folder:              defaultFolder,
		ignoreDefaultFormat: ignoreDefaultFormat,
		customFormatPath:    customFormatPath,
		Path:                defaultFolder + "/" + fileName,
	}
}

func (f *File) Investigate() (*File, error) {
	if f.customFormatPath != "" {
		ok, err := f.investigateCustomRules()
		if err != nil {
			return &File{}, err
		}
		if ok {
			return f, nil
		}
	}
	if f.ignoreDefaultFormat {
		f.Skip = true
		return f, nil
	}
	for k, d := range DefaultMime {
		for _, z := range d {
			if strings.HasSuffix(f.FileName, "."+z) {
				f.Extension = z
				f.Folder = k
				f.Path = k + "/" + f.FileName
				return f, nil
			}
		}
	}
	return f, nil
}

func (f *File) investigateCustomRules() (bool, error) {
	customRules, err := UnmarshalCustomRules(f.customFormatPath)
	if err != nil {
		return false, err
	}
	for i, d := range customRules {
		if (d == CustomRules{}) {
			fmt.Println(fmt.Sprintln("Warning: rule at index", i, "being skipped due to empty value"))
			continue
		}
		if strings.HasSuffix(f.FileName, d.Format) {
			f.Extension, f.Folder = d.Format, d.Folder
			if d.MaxFileSize != 0 && (d.MaxFileSize < int(f.FileSize)/(1024)) {
				if d.PermitSizeOverflow {
					f.Folder = d.BiggerSizeFolder
				} else {
					f.Skip = true
				}
			}
			f.Path = f.Folder + "/" + f.FileName
			return true, nil
		}
	}
	return false, nil
}

func UnmarshalCustomRules(customFormatPath string) ([]CustomRules, error) {
	var result []CustomRules
	if !strings.HasSuffix(customFormatPath, ".json") {
		return result, errors.New("format error: file type must be JSON")
	}
	dByte, err := ioutil.ReadFile(customFormatPath)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(dByte, &result); err != nil {
		return result, err
	}
	return result, nil
}
