run:
	docker run -it -v ~/kafka/sampleTopic:/sampleTopic samplekafka:1.0
build:
	docker build -t samplekafka:1.0 .

