package graph
import (
    "strings"
    "errors"
)
// 
type ResultantAppState int

const (
    NOTAPPLICABLE_RESULTANTAPPSTATE ResultantAppState = iota
    INSTALLED_RESULTANTAPPSTATE
    FAILED_RESULTANTAPPSTATE
    NOTINSTALLED_RESULTANTAPPSTATE
    UNINSTALLFAILED_RESULTANTAPPSTATE
    PENDINGINSTALL_RESULTANTAPPSTATE
    UNKNOWN_RESULTANTAPPSTATE
)

func (i ResultantAppState) String() string {
    return []string{"NOTAPPLICABLE", "INSTALLED", "FAILED", "NOTINSTALLED", "UNINSTALLFAILED", "PENDINGINSTALL", "UNKNOWN"}[i]
}
func ParseResultantAppState(v string) (interface{}, error) {
    result := NOTAPPLICABLE_RESULTANTAPPSTATE
    switch strings.ToUpper(v) {
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_RESULTANTAPPSTATE
        case "INSTALLED":
            result = INSTALLED_RESULTANTAPPSTATE
        case "FAILED":
            result = FAILED_RESULTANTAPPSTATE
        case "NOTINSTALLED":
            result = NOTINSTALLED_RESULTANTAPPSTATE
        case "UNINSTALLFAILED":
            result = UNINSTALLFAILED_RESULTANTAPPSTATE
        case "PENDINGINSTALL":
            result = PENDINGINSTALL_RESULTANTAPPSTATE
        case "UNKNOWN":
            result = UNKNOWN_RESULTANTAPPSTATE
        default:
            return 0, errors.New("Unknown ResultantAppState value: " + v)
    }
    return &result, nil
}
func SerializeResultantAppState(values []ResultantAppState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
