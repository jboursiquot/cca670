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
    "Description": "Customer managed KMI"
}
aws cloudformation deploy \
                --stack-name cca670-p2-cust-kmi \
                --template-file deploy.yaml \
                --parameter-overrides \
                                BucketName=cca670-p2-cust-kmi

Waiting for changeset to be created..
Waiting for stack create/update to complete
Successfully created/updated stack - cca670-p2-cust-kmi
```

```
$ cat datafile
TEST DATA
```

```
$ make upload 
./bin/encrypt \
                --passphrase test \
                --file datafile
aws s3 cp datafile.encrypted s3://cca670-p2-cust-kmi
upload: ./datafile.encrypted to s3://cca670-p2-cust-kmi/datafile.encrypted
```

```
$ cat datafile.encrypted 
f�ؿ?O<M��/n��m2�-�W����3B�~'E'���% 
```

```
$ make download
aws s3 cp s3://cca670-p2-cust-kmi/datafile.encrypted .
download: s3://cca670-p2-cust-kmi/datafile.encrypted to ./datafile.encrypted
/Applications/Xcode.app/Contents/Developer/usr/bin/make decrypt
./bin/decrypt \
                --passphrase test \
                --file datafile.encrypted \
                --outfile datafile.decrypted
```

```
$ cat datafile.decrypted 
TEST DATA
```