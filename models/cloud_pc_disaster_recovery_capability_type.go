package models
import (
    "errors"
)
type CloudPcDisasterRecoveryCapabilityType int

const (
    NONE_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE CloudPcDisasterRecoveryCapabilityType = iota
    FAILOVER_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
    FAILBACK_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
)

func (i CloudPcDisasterRecoveryCapabilityType) String() string {
    return []string{"none", "failover", "failback", "unknownFutureValue"}[i]
}
func ParseCloudPcDisasterRecoveryCapabilityType(v string) (any, error) {
    result := NONE_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
    switch v {
        case "none":
            result = NONE_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
        case "failover":
            result = FAILOVER_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
        case "failback":
            result = FAILBACK_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYCAPABILITYTYPE
        default:
            return 0, errors.New("Unknown CloudPcDisasterRecoveryCapabilityType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDisasterRecoveryCapabilityType(values []CloudPcDisasterRecoveryCapabilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcDisasterRecoveryCapabilityType) isMultiValue() bool {
    return false
}
