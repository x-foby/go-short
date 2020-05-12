package webserver

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Settings описывает настройки веб-сервера
type Settings struct {
	Host              string `json:"host,omitempty" yaml:"host,omitempty"`
	Port              int    `json:"port,omitempty" yaml:"port,omitempty"`
	ReadTimeout       int64  `json:"readTimeout,omitempty" yaml:"readTimeout,omitempty"`
	ReadHeaderTimeout int64  `json:"readHeaderTimeout,omitempty" yaml:"readHeaderTimeout,omitempty"`
	MaxHeaderBytes    int    `json:"maxHeaderBytes,omitempty" yaml:"maxHeaderBytes,omitempty"`
}

// Ошибки
var (
	ErrNoSettings = errors.New("go-short: Не передан указатель на объект с настройками")
)

var settings Settings

// GetSettings возвращает указатель на настройки БД
func GetSettings() *Settings {
	return &settings
}

// Start запускает http-сервер
func Start(handler http.HandlerFunc, logger func(s *Settings, v ...interface{})) error {
	checkSettings()

	http := &http.Server{
		Addr:           settings.Host + ":" + strconv.Itoa(settings.Port),
		Handler:        http.HandlerFunc(handler),
		ReadTimeout:    time.Duration(settings.ReadTimeout) * time.Millisecond,
		MaxHeaderBytes: settings.MaxHeaderBytes,
	}

	logger(&settings)

	if err := http.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Print выводит ответ сервера
func Print(w http.ResponseWriter, status int, data []byte) {
	w.WriteHeader(status)
	w.Write(data)
}

func checkSettings() {
	if strings.TrimSpace(settings.Host) == "" {
		settings.Host = "localhost"
	}

	if settings.Port < 1 || settings.Port > 65535 {
		settings.Port = 65432
	}
}
