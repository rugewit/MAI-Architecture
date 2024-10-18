### Лабораторная работа № 6

#### Чернобаев Андрей Александрович М8О-114М-23

### Цель

Получение практических навыков в построении отказоустойчивых приложений.

### Задание

Необходимо реализовать сервис API Gateway, который будет получать JWT токен в
User Service и вызывать сервисы согласно варианту задания.
Для вызова сервисов, согласно варианту задания реализовать паттерн Circuit Breaker.

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

2024/10/17 22:44:15 Request failed: dial tcp4 172.18.0.6:8082: connect: connection refused
2024/10/17 22:44:15 Service unavailable due to circuit breaker
2024/10/17 22:44:19 Request failed: dialing to the given TCP address timed out
2024/10/17 22:44:19 Service unavailable due to circuit breaker
2024/10/17 22:44:23 Request failed: dialing to the given TCP address timed out
2024/10/17 22:44:23 Service unavailable due to circuit breaker
2024/10/17 22:44:27 Request failed: dialing to the given TCP address timed out
2024/10/17 22:44:27 Service unavailable due to circuit breaker
2024/10/17 22:44:31 Request failed: dialing to the given TCP address timed out
2024/10/17 22:44:31 Circuit breaker transitioned to OPEN state
2024/10/17 22:44:31 Service unavailable due to circuit breaker
2024/10/17 22:44:31 Service unavailable due to circuit breaker
2024/10/17 22:44:32 Service unavailable due to circuit breaker
2024/10/17 22:44:33 Service unavailable due to circuit breaker
2024/10/17 22:44:34 CircuitBreaker is now HalfOpen
2024/10/17 22:44:37 Request failed: dialing to the given TCP address timed out
2024/10/17 22:44:37 Circuit breaker transitioned to OPEN state
2024/10/17 22:44:37 Service unavailable due to circuit breaker
2024/10/17 22:44:38 Service unavailable due to circuit breaker
2024/10/17 22:45:37 CircuitBreaker is now HalfOpen
2024/10/17 22:45:39 CircuitBreaker is now Closed

