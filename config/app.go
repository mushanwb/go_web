package config

func init() {
	Add("app", StrMap{
		// 应用名称，暂时没有使用到
		"name": Env("APP_NAME", "GoWeb"),

		// 当前环境，用以区分多环境
		"env": Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": Env("APP_DEBUG", false),

		// 应用服务端口
		"port": Env("APP_PORT", "3000"),
	})
}
