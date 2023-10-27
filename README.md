<h1 align="center">  Вычислитель арифметической прогресии </h1>

## ![Typing SVG](https://readme-typing-svg.herokuapp.com?color=%2336BCF7&lines=Что+реализовано+в+проекте?)

### `cmd/APCServer - package main`
Главный пакет отражающий все зависимости приложения
### `internal/apc - package apc`
Описывает принцип вычисления арифметической програссии<br> 
Содержит структуру ```AP``` со следующими полями:<br>
- `N1 - Стартовый элемент`
- `D  - шаг арифметической прогрессии`
```go
    type AP struct {
        N1 float64 `json:"n1"` // first elem
        D  float64 `json:"d"`  // delta
    }

    func (ap *AP) Count() {
        ap.N1 = ap.N1 + ap.D
    }
```
### `intenal/config - package config`
Инициализирует конфигурацию сервиса<br>
Содержит структуру `Config` со следющими полями:<br>
- `cFile      - путь до файла конфигурации`
- `ServerAddr - адрес на котором будет запщен сервер`
- `QMax - максимальное количество исполняемых одновременно задач`
- `StoragePath - путь по которому будет храниться файл с данными`
- `StorageInterval - интервал сохранения данных на диск`

```go
    type Config struct {
        cFile      string `env:"PATH"` // config file
        ServerAddr string `env:"ADDRESS" json:"server_addr,omitempty"`
        QMax   int    `env:"QMAX" json:"q_max,omitempty"`
      
        StoragePath     string `env:"STORAGE" json:"storage_path"`
        StorageInterval int    `env:"STORE_INTERVAL" json:"storage_interval"`
    }
```
### `intenal/scheduler - package scheduler`
Описывает работу планировщика задач <br>
Имеет следующие структуры и методы: <br>
### `Scheduler` - **oтвечает за планирование задач**.
___
```go
// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qMaxCount int32   // total data in queue
	runQ      []*Task // Queue of tasks
	qCount    int32

	lock sync.Mutex

	// for storaging tasks
	storageInterval time.Duration
	file            *os.File
}
```

Вся идея сводиться к тому, что есть **глобальная очередь задач** (`runQ`), по которой идёт планировщик <br>
и добляет задачи на выполнение при условии, что текущее количество одновремеено выполняемых задач (`qCount`) <br>
не превышает максимально допустимого (`qMaxCount`).

Также в структуре представлены поля:
+ `storageInterval time.Duration` - интервал сохранения данных на диск
+ `file            *os.File` - файл в который происходит сохранение данных

## Запуск: 
```bash
go run APCServer.go
```
### или
```bash
go build APCServer
./APCServer
```
    
## Флаги:
    APCServer - h
+ `-a` - адрес сервера (стандартое значение: "localhost:8080")
+ `-file` - пуь до файла, в который будут сохраняться данные
+ `-max ` - максимальное колиество задач, которое сервис обрабытвать одновременно (стандартое значение: 6)
+ `-p`  - путь до файла конфигурации 
+ `-store` - интервал, с которым данные будут выгружаться на диск (стандартое значение: 30s)