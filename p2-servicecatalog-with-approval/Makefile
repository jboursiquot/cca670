PROJECT = $(shell basename $(CURDIR))
TEMPLATE = deploy.yaml
STACK_NAME = cca670-$(PROJECT)
APPROVER_EMAIL ?= jboursiquot@gmail.com

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
				EmailID=$(APPROVER_EMAIL)
				
describe:
	aws cloudformation describe-stacks \
		--stack-name $(STACK_NAME)

destroy:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)
