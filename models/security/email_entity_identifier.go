package security
import (
    "errors"
)
// 
type EmailEntityIdentifier int

const (
    NETWORKMESSAGEID_EMAILENTITYIDENTIFIER EmailEntityIdentifier = iota
    RECIPIENTEMAILADDRESS_EMAILENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER
)

func (i EmailEntityIdentifier) String() string {
    return []string{"networkMessageId", "recipientEmailAddress", "unknownFutureValue"}[i]
}
func ParseEmailEntityIdentifier(v string) (any, error) {
    result := NETWORKMESSAGEID_EMAILENTITYIDENTIFIER
    switch v {
        case "networkMessageId":
            result = NETWORKMESSAGEID_EMAILENTITYIDENTIFIER
        case "recipientEmailAddress":
            result = RECIPIENTEMAILADDRESS_EMAILENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown EmailEntityIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeEmailEntityIdentifier(values []EmailEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
