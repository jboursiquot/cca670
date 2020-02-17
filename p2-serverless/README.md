
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
            "Description": "Name of S3 bucket to host site."
        }
    ],
    "Description": "This stack sets up an S3-based website."
}
aws cloudformation deploy \
                --stack-name cca670-p2-serverless-wildrydes \
                --template-file deploy.yaml \
                --capabilities CAPABILITY_NAMED_IAM \
                --parameter-overrides \
                                BucketName=cca670-p2-serverless-wildrydes

Waiting for changeset to be created..
Waiting for stack create/update to complete
Successfully created/updated stack - cca670-p2-serverless-wildrydes

```
$ make sync
aws s3 sync \
                --acl public-read \
    ./website s3://cca670-p2-serverless-wildrydes
upload: website/apply.html to s3://cca670-p2-serverless-wildrydes/apply.html
upload: website/css/message.css to s3://cca670-p2-serverless-wildrydes/css/message.css
upload: website/css/mapbox-gl.css to s3://cca670-p2-serverless-wildrydes/css/mapbox-gl.css
upload: website/css/font.css to s3://cca670-p2-serverless-wildrydes/css/font.css
upload: website/css/ride.css to s3://cca670-p2-serverless-wildrydes/css/ride.css
upload: website/css/normalize.css to s3://cca670-p2-serverless-wildrydes/css/normalize.css
upload: website/error.html to s3://cca670-p2-serverless-wildrydes/error.html
upload: website/css/index.css to s3://cca670-p2-serverless-wildrydes/css/index.css
upload: website/css/bootstrap.min.css to s3://cca670-p2-serverless-wildrydes/css/bootstrap.min.css
upload: website/faq.html to s3://cca670-p2-serverless-wildrydes/faq.html
upload: website/favicon.ico to s3://cca670-p2-serverless-wildrydes/favicon.ico
upload: website/fonts/glyphicons-halflings-regular.eot to s3://cca670-p2-serverless-wildrydes/fonts/glyphicons-halflings-regular.eot
upload: website/fonts/fairplex-wide-n4.woff to s3://cca670-p2-serverless-wildrydes/fonts/fairplex-wide-n4.woff
upload: website/fonts/fairplex-wide-n7.woff to s3://cca670-p2-serverless-wildrydes/fonts/fairplex-wide-n7.woff
upload: website/fonts/glyphicons-halflings-regular.ttf to s3://cca670-p2-serverless-wildrydes/fonts/glyphicons-halflings-regular.ttf
upload: website/fonts/glyphicons-halflings-regular.svg to s3://cca670-p2-serverless-wildrydes/fonts/glyphicons-halflings-regular.svg
upload: website/images/logo.png to s3://cca670-p2-serverless-wildrydes/images/logo.png
upload: website/images/loading.gif to s3://cca670-p2-serverless-wildrydes/images/loading.gif
upload: website/css/main.css to s3://cca670-p2-serverless-wildrydes/css/main.css
upload: website/images/background.png to s3://cca670-p2-serverless-wildrydes/images/background.png
upload: website/fonts/glyphicons-halflings-regular.woff to s3://cca670-p2-serverless-wildrydes/fonts/glyphicons-halflings-regular.woff
upload: website/fonts/glyphicons-halflings-regular.woff2 to s3://cca670-p2-serverless-wildrydes/fonts/glyphicons-halflings-regular.woff2
upload: website/images/bucephalus.png to s3://cca670-p2-serverless-wildrydes/images/bucephalus.png
upload: website/images/bbd3207c.png to s3://cca670-p2-serverless-wildrydes/images/bbd3207c.png
upload: website/images/star-pattern.png to s3://cca670-p2-serverless-wildrydes/images/star-pattern.png
upload: website/images/spinning-gears.gif to s3://cca670-p2-serverless-wildrydes/images/spinning-gears.gif
upload: website/images/unicorn-icon.png to s3://cca670-p2-serverless-wildrydes/images/unicorn-icon.png
upload: website/images/unicorn-logo.png to s3://cca670-p2-serverless-wildrydes/images/unicorn-logo.png
upload: website/images/wr-home-W.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-W.png
upload: website/images/rocinante.png to s3://cca670-p2-serverless-wildrydes/images/rocinante.png
upload: website/images/shadowfox.png to s3://cca670-p2-serverless-wildrydes/images/shadowfox.png
upload: website/images/wr-home-Xiaomi.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-Xiaomi.png
upload: website/images/wr-home-apple.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-apple.png
upload: website/images/wr-home-blackberry.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-blackberry.png
upload: website/images/wr-faq-header.jpg to s3://cca670-p2-serverless-wildrydes/images/wr-faq-header.jpg
upload: website/images/wr-home-block-1.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-block-1.png
upload: website/images/wr-home-block-2.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-block-2.png
upload: website/images/unicorn-map-bg.png to s3://cca670-p2-serverless-wildrydes/images/unicorn-map-bg.png
upload: website/images/wr-home-block-4.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-block-4.png
upload: website/images/wr-home-block-3.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-block-3.png
upload: website/images/wr-home-about.jpg to s3://cca670-p2-serverless-wildrydes/images/wr-home-about.jpg
upload: website/images/wr-home-facebook.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-facebook.png
upload: website/images/wr-home-google.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-google.png
upload: website/images/wr-home-instagram.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-instagram.png
upload: website/images/wr-apply-header.png to s3://cca670-p2-serverless-wildrydes/images/wr-apply-header.png
upload: website/images/wr-home-twitter.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-twitter.png
upload: website/images/wr-home-wechat.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-wechat.png
upload: website/images/wr-home-weibo.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-weibo.png
upload: website/images/wr-investors-1.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-1.png
upload: website/images/wr-home-top.jpg to s3://cca670-p2-serverless-wildrydes/images/wr-home-top.jpg
upload: website/images/wr-investors-3.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-3.png
upload: website/images/wr-investors-4.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-4.png
upload: website/images/wr-home-quote.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-quote.png
upload: website/images/wr-investors-2.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-2.png
upload: website/images/unicorn-silhouette.png to s3://cca670-p2-serverless-wildrydes/images/unicorn-silhouette.png
upload: website/images/wr-investors-5.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-5.png
upload: website/images/wr-investors-pcp.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-pcp.png
upload: website/images/wr-investors-awesome.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-awesome.png
upload: website/images/wr-logo-black.png to s3://cca670-p2-serverless-wildrydes/images/wr-logo-black.png
upload: website/images/wr-investors-thebarn.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-thebarn.png
upload: website/images/wr-logo-white.png to s3://cca670-p2-serverless-wildrydes/images/wr-logo-white.png
upload: website/images/wr-investors-header.png to s3://cca670-p2-serverless-wildrydes/images/wr-investors-header.png
upload: website/images/wr-unicorn-header.png to s3://cca670-p2-serverless-wildrydes/images/wr-unicorn-header.png
upload: website/index.html to s3://cca670-p2-serverless-wildrydes/index.html
upload: website/images/wr-unicorn-three.png to s3://cca670-p2-serverless-wildrydes/images/wr-unicorn-three.png
upload: website/js/cognito-auth.js to s3://cca670-p2-serverless-wildrydes/js/cognito-auth.js
upload: website/images/wr-unicorn-two.png to s3://cca670-p2-serverless-wildrydes/images/wr-unicorn-two.png
upload: website/js/esri-map.js to s3://cca670-p2-serverless-wildrydes/js/esri-map.js
upload: website/js/main.js to s3://cca670-p2-serverless-wildrydes/js/main.js
upload: website/investors.html to s3://cca670-p2-serverless-wildrydes/investors.html
upload: website/js/config.js to s3://cca670-p2-serverless-wildrydes/js/config.js
upload: website/js/ride.js to s3://cca670-p2-serverless-wildrydes/js/ride.js
upload: website/js/vendor.js to s3://cca670-p2-serverless-wildrydes/js/vendor.js
upload: website/js/vendor/amazon-cognito-identity.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/amazon-cognito-identity.min.js
upload: website/js/vendor/html5shiv.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/html5shiv.min.js
upload: website/js/vendor/modernizr.js to s3://cca670-p2-serverless-wildrydes/js/vendor/modernizr.js
upload: website/js/vendor/aws-cognito-sdk.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/aws-cognito-sdk.min.js
upload: website/js/vendor/respond.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/respond.min.js
upload: website/images/wr-unicorn-one.png to s3://cca670-p2-serverless-wildrydes/images/wr-unicorn-one.png
upload: website/js/vendor/jquery-3.1.0.js to s3://cca670-p2-serverless-wildrydes/js/vendor/jquery-3.1.0.js
upload: website/js/vendor/moment.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/moment.min.js
upload: website/js/vendor/bootstrap.min.js to s3://cca670-p2-serverless-wildrydes/js/vendor/bootstrap.min.js
upload: website/ride.html to s3://cca670-p2-serverless-wildrydes/ride.html
upload: website/register.html to s3://cca670-p2-serverless-wildrydes/register.html
upload: website/robots.txt to s3://cca670-p2-serverless-wildrydes/robots.txt
upload: website/unicorns.html to s3://cca670-p2-serverless-wildrydes/unicorns.html
upload: website/signin.html to s3://cca670-p2-serverless-wildrydes/signin.html
upload: website/js/vendor/unicorn-icon to s3://cca670-p2-serverless-wildrydes/js/vendor/unicorn-icon
upload: website/verify.html to s3://cca670-p2-serverless-wildrydes/verify.html
upload: website/images/wr-home-quote.jpg to s3://cca670-p2-serverless-wildrydes/images/wr-home-quote.jpg
upload: website/images/wr-home-kraken.png to s3://cca670-p2-serverless-wildrydes/images/wr-home-kraken.png
upload: website/css/bootstrap.min.css.map to s3://cca670-p2-serverless-wildrydes/css/bootstrap.min.css.map
```


