package config

import "testing"

var errorString = "Expected %v, got %v"

type Config struct {
	StringProperty *string                 `json:"stringProperty"`
	IntProperty    *int                    `json:"intProperty"`
	BoolProperty   *bool                   `json:"boolProperty"`
	SliceProperty  *[]interface{}          `json:"sliceProperty"`
	MapProperty    *map[string]interface{} `json:"mapProperty"`
}

func TestSet(t *testing.T) {
	var cfg Config

	Set(cfg)

	if config != cfg {
		t.Errorf(errorString, cfg, cfg)
	}
}

func TestSetFilename(t *testing.T) {
	fn := "./config.json"
	SetFilename(fn)

	if filename != fn {
		t.Errorf(errorString, fn, filename)
	}
}

func TestWriteToFile(t *testing.T) {
	var cfg Config
	var err error
	Set(&cfg)

	SetFilename("")

	stringProperty := "value"
	intProperty := 123
	boolProperty := true
	sliceProperty := []interface{}{1, 2}
	mapProperty := map[string]interface{}{"key": "value"}

	cfg.StringProperty = &stringProperty
	cfg.IntProperty = &intProperty
	cfg.BoolProperty = &boolProperty
	cfg.SliceProperty = &sliceProperty
	cfg.MapProperty = &mapProperty

	err = WriteToFile()
	if err != ErrNoFileName {
		t.Errorf(errorString, nil, err)
	}

	SetFilename("./config.json")

	err = WriteToFile()
	if err != nil {
		t.Errorf(errorString, nil, err)
	}

	Set(make(chan int))

	err = WriteToFile()
	if err == nil {
		t.Errorf(errorString, "error", err)
	}
}

func TestReadFromFile(t *testing.T) {
	var cfg Config
	var err error
	Set(&cfg)

	SetFilename("")

	err = ReadFromFile()
	if err != ErrNoFileName {
		t.Errorf(errorString, nil, err)
	}

	SetFilename("./config.json")

	err = ReadFromFile()
	if err != nil {
		t.Errorf(errorString, nil, err)
	}

	SetFilename("./config2.json")

	err = ReadFromFile()
	if err == nil {
		t.Errorf(errorString, "error", err)
	}
}

func TestUpdate(t *testing.T) {
	var cfg Config
	Set(&cfg)

	if err := Update([]byte(`{"stringProperty":"value","intProperty":123,"boolProperty":true,"sliceProperty":[1,2],"mapProperty":{"key":"value"}}`)); err != nil {
		t.Errorf(errorString, nil, err)
	}
}

func TestBackup(t *testing.T) {
	SetFilename("")
	if err := Backup(); err != ErrNoFileName {
		t.Errorf(errorString, ErrNoFileName, err)
	}

	SetFilename("./config.json")
	if err := Backup(); err != nil {
		t.Errorf(errorString, nil, err)
	}
}

func TestRestore(t *testing.T) {
	SetFilename("")
	if err := Restore(); err != ErrNoFileName {
		t.Errorf(errorString, ErrNoFileName, err)
	}

	SetFilename("./config.json")
	if err := Restore(); err != nil {
		t.Errorf(errorString, nil, err)
	}
}

// func readFromFile(fn string) error {
// 	if err := checkFilename(); err != nil {
// 		return err
// 	}

// 	buf, err := ioutil.ReadFile(fn)
// 	if err != nil {
// 		return err
// 	}

// 	return Update(buf)
// }

// func writeToFile(fn string) error {
// 	if err := checkFilename(); err != nil {
// 		return err
// 	}

// 	buf, err := json.MarshalIndent(config, "", "  ")
// 	if err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(fn, buf, 0755)
// }

// func checkFilename() error {
// 	if filename == "" {
// 		return ErrNoFileName
// 	}

// 	return nil
// }
