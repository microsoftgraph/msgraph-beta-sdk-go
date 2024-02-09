package security
import (
    "errors"
    "math"
    "strings"
)
type DisableUserEntityIdentifier int

const (
    ACCOUNTSID_DISABLEUSERENTITYIDENTIFIER = 1
    INITIATINGPROCESSACCOUNTSID_DISABLEUSERENTITYIDENTIFIER = 2
    REQUESTACCOUNTSID_DISABLEUSERENTITYIDENTIFIER = 4
    ONPREMSID_DISABLEUSERENTITYIDENTIFIER = 8
    UNKNOWNFUTUREVALUE_DISABLEUSERENTITYIDENTIFIER = 16
)

func (i DisableUserEntityIdentifier) String() string {
    var values []string
    options := []string{"accountSid", "initiatingProcessAccountSid", "requestAccountSid", "onPremSid", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := DisableUserEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
