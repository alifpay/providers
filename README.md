# providers
Standard Protocol API, Протокол взаимодействия с провайдерами

This is example for Standard Protocol API for providers to integrate with alif.

```
//if you want run api in docker container

docker build -t alif-api:latest .

docker run -p 8088:8088 -d --name alif-api --restart unless-stopped --env VERSION=v1 alif-api:latest

```