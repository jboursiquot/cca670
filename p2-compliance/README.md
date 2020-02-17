```
$ make repo
aws codecommit create-repository\
                --repository-name DevSecOpsExampleRepo \
                --repository-description "My DevSecOps example repository"
{
    "repositoryMetadata": {
        "accountId": "191848360966",
        "repositoryId": "82144cca-9186-4c23-82ec-5aaaec9b2f1a",
        "repositoryName": "DevSecOpsExampleRepo",
        "repositoryDescription": "My DevSecOps example repository",
        "lastModifiedDate": 1581895060.325,
        "creationDate": 1581895060.325,
        "cloneUrlHttp": "https://git-codecommit.us-east-1.amazonaws.com/v1/repos/DevSecOpsExampleRepo",
        "cloneUrlSsh": "ssh://git-codecommit.us-east-1.amazonaws.com/v1/repos/DevSecOpsExampleRepo",
        "Arn": "arn:aws:codecommit:us-east-1:191848360966:DevSecOpsExampleRepo"
    }
}
```