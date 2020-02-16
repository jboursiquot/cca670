```
$ make deploy
aws cloudformation validate-template \
                --template-body file://deploy.yaml
{
    "Parameters": [
        {
            "ParameterKey": "EmailID",
            "DefaultValue": "abc.xyz@email.com",
            "NoEcho": false,
            "Description": "Enter Approvers Email ID"
        }
    ],
    "Description": "approval workflow for service catalog product launch",
    "Capabilities": [
        "CAPABILITY_IAM"
    ],
    "CapabilitiesReason": "The following resource(s) require capabilities: [AWS::IAM::Role]"
}
aws cloudformation deploy \
                --stack-name cca670-p2-servicecatalog-with-approval \
                --template-file deploy.yaml \
                --capabilities CAPABILITY_IAM \
                --parameter-overrides \
                                EmailID=user@example.com

Waiting for changeset to be created..
Waiting for stack create/update to complete
Successfully created/updated stack - cca670-p2-servicecatalog-with-approval
```