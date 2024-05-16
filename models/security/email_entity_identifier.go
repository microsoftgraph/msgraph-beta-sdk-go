package security
import (
    "math"
    "strings"
)
type EmailEntityIdentifier int

const (
    NETWORKMESSAGEID_EMAILENTITYIDENTIFIER = 1
    RECIPIENTEMAILADDRESS_EMAILENTITYIDENTIFIER = 2
    UNKNOWNFUTUREVALUE_EMAILENTITYIDENTIFIER = 4
)

func (i EmailEntityIdentifier) String() string {
    var values []string
    options := []string{"networkMessageId", "recipientEmailAddress", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := EmailEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
