PROJECT = $(shell basename $(CURDIR))
STACK_NAME = cca670-ballotonline-p1-$(PROJECT)
TEMPLATE = template.yaml
CFN_BUCKET = $(STACK_NAME)
DOMAIN = goserverless.info
SENDER = hello@$(DOMAIN)

cfn-resources-bucket:
	aws s3 mb s3://$(CFN_BUCKET)

cfn-resources-upload:
	aws s3 sync ./custom-cfn-resources s3://$(CFN_BUCKET)

cfn-resources-list:
	aws s3 ls s3://$(CFN_BUCKET)

deploy:
	aws cloudformation validate-template \
  	--template-body file://$(TEMPLATE)
	aws cloudformation deploy \
					--stack-name $(STACK_NAME) \
					--template-file $(TEMPLATE) \
					--capabilities CAPABILITY_IAM \
					--parameter-overrides \
									CfnBucket=$(CFN_BUCKET) \
									Domain=$(DOMAIN)
	$(MAKE) active-rule-set

activate-rule-set:
	aws ses set-active-receipt-rule-set \
		--rule-set-name `aws ses list-receipt-rule-sets \
			--query 'RuleSets[0].Name' \
			--output text`

destroy:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME)

describe:
	aws cloudformation describe-stacks \
		--stack-name $(STACK_NAME)

events:
	aws cloudformation describe-stack-events \
		--stack-name $(STACK_NAME)