package config

import "github.com/cfx/warehouses/app"

type HostPath struct {
	Host   string
	Path   string
	Params map[string]string
}

type TgChatConfig struct {
	Token  string
	ChatID int64
}
type File struct {
	Name string
}

func NewFile(name string) *File {
	return &File{
		Name: name,
	}
}
func (f *File) HostPathWithParams(key string) (*HostPath, error) {
	blockTokenData := new(HostPath)
	err := app.Config(f.Name).UnmarshalKey(key, blockTokenData)
	if err != nil {
		return nil, err
	}
	return blockTokenData, nil
}

func (f *File) MapStrToStr(key string) (map[string]string, error) {
	data := make(map[string]string)
	err := app.Config(f.Name).UnmarshalKey(key, &data)
	return data, err
}

// 读取tg 配置
func (f *File) TgChatConfig(key string) (*TgChatConfig, error) {
	data := &TgChatConfig{}
	err := app.Config(f.Name).UnmarshalKey(key, &data)
	return data, err
}

// 读取配置slice 配置
func (f *File) ReadSliceString(key string) ([]string, error) {
	data := make([]string, 0)
	err := app.Config(f.Name).UnmarshalKey(key, &data)
	return data, err
}
