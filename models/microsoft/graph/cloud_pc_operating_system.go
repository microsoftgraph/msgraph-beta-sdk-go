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
    switch strings.ToUpper(v) {
        case "WINDOWS10":
            return WINDOWS10_CLOUDPCOPERATINGSYSTEM, nil
        case "WINDOWS11":
            return WINDOWS11_CLOUDPCOPERATINGSYSTEM, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCOPERATINGSYSTEM, nil
    }
    return 0, errors.New("Unknown CloudPcOperatingSystem value: " + v)
}
func SerializeCloudPcOperatingSystem(values []CloudPcOperatingSystem) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
