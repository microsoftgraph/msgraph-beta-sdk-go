package models
// Userless enrollment block status, indicating whether the next device enrollment will be blocked.
type WindowsAutopilotUserlessEnrollmentStatus int

const (
    // Unknown userless enrollment block status. Next userless enrollment may fail. This is the default value.
    UNKNOWN_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS WindowsAutopilotUserlessEnrollmentStatus = iota
    // Indicates next userless enrollment can proceed.
    ALLOWED_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
    // Indicates next userless enrollment cannot proceed without resetting the windowsAutopilotUserlessEnrollmentStatus.
    BLOCKED_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
)

func (i WindowsAutopilotUserlessEnrollmentStatus) String() string {
    return []string{"unknown", "allowed", "blocked", "unknownFutureValue"}[i]
}
func ParseWindowsAutopilotUserlessEnrollmentStatus(v string) (any, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
        case "allowed":
            result = ALLOWED_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
        case "blocked":
            result = BLOCKED_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTUSERLESSENROLLMENTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsAutopilotUserlessEnrollmentStatus(values []WindowsAutopilotUserlessEnrollmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsAutopilotUserlessEnrollmentStatus) isMultiValue() bool {
    return false
}
