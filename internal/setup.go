package internal

import (
	"context"
	"path/filepath"
	"xiaren/internal/config"
	"xiaren/internal/log"
	"xiaren/internal/resource"
	"xiaren/internal/saving"
	"xiaren/internal/setting"
)

func Setup(ctx context.Context, resourcePath string) error {
	var err error

	// init config
	err = config.LoadConfig(filepath.Join(".", "configs", "config.toml"))
	if err != nil {
		return err
	}

	// init load
	err = saving.InitGData(config.GetConfig().Title)
	if err != nil {
		return err
	}

	// init log
	err = log.InitLogger(config.GetConfig().Env)
	if err != nil {
		return err
	}
	logger := log.GetLog(context.Background())
	// init resources
	err = resource.InitResource(resourcePath)
	if err != nil {
		logger.Error("Error when loading resources")
		return err
	}
	// load setting
	err = setting.LoadSetting(ctx)
	if err != nil {
		// cannot read from file, use default setting
		logger.Error("Error when loading settings in main")
		return err
	}
	return nil
}

func GetSetting() *setting.Setting {
	return setting.GetSetting()
}
