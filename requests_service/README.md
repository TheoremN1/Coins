# RequestsService
Сервис для работы с заявками.

## Запросы и примеры
### /api/coinsrequests
GET

http://localhost:9003/api/coinsrequests

POST

http://localhost:9003/api/coinsrequests

{
  "userId": "1",
  "userMessage": "Я заслужил!",
  "achievementId": "1"
}

PUT

http://localhost:9003/api/coinsrequests

{
  "id" : "1",
  "userId": "1",
  "userMessage": "Я заслужил!",
  "achievementId": "1",
  "hrId": "2",
  "hrMessage": ")))",
  "statusKey": "denied"
}

DELETE

http://localhost:9003/api/coinsrequests?id=1

### /api/merchrequests
GET

http://localhost:9003/api/merchrequests

POST

http://localhost:9003/api/merchrequests

{
  "userId": "1",
  "userMessage": "Хочу вот эту кепку",
  "merchId": "1"
}

PUT

http://localhost:9003/api/merchrequests

{
  "id" : "1",
  "userId": "1",
  "userMessage": "Хочу вот эту кепку",
  "merchId": "1",
  "hrId": "2",
  "hrMessage": "Приходите за этой кепкой в понедельник в наш отдел",
  "statusKey": "accept"
}

DELETE

http://localhost:9003/api/merchrequests?id=1
