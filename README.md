
Сборка:
  - через Makefile:
    make run
    (команда создаст папку ./.bin и забилдит туда бинарник, потом его запустит)
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


  скриншоты работы:

успешные запросы:
- get
![get](https://github.com/user-attachments/assets/cfaa3a20-3492-429a-8c33-56cffc4f63d3)

-post
![post](https://github.com/user-attachments/assets/4394a0c4-976d-43cb-95d1-fc665864ac4f)

-put
![put](https://github.com/user-attachments/assets/7aea60a8-0e42-456e-9f4e-9fa0e07109b9)

-delete
![delete](https://github.com/user-attachments/assets/3ffb13e7-26b7-4e7d-b015-94d69a510e3c)

Таблица
- до запросов выше
![Таблица до](https://github.com/user-attachments/assets/35add1df-4713-4dbc-8df6-26e2110fd2cd)

- после запросов выше
![Таблица после](https://github.com/user-attachments/assets/399996f0-fdac-474a-8e32-ecba3a3c27c1)






