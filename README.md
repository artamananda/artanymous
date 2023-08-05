# ARTANYMOUS RESTful API
An application like secreto.site. Anonymous question.
## API Specs
Servers:
```
http:localhost:3000/api
```

### Questions
#### List all Questions

Endpoint:
```
GET /questions
```

Request Header:
- X-API-Key : Key (Mandatory)

Response:
```
{
    "code" : 200,
    "status" : "OK",
    "data" : [
        {
            "id" : 1,
            "question" : "string"
        }
    ]
}
```

#### Create new Questions

Endpoint:
```
POST /questions
```

Request Header:
- X-API-Key : Key (Mandatory)

Request Body:
```
{
    "question" : "string"
}
```

Response:
```
{
    "code" : 200,
    "status" : "OK",
    "data" : [
        {
            "id" : 1,
            "question" : "string"
        }
    ]
}
```

#### Delete question by Id
Endpoint:
```
DELETE /questions/{questionId}
```

Request Header:
- X-API-Key : Key (Mandatory)

Response:
```
{
    "code" : 200,
    "status" : "OK"
}
```