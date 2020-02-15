```
$ make deploy
aws cloudformation validate-template \
                --template-body file://deploy.yaml
{
    "Parameters": [
        {
            "ParameterKey": "BucketName",
            "DefaultValue": "unset",
            "NoEcho": false,
            "Description": "Name of S3 bucket to put objects."
        }
    ],
    "Description": "KMS customer managed CMK for AWS services"
}

aws cloudformation deploy \
                --stack-name cca670-p2-kms \
                --template-file deploy.yaml \
                --parameter-overrides \
                                BucketName=cca670-p2-kms

Waiting for changeset to be created..
```

```
$ make test
aws s3 cp datafile.in.txt s3://cca670-p2-kms
upload: ./datafile.in.txt to s3://cca670-p2-kms/datafile.in.txt
aws s3 cp s3://cca670-p2-kms/datafile.in.txt datafile.out.txt
download: s3://cca670-p2-kms/datafile.in.txt to ./datafile.out.txt
```