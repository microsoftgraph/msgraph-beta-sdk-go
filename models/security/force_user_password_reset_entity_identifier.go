package security
import (
    "errors"
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
    return []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}[i]
}
func ParseForceUserPasswordResetEntityIdentifier(v string) (any, error) {
    result := ACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
    switch v {
        case "accountSid":
            result = ACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
        case "initiatingProcessAccountSid":
            result = INITIATINGPROCESSACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
        case "requestAccountSid":
            result = REQUESTACCOUNTSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
        case "onPremSid":
            result = ONPREMSID_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FORCEUSERPASSWORDRESETENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown ForceUserPasswordResetEntityIdentifier value: " + v)
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
