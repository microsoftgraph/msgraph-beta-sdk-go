package graph
import (
    "strings"
    "errors"
)
// 
type ITunesPairingMode int

const (
    DISALLOW_ITUNESPAIRINGMODE ITunesPairingMode = iota
    ALLOW_ITUNESPAIRINGMODE
    REQUIRESCERTIFICATE_ITUNESPAIRINGMODE
)

func (i ITunesPairingMode) String() string {
    return []string{"DISALLOW", "ALLOW", "REQUIRESCERTIFICATE"}[i]
}
func ParseITunesPairingMode(v string) (interface{}, error) {
    result := DISALLOW_ITUNESPAIRINGMODE
    switch strings.ToUpper(v) {
        case "DISALLOW":
            result = DISALLOW_ITUNESPAIRINGMODE
        case "ALLOW":
            result = ALLOW_ITUNESPAIRINGMODE
        case "REQUIRESCERTIFICATE":
            result = REQUIRESCERTIFICATE_ITUNESPAIRINGMODE
        default:
            return 0, errors.New("Unknown ITunesPairingMode value: " + v)
    }
    return &result, nil
}
func SerializeITunesPairingMode(values []ITunesPairingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
