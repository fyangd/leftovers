package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroup struct {
	client     securityGroupsClient
	id         *string
	identifier string
}

func NewSecurityGroup(client securityGroupsClient, id, groupName *string, tags []*awsec2.Tag) SecurityGroup {
	identifier := *groupName

	var extra []string
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		identifier = fmt.Sprintf("%s (%s)", *groupName, strings.Join(extra, ", "))
	}

	return SecurityGroup{
		client:     client,
		id:         id,
		identifier: identifier,
	}
}
