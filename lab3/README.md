### Лабораторная работа № 2

#### Чернобаев Андрей Александрович М8О-114М-23

### Цель

Получение практических навыков в построении сервисов, работающих с
реляционными данными.

### Задание

Разработать приложение осуществляющее хранение данных о пользователях в реляционной СУБД. Для выявленных в предыдущем задании вызовов между сервисами создайте REST интерфейс.
Должны выполняться следующие условия:

- Данные должны храниться в СУБД PostgreSQL;
- Должна(ы) быть создана(ы) таблица(ы) для сущности, соответствующей клиенту/пользователю;
- Интерфейс к сущностям должен предоставляться в соответствии со стилем REST;
- API должен быть специфицирован в OpenAPI 3.0 (должен хранится в index.yaml);
- Должен быть создан скрипт по созданию базы данных и таблиц(ы), а также наполнению СУБД тестовыми значениями;
- Для сущности, отвечающей за хранение данных о пользователе (клиенте), должен быть реализован интерфейс поиска по маске фамилии и имени, а также стандартные CRUD операции.
- Данные о пользователе должны включать логин и пароль. Пароль должен храниться в закрытом виде (хэширован)

Рекомендуемая последовательность выполнения работы:

- Создайте схему БД
- Создайте таблицу(ы)
- Создайте скрипт для первичного наполнение БД и выполните
- Реализуйте REST-сервис
- Сделайте спецификацию с OpenAPI с помощью Postman и сохраните ее в index.yml
- Протестируйте сервис
- Создайте Dockerfile для вашего сервиса
- Протестируйте его работу в Docker
- Опубликуйте на github проект

### Решение

#### Запуск

```bash
docker compose up
```

url: http://localhost:8081/

swagger url: http://localhost:8081/swagger/index.html

#### Скрипт по созданию БД и наполнению тестовыми значениями

```postgresql
CREATE TABLE Users (
                       Id VARCHAR(255) PRIMARY KEY,
                       Name VARCHAR(255),
                       Lastname VARCHAR(255),
                       Password VARCHAR(255),
                       CreationDate TIMESTAMP
);

INSERT INTO Users (Id, Name, Lastname, Password, CreationDate)
VALUES
('6c6f5595-bacb-4e13-9a35-e880602a7200', 'Giovanny', 'Champlin', '$2a$10$XTnJFSzVsN8wlRkZSsuYiuCpscsRlD492rl.gyihGEEIDnnhjookW', '2024-04-17 07:38:23'),
('26fc1fde-676c-46ac-a155-412778f2673c', 'Ashlynn', 'Jast', '$2a$10$7yW8gRy3XF0AegcRAwSHKecqbd9tOFrUfx5qbg2eWV8hz0WiUcF1C', '2024-04-17 06:38:23'),
('6a118da2-8794-4c18-b39e-2220bd7261cc', 'Velma', 'Bashirian', '$2a$10$SmA5kbfripaxwvEXokH7qeM24rNLzzlngGjfFGijZuru8hsbg8Zb.', '2024-04-16 23:38:23'),
('37d4897b-752e-4229-acaa-69365b9263de', 'Isidro', 'Emard', '$2a$10$J7oDLad6QJFb4buGMg/gVOT7xlUsYexKsQVxBDeSsK3nuc19tYNV.', '2024-04-17 06:38:23'),
('5eaabb0b-ecac-4fe3-9c45-5a96d1b768f4', 'Destinee', 'Grady', '$2a$10$aRA/ZTiShCRxlCFkc9ecHu5oxWF.YS0icsSKRGcpz2HxsBrtbqUA.', '2024-04-17 09:38:24'),
('ce07a7ac-ebd2-453c-b955-ca766345913b', 'Stacey', 'Reilly', '$2a$10$kWBF4L2JbTtvrfrbqjGlTe0Gf9OZouP5AIMW2qKe0l6FSzUgWpmnq', '2024-04-17 02:38:24'),
('d49e1d69-e19e-4f74-8f7c-2a09cb6ca1c6', 'Celine', 'Fadel', '$2a$10$ha3eYfcT5aTeNHi72Twr8eDD3B.MkzCv/pGeX/8wuTqMXa3QCs2qS', '2024-04-17 05:38:24'),
('90058aa7-0ac3-48e3-9ddf-0da06b808d79', 'Destinee', 'Brown', '$2a$10$wFk5HrthYLs0E4ggar5/weKdDHsZROxtoKtYEqgKgFolUaEquBvm.', '2024-04-17 06:38:24'),
('4aaec7ee-64d2-4ceb-b956-e00fd1003110', 'Eleonore', 'Leuschke', '$2a$10$7xgYm0tHfwIosDjDEonUJ.kpVGfY.WHJH9/zGQ4YZWmwNvhZW.ype', '2024-04-17 00:38:24'),
('90148d40-5ee5-410c-af04-2c7cddcbd896', 'Russ', 'Shields', '$2a$10$5ch2KGJSGfdKRrkZTLSGWOBtGtTVyEG5eg5u8vC8t92Mb2gv.tX42', '2024-04-17 05:38:24');
```

