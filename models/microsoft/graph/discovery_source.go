package graph
import (
    "strings"
    "errors"
)
// 
type DiscoverySource int

const (
    UNKNOWN_DISCOVERYSOURCE DiscoverySource = iota
    ADMINIMPORT_DISCOVERYSOURCE
    DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE
)

func (i DiscoverySource) String() string {
    return []string{"UNKNOWN", "ADMINIMPORT", "DEVICEENROLLMENTPROGRAM"}[i]
}
func ParseDiscoverySource(v string) (interface{}, error) {
    result := UNKNOWN_DISCOVERYSOURCE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DISCOVERYSOURCE
        case "ADMINIMPORT":
            result = ADMINIMPORT_DISCOVERYSOURCE
        case "DEVICEENROLLMENTPROGRAM":
            result = DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE
        default:
            return 0, errors.New("Unknown DiscoverySource value: " + v)
    }
    return &result, nil
}
func SerializeDiscoverySource(values []DiscoverySource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
