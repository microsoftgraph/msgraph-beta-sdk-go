package models
import (
    "errors"
    "strings"
)
// 
type AwsSecurityToolWebServices int

const (
    MACIE_AWSSECURITYTOOLWEBSERVICES AwsSecurityToolWebServices = iota
    WAFSHIELD_AWSSECURITYTOOLWEBSERVICES
    CLOUDTRAIL_AWSSECURITYTOOLWEBSERVICES
    INSPECTOR_AWSSECURITYTOOLWEBSERVICES
    SECURITYHUB_AWSSECURITYTOOLWEBSERVICES
    DETECTIVE_AWSSECURITYTOOLWEBSERVICES
    GUARDDUTY_AWSSECURITYTOOLWEBSERVICES
    UNKNOWNFUTUREVALUE_AWSSECURITYTOOLWEBSERVICES
)

func (i AwsSecurityToolWebServices) String() string {
    var values []string
    for p := AwsSecurityToolWebServices(1); p <= UNKNOWNFUTUREVALUE_AWSSECURITYTOOLWEBSERVICES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"macie", "wafShield", "cloudTrail", "inspector", "securityHub", "detective", "guardDuty", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAwsSecurityToolWebServices(v string) (any, error) {
    var result AwsSecurityToolWebServices
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "macie":
                result |= MACIE_AWSSECURITYTOOLWEBSERVICES
            case "wafShield":
                result |= WAFSHIELD_AWSSECURITYTOOLWEBSERVICES
            case "cloudTrail":
                result |= CLOUDTRAIL_AWSSECURITYTOOLWEBSERVICES
            case "inspector":
                result |= INSPECTOR_AWSSECURITYTOOLWEBSERVICES
            case "securityHub":
                result |= SECURITYHUB_AWSSECURITYTOOLWEBSERVICES
            case "detective":
                result |= DETECTIVE_AWSSECURITYTOOLWEBSERVICES
            case "guardDuty":
                result |= GUARDDUTY_AWSSECURITYTOOLWEBSERVICES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_AWSSECURITYTOOLWEBSERVICES
            default:
                return 0, errors.New("Unknown AwsSecurityToolWebServices value: " + v)
        }
    }
    return &result, nil
}
func SerializeAwsSecurityToolWebServices(values []AwsSecurityToolWebServices) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsSecurityToolWebServices) isMultiValue() bool {
    return true
}