#### REST API

url: http://localhost:8081/swagger/index.html

![user api image](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/1.png)

#### OpenAPI

shop/docs/swagger.yaml

```yaml
basePath: /
definitions:
  models.PatternSearchRequest:
    properties:
      lastNamePattern:
        example: '%_%'
        type: string
      namePattern:
        example: '%_%'
        type: string
    type: object
  models.SignUpUser:
    properties:
      lastname:
        example: Ivanov
        type: string
      name:
        example: Alex
        type: string
      password:
        example: qwerty
        type: string
    type: object
  models.User:
    properties:
      creationDate:
        type: string
      id:
        type: string
      lastname:
        type: string
      name:
        type: string
      password:
        type: string
    type: object
host: localhost:8081
info:
  contact: {}
  title: Shop
  version: "1.0"
paths:
  /users:
    get:
      consumes:
      - application/json
      description: Get Users
      operationId: get-users
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            items:
              $ref: '#/definitions/models.User'
            type: array
        "400":
          description: Bad request
          schema: {}
      summary: Get Users
      tags:
      - User API
  /users/:
    post:
      consumes:
      - application/json
      description: Create an user
      operationId: create-user
      parameters:
      - description: user
        in: body
        name: input
        required: true
        schema:
          $ref: '#/definitions/models.SignUpUser'
      produces:
      - application/json
      responses:
        "201":
          description: created
          schema:
            $ref: '#/definitions/models.User'
        "400":
          description: Bad request
          schema: {}
      summary: Create an user
      tags:
      - User API
  /users/{id}:
    delete:
      consumes:
      - application/json
      description: Delete a user
      operationId: delete-user
      parameters:
      - description: user ID
        in: path
        name: id
        required: true
        type: string
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            type: string
        "404":
          description: Not found
          schema: {}
      summary: Delete a user
      tags:
      - User API
    get:
      consumes:
      - application/json
      description: Get an user
      operationId: get-user
      parameters:
      - description: user ID
        in: path
        name: id
        required: true
        type: string
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/models.User'
        "400":
          description: Bad request
          schema: {}
        "404":
          description: Not found
          schema: {}
      summary: Get an user
      tags:
      - User API
    put:
      consumes:
      - application/json
      description: Update an user
      operationId: update-user
      parameters:
      - description: user ID
        in: path
        name: id
        required: true
        type: string
      - description: updated user
        in: body
        name: input
        required: true
        schema:
          $ref: '#/definitions/models.SignUpUser'
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/models.User'
        "400":
          description: Bad request
          schema: {}
        "404":
          description: Not found
          schema: {}
      summary: Update an user
      tags:
      - User API
  /users/pattern-search:
    post:
      consumes:
      - application/json
      description: Pattern Search. % The percent sign represents zero, one, or multiple
        characters. _ The underscore sign represents one, single character
      operationId: pattern-search-users
      parameters:
      - description: pattern search request
        in: body
        name: input
        required: true
        schema:
          $ref: '#/definitions/models.PatternSearchRequest'
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            items:
              $ref: '#/definitions/models.User'
            type: array
        "400":
          description: Bad request
          schema: {}
      summary: Pattern Search
      tags:
      - User API
swagger: "2.0"
```

#### Хэширование пароля

Пароли хэшированы с помощью bcrypt (Provos and Mazières's bcrypt adaptive hashing algorithm)



#### Демонстрация работы

**Получение всех пользователей**

![get all user image](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/2.png)

**Добавление нового пользователя**

![add a new user image](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/3.png)

**Поиск по маске**

![users pattern search](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/4.png)

**Получение пользователя**

![get an user](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/5.png)

**Обновление пользователя**

![update an user](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/6.png)

**Удаление пользователя**

![delete an user](https://github.com/rugewit/MAI-Architecture/blob/main/lab2/images/7.png)