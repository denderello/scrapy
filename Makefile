PROJECT=scrapy
OWNER=denderello

DOCKER_IMAGE=$(OWNER)/$(PROJECT)

$(PROJECT):
	go build .

container:
	docker build -t $(DOCKER_IMAGE) .
