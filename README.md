# go-assume

A simple CLI for assuming a role and setting the session information to the default environment variables. Intended for automated usage like CI/CD pipelines

Exports 3 environment variables:
 - `AWS_ACCESS_KEY_ID`
 - `AWS_SECRET_ACCESS_KEY`
 - `AWS_ACCESS_KEY_ID`

## Usage

`eval "$(go-assume -region $REGION -key $ACCESS_KEY_ID -secret $SECRET_ACCESS_KEY -role-arn $ROLE_ARN)"`
