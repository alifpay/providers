# providers
move to https://alifcapital.github.io/providers

Standard Protocol API, Протокол взаимодействия с провайдерами
https://alifpay.github.io

This is the Standard Protocol API for service providers who want their products or services to be available on Alif.mobi. By following this standard, providers enable their clients to make payments seamlessly via the platform. Providers need to implement their API according to these guidelines, ensuring that integration is straightforward with minimal setup—primarily through configuration settings.

This is example for Standard Protocol API for providers.

```
//if you want run api in docker container

docker build -t alif-api:latest .

docker run -p 8088:8088 -d --name alif-api --restart unless-stopped --env VERSION=v1 alif-api:latest

```
