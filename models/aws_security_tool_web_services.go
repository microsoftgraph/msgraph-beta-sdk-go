package models
import (
    "errors"
    "math"
    "strings"
)
type AwsSecurityToolWebServices int

const (
    MACIE_AWSSECURITYTOOLWEBSERVICES = 1
    WAFSHIELD_AWSSECURITYTOOLWEBSERVICES = 2
    CLOUDTRAIL_AWSSECURITYTOOLWEBSERVICES = 4
    INSPECTOR_AWSSECURITYTOOLWEBSERVICES = 8
    SECURITYHUB_AWSSECURITYTOOLWEBSERVICES = 16
    DETECTIVE_AWSSECURITYTOOLWEBSERVICES = 32
    GUARDDUTY_AWSSECURITYTOOLWEBSERVICES = 64
    UNKNOWNFUTUREVALUE_AWSSECURITYTOOLWEBSERVICES = 128
)

func (i AwsSecurityToolWebServices) String() string {
    var values []string
    options := []string{"macie", "wafShield", "cloudTrail", "inspector", "securityHub", "detective", "guardDuty", "unknownFutureValue"}
    for p := 0; p < 8; p++ {
        mantis := AwsSecurityToolWebServices(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
