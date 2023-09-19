package security
import (
    "errors"
    "strings"
)
// 
type StopAndQuarantineFileEntityIdentifier int

const (
    DEVICEID_STOPANDQUARANTINEFILEENTITYIDENTIFIER StopAndQuarantineFileEntityIdentifier = iota
    SHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
    INITIATINGPROCESSSHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
    UNKNOWNFUTUREVALUE_STOPANDQUARANTINEFILEENTITYIDENTIFIER
)

func (i StopAndQuarantineFileEntityIdentifier) String() string {
    var values []string
    for p := StopAndQuarantineFileEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_STOPANDQUARANTINEFILEENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"deviceId", "sha1", "initiatingProcessSHA1", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseStopAndQuarantineFileEntityIdentifier(v string) (any, error) {
    var result StopAndQuarantineFileEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "deviceId":
                result |= DEVICEID_STOPANDQUARANTINEFILEENTITYIDENTIFIER
            case "sha1":
                result |= SHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
            case "initiatingProcessSHA1":
                result |= INITIATINGPROCESSSHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_STOPANDQUARANTINEFILEENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown StopAndQuarantineFileEntityIdentifier value: " + v)
        }
    }
    return &result, nil
}
func SerializeStopAndQuarantineFileEntityIdentifier(values []StopAndQuarantineFileEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StopAndQuarantineFileEntityIdentifier) isMultiValue() bool {
    return true
}
