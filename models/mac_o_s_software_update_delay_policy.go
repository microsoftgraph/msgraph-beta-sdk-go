package models
import (
    "errors"
    "strings"
)
// Flag enum to determine whether to delay software updates for macOS.
type MacOSSoftwareUpdateDelayPolicy int

const (
    // Software update delays will not be enforced.
    NONE_MACOSSOFTWAREUPDATEDELAYPOLICY MacOSSoftwareUpdateDelayPolicy = iota
    // Force delays for OS software updates.
    DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Force delays for app software updates.
    DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Force delays for major OS software updates.
    DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
)

func (i MacOSSoftwareUpdateDelayPolicy) String() string {
    var values []string
    for p := MacOSSoftwareUpdateDelayPolicy(1); p <= DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "delayOSUpdateVisibility", "delayAppUpdateVisibility", "unknownFutureValue", "delayMajorOsUpdateVisibility"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMacOSSoftwareUpdateDelayPolicy(v string) (any, error) {
    var result MacOSSoftwareUpdateDelayPolicy
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayOSUpdateVisibility":
                result |= DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayAppUpdateVisibility":
                result |= DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayMajorOsUpdateVisibility":
                result |= DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            default:
                return 0, errors.New("Unknown MacOSSoftwareUpdateDelayPolicy value: " + v)
        }
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateDelayPolicy(values []MacOSSoftwareUpdateDelayPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSSoftwareUpdateDelayPolicy) isMultiValue() bool {
    return true
}
