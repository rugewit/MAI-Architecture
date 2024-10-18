### Лабораторная работа № 3

#### Чернобаев Андрей Александрович М8О-114М-23

### Цель

Получение практических навыков в построении сервисов, работающих с документо-ориентированными базами данных.

### Задание

Разработать приложение осуществляющее хранение данных согласно варианта
задания в MongoDB. Для описанных в архитектуре вызовов между сервисами создайте
REST интерфейс.

Должны выполняться следующие условия:

- Данные должны храниться в СУБД MongoDB;
- Интерфейс к сущностям должен предоставляться в соответствии со стилем REST;
- API должен быть специфицирован в OpenAPI 3.0 (должен хранится в index.yaml);
- Должен быть создан скрипт по созданию базы данных и таблиц, а также наполнению СУБД тестовыми значениями;
- Для каждой коллекции должен быть создан отдельный сервис, реализующий CRUD операции.~~

### Запуск

```bash
docker compose up
```

или с доп. параметрами для docker'a

```bash
docker-compose down && docker-compose up -d --force-recreate --build
```

url: http://localhost:8081/

#### Демонстрация работы

**Получение всех пользователей**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/1.png)

**Получение пользователя**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/2.png)

**Добавление пользователя**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/3.png)

**Поиск по маске**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/4.png)

**Обновление пользователя**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/5.png)

**Удаление пользователя**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab3/images/6.png)