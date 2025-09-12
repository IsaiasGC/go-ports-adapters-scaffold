#!/bin/bash

# Inicia Kafka en segundo plano
/etc/confluent/docker/run &

# Espera a que Kafka esté listo
while ! nc -z kafka 9092; do   
  echo "Retrying in 5 seconds..."
  sleep 5
done

message='{ "event": "init test" }'
# Envia el mensaje a Kafka
echo ${message} | kafka-console-producer --broker-list kafka:9092 --topic dev-topic

# Espera a que Kafka termine
wait