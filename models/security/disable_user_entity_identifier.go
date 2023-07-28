package security
import (
    "errors"
)
// 
type DisableUserEntityIdentifier int

const (
    ACCOUNTSID_DISABLEUSERENTITYIDENTIFIER DisableUserEntityIdentifier = iota
    INITIATINGPROCESSACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
    REQUESTACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
    ONPREMSID_DISABLEUSERENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_DISABLEUSERENTITYIDENTIFIER
)

func (i DisableUserEntityIdentifier) String() string {
    return []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}[i]
}
func ParseDisableUserEntityIdentifier(v string) (any, error) {
    result := ACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
    switch v {
        case "accountSid":
            result = ACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
        case "initiatingProcessAccountSid":
            result = INITIATINGPROCESSACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
        case "requestAccountSid":
            result = REQUESTACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
        case "onPremSid":
            result = ONPREMSID_DISABLEUSERENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DISABLEUSERENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown DisableUserEntityIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeDisableUserEntityIdentifier(values []DisableUserEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
