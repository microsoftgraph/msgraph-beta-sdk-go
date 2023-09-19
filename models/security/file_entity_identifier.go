package security
import (
    "errors"
    "strings"
)
// 
type FileEntityIdentifier int

const (
    SHA1_FILEENTITYIDENTIFIER FileEntityIdentifier = iota
    INITIATINGPROCESSSHA1_FILEENTITYIDENTIFIER
    SHA256_FILEENTITYIDENTIFIER
    INITIATINGPROCESSSHA256_FILEENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_FILEENTITYIDENTIFIER
)

func (i FileEntityIdentifier) String() string {
    var values []string
    for p := FileEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_FILEENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"sha1", "initiatingProcessSHA1", "sha256", "initiatingProcessSHA256", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseFileEntityIdentifier(v string) (any, error) {
    var result FileEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "sha1":
                result |= SHA1_FILEENTITYIDENTIFIER
            case "initiatingProcessSHA1":
                result |= INITIATINGPROCESSSHA1_FILEENTITYIDENTIFIER
            case "sha256":
                result |= SHA256_FILEENTITYIDENTIFIER
            case "initiatingProcessSHA256":
                result |= INITIATINGPROCESSSHA256_FILEENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_FILEENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown FileEntityIdentifier value: " + v)
        }
    }
    return &result, nil
}
func SerializeFileEntityIdentifier(values []FileEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FileEntityIdentifier) isMultiValue() bool {
    return true
}
