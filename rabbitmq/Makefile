runRabbitmq:
	# docker network create learn-rabbitmq
	docker run  --rm --net learn-rabbitmq --env-file ./.env --hostname rabbitmq -p 15672:15672 -p 5672:5672 --name learn-rabbitmq rabbitmq:3.9-management

removeRabbitmq:
	docker rm -f learn-rabbitmq

buildPublisher:
	docker build ./publisher/ -t rabbitmq-publisher

runPublisher:
	docker run -d  --rm --net learn-rabbitmq  --env-file ./.env rabbitmq-publisher 

removePublisher:
	docker rm -f duckhue01/rabbitmq-publisher

buildConsumer:
	docker build ./consumer/ -t rabbitmq-consumer

runConsumer:
	docker run -d  --rm --net learn-rabbitmq --env-file ./.env rabbitmq-consumer 

removeConsumer:
	docker rm -f duckhue01/rabbitmq-consumer