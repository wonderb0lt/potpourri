# potpourri
A fun doodle. Reddit API -> Go Consumer -> RabbitMQ -> Go Procesor.

Just for my own entertainment purposes.

## What it does

A Go client polls reddit and then sleeps for a while, sending new posts to a Go worker via RabbitMQ.

## What I want to use

* [GiantSwarm](https://giantswarm.io)
* Docker (duh)
* https://github.com/jzelinskie/geddit
* https://github.com/streadway/amq
* [Celery](http://www.celeryproject.org/)

## Learning goals

* How to use GiantSwarm
* Get some more Docker experience
* Write some Go for the first time since Uni
