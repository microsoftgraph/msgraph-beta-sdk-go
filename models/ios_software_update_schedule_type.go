package models
// Update schedule type for iOS software updates.
type IosSoftwareUpdateScheduleType int

const (
    // Update outside of active hours.
    UPDATEOUTSIDEOFACTIVEHOURS_IOSSOFTWAREUPDATESCHEDULETYPE IosSoftwareUpdateScheduleType = iota
    // Always update.
    ALWAYSUPDATE_IOSSOFTWAREUPDATESCHEDULETYPE
    // Update during time windows.
    UPDATEDURINGTIMEWINDOWS_IOSSOFTWAREUPDATESCHEDULETYPE
    // Update outside of time windows.
    UPDATEOUTSIDEOFTIMEWINDOWS_IOSSOFTWAREUPDATESCHEDULETYPE
)

func (i IosSoftwareUpdateScheduleType) String() string {
    return []string{"updateOutsideOfActiveHours", "alwaysUpdate", "updateDuringTimeWindows", "updateOutsideOfTimeWindows"}[i]
}
func ParseIosSoftwareUpdateScheduleType(v string) (any, error) {
    result := UPDATEOUTSIDEOFACTIVEHOURS_IOSSOFTWAREUPDATESCHEDULETYPE
    switch v {
        case "updateOutsideOfActiveHours":
            result = UPDATEOUTSIDEOFACTIVEHOURS_IOSSOFTWAREUPDATESCHEDULETYPE
        case "alwaysUpdate":
            result = ALWAYSUPDATE_IOSSOFTWAREUPDATESCHEDULETYPE
        case "updateDuringTimeWindows":
            result = UPDATEDURINGTIMEWINDOWS_IOSSOFTWAREUPDATESCHEDULETYPE
        case "updateOutsideOfTimeWindows":
            result = UPDATEOUTSIDEOFTIMEWINDOWS_IOSSOFTWAREUPDATESCHEDULETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIosSoftwareUpdateScheduleType(values []IosSoftwareUpdateScheduleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IosSoftwareUpdateScheduleType) isMultiValue() bool {
    return false
}
