package security
import (
    "errors"
    "math"
    "strings"
)
// 
type MarkUserAsCompromisedEntityIdentifier int

const (
    ACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER = 1
    INITIATINGPROCESSACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER = 2
    SERVICEPRINCIPALID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER = 4
    RECIPIENTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER = 8
    UNKNOWNFUTUREVALUE_MARKUSERASCOMPROMISEDENTITYIDENTIFIER = 16
)

func (i MarkUserAsCompromisedEntityIdentifier) String() string {
    var values []string
    options := []string{"accountObjectId", "initiatingProcessAccountObjectId", "servicePrincipalId", "recipientObjectId", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := MarkUserAsCompromisedEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMarkUserAsCompromisedEntityIdentifier(v string) (any, error) {
    var result MarkUserAsCompromisedEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "accountObjectId":
                result |= ACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
            case "initiatingProcessAccountObjectId":
                result |= INITIATINGPROCESSACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
            case "servicePrincipalId":
                result |= SERVICEPRINCIPALID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
            case "recipientObjectId":
                result |= RECIPIENTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown MarkUserAsCompromisedEntityIdentifier value: " + v)
        }
    }
    return &result, nil
}
func SerializeMarkUserAsCompromisedEntityIdentifier(values []MarkUserAsCompromisedEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MarkUserAsCompromisedEntityIdentifier) isMultiValue() bool {
    return true
}
