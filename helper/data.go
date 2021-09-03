package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jormin/gacode/entity"
	"github.com/mitchellh/go-homedir"
)

// GetDatafilePath Get path of data file
func GetDatafilePath() string {
	home, _ := homedir.Dir()
	return fmt.Sprintf("%s/gcode", home)
}

// ReadData 读取数据
func ReadData() (*entity.Data, error) {
	path := GetDatafilePath()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data *entity.Data
	if string(b) == "" {
		data = NewData()
	} else {
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, err
		}
	}
	if data.Accounts == nil {
		data.Accounts = []entity.Account{}
	}
	return data, err
}

// WriteData 写入数据
func WriteData(data *entity.Data) error {
	path := GetDatafilePath()
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0777)
	return err
}

// Get new data
func NewData() *entity.Data {
	return &entity.Data{
		Accounts: []entity.Account{},
	}
}
