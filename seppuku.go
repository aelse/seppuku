/*
Package seppuku is a tool for requesting self-termination of an EC2 instance in an autoscale group.

Running the seppuku command on an EC2 instance will set the instance health status to Unhealthy,
causing it to be destroyed and replaced according to autoscale policy.
*/
package seppuku

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func Seppuku() error {
	mdSvc := ec2metadata.New(session.New())
	doc, err := mdSvc.GetInstanceIdentityDocument()
	if err != nil {
		return err
	}
	svc := autoscaling.New(session.New(), &aws.Config{Region: aws.String(doc.Region)})
	input := &autoscaling.SetInstanceHealthInput{
		HealthStatus: aws.String("Unhealthy"),
		InstanceId:   aws.String(doc.InstanceID),
	}
	if _, err = svc.SetInstanceHealth(input); err != nil {
		return err
	}
	return nil
}
