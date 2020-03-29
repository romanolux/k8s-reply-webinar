package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
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
	fmt.Printf("Instance Name: %s /n", k8sdynamo.Name)
	fmt.Printf("Instance Values: %s %d %d /n", tableName, readCapacity, writeCapacity)
	callAWSClient(tableName, readCapacity, writeCapacity)
	// fmt.Println("Object Created")
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	fmt.Println("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
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
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		// os.Exit(1)
		return
	}
	fmt.Println(res)
	fmt.Println("Created the table", tableName)
}
