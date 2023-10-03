package models
import (
    "errors"
    "strings"
)
// 
type AwsRoleTrustEntityType int

const (
    NONE_AWSROLETRUSTENTITYTYPE AwsRoleTrustEntityType = iota
    SERVICE_AWSROLETRUSTENTITYTYPE
    SSO_AWSROLETRUSTENTITYTYPE
    CROSSACCOUNT_AWSROLETRUSTENTITYTYPE
    WEBIDENTITY_AWSROLETRUSTENTITYTYPE
    UNKNOWNFUTUREVALUE_AWSROLETRUSTENTITYTYPE
)

func (i AwsRoleTrustEntityType) String() string {
    var values []string
    for p := AwsRoleTrustEntityType(1); p <= UNKNOWNFUTUREVALUE_AWSROLETRUSTENTITYTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "service", "sso", "crossAccount", "webIdentity", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAwsRoleTrustEntityType(v string) (any, error) {
    var result AwsRoleTrustEntityType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_AWSROLETRUSTENTITYTYPE
            case "service":
                result |= SERVICE_AWSROLETRUSTENTITYTYPE
            case "sso":
                result |= SSO_AWSROLETRUSTENTITYTYPE
            case "crossAccount":
                result |= CROSSACCOUNT_AWSROLETRUSTENTITYTYPE
            case "webIdentity":
                result |= WEBIDENTITY_AWSROLETRUSTENTITYTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_AWSROLETRUSTENTITYTYPE
            default:
                return 0, errors.New("Unknown AwsRoleTrustEntityType value: " + v)
        }
    }
    return &result, nil
}
func SerializeAwsRoleTrustEntityType(values []AwsRoleTrustEntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsRoleTrustEntityType) isMultiValue() bool {
    return true
}
