ensure:
	dep ensure
	git apply patch/* --directory=vendor/github.com/ethereum/go-ethereum

image:
	docker build . -t statusteam/boothealth:latest

push:
	docker push statusteam/boothealth:latest
