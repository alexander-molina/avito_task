# antibot-developer-trainee

My implementation for https://github.com/avito-tech/antibot-developer-trainee

## Work locally

### Via make

- make test (run tests)
- make run (run app)
- make build (build app)
- make install (install app)

## Work with docker

- docker-compose up --build

## API

| Endpoint                        | Method | Body |
| ------------------------------- | ------ | ---- |
| [/](#/)                         | GET    | NO   |
| [/limits/reset](#/limits/reset) | POST   | YES  |

### /

Path to make requests and handle limitations

### /limits/reset

Path to reset limitations for many addresses

#### Body example

```json
{
  "addresses": ["203.0.113.195", "70.41.3.18", "150.172.238.178"]
}
```
