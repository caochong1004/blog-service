package global

import (
	"github.com/longjoy/blog-service/pkg/logger"
	"github.com/longjoy/blog-service/pkg/setting"

)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
	JWTSetting *setting.JWTSettingS
	EmailSetting *setting.EmailSettingS
)

