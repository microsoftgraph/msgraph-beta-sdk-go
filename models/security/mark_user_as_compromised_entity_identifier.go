package security
import (
    "errors"
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
    return []string{"accountObjectId", "initiatingProcessAccountObjectId", "servicePrincipalId", "recipientObjectId", "unknownFutureValue"}[i]
}
func ParseMarkUserAsCompromisedEntityIdentifier(v string) (any, error) {
    result := ACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
    switch v {
        case "accountObjectId":
            result = ACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
        case "initiatingProcessAccountObjectId":
            result = INITIATINGPROCESSACCOUNTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
        case "servicePrincipalId":
            result = SERVICEPRINCIPALID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
        case "recipientObjectId":
            result = RECIPIENTOBJECTID_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MARKUSERASCOMPROMISEDENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown MarkUserAsCompromisedEntityIdentifier value: " + v)
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
