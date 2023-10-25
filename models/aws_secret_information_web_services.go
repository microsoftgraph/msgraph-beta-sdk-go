package models
import (
    "errors"
    "strings"
)
// 
type AwsSecretInformationWebServices int

const (
    SECRETSMANAGER_AWSSECRETINFORMATIONWEBSERVICES AwsSecretInformationWebServices = iota
    CERTIFICATEAUTHORITY_AWSSECRETINFORMATIONWEBSERVICES
    CLOUDHSM_AWSSECRETINFORMATIONWEBSERVICES
    CERTIFICATEMANAGER_AWSSECRETINFORMATIONWEBSERVICES
    UNKNOWNFUTUREVALUE_AWSSECRETINFORMATIONWEBSERVICES
)

func (i AwsSecretInformationWebServices) String() string {
    var values []string
    for p := AwsSecretInformationWebServices(1); p <= UNKNOWNFUTUREVALUE_AWSSECRETINFORMATIONWEBSERVICES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"secretsManager", "certificateAuthority", "cloudHsm", "certificateManager", "unknownFutureValue"}[p])
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
