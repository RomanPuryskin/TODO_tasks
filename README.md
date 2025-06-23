
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
![image](https://github.com/user-attachments/assets/99a8bc1c-6215-44a5-8450-c142a4450353)

-post
![image](https://github.com/user-attachments/assets/1fe9f82a-caf3-440d-b49a-071acbd861db)

-put
![image](https://github.com/user-attachments/assets/07c6b8af-6b0c-4c9e-9ecc-a27fed12b8fc)

-delete
![image](https://github.com/user-attachments/assets/56589d39-948f-404c-a36d-a83af40084ef)

Таблица
- до запросов выше
![image](https://github.com/user-attachments/assets/c0bc8acd-45da-4f32-8fcd-6918e4ad4cba)


- после запросов выше
![image](https://github.com/user-attachments/assets/1baf42ed-299b-4401-8a1b-3ad55cb42846)







