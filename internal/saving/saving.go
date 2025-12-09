package saving

import (
	"sync"

	gdata "github.com/quasilyte/gdata/v2"
)

const (
	SETTING_OBJ string = "setting"
	LOG_OBJ     string = "log"
	SAVE_OBJ    string = "save"
)

var _gdataInstance *gdata.Manager

func InitGData(appName string) error {
	return sync.OnceValue(func() error {
		var err error
		_gdataInstance, err = gdata.Open(
			gdata.Config{
				AppName: appName})
		if err != nil {
			return err
		}
		return nil
	})()
}

func GetGData() *gdata.Manager {
	return _gdataInstance
}
