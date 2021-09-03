package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/jormin/gacode/entity"
)

// GetDatafilePath Get path of data file
func GetDatafilePath() string {
	u, _ := user.Current()
	return fmt.Sprintf("%s/gacode.json", u.HomeDir)
}

// ReadData read data from file
func ReadData() (*entity.Data, error) {
	path := GetDatafilePath()
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			_, _ = os.Create(path)
		}
		err = nil
	}
	b, _ := ioutil.ReadFile(path)
	var data *entity.Data
	if string(b) == "" {
		data = NewData()
	} else {
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, err
		}
	}
	return data, err
}

// WriteData write data to file
func WriteData(data *entity.Data) error {
	path := GetDatafilePath()
	b, _ := json.Marshal(data)
	err := ioutil.WriteFile(path, b, 0777)
	return err
}

// NewData Get new data
func NewData() *entity.Data {
	return &entity.Data{
		Accounts: []*entity.Account{},
	}
}
