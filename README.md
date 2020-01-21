# cca670

## Compliant Glacier Vault

```
$ make create-vault 
aws glacier create-vault \
  --vault-name cca670-ballotonline-p1-compliance-vault \
  --account-id -
{
    "location": "/191848360966/vaults/cca670-ballotonline-p1-compliance-vault"
}


$ make policy-from-template
sed s#VAULT_ARN#arn:aws:glacier:us-east-1:191848360966:vaults/cca670-ballotonline-p1-compliance-vault#g \
                        policy-template.json > policy.json


$ make apply-policy 
aws glacier complete-vault-lock \
    --account-id - \
    --vault-name cca670-ballotonline-p1-compliance-vault \
    --lock-id `aws glacier initiate-vault-lock \
        --account-id - \
        --vault-name cca670-ballotonline-p1-compliance-vault \
        --policy file://policy.json \
        --output text \
        --query 'lockId'`


$ make describe-vault 
aws glacier describe-vault \
  --vault-name cca670-ballotonline-p1-compliance-vault \
  --account-id -
{
    "VaultARN": "arn:aws:glacier:us-east-1:191848360966:vaults/cca670-ballotonline-p1-compliance-vault",
    "VaultName": "cca670-ballotonline-p1-compliance-vault",
    "CreationDate": "2020-01-21T03:29:05.327Z",
    "NumberOfArchives": 0,
    "SizeInBytes": 0
}

```

## SES Delivery Pipeline

```
$ make cfn-resources-bucket
aws s3 mb s3://cca670-ballotonline-p1-ses
make_bucket: cca670-ballotonline-p1-ses


$ make cfn-resources-upload 
aws s3 sync ./custom-cfn-resources s3://cca670-ballotonline-p1-ses
upload: custom-cfn-resources/aws-cfn-ses-domain-0.3.cf.yaml to s3://cca670-ballotonline-p1-ses/aws-cfn-ses-domain-0.3.cf.yaml
upload: custom-cfn-resources/aws-cfn-ses-domain-0.3.lambda.zip to s3://cca670-ballotonline-p1-ses/aws-cfn-ses-domain-0.3.lambda.zip


$ make cfn-resources-list  
aws s3 ls s3://cca670-ballotonline-p1-ses
2020-01-20 14:45:06       4049 aws-cfn-ses-domain-0.3.cf.yaml
2020-01-20 14:45:06      20077 aws-cfn-ses-domain-0.3.lambda.zip

$ make deploy
aws cloudformation validate-template \
        --template-body file://template.yaml
{
    "Parameters": [
        {
            "ParameterKey": "CfnBucket",
            "NoEcho": false,
            "Description": "Bucket where the nested CF stack resides to help with this automation."
        },
        {
            "ParameterKey": "Domain",
            "NoEcho": false,
            "Description": "Domain name to verify and use for outbound emails."
        }
    ],
    "Capabilities": [
        "CAPABILITY_NAMED_IAM",
        "CAPABILITY_AUTO_EXPAND"
    ],
    "CapabilitiesReason": "The following resource(s) require capabilities: [AWS::CloudFormation::Stack]"
}
aws cloudformation deploy \
    --stack-name cca670-ballotonline-p1-ses \
    --template-file template.yaml \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides \
        CfnBucket=cca670-ballotonline-p1-ses \
        Domain=goserverless.info


$ make activate-rule-set 
aws ses set-active-receipt-rule-set \
    --rule-set-name `aws ses list-receipt-rule-sets \
      --query 'RuleSets[0].Name' \
      --output text`
```