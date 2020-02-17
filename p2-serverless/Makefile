PROJECT = $(shell basename $(CURDIR))
TEMPLATE = deploy.yaml
STACK_NAME = cca670-$(PROJECT)-wildrydes
BUCKET ?= $(STACK_NAME)

.PHONY: vars describe deploy destroy sync

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
		--capabilities CAPABILITY_NAMED_IAM \
		--parameter-overrides \
				WebsiteBucket=$(BUCKET)
				

describe:
	aws cloudformation describe-stacks \
		--stack-name $(STACK_NAME)

destroy:
	-aws s3 rm s3://$(BUCKET) --recursive
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)

sync:
	aws s3 sync \
		--acl public-read \
    ./website s3://$(BUCKET)