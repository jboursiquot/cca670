```
$ make deploy
aws cloudformation validate-template \
                --template-body file://deploy.yaml
{
    "Parameters": [
        {
            "ParameterKey": "NetAppSyncBucketName",
            "DefaultValue": "unset",
            "NoEcho": false,
            "Description": "Name of S3 bucket for NetApp to sync with."
        },
        {
            "ParameterKey": "NetAppUserPassword",
            "DefaultValue": "OverrideMePlease2020",
            "NoEcho": false
        }
    ],
    "Description": "Stack sets up VPCs.",
    "Capabilities": [
        "CAPABILITY_IAM"
    ],
    "CapabilitiesReason": "The following resource(s) require capabilities: [AWS::IAM::User]"
}
aws cloudformation deploy \
                --stack-name cca670-p2-netapp \
                --template-file deploy.yaml \
                --capabilities CAPABILITY_IAM \
                --parameter-overrides \
                                NetAppSyncBucketName=cca670-p2-netapp \
                                NetAppUserPassword=`aws secretsmanager \
                                        get-random-password \
                                        --password-length 20 \
                                        --require-each-included-type \
                                        --output text`

Waiting for changeset to be created..
Waiting for stack create/update to complete
Successfully created/updated stack - cca670-p2-netapp
```