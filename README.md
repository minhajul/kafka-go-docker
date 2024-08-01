# Kafka Go Docker
Sample project for practicing kafka cluster in docker and writing Go code to produce and consume message.

## Usage

To get started, make sure you have [Docker installed](https://docs.docker.com/docker-for-mac/install/) on your system, and then clone this repository.

From the project's root directory run `docker-compose up`. If you want to run this in background, run `docker-compose up -d`. Run `go run producer.go` and `go run comsumer.go` to boot up the producer and consumer repectively.

You will see the below message in you consurmer tab.
```
Message on test_topic[0]@15: Welcome
Message on test_topic[0]@16: to
Message on test_topic[0]@17: Kafka
Message on test_topic[0]@18: with
Message on test_topic[0]@19: Golang
```
