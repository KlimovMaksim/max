## Описание
Банковское приложение, управляющее счетами, транзакциями и пользователями. Аутентификация реализована с помощью JWT
## Настройка
1. Поднять бд в docker
```bash
docker compose up -d
```
2. Запуск сервиса
```bash
go run cmd/api/main.go
```
3. Сервис запускается на порту 8080. Пример обращения к сервису: localhost:8080/register
4. Для всех запросов, кроме регистрации и аутентификации, реализована защита с помощью JWT.
## Функциональность
1. Регистрация и логин
2. Управление счетами (/api/accounts)
3. Управление картами (/api/cards)
4. Управление транзакциями (/api/transfer)
5. Управление кредитами (/api/credits)
6. Управление аналитикой (/api/analytics)