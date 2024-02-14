package models
import (
    "errors"
)
type AwsStatementEffect int

const (
    ALLOW_AWSSTATEMENTEFFECT AwsStatementEffect = iota
    DENY_AWSSTATEMENTEFFECT
    UNKNOWNFUTUREVALUE_AWSSTATEMENTEFFECT
)

func (i AwsStatementEffect) String() string {
    return []string{"allow", "deny", "unknownFutureValue"}[i]
}
func ParseAwsStatementEffect(v string) (any, error) {
    result := ALLOW_AWSSTATEMENTEFFECT
    switch v {
        case "allow":
            result = ALLOW_AWSSTATEMENTEFFECT
        case "deny":
            result = DENY_AWSSTATEMENTEFFECT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AWSSTATEMENTEFFECT
        default:
            return 0, errors.New("Unknown AwsStatementEffect value: " + v)
    }
    return &result, nil
}
func SerializeAwsStatementEffect(values []AwsStatementEffect) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsStatementEffect) isMultiValue() bool {
    return false
}
