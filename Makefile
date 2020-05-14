validate:
	swagger validate ./swagger/swagger.yml
gen: validate
	swagger generate server \
		--target=./swagger \
		--spec=./swagger/swagger.yml \
		--exclude-main \
		--name=hello


install:
	go install -v

.PHONY: install
