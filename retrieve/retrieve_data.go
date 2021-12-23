package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func retrieveAllRecords() {
	var err error
	var items []map[string]*dynamodb.AttributeValue
	var scanResult *dynamodb.ScanOutput

	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("MedicationUse"),
	}

	if scanResult, err = svc.Scan(scanInput); err != nil {
		panic(err)
	}

	keyExists := true

	for keyExists {
		if scanResult.LastEvaluatedKey != nil {
			for index := range scanResult.Items {
				dynamodbattribute.UnmarshalMap(scanResult.Items[index], items)
			}
	
			scanInput = &dynamodb.ScanInput{
				TableName: aws.String("MedicationUse"),
				ExclusiveStartKey: scanResult.LastEvaluatedKey,
			}
		} else {
			for index := range scanResult.Items {
				dynamodbattribute.UnmarshalMap(scanResult.Items[index], items)
			}
			keyExists = false
		}
	}

	fmt.Sprintf("the scan result is %v", items)
}