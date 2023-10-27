<h1 align="center">  Вычислитель арифметической прогресии </h1>


## Что это ?

Это сервис, который высчитывает элементы арифметической прогрессии. <br>

**Особенности:** 
+ сервис запускает ограниченныое количество одновременных задач, максимально возможное число<br>
задач передаётся при запуске
+ у сервиса два обработчика:
  + Добавить задачу
  + Вернуть отсортированный список задач
___

## ![Typing SVG](https://readme-typing-svg.herokuapp.com?color=%2336BCF7&lines=Что+реализовано+в+проекте?)

### [`cmd/APCServer - package main`](https://github.com/EgorKo25/APC/blob/main/cmd/APCServer/APCServer.go "GO to code")  
Главный пакет отражающий все зависимости приложения
### [`internal/apc - package apc`](https://github.com/EgorKo25/APC/blob/main/internal/apc/apc.go "GO to code")
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
### [`intenal/config - package config`](https://github.com/EgorKo25/APC/tree/main/internal/config/config.go "GO to code")
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
        QMax       int    `env:"QMAX" json:"q_max,omitempty"`
      
        StoragePath     string `env:"STORAGE" json:"storage_path"`
        StorageInterval int    `env:"STORE_INTERVAL" json:"storage_interval"`
    }
```
### [`intenal/scheduler - package scheduler`](https://github.com/EgorKo25/APC/blob/main/internal/scheduler/scheduler.go "GO to code")
Описывает работу планировщика задач <br>
Имеет следующие структуры и методы: <br>
#### `Scheduler` - **oтвечает за планирование задач**.
___
```go
// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qMaxCount int32   // total data in queue
	runQ      []*Task // Queue of tasks
	qCount    int32

	lock sync.Mutex

	// for storaging of tasks
	storageInterval time.Duration
	file            *os.File
}
```

Вся идея сводиться к тому, что есть **глобальная очередь задач** (`runQ`), по которой идёт планировщик <br>
и добляет задачи на выполнение при условии, что текущее количество одновремеено выполняемых <br>задач (`qCount`) 
не превышает максимально допустимого (`qMaxCount`).

Также в структуре представлены поля:
+ `storageInterval time.Duration` - интервал сохранения данных на диск
+ `file            *os.File` - файл в который происходит сохранение данных

#### Планировщик имеет следующие методы:
+ **InsertTask** Добавляет задачу в очередь
```go 
func (s *Scheduler) InsertTask(t *Task) 
```
+ **GetSortQueue** возвращает остортированную по статусу выполнения очередь задач
```go 
func (s *Scheduler) GetSortQueue() []*Task
```
+ **WriteAll** записывает текущую отсортированную очередь `runQ` в `file`
```go 
func (s *Scheduler) WriteAll() (err error) 
```
+ **Run** запускает планирощик 
```go 
func (s *Scheduler) Run()
```
#### [`Task` - **описывает типовую задачу**.](https://github.com/EgorKo25/APC/blob/main/internal/scheduler/task.go "GO to code")
```go
type Task struct {
	apc.AP
	I       time.Duration `json:"interval,omitempty"` // interval between iter
	Iter    int           `json:"iteration,omitempty"` // number of iter
	IterNow int           `json:"iter_now"`

	Status string        `json:"status,omitempty"`
	TTL    time.Duration `json:"ttl,omitempty"` // times to life before finished

	Create time.Time `json:"create,omitempty"`
	Start  time.Time `json:"start,omitempty"`
	Finish time.Time `json:"finish,omitempty"`
}
```
Содержит поля:
+ `I` - интервал между итерациями
+ `Iter` - количество итераций
+ `IterNow` - текущая итерация
+ `TTL` - время хранения результата после вычисления
+ `Status` - статус задачи `Run`/`Wait`/`Finished`
### [`intenal/server - package server`](https://github.com/EgorKo25/APC/tree/main/internal/server/ "GO to code")
___
**Немного о сервере**
+ в качестве фраемворка был выбран `chi`<br>
  + имеется два обработчика 
    + ***POST /set/ -*** [`SetTaskToQueue`](https://github.com/EgorKo25/APC/blob/main/internal/server/handler/handler.go "GO to code") -
    Принимает массив задач в формате `JSON` и добавляет их в очередь

    ```json
    [
       {
         "n1": 1,
         "d": 3,
         "ttl": 30,
         "interval": 5,
         "iteration": 35
       },
       {
         "n1": 5,
         "d": 89,
         "ttl": 3,
         "interval": 1,
         "iteration": 100
        }
     ]
    ```
  + ***GET /get/ -*** [`GetTasksList`](https://github.com/EgorKo25/APC/blob/main/internal/server/handler/handler.go "GO to code") - 
  возвращает отсортированную очередь задач в формате `JSON`

    ```json
      [
        {
          "n1": 5,
          "d": 89,
          "ttl": 3,
          "interval": 1,
          "iteration": 100,
          "iter_now": 2,
          "status": "Finished",
          "create": "time",
          "start": "time",
          "finish": "time"
        },
        {
         "n1": 1,
         "d": 3,
         "ttl": 30,
         "interval": 5,
         "iteration": 35,
         "iter_now": 2,
         "status": "Run",
         "create": "time",
         "start": "time"
        },
        {
          "n1": 5,
          "d": 89,
          "ttl": 3,
          "interval": 1,
          "iteration": 100,
          "iter_now": 2,
          "status": "Run",
          "create": "time",
          "start": "time"
        },
        {
          "n1": 5,
          "d": 89,
          "ttl": 3,
          "interval": 1,
          "iteration": 100,
          "iter_now": 2,
          "status": "Wait",
          "create": "time"
        }
      ] 
      ```
## Запуск
___
```bash
go run APCServer.go
```
### или
```bash
go build APCServer
./APCServer
```
    
## Флаги
___
```bash
./APCServer - h
```
+ `-a` - адрес сервера (стандартое значение: "localhost:8080")
+ `-file` - пуь до файла, в который будут сохраняться данные
+ `-max ` - максимальное колиество задач, которое сервис обрабытвать одновременно (стандартое значение: 6)
+ `-p`  - путь до файла конфигурации 
+ `-store` - интервал, с которым данные будут выгружаться на диск (стандартое значение: 30s)