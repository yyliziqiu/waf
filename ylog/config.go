package ylog

func Initialize(c Config) {
	SetLevelString(c.Level)
	SetServiceId(c.ServiceId)
}
