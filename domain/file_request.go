package domain

import "io/fs"

type FileItem struct {
	Size    int64       `json:"size"`
	Mode    fs.FileMode `json:"mode"`
	IsDir   bool        `json:"isDir"`
	Name    string      `json:"name"`
	ModTime string      `json:"modTime"`
}

type ListRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DownloadRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DeleteRequest struct {
	IsDir bool   `json:"isDir"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}
