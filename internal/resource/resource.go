package resource

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"sync"

	"xiaren/internal/constant"
	"xiaren/internal/log"
)

type Resource struct {
	path string
}

func (r *Resource) GetAsset(ctx context.Context, entry string) (io.Reader, error) {
	logger := log.GetLog(context.WithValue(ctx, constant.CTX_LOGGER, "resource.GetAsset()"))
	logger.Debug("starting getting an asset ", "entry", entry)
	if r.isPathDir() {
		logger.Debug("loading from a dir")
		entryPath := filepath.Join(r.path, entry)
		file, err := os.Open(entryPath)
		if err != nil {
			logger.Error("cannot load file", "entry", entry, "error", err)
			return nil, err
		}
		logger.Debug("loading file done", "entry", entry)
		return io.Reader(file), nil
	}
	logger.Debug("loading from a file")
	// todo: load from archive
	logger.Debug("loading file done", "entry", entry)
	return nil, nil
}

func (r *Resource) isPathDir() bool {
	fileStat, _ := os.Stat(r.path)
	return fileStat.IsDir()
}

// Singleton instance
var _resourceInstance *Resource

func InitResource(path string) error {
	return sync.OnceValue(func() error {
		// if path is invalid, return error
		_, err := os.Stat(path)
		if err != nil {
			return err
		}
		_resourceInstance = &Resource{
			path: path,
		}
		return nil
	})()
}

func GetResource() *Resource {
	return _resourceInstance
}
