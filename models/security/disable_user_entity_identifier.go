package security
import (
    "errors"
    "strings"
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
    var values []string
    for p := DisableUserEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_DISABLEUSERENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDisableUserEntityIdentifier(v string) (any, error) {
    var result DisableUserEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "accountSid":
                result |= ACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
            case "initiatingProcessAccountSid":
                result |= INITIATINGPROCESSACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
            case "requestAccountSid":
                result |= REQUESTACCOUNTSID_DISABLEUSERENTITYIDENTIFIER
            case "onPremSid":
                result |= ONPREMSID_DISABLEUSERENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DISABLEUSERENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown DisableUserEntityIdentifier value: " + v)
        }
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
func (i DisableUserEntityIdentifier) isMultiValue() bool {
    return true
}
