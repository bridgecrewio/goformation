package resources


// AWSWAFXssMatchSet_XssMatchTuple AWS CloudFormation Resource (AWS::WAF::XssMatchSet.XssMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html
type AWSWAFXssMatchSet_XssMatchTuple struct {
    
    // FieldToMatch AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch
    FieldToMatch *AWSWAFXssMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`
    
    // TextTransformation AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-texttransformation
    TextTransformation string `json:"TextTransformation,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
    return "AWS::WAF::XssMatchSet.XssMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFXssMatchSet_XssMatchTuple) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}