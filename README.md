# Duckhunt
A photohunt messenger app

## How to run

* In Go
    ```shell
    go get -u github.com/basraven/duckhunt
    go run github.com/basraven/duckhunt
    ```
* In Docker
    ```shell
    cd docker
    docker-compose up
    ```
* In Kubernetes
    ```shell
    kubectl apply -f docker/k8s/duckhunt.yaml
    ```
### Env variables
duckhunt requires the following environmental variables to be set to work:
```
TELEGRAM_API_KEY=
CHAT_ID=
```