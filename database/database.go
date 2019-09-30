package database

import (
	"database/sql"
	"fmt"
)

// Command опсиывает настройки запроса
type Command struct {
	Query string        `json:"query,omitempty"`
	Args  []interface{} `json:"args,omitempty"`
}

// ConnectionSetting опсиывает настройки подключения к БД
type ConnectionSetting struct {
	Driver                 string                 `json:"driver,omitempty"`
	ConnectionStringParams map[string]interface{} `json:"connectionStringParams,omitempty"`
	AfterConnection        []Command              `json:"afterConnection,omitempty"`
}

// DriverSetting опсиывает настройки драйвер к БД
type DriverSetting struct {
	GetConnectionString func(s map[string]interface{}) (string, error)
	AfterConnection     func(db *sql.DB, cs ConnectionSetting) error
}

// Settings описывает настройки БД по умолчанию и пул соединений
type Settings struct {
	Pool map[string]ConnectionSetting `json:"pool"`
}

var settings Settings
var pool = make(map[string]*sql.DB)
var drivers = make(map[string]DriverSetting)

// GetSettings возвращает указатель на настройки БД
func GetSettings() *Settings {
	return &settings
}

// RegisterDriver регистрирует драйвер БД и задаёт функцию, возвращающую connection string для подключения
func RegisterDriver(name string, driverSetting DriverSetting) {
	drivers[name] = driverSetting
}

// Open устанавливает соединение с базой данных из пула, если соединение ещё не установлено, и возвращает ссылку на него
func Open(name string) (*sql.DB, error) {
	connectionSettings, ok := settings.Pool[name]
	if !ok {
		return nil, fmt.Errorf("%q отсутствует в пуле соединений", name)
	}

	var err error
	pool[name], err = getConn(pool[name], connectionSettings)

	return pool[name], err
}

// OpenCustom устанавливает соединение с базой данных из пула по произвольным параметрам, добавленным к текущим настройкам соединения,
// если соединение ещё не установлено, и возвращает ссылку на него
func OpenCustom(name, id string, customParams map[string]interface{}) (*sql.DB, error) {
	connectionSettings, ok := settings.Pool[name]
	if !ok {
		return nil, fmt.Errorf("%q отсутствует в пуле соединений", name)
	}

	cs := connectionSettings
	cs.ConnectionStringParams = make(map[string]interface{})
	for k, v := range connectionSettings.ConnectionStringParams {
		if param, ok := customParams[k]; ok {
			cs.ConnectionStringParams[k] = param
		} else {
			cs.ConnectionStringParams[k] = v
		}
	}

	for k, v := range customParams {
		if _, ok := cs.ConnectionStringParams[k]; !ok {
			cs.ConnectionStringParams[k] = v
		}
	}

	var err error
	connName := name + id
	pool[connName], err = getConn(pool[connName], cs)

	return pool[connName], err
}

// CloseAll закрывает все соединения
func CloseAll() {
	for _, conn := range pool {
		if err := conn.Ping(); err != nil {
			conn.Close()
		}
	}
}

func getConn(conn *sql.DB, cs ConnectionSetting) (*sql.DB, error) {
	if conn == nil {
		return openConn(cs)
	}

	if err := conn.Ping(); err != nil {
		return openConn(cs)
	}

	return conn, nil
}

func openConn(s ConnectionSetting) (*sql.DB, error) {
	ds, ok := drivers[s.Driver]
	if !ok {
		return nil, fmt.Errorf("Драйвер для %q не зарегистрирован", s.Driver)
	}

	cs, err := ds.GetConnectionString(s.ConnectionStringParams)
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open(s.Driver, cs)
	if err != nil {
		return nil, err
	}

	if ds.AfterConnection != nil {
		ds.AfterConnection(conn, s)
	}

	for _, q := range s.AfterConnection {
		if _, err := conn.Exec(q.Query, q.Args...); err != nil {
			return nil, err
		}
	}

	return conn, err
}