```
$ make describe 
aws cloudformation describe-stacks \
                --stack-name cca670-p2-serverless-wildrydes
{
    "Stacks": [
        {
            "StackId": "arn:aws:cloudformation:us-east-1:191848360966:stack/cca670-p2-serverless-wildrydes/96dd0b10-51c4-11ea-94ad-0af228e676ce",
            "StackName": "cca670-p2-serverless-wildrydes",
            "ChangeSetId": "arn:aws:cloudformation:us-east-1:191848360966:changeSet/awscli-cloudformation-package-deploy-1581975268/40180732-ef6b-4a7e-88ef-addb4ab7da7d",
            "Description": "This stack sets up an S3-based website.",
            "Parameters": [
                {
                    "ParameterKey": "WebsiteBucket",
                    "ParameterValue": "cca670-p2-serverless-wildrydes"
                }
            ],
            "CreationTime": "2020-02-17T20:32:17.631Z",
            "LastUpdatedTime": "2020-02-17T21:34:34.159Z",
            "RollbackConfiguration": {},
            "StackStatus": "UPDATE_COMPLETE",
            "DisableRollback": false,
            "NotificationARNs": [],
            "Capabilities": [
                "CAPABILITY_NAMED_IAM"
            ],
            "Outputs": [
                {
                    "OutputKey": "CognitoUserPoolClientID",
                    "OutputValue": "5kf77be9fjpkjpud3lp6sob1an",
                    "Description": "UserPool Client ID"
                },
                {
                    "OutputKey": "CognitoUserPoolID",
                    "OutputValue": "us-east-1_qnwS4Cm7V",
                    "Description": "UserPool ID"
                },
                {
                    "OutputKey": "S3WebsiteBucket",
                    "OutputValue": "cca670-p2-serverless-wildrydes",
                    "Description": "S3 Bucket"
                },
                {
                    "OutputKey": "WildRydesApiInvokeUrl",
                    "OutputValue": "https://a0e8tnqg6j.execute-api.us-east-1.amazonaws.com/prod",
                    "Description": "URL for the deployed wild rydes API",
                    "ExportName": "WildRydesApiUrl"
                }
            ],
            "Tags": [],
            "EnableTerminationProtection": false,
            "DriftInformation": {
                "StackDriftStatus": "NOT_CHECKED"
            }
        }
    ]
}
```