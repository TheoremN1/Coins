# ProductServuce
Сервис для работы с мерчем и достижениями.

## Запросы и примеры
### /api/achievements
GET

http://localhost:9002/api/achievements

POST

http://localhost:9002/api/achievements

{
  "name": "AchName1",
  "description": "AchDesc1",
  "reward": 10
}

PUT

http://localhost:9002/api/achievements

{
  "id" : 1
  "name": "AchNewName1",
  "description": "AchNewDesc1",
  "reward": 10
}

DELETE

http://localhost:9002/api/achievements?id=1

### /api/merch
GET

http://localhost:9002/api/merch

POST

http://localhost:9002/api/merch

{
  "name": "MerchName1",
  "description": "MerchDesc1",
  "price": 10
}

PUT

http://localhost:9002/api/merch

{
  "id": 1
  "name": "MerchNewName1",
  "description": "MerchNewDesc1",
  "price": 10
}

DELETE

http://localhost:9002/api/merch?id=1
