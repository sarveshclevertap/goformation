package resources

// AWS::ElasticLoadBalancing::LoadBalancer.Policies AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
type AWSElasticLoadBalancingLoadBalancerPolicies struct {
    
    // Attributes AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-attributes
    Attributes []string `json:"Attributes"`
    
    // InstancePorts AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-instanceports
    InstancePorts []string `json:"InstancePorts"`
    
    // LoadBalancerPorts AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-loadbalancerports
    LoadBalancerPorts []string `json:"LoadBalancerPorts"`
    
    // PolicyName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policyname
    PolicyName string `json:"PolicyName"`
    
    // PolicyType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policytype
    PolicyType string `json:"PolicyType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancerPolicies) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancing::LoadBalancer.Policies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancerPolicies) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}