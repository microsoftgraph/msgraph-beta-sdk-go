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
    switch strings.ToUpper(v) {
        case "DISALLOW":
            return DISALLOW_ITUNESPAIRINGMODE, nil
        case "ALLOW":
            return ALLOW_ITUNESPAIRINGMODE, nil
        case "REQUIRESCERTIFICATE":
            return REQUIRESCERTIFICATE_ITUNESPAIRINGMODE, nil
    }
    return 0, errors.New("Unknown ITunesPairingMode value: " + v)
}
func SerializeITunesPairingMode(values []ITunesPairingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
