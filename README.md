# go-assume

A simple CLI for assuming a role and setting the session information to the default environment variables. Intended for automated usage like CI/CD pipelines

## Usage

go-assume -region $REGION -key $ACCESS_KEY_ID -secret $SECRET_ACCESS_KEY -role-arn $ROLE_ARN
