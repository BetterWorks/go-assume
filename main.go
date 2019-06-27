package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	region := flag.String("region", "", "AWS Region")
	accessKeyId := flag.String("key", "", "AWS Access Key Id")
	secretAccessKey := flag.String("secret", "", "AWS Secret Access Key")
	assumeRoleArn := flag.String("role-arn", "", "AWS Assume Role Arn")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	sessionName := strconv.FormatUint(rand.Uint64(), 10)

	svc := sts.New(session.New(&aws.Config{
		Region:      aws.String(*region),
		Credentials: credentials.NewStaticCredentials(*accessKeyId, *secretAccessKey, ""),
	}))

	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(*assumeRoleArn),
		RoleSessionName: aws.String(sessionName),
	}

	result, err := svc.AssumeRole(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeMalformedPolicyDocumentException:
				log.Fatal(sts.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case sts.ErrCodePackedPolicyTooLargeException:
				log.Fatal(sts.ErrCodePackedPolicyTooLargeException, aerr.Error())
			case sts.ErrCodeRegionDisabledException:
				log.Fatal(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatal(err.Error())
		}
		return
	}

	os.Setenv("AWS_SESSION_TOKEN", *result.Credentials.SessionToken)
	os.Setenv("AWS_SECRET_ACCESS_KEY", *result.Credentials.SecretAccessKey)
	os.Setenv("AWS_ACCESS_KEY_ID", *result.Credentials.AccessKeyId)
}
