package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := costexplorer.NewFromConfig(cfg, func(o *costexplorer.Options) {
		o.Region = "ap-northeast-1"
	})

	// TimePeriod
	// 現在時刻の取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().UTC().In(jst)
	dayBefore := now.AddDate(0, 0, -31)

	nowDate := now.Format("2006-01-02")
	dateBefore := dayBefore.Format("2006-01-02")

	// 昨日から今日まで
	timePeriod := &types.DateInterval{
		Start: aws.String(dateBefore),
		End:   aws.String(nowDate),
	}

	// GroupBy
	// group := types.GroupDefinition{
	// 	Key:  aws.String("SERVICE"),
	// 	Type: types.GroupDefinitionTypeDimension,
	// }
	// groups := []*types.GroupDefinition{&group}

	output, err := client.GetCostAndUsage(context.TODO(), &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
		TimePeriod:  timePeriod,
		// GroupBy:     groups,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)
}
