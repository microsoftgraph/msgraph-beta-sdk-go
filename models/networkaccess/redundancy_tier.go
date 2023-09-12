package networkaccess
import (
    "errors"
)
// 
type RedundancyTier int

const (
    NOREDUNDANCY_REDUNDANCYTIER RedundancyTier = iota
    ZONEREDUNDANCY_REDUNDANCYTIER
    UNKNOWNFUTUREVALUE_REDUNDANCYTIER
)

func (i RedundancyTier) String() string {
    return []string{"noRedundancy", "zoneRedundancy", "unknownFutureValue"}[i]
}
func ParseRedundancyTier(v string) (any, error) {
    result := NOREDUNDANCY_REDUNDANCYTIER
    switch v {
        case "noRedundancy":
            result = NOREDUNDANCY_REDUNDANCYTIER
        case "zoneRedundancy":
            result = ZONEREDUNDANCY_REDUNDANCYTIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REDUNDANCYTIER
        default:
            return 0, errors.New("Unknown RedundancyTier value: " + v)
    }
    return &result, nil
}
func SerializeRedundancyTier(values []RedundancyTier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RedundancyTier) isMultiValue() bool {
    return false
}
