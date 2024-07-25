# UserService
Сервис для работы с пользователями, их балансом и ролью.

## Запросы и примеры
### /api/users
GET

http://localhost:9004/api/users?id=1

POST

http://localhost:9004/api/users

{
  "name": "Maxim",
  "surname": "Bolovin",
  "login": "chikibambony",
  "password": "123",
  "roleKey": "user"
}

PUT

http://localhost:9004/api/users

{
  "id": 1,
  "name": "Maxim",
  "surname": "Bolovin",
  "login": "bambonychiki",
  "password": "321",
  "roleKey": "user"
}

DELETE

http://localhost:9004/api/users?id=1

### /api/balance
GET

http://localhost:9004/api/balance?id=1

PUT

http://localhost:9004/api/balance?id=1&action=plus&amount=50

http://localhost:9004/api/balance?id=1&action=minus&amount=30

### /api/role
GET

http://localhost:9004/api/role?id=1

PUT

http://localhost:9004/api/role?id=1&role=hr
