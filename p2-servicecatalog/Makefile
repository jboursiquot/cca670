PROJECT = $(shell basename $(CURDIR))
TEMPLATE = deploy.json
STACK_NAME = cca670-$(PROJECT)
KEY_NAME ?= cca670

.PHONY: vars describe deploy destroy

vars:
	@echo PROJECT = $(PROJECT)
	@echo TEMPLATE = $(TEMPLATE)
	@echo STACK_NAME = $(STACK_NAME)

validate:
	aws cloudformation validate-template \
		--template-body file://$(TEMPLATE)

deploy: validate
	aws cloudformation deploy \
		--stack-name $(STACK_NAME) \
		--template-file $(TEMPLATE) \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides \
				KeyName=$(KEY_NAME) \
				SSHLocation=`curl https://api.ipify.org`/32
				
describe:
	aws cloudformation describe-stacks \
		--stack-name $(STACK_NAME)

destroy:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)
