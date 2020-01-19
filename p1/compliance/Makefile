VAULT_NAME ?= cca670-ballotonline-p1-compliance-vault
POLICY_FILE ?= policy.json

create-vault:
	aws glacier create-vault \
		--vault-name $(VAULT_NAME) \
		--account-id -

describe-vault:
	aws glacier describe-vault \
		--vault-name $(VAULT_NAME) \
		--account-id -

delete-vault:
	aws glacier delete-vault \
		--vault-name $(VAULT_NAME)
		--account-id -

policy-from-template:
		sed s#VAULT_ARN#$(shell aws glacier describe-vault --vault-name $(VAULT_NAME) --account-id - --output text --query "VaultARN")#g \
			policy-template.json > policy.json

apply-policy:
	aws glacier complete-vault-lock \
    --account-id - \
    --vault-name $(VAULT_NAME) \
    --lock-id `aws glacier initiate-vault-lock \
    	--account-id - \
    	--vault-name $(VAULT_NAME) \
    	--policy file://$(POLICY_FILE) \
			--output text \
			--query 'lockId'`
