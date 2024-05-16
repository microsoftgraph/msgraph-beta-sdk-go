package security
import (
    "math"
    "strings"
)
type StopAndQuarantineFileEntityIdentifier int

const (
    DEVICEID_STOPANDQUARANTINEFILEENTITYIDENTIFIER = 1
    SHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER = 2
    INITIATINGPROCESSSHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER = 4
    UNKNOWNFUTUREVALUE_STOPANDQUARANTINEFILEENTITYIDENTIFIER = 8
)

func (i StopAndQuarantineFileEntityIdentifier) String() string {
    var values []string
    options := []string{"deviceId", "sha1", "initiatingProcessSHA1", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := StopAndQuarantineFileEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
