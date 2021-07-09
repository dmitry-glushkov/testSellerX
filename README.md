## Запуск

Сначала запускается контейнер с postgresql:

```
docker-compose up
```
(Скорее всего нужно будет нажать ctrl+z чтобы выйти из логов докера)

Далее сам сервер:

```
make
```

------

## HTTP запросы 

Все необходимые запросы можно выполнить с помощью команд, описанных в тестовом задании, но индексы нужно передавать 
не строкой, а числом, т.е. вместо 
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": ["<USER_ID_1>", "<USER_ID_2>"]}' \
  http://localhost:9000/chats/add
```
использовать 
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [<USER_ID_1>, <USER_ID_2>]}' \
  http://localhost:9000/chats/add
```
