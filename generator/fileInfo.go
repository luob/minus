package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	os.FileInfo
	baseDir string
}

func newFileInfo(baseDir string, fi os.FileInfo) *fileInfo {
	return &fileInfo{
		FileInfo: fi,
		baseDir:  baseDir,
	}
}

func (f *fileInfo) name() string {
	return f.FileInfo.Name()
}

func (f *fileInfo) isDir() bool {
	return f.FileInfo.IsDir()
}

func (f *fileInfo) ext() string {
	return filepath.Ext(f.name())
}

func (f *fileInfo) nameWithoutExt() string {
	return strings.TrimSuffix(f.name(), f.ext())
}

func (f *fileInfo) absFileName() string {
	return path.Join(f.baseDir, f.name())
}
