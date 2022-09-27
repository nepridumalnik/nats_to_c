# Репозиторий с кодом и примерами:
https://github.com/nats-io/nats.go

# Описание:
Данный репозиторий является примером оборачивания nats go клиента в C код

# Загрузка go клиента:
```
go get github.com/nats-io/nats.go/@latest
```

# Сборка проекта:
```
mkdir build
cd build
cmake ..
cmake --build .
```

# Подключение go библиотеки к С++ коду:
Go библиотека собирается в статическую C библиотеку, необходимо добавить заголовочные файлы в область видимости сборочной системы:
```
include_directories(${GO_NATS_INCLUDE})
```

Затем слинковать с библиотекой:
```
target_link_libraries(${TARGET} ${GO_NATS_LIBRARIES})
```

# Запуск докера с nats сервером:
```
docker run -p 4222:4222 -ti nats
```
