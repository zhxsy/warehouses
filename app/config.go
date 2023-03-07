package app

import (
	"errors"
	"fmt"
	"github.com/cfx/warehouses/helper"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

var configs map[string]*viper.Viper

func InitConfig(apEnv, pathDir string) {
	//初始化
	configs = make(map[string]*viper.Viper)

	//获取全部配置文件
	env := GetEnv()
	envPath, ok := GetEnvPath(apEnv)
	if ok {
		pathDir = envPath + pathDir
	}
	var p = fmt.Sprintf("%s/%s", pathDir, env)

	files, err := ioutil.ReadDir(p)
	if err != nil {
		cp, _ := helper.GetCurrentPath()
		p = cp + p
		files, err = ioutil.ReadDir(p)
		if err != nil {
			panic(errors.New(fmt.Sprintf("init_config_error,path: %s,error: %s", p, err.Error())))
		}
	}

	//遍历所有配置文件
	for _, f := range files {
		//处理文件名

		fa := path.Base(f.Name())
		fs := path.Ext(f.Name())
		fp := fa[0 : len(fa)-len(fs)]

		//设置路径和格式
		config := viper.New()
		config.AddConfigPath(fp)
		config.SetConfigType("json")
		config.SetConfigFile(p + "/" + f.Name())

		// 读取配置
		if err := config.ReadInConfig(); err != nil {
			panic(errors.New(fmt.Sprintf("init_config_error,path: %s,file: %s, err: %s", p, f.Name(), err.Error())))
		}

		//装载
		configs[fp] = config
	}

	fmt.Printf("init_config_success, path: %s\n", p)
}

func Config(name string) *viper.Viper {
	c, ok := configs[name]
	if ok {
		return c
	}

	Log().WithField("name", name).Error("config_not_init")
	return nil
}

func InitYamlConfig(apEnv, pathDir string) {
	//初始化
	configs = make(map[string]*viper.Viper)

	//获取全部配置文件
	env := GetEnv()
	envPath, ok := GetEnvPath(apEnv)
	if ok {
		pathDir = envPath + pathDir
	}
	var p = fmt.Sprintf("%s/%s", pathDir, env)

	files, err := ioutil.ReadDir(p)
	if err != nil {
		cp, _ := helper.GetCurrentPath()
		p = cp + p
		files, err = ioutil.ReadDir(p)
		if err != nil {
			panic(errors.New(fmt.Sprintf("init_config_error,path: %s,error: %s", p, err.Error())))
		}
	}

	//遍历所有配置文件
	for _, f := range files {
		//处理文件名

		fa := path.Base(f.Name())
		fs := path.Ext(f.Name())
		fp := fa[0 : len(fa)-len(fs)]

		//设置路径和格式
		config := viper.New()
		config.AddConfigPath(fp)
		config.SetConfigType("yaml")
		config.SetConfigFile(p + "/" + f.Name())

		// 读取配置
		if err := config.ReadInConfig(); err != nil {
			panic(errors.New(fmt.Sprintf("init_config_error,path: %s,file: %s, err: %s", p, f.Name(), err.Error())))
		}

		//装载
		configs[fp] = config
	}

	fmt.Printf("init_config_success, path: %s\n", p)
}
