package models
import (
    "errors"
)
// 
type CloudPcPowerState int

const (
    RUNNING_CLOUDPCPOWERSTATE CloudPcPowerState = iota
    POWEREDOFF_CLOUDPCPOWERSTATE
    UNKNOWNFUTUREVALUE_CLOUDPCPOWERSTATE
)

func (i CloudPcPowerState) String() string {
    return []string{"running", "poweredOff", "unknownFutureValue"}[i]
}
func ParseCloudPcPowerState(v string) (any, error) {
    result := RUNNING_CLOUDPCPOWERSTATE
    switch v {
        case "running":
            result = RUNNING_CLOUDPCPOWERSTATE
        case "poweredOff":
            result = POWEREDOFF_CLOUDPCPOWERSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPOWERSTATE
        default:
            return 0, errors.New("Unknown CloudPcPowerState value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcPowerState(values []CloudPcPowerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcPowerState) isMultiValue() bool {
    return false
}
