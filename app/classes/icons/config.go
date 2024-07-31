package icons

// Configuration ...
type Configuration struct {
	WebPathPrefix string
	ResPathPrefix string
	Priority      int // 注册信息的优先级
}

// ConfigProvider ...
type ConfigProvider interface {
	Configuration() *Configuration
}

////////////////////////////////////////////////////////////////////////////////

// ConfigProviderImpl ...
type ConfigProviderImpl struct {

	//starter:component

	_as func(ConfigProvider) //starter:as("#")

	WebPathPrefix string //starter:inject("${mimetypes.icons.web-path-prefix}")
	ResPathPrefix string //starter:inject("${mimetypes.icons.res-path-prefix}")
	RegPriority   int    //starter:inject("${mimetypes.icons.registration-priority}")

}

func (inst *ConfigProviderImpl) _impl() ConfigProvider { return inst }

// Configuration ...
func (inst *ConfigProviderImpl) Configuration() *Configuration {
	return &Configuration{
		WebPathPrefix: inst.WebPathPrefix,
		ResPathPrefix: inst.ResPathPrefix,
		Priority:      inst.RegPriority,
	}
}
