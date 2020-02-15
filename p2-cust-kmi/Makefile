PROJECT = $(shell basename $(CURDIR))
TEMPLATE = deploy.yaml
STACK_NAME = cca670-$(PROJECT)
BUCKET ?= $(STACK_NAME)

TEST_PASSPHRASE ?= test
TEST_FILE ?= datafile

.PHONY: vars describe deploy destroy upload download encrypt decrypt

vars:
	@echo PROJECT = $(PROJECT)
	@echo TEMPLATE = $(TEMPLATE)
	@echo STACK_NAME = $(STACK_NAME)
	@echo BUCKET = $(BUCKET)

validate:
	aws cloudformation validate-template \
		--template-body file://$(TEMPLATE)

deploy: validate
	aws cloudformation deploy \
		--stack-name $(STACK_NAME) \
		--template-file $(TEMPLATE) \
		--parameter-overrides \
				BucketName=$(BUCKET)
				
describe:
	aws cloudformation describe-stacks \
		--stack-name $(STACK_NAME)

destroy:
	-aws s3 rm s3://$(BUCKET) --recursive
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)

### CUSTOM TOOLING FOR ENCRYPT/DECRYPT ###

build:
	go build -o bin/encrypt ./cmd/encrypt
	go build -o bin/decrypt ./cmd/decrypt

encrypt:
	./bin/encrypt \
		--passphrase $(TEST_PASSPHRASE) \
		--file $(TEST_FILE)

decrypt:
	./bin/decrypt \
		--passphrase $(TEST_PASSPHRASE) \
		--file $(TEST_FILE).encrypted \
		--outfile $(TEST_FILE).decrypted

### TEST ###

upload: encrypt
	aws s3 cp datafile.encrypted s3://$(BUCKET)

download:
	aws s3 cp s3://$(BUCKET)/datafile.encrypted .
	$(MAKE) decrypt