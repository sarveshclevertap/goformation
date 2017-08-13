package resources

// AWS::RDS::DBSecurityGroup.Ingress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
type AWSRDSDBSecurityGroupIngress struct {
    
    // CIDRIP AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-cidrip
    CIDRIP string `json:"CIDRIP"`
    
    // EC2SecurityGroupId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupid
    EC2SecurityGroupId string `json:"EC2SecurityGroupId"`
    
    // EC2SecurityGroupName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupname
    EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
    
    // EC2SecurityGroupOwnerId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupownerid
    EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSecurityGroupIngress) AWSCloudFormationType() string {
    return "AWS::RDS::DBSecurityGroup.Ingress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBSecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}