package models
import (
    "math"
    "strings"
)
type TokenFormat int

const (
    SAML_TOKENFORMAT = 1
    JWT_TOKENFORMAT = 2
    UNKNOWNFUTUREVALUE_TOKENFORMAT = 4
)

func (i TokenFormat) String() string {
    var values []string
    options := []string{"saml", "jwt", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := TokenFormat(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseTokenFormat(v string) (any, error) {
    var result TokenFormat
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "saml":
                result |= SAML_TOKENFORMAT
            case "jwt":
                result |= JWT_TOKENFORMAT
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_TOKENFORMAT
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeTokenFormat(values []TokenFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TokenFormat) isMultiValue() bool {
    return true
}
