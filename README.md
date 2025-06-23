# TODO_tasks

Сборка:
  - через Makefile:
    make run
    (команда создаст папку .bin и забилдит туда бинарник, потом его запустит)
  - забилдить main.go напрямую
    go run cmd/main.go

Данные моего локального .env файла (приведены просто для примера, так делать нельзя):
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456toy
DB_NAME=TODO_tasks

Документация: 
  документация уже сгенирована, хранится в папке ./docs и доступна по эндпоинту http://localhost:3000/swagger/index.html после запуска проекта
  для повторной генерации документации команда: make swag
