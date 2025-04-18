// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security
import (
    "math"
    "strings"
)
type RecipientType int

const (
    USER_RECIPIENTTYPE = 1
    ROLEGROUP_RECIPIENTTYPE = 2
    UNKNOWNFUTUREVALUE_RECIPIENTTYPE = 4
)

func (i RecipientType) String() string {
    var values []string
    options := []string{"user", "roleGroup", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := RecipientType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseRecipientType(v string) (any, error) {
    var result RecipientType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "user":
                result |= USER_RECIPIENTTYPE
            case "roleGroup":
                result |= ROLEGROUP_RECIPIENTTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_RECIPIENTTYPE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeRecipientType(values []RecipientType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RecipientType) isMultiValue() bool {
    return true
}
