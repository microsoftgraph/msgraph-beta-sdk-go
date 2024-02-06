package security
import (
    "errors"
    "math"
    "strings"
)
// 
type FileEntityIdentifier int

const (
    SHA1_FILEENTITYIDENTIFIER = 1
    INITIATINGPROCESSSHA1_FILEENTITYIDENTIFIER = 2
    SHA256_FILEENTITYIDENTIFIER = 4
    INITIATINGPROCESSSHA256_FILEENTITYIDENTIFIER = 8
    UNKNOWNFUTUREVALUE_FILEENTITYIDENTIFIER = 16
)

func (i FileEntityIdentifier) String() string {
    var values []string
    options := []string{"sha1", "initiatingProcessSHA1", "sha256", "initiatingProcessSHA256", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := FileEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
