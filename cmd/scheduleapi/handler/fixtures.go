package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoJSON = dynamodb.QueryOutput{
	LastEvaluatedKey: map[string]*dynamodb.AttributeValue{
		"DESTINATION": {
			S: aws.String("Indian Creek"),
		},
		"DIRECTION": {
			S: aws.String("E"),
		},
		"EVENT_TIME": {
			S: aws.String("6/27/2019 11:22:20 AM"),
		},
		"LINE": {
			S: aws.String("BLUE"),
		},
		"NEXT_ARR": {
			S: aws.String("11:23:37 AM"),
		},
		"PrimaryKey": {
			S: aws.String("FIVE POINTS STATION_Indian Creek_2019-06-27"),
		},
		"SortKey": {
			S: aws.String("6/27/2019 11:22:20 AM_104026"),
		},
		"STATION": {
			S: aws.String("FIVE POINTS STATION"),
		},
		"TRAIN_ID": {
			S: aws.String("104026"),
		},
		"TTL": {
			N: aws.String("1564240951"),
		},
		"WAITING_SECONDS": {
			S: aws.String("66"),
		},
		"WAITING_TIME": {
			S: aws.String("Arriving"),
		},
	},
	Items: []map[string]*dynamodb.AttributeValue{
		map[string]*dynamodb.AttributeValue{
			"DESTINATION": {
				S: aws.String("Indian Creek"),
			},
			"DIRECTION": {
				S: aws.String("E"),
			},
			"EVENT_TIME": {
				S: aws.String("6/27/2019 11:22:20 AM"),
			},
			"LINE": {
				S: aws.String("BLUE"),
			},
			"NEXT_ARR": {
				S: aws.String("11:23:37 AM"),
			},
			"PrimaryKey": {
				S: aws.String("FIVE POINTS STATION_Indian Creek_2019-06-27"),
			},
			"SortKey": {
				S: aws.String("6/27/2019 11:22:20 AM_104026"),
			},
			"STATION": {
				S: aws.String("FIVE POINTS STATION"),
			},
			"TRAIN_ID": {
				S: aws.String("104026"),
			},
			"TTL": {
				N: aws.String("1564240951"),
			},
			"WAITING_SECONDS": {
				S: aws.String("66"),
			},
			"WAITING_TIME": {
				S: aws.String("Arriving"),
			},
		},
		map[string]*dynamodb.AttributeValue{
			"DESTINATION": {
				S: aws.String("Indian Creek"),
			},
			"DIRECTION": {
				S: aws.String("E"),
			},
			"EVENT_TIME": {
				S: aws.String("6/27/2019 11:22:20 AM"),
			},
			"LINE": {
				S: aws.String("BLUE"),
			},
			"NEXT_ARR": {
				S: aws.String("11:23:37 AM"),
			},
			"PrimaryKey": {
				S: aws.String("FIVE POINTS STATION_Indian Creek_2019-06-27"),
			},
			"SortKey": {
				S: aws.String("6/27/2019 11:22:20 AM_104026"),
			},
			"STATION": {
				S: aws.String("FIVE POINTS STATION"),
			},
			"TRAIN_ID": {
				S: aws.String("104026"),
			},
			"TTL": {
				N: aws.String("1564240951"),
			},
			"WAITING_SECONDS": {
				S: aws.String("66"),
			},
			"WAITING_TIME": {
				S: aws.String("Arriving"),
			},
		},
		map[string]*dynamodb.AttributeValue{
			"DESTINATION": {
				S: aws.String("Indian Creek"),
			},
			"DIRECTION": {
				S: aws.String("E"),
			},
			"EVENT_TIME": {
				S: aws.String("6/27/2019 11:22:20 AM"),
			},
			"LINE": {
				S: aws.String("BLUE"),
			},
			"NEXT_ARR": {
				S: aws.String("11:23:37 AM"),
			},
			"PrimaryKey": {
				S: aws.String("FIVE POINTS STATION_Indian Creek_2019-06-27"),
			},
			"SortKey": {
				S: aws.String("6/27/2019 11:22:20 AM_104026"),
			},
			"STATION": {
				S: aws.String("FIVE POINTS STATION"),
			},
			"TRAIN_ID": {
				S: aws.String("104026"),
			},
			"TTL": {
				N: aws.String("1564240951"),
			},
			"WAITING_SECONDS": {
				S: aws.String("66"),
			},
			"WAITING_TIME": {
				S: aws.String("Arriving"),
			},
		},
	},
}
