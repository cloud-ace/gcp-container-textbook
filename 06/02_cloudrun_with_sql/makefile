PROJECT_NAME=ca-seno-test
PACKAGE_NAME=with-sql
CONTAINER_PATH=gcr.io/${PROJECT_NAME}/${PACKAGE_NAME}

build:
	docker build -t ${CONTAINER_PATH} --target builder -t cloudrun_sample:builder --cache-from cloudrun_sample:builder ./
	docker build -t ${CONTAINER_PATH} --cache-from cloudrun_sample:builder ./

run:
	docker run --rm -l cloudrun_sample -p 8080:8080 ${CONTAINER_PATH}

brun:
	make build
	make run

upload:
	docker push ${CONTAINER_PATH}

deploy:
	make build
	make upload
	gcloud beta run deploy ${PACKAGE_NAME} --image $(CONTAINER_PATH) --platform managed --region asia-northeast1

