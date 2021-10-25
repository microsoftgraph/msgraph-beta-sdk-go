package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_RESULTANTAPPSTATE, nil
        case "INSTALLED":
            return INSTALLED_RESULTANTAPPSTATE, nil
        case "FAILED":
            return FAILED_RESULTANTAPPSTATE, nil
        case "NOTINSTALLED":
            return NOTINSTALLED_RESULTANTAPPSTATE, nil
        case "UNINSTALLFAILED":
            return UNINSTALLFAILED_RESULTANTAPPSTATE, nil
        case "PENDINGINSTALL":
            return PENDINGINSTALL_RESULTANTAPPSTATE, nil
        case "UNKNOWN":
            return UNKNOWN_RESULTANTAPPSTATE, nil
    }
    return 0, errors.New("Unknown ResultantAppState value: " + v)
}
func SerializeResultantAppState(values []ResultantAppState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
