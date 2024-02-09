package models
import (
    "errors"
    "math"
    "strings"
)
type AwsSecretInformationWebServices int

const (
    SECRETSMANAGER_AWSSECRETINFORMATIONWEBSERVICES = 1
    CERTIFICATEAUTHORITY_AWSSECRETINFORMATIONWEBSERVICES = 2
    CLOUDHSM_AWSSECRETINFORMATIONWEBSERVICES = 4
    CERTIFICATEMANAGER_AWSSECRETINFORMATIONWEBSERVICES = 8
    UNKNOWNFUTUREVALUE_AWSSECRETINFORMATIONWEBSERVICES = 16
)

func (i AwsSecretInformationWebServices) String() string {
    var values []string
    options := []string{"secretsManager", "certificateAuthority", "cloudHsm", "certificateManager", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := AwsSecretInformationWebServices(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAwsSecretInformationWebServices(v string) (any, error) {
    var result AwsSecretInformationWebServices
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "secretsManager":
                result |= SECRETSMANAGER_AWSSECRETINFORMATIONWEBSERVICES
            case "certificateAuthority":
                result |= CERTIFICATEAUTHORITY_AWSSECRETINFORMATIONWEBSERVICES
            case "cloudHsm":
                result |= CLOUDHSM_AWSSECRETINFORMATIONWEBSERVICES
            case "certificateManager":
                result |= CERTIFICATEMANAGER_AWSSECRETINFORMATIONWEBSERVICES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_AWSSECRETINFORMATIONWEBSERVICES
            default:
                return 0, errors.New("Unknown AwsSecretInformationWebServices value: " + v)
        }
    }
    return &result, nil
}
func SerializeAwsSecretInformationWebServices(values []AwsSecretInformationWebServices) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsSecretInformationWebServices) isMultiValue() bool {
    return true
}
