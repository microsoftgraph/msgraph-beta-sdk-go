package models
import (
    "errors"
)
type AwsRoleType int

const (
    SYSTEM_AWSROLETYPE AwsRoleType = iota
    CUSTOM_AWSROLETYPE
    UNKNOWNFUTUREVALUE_AWSROLETYPE
)

func (i AwsRoleType) String() string {
    return []string{"system", "custom", "unknownFutureValue"}[i]
}
func ParseAwsRoleType(v string) (any, error) {
    result := SYSTEM_AWSROLETYPE
    switch v {
        case "system":
            result = SYSTEM_AWSROLETYPE
        case "custom":
            result = CUSTOM_AWSROLETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AWSROLETYPE
        default:
            return 0, errors.New("Unknown AwsRoleType value: " + v)
    }
    return &result, nil
}
func SerializeAwsRoleType(values []AwsRoleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsRoleType) isMultiValue() bool {
    return false
}
