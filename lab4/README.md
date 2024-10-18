### Лабораторная работа № 4

#### Чернобаев Андрей Александрович М8О-114М-23

### Цель

Получение практических навыков в обеспечении безопасности систем.

### Задание

Необходимо, что бы сервисы работающие с данными обрабатывали только запросы,
содержащие JWT токен, аутентифицирующий пользователя.
Сервис, работы с пользователями должен реализовать endpoint «auth» который бы
осуществлял basic authentication и выдавал JWT токен.

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

**Попытка получить доступ (неудачная из-за отсутствия токена)**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab4/images/1.png)

**Регистрация нового пользователя**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab4/images/2.png)

**Заходим в аккаунт и получаем токен**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab4/images/3.png)

**Успешная попытка получения доступа**

![](https://github.com/rugewit/MAI-Architecture/blob/main/lab4/images/4.png)
