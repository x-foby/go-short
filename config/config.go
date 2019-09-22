package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	// config имеет тип `interface{}` для того чтоб было возможно
	// реализовать нужные механики (чтение, запись, маршаллинг, анмаршаллинг и т.д.),
	// абстрагировавшись от конкретной структуры
	// Реально работа будет вестись со структурой, установленной через `Set(interface{})`
	config interface{}

	// filename хранит путь к файлу, в котором хранятся настройки.
	// Значение `filename` необходимо установить через метод `SetFilename(string)` до вызова
	// чтения
	filename string
)

// Ошибки
var (
	ErrNoFileName = errors.New("Не указан путь к файлу конфигураций")
)

// Set устанавливает ссылку на структуру кофигурации
func Set(c interface{}) {
	config = c
}

// SetFilename устанавливает путь к файлу с конфигурацией
func SetFilename(fn string) {
	filename = fn
}

// ReadFromFile читает кофигурации из файла и обновляет их
func ReadFromFile() error {
	return readFromFile(filename)
}

// WriteToFile записывает текущие кофигурации в файл
func WriteToFile() error {
	return writeToFile(filename)
}

// Update обновляет кофигурации
func Update(buf []byte) error {
	return json.Unmarshal(buf, &config)
}

// Backup сохраняет текущюю конфигурацию
func Backup() error {
	return writeToFile(filename + "~")
}

// Restore сохраняет текущюю конфигурацию
func Restore() error {
	if err := readFromFile(filename + "~"); err != nil {
		return err
	}

	return writeToFile(filename)
}

func readFromFile(fn string) error {
	if err := checkFilename(); err != nil {
		return err
	}

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	return Update(buf)
}

func writeToFile(fn string) error {
	if err := checkFilename(); err != nil {
		return err
	}

	buf, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fn, buf, 0755)
}

func checkFilename() error {
	if filename == "" {
		return ErrNoFileName
	}

	return nil
}
