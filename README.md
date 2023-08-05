# ARTANYMOUS RESTful API
An application like secreto.site. Anonymous message API.
## API Specs
Servers:
```
http://localhost:3000
```

### Messages
#### List all Messages

Endpoint:
```
GET /api/messages
```

Request Header:
- X-API-Key : Key (Mandatory)

Response:
```json
{
    "code" : 200,
    "status" : "OK",
    "data" : [
        {
            "id" : "69bcd738-b4eb-4f24-a838-25a0d7351ea4",
            "question" : "Hello World",
            "created_at": "2023-08-05T23:11:50.749697Z"
        }
    ]
}
```

#### Create new Messages

Endpoint:
```
POST /api/messages
```

Request Header:
- X-API-Key : Key (Mandatory)

Request Body:
```json
{
    "question" : "Hello World"
}
```

Response:
```json
{
    "code" : 200,
    "status" : "OK",
    "data" : [
        {
            "id": "69bcd738-b4eb-4f24-a838-25a0d7351ea4",
            "question": "Hello World"
        }
    ]
}
```

#### Find Messages by Id

Endpoint:
```
GET /api/messages/:messageId
```

Request Header:
- X-API-Key : Key (Mandatory)

Response:
```json
{
    "code" : 200,
    "status" : "OK",
    "data" : [
        {
            "id" : "69bcd738-b4eb-4f24-a838-25a0d7351ea4",
            "question" : "Hello World",
            "created_at": "2023-08-05T23:11:50.749697Z"
        }
    ]
}
```

#### Delete message by Id
Endpoint:
```
DELETE /api/messages/:messageId
```

Request Header:
- X-API-Key : Key (Mandatory)

Response:
```json
{
    "code" : 200,
    "status" : "OK"
}
```