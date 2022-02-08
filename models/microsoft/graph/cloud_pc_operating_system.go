package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcOperatingSystem int

const (
    WINDOWS10_CLOUDPCOPERATINGSYSTEM CloudPcOperatingSystem = iota
    WINDOWS11_CLOUDPCOPERATINGSYSTEM
    UNKNOWNFUTUREVALUE_CLOUDPCOPERATINGSYSTEM
)

func (i CloudPcOperatingSystem) String() string {
    return []string{"WINDOWS10", "WINDOWS11", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcOperatingSystem(v string) (interface{}, error) {
    result := WINDOWS10_CLOUDPCOPERATINGSYSTEM
    switch strings.ToUpper(v) {
        case "WINDOWS10":
            result = WINDOWS10_CLOUDPCOPERATINGSYSTEM
        case "WINDOWS11":
            result = WINDOWS11_CLOUDPCOPERATINGSYSTEM
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCOPERATINGSYSTEM
        default:
            return 0, errors.New("Unknown CloudPcOperatingSystem value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOperatingSystem(values []CloudPcOperatingSystem) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
