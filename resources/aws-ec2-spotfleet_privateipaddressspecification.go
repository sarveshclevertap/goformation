package resources

// AWS::EC2::SpotFleet.PrivateIpAddressSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html
type AWSEC2SpotFleetPrivateIpAddressSpecification struct {
    
    // Primary AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
    Primary bool `json:"Primary"`
    
    // PrivateIpAddress AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
    PrivateIpAddress string `json:"PrivateIpAddress"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleetPrivateIpAddressSpecification) AWSCloudFormationType() string {
    return "AWS::EC2::SpotFleet.PrivateIpAddressSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleetPrivateIpAddressSpecification) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}