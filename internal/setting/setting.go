package setting

import (
	"context"

	"github.com/BurntSushi/toml"

	"xiaren/internal/constant"
	"xiaren/internal/log"
	"xiaren/internal/saving"
)

type Setting struct {
	GameSetting   *GameSetting
	AudioSetting  *AudioSetting
	WindowSetting *WindowSetting
}

type GameSetting struct {
	Title string
}

type AudioSetting struct {
	EnableSound bool
	MusicVolume float64
	SoundVolume float64
	VoiceVolume float64
}

type WindowSetting struct {
	Resizable   bool
	Fullscreen  bool
	ResolutionX int
	ResolutionY int
}

func newDefaultSetting() *Setting {
	return &Setting{
		GameSetting: &GameSetting{
			Title: "XiaRen",
		},
		AudioSetting: &AudioSetting{
			EnableSound: true,
			MusicVolume: 0.5,
			SoundVolume: 0.5,
			VoiceVolume: 0.0,
		},
		WindowSetting: &WindowSetting{
			Resizable:   false,
			Fullscreen:  false,
			ResolutionX: 1280,
			ResolutionY: 720,
		},
	}
}

var _setting *Setting

func LoadSetting(ctx context.Context) error {
	logger := log.GetLog(context.WithValue(ctx, constant.CTX_LOGGER, "setting.LoadSetting()"))
	logger.Debug("Loading setting")
	var loadedSetting Setting
	gdata := saving.GetGData()

	if !gdata.ObjectExists(saving.SETTING_OBJ) && !gdata.ObjectPropExists(saving.SETTING_OBJ, "setting.toml") {
		// 配置文件不存在，使用默认配置
		logger.Debug("setting.toml not found, using default setting")
		_setting = newDefaultSetting()
		return SaveSetting(ctx)
	}
	data, err := gdata.LoadObjectProp(saving.SETTING_OBJ, "setting.toml")
	if err != nil {
		logger.Warn("Failed to load setting.toml", "error", err)
		_setting = newDefaultSetting()
		return SaveSetting(ctx)
	}

	// 尝试解析TOML配置
	if _, err := toml.Decode(string(data), &loadedSetting); err != nil {
		logger.Warn("Failed to decode setting.toml", "error", err)
		_setting = newDefaultSetting()
		return SaveSetting(ctx)
	}

	// 成功读取配置
	_setting = &loadedSetting
	logger.Debug("Setting loaded successfully")
	return nil
}

func SaveSetting(ctx context.Context) error {
	logger := log.GetLog(context.WithValue(ctx, constant.CTX_LOGGER, "setting.SaveSetting()"))
	logger.Debug("Saving setting")

	bytes, err := toml.Marshal(GetSetting())
	if err != nil {
		logger.Error("Failed to marshal setting", "error", err)
		return err
	}
	gdata := saving.GetGData()
	err = gdata.SaveObjectProp(saving.SETTING_OBJ, "setting.toml", bytes)
	if err != nil {
		logger.Error("Failed to save setting", "error", err)
		return err
	}
	return nil
}

func GetSetting() *Setting {
	if _setting == nil {
		_setting = newDefaultSetting()
	}
	return _setting
}
