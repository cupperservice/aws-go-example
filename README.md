## 準備
以下が必要。
* VSCode + Remote - Containers
* Docker + Docker Compose

2. docker-compose.ymlを編集
以下の環境変数を設定する。
* AWS_ACCESS_KEY_ID : IAM ユーザーまたはロールに関連付けられる AWS アクセスキー
* AWS_SECRET_ACCESS_KEY : アクセスキーに関連付けられるシークレットキー
* AWS_DEFAULT_REGION : デフォルトリージョン
* AWS_REGION : リージョン
* S3_BACKET_NAME : S3バケット名
* FUNC_NAME : Lambda関数名

VSCode Remote Containersで、本プロジェクトを開くと、Docker Composeを使用して、Dockerコンテナが起動する。  
以降、Dockerコンテナで作業する。

-----

1. テンプレートの作成
以下のコマンドで、雛形を作成する。
```
# sam init --runtime go1.x --name sample-app
Which template source would you like to use?
        1 - AWS Quick Start Templates
        2 - Custom Template Location
Choice: 1

Allow SAM CLI to download AWS-provided quick start templates from Github [Y/n]: Y

-----------------------
Generating application:
-----------------------
Name: sample-app
Runtime: go1.x
Dependency Manager: mod
Application Template: hello-world
Output Directory: .

Next steps can be found in the README file at ./sample-app/README.md
```

2. Makefileの編集
作成した関数をパッケージして、AWS Lambdaに配備するコマンドをMakefileに追加する。

```
package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket ${S3_BACKET_NAME}

deploy:
	sam deploy --template-file output-template.yaml --stack-name ${FUNC_NAME} --capabilities CAPABILITY_IAM
```

3. Lambda関数の作成
main.goを修正する。

4. 必要なライブラリを取得
```
# go get github.com/aws/aws-lambda-go/lambda
# go get github.com/aws/aws-sdk-go
# go get github.com/ashwanthkumar/slack-go-webhook
```

4. template.yamlの編集

5. S3バケットの作成
CloudFormationのテンプレートを保存するS3バケットを作成する。

```
# aws s3 mb s3://${S3_BACKET_NAME}
```
6. パッケージの作成
以下のコマンドでテンプレートを作成する。

```
# make package
sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket cupper-sam-template-store

        SAM CLI now collects telemetry to better understand customer needs.

        You can OPT OUT and disable telemetry collection by setting the
        environment variable SAM_CLI_TELEMETRY=0 in your shell.
        Thanks for your help!

        Learn More: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-telemetry.html

2019-11-22 05:59:12 Found credentials in environment variables.
Uploading to 8146823f7fd54ea7609a71b3fc649766  1173 / 1173.0  (100.00%)Successfully packaged artifacts and wrote output template to file output-template.yaml.
Execute the following command to deploy the packaged template
sam deploy --template-file /go/src/github.com/yumemi-kkawashima/sample-app/output-template.yaml --stack-name <YOUR STACK NAME>
```
