package security
import (
    "errors"
    "strings"
)
// 
type ForceUserPasswordResetEntityIdentifier int

const (
    ACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER ForceUserPasswordResetEntityIdentifier = iota
    INITIATINGPROCESSACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
    REQUESTACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
    ONPREMSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
)

func (i ForceUserPasswordResetEntityIdentifier) String() string {
    var values []string
    for p := ForceUserPasswordResetEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_FORCEUSERPASSWORDRESETENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}[p])
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
                return 0, errors.New("Unknown ForceUserPasswordResetEntityIdentifier value: " + v)
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
