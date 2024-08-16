package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Recurring struct {
	Days    int       `json:"days"`
	Start   time.Time `json:"start"`
	Name    string    `json:"name"`
	Project string    `json:"project"`
}

type Period struct {
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Name    string    `json:"name"`
	Project string    `json:"project"`
}

type All struct {
	Recurrings []Recurring `json:"recurrings"`
	Periods    []Period    `json:"periods"`
}

func LoadAll(path string) (All, error) {
	file, err := os.Open(path)
	if err != nil {
		return All{}, fmt.Errorf("could not open file: %w", err)
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return All{}, fmt.Errorf("could not read file: %w", err)
	}

	all := All{}
	if err := json.Unmarshal(data, &all); err != nil {
		return All{}, fmt.Errorf("could not parse file: %w", err)
	}

	return all, nil
}
