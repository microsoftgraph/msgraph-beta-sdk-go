package models
// Update schedule type for macOS software updates.
type MacOSSoftwareUpdateScheduleType int

const (
    // Always update.
    ALWAYSUPDATE_MACOSSOFTWAREUPDATESCHEDULETYPE MacOSSoftwareUpdateScheduleType = iota
    // Update during time windows.
    UPDATEDURINGTIMEWINDOWS_MACOSSOFTWAREUPDATESCHEDULETYPE
    // Update outside of time windows.
    UPDATEOUTSIDEOFTIMEWINDOWS_MACOSSOFTWAREUPDATESCHEDULETYPE
)

func (i MacOSSoftwareUpdateScheduleType) String() string {
    return []string{"alwaysUpdate", "updateDuringTimeWindows", "updateOutsideOfTimeWindows"}[i]
}
func ParseMacOSSoftwareUpdateScheduleType(v string) (any, error) {
    result := ALWAYSUPDATE_MACOSSOFTWAREUPDATESCHEDULETYPE
    switch v {
        case "alwaysUpdate":
            result = ALWAYSUPDATE_MACOSSOFTWAREUPDATESCHEDULETYPE
        case "updateDuringTimeWindows":
            result = UPDATEDURINGTIMEWINDOWS_MACOSSOFTWAREUPDATESCHEDULETYPE
        case "updateOutsideOfTimeWindows":
            result = UPDATEOUTSIDEOFTIMEWINDOWS_MACOSSOFTWAREUPDATESCHEDULETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateScheduleType(values []MacOSSoftwareUpdateScheduleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSSoftwareUpdateScheduleType) isMultiValue() bool {
    return false
}
