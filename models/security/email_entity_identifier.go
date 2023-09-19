package security
import (
    "errors"
    "strings"
)
// 
type EmailEntityIdentifier int

const (
    NETWORKMESSAGEID_EMAILENTITYIDENTIFIER EmailEntityIdentifier = iota
    RECIPIENTEMAILADDRESS_EMAILENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER
)

func (i EmailEntityIdentifier) String() string {
    var values []string
    for p := EmailEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"networkMessageId", "recipientEmailAddress", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseEmailEntityIdentifier(v string) (any, error) {
    var result EmailEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "networkMessageId":
                result |= NETWORKMESSAGEID_EMAILENTITYIDENTIFIER
            case "recipientEmailAddress":
                result |= RECIPIENTEMAILADDRESS_EMAILENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown EmailEntityIdentifier value: " + v)
        }
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
func (i EmailEntityIdentifier) isMultiValue() bool {
    return true
}
