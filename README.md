# Kafka Go Docker
A sample project for practicing Kafka clusters, Docker, and Go, focusing on producing and consuming messages.

## Usage

To get started, make sure you have [Docker installed](https://docs.docker.com/docker-for-mac/install/) on your system, and then clone this repository.

From the project's root directory run `docker-compose up`. If you want to run this in background, run `docker-compose up -d`. Run `go run producer.go` and `go run comsumer.go` in the different tab to boot up the producer and consumer respectively.

You will see the below message in you consumer tab.
```
Message on test_topic[0]@1: Welcome
Message on test_topic[0]@2: to
Message on test_topic[0]@3: Kafka
Message on test_topic[0]@4: with
Message on test_topic[0]@5: Golang
```
