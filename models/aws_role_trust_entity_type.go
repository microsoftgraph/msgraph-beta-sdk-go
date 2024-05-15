package models
import (
    "math"
    "strings"
)
type AwsRoleTrustEntityType int

const (
    NONE_AWSROLETRUSTENTITYTYPE = 1
    SERVICE_AWSROLETRUSTENTITYTYPE = 2
    SSO_AWSROLETRUSTENTITYTYPE = 4
    CROSSACCOUNT_AWSROLETRUSTENTITYTYPE = 8
    WEBIDENTITY_AWSROLETRUSTENTITYTYPE = 16
    UNKNOWNFUTUREVALUE_AWSROLETRUSTENTITYTYPE = 32
)

func (i AwsRoleTrustEntityType) String() string {
    var values []string
    options := []string{"none", "service", "sso", "crossAccount", "webIdentity", "unknownFutureValue"}
    for p := 0; p < 6; p++ {
        mantis := AwsRoleTrustEntityType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
