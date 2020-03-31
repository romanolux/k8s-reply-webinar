package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	stormv1 "github.com/romanolux/k8s-reply-webinar/pkg/apis/k8dynamo/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	fmt.Println("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	k8sdynamo, _ := obj.(*stormv1.K8dynamo)

	tableName := k8sdynamo.Spec.TableName
	readCapacity := int64(*k8sdynamo.Spec.ReadUnit)
	writeCapacity := int64(*k8sdynamo.Spec.WriteUnit)
	// fmt.Printf("%s %d %d", tableName, readCapacity, writeCapacity)
	fmt.Println("Start Object Creation")
	fmt.Printf("Instance Name: %s \n", k8sdynamo.Name)
	fmt.Printf("Instance Values: %s %d %d \n", tableName, readCapacity, writeCapacity)
	callAWSClient(k8sdynamo.Name, readCapacity, writeCapacity)
	// fmt.Println("Object Created")
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {

	tableName := strings.Split(obj.(string), "/")[1]
	fmt.Println("Object deleted")
	fmt.Println(tableName)
	// fmt.Println(tableName)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "saml"),
	})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	}
	result, err := svc.DeleteTable(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	// if aerr, ok := err.(awserr.Error); ok {
	// 	switch aerr.Code() {
	// 	case dynamodb.ErrCodeResourceInUseException:
	// 		fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
	// 	case dynamodb.ErrCodeResourceNotFoundException:
	// 		fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
	// 	case dynamodb.ErrCodeLimitExceededException:
	// 		fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
	// 	case dynamodb.ErrCodeInternalServerError:
	// 		fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
	// 	default:
	// 		fmt.Println(aerr.Error())
	// 	}
	// } else {
	// 	// Print the error, cast err to awserr.Error to get the Code and
	// 	// Message from an error.
	// 	fmt.Println(err.Error())
	// }
	// return
	// }
	fmt.Println(result)
	fmt.Println("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	k8sdynamo, _ := objNew.(*stormv1.K8dynamo)

	tableName := k8sdynamo.Spec.TableName
	readCapacity := int64(*k8sdynamo.Spec.ReadUnit)
	writeCapacity := int64(*k8sdynamo.Spec.WriteUnit)
	// fmt.Printf("%s %d %d", tableName, readCapacity, writeCapacity)
	fmt.Println("Start Object Creation")
	fmt.Printf("Instance Name: %s \n", k8sdynamo.Name)
	fmt.Printf("Instance Values: %s %d %d \n", tableName, readCapacity, writeCapacity)
	// callAWSClient(k8sdynamo.Name, readCapacity, writeCapacity)
	// fmt.Println("Object Created")
	fmt.Println("TestHandler.ObjectUpdated")
}

func callAWSClient(tableName string, read, write int64) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "saml"),
	})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Value"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Value"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(read),
			WriteCapacityUnits: aws.Int64(write),
		},
		TableName: aws.String(tableName),
	}

	res, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable")
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				input := &dynamodb.UpdateTableInput{
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(read),
						WriteCapacityUnits: aws.Int64(write),
					},
					TableName: aws.String(tableName),
				}
				res, _ := svc.UpdateTable(input)
				fmt.Println(res)
			default:
				fmt.Printf("Got error calling CreateTable: %s \n", awsErr.Code())
			}
		}

		// os.Exit(1)
		return
	}
	fmt.Println(res)
	fmt.Println("Created the table", tableName)
}
