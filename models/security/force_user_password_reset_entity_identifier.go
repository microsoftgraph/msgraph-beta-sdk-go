package security
import (
    "math"
    "strings"
)
type ForceUserPasswordResetEntityIdentifier int

const (
    ACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER = 1
    INITIATINGPROCESSACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER = 2
    REQUESTACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER = 4
    ONPREMSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER = 8
    UNKNOWNFUTUREVALUE_FORCEUSERPASSWORDRESETENTITYIDENTIFIER = 16
)

func (i ForceUserPasswordResetEntityIdentifier) String() string {
    var values []string
    options := []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := ForceUserPasswordResetEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseForceUserPasswordResetEntityIdentifier(v string) (any, error) {
    var result ForceUserPasswordResetEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "accountSid":
                result |= ACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
            case "initiatingProcessAccountSid":
                result |= INITIATINGPROCESSACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
            case "requestAccountSid":
                result |= REQUESTACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
            case "onPremSid":
                result |= ONPREMSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeForceUserPasswordResetEntityIdentifier(values []ForceUserPasswordResetEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ForceUserPasswordResetEntityIdentifier) isMultiValue() bool {
    return true
}
