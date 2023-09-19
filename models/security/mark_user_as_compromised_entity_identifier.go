package security
import (
    "errors"
    "strings"
)
// 
type MarkUserAsCompromisedEntityIdentifier int

const (
    ACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER MarkUserAsCompromisedEntityIdentifier = iota
    INITIATINGPROCESSACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
    SERVICEPRINCIPALID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
    RECIPIENTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
)

func (i MarkUserAsCompromisedEntityIdentifier) String() string {
    var values []string
    for p := MarkUserAsCompromisedEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_MARKUSERASCOMPROMISEDENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"accountObjectId", "initiatingProcessAccountObjectId", "servicePrincipalId", "recipientObjectId", "unknownFutureValue"}[p])
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
