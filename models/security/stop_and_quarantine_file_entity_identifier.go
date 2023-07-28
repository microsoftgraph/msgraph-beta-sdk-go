package security
import (
    "errors"
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
    return []string{"deviceId", "sha1", "initiatingProcessSHA1", "unknownFutureValue"}[i]
}
func ParseStopAndQuarantineFileEntityIdentifier(v string) (any, error) {
    result := DEVICEID_STOPANDQUARANTINEFILEENTITYIDENTIFIER
    switch v {
        case "deviceId":
            result = DEVICEID_STOPANDQUARANTINEFILEENTITYIDENTIFIER
        case "sha1":
            result = SHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
        case "initiatingProcessSHA1":
            result = INITIATINGPROCESSSHA1_STOPANDQUARANTINEFILEENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_STOPANDQUARANTINEFILEENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown StopAndQuarantineFileEntityIdentifier value: " + v)
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
