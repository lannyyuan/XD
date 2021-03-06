package config

import (
	"os"
	"path/filepath"
	"xd/lib/configparser"
	"xd/lib/storage"
	"xd/lib/util"
)

// EnvRootDir is the name of the environmental variable to set the root storage directory at runtime
const EnvRootDir = "XD_HOME"

type StorageConfig struct {
	// downloads directory
	Downloads string
	// metadata directory
	Meta string
	// root directory
	Root string
}

func (cfg *StorageConfig) Load(s *configparser.Section) error {

	cfg.Root = "storage"
	if s != nil {
		cfg.Root = s.Get("rootdir", cfg.Root)
	}

	cfg.Meta = filepath.Join(cfg.Root, "metadata")

	cfg.Downloads = filepath.Join(cfg.Root, "downloads")
	if s != nil {
		cfg.Downloads = s.Get("downloads", cfg.Downloads)
	}
	return nil
}

func (cfg *StorageConfig) Save(s *configparser.Section) error {

	s.Add("rootdir", cfg.Root)
	s.Add("metadata", cfg.Meta)
	s.Add("downloads", cfg.Downloads)
	return nil
}

func (cfg *StorageConfig) LoadEnv() {
	dir := os.Getenv(EnvRootDir)
	if dir != "" {
		cfg.Root = dir
	}
}

func (cfg *StorageConfig) CreateStorage() storage.Storage {
	util.EnsureDir(cfg.Root)
	util.EnsureDir(cfg.Downloads)
	util.EnsureDir(cfg.Meta)
	return &storage.FsStorage{
		DataDir: cfg.Downloads,
		MetaDir: cfg.Meta,
	}
}
