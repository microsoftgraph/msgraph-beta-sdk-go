package models
type AwsAccessType int

const (
    PUBLIC_AWSACCESSTYPE AwsAccessType = iota
    RESTRICTED_AWSACCESSTYPE
    CROSSACCOUNT_AWSACCESSTYPE
    PRIVATE_AWSACCESSTYPE
    UNKNOWNFUTUREVALUE_AWSACCESSTYPE
)

func (i AwsAccessType) String() string {
    return []string{"public", "restricted", "crossAccount", "private", "unknownFutureValue"}[i]
}
func ParseAwsAccessType(v string) (any, error) {
    result := PUBLIC_AWSACCESSTYPE
    switch v {
        case "public":
            result = PUBLIC_AWSACCESSTYPE
        case "restricted":
            result = RESTRICTED_AWSACCESSTYPE
        case "crossAccount":
            result = CROSSACCOUNT_AWSACCESSTYPE
        case "private":
            result = PRIVATE_AWSACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AWSACCESSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAwsAccessType(values []AwsAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsAccessType) isMultiValue() bool {
    return false
}
