package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	fi      os.FileInfo
	baseDir string
}

func newFileInfo(baseDir string, fi os.FileInfo) *fileInfo {
	return &fileInfo{
		fi:      fi,
		baseDir: baseDir,
	}
}

func (f fileInfo) name() string {
	return f.fi.Name()
}

func (f *fileInfo) getExt() string {
	return filepath.Ext(f.name())
}

func (f *fileInfo) getNameWithoutExt() string {
	return strings.TrimSuffix(f.name(), f.getExt())
}

func (f *fileInfo) getAbsFileName() string {
	return path.Join(f.baseDir, f.name())
}
