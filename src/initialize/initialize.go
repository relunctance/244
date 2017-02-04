package initialize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func InitConfig(configPath string) (*Configuration, error) {
	conf, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("fetch config path faild, err: %v", err)
	}
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, fmt.Errorf("read config file:[%s]  faild, err: %v", conf, err)
	}
	config := new(Configuration)
	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("json_decode config file faild, err: %v", err)
	}
	return config, nil

}
