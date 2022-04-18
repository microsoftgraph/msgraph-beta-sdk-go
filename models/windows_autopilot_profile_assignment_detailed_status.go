package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WindowsAutopilotProfileAssignmentDetailedStatus int

const (
    NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS WindowsAutopilotProfileAssignmentDetailedStatus = iota
    HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
)

func (i WindowsAutopilotProfileAssignmentDetailedStatus) String() string {
    return []string{"NONE", "HARDWAREREQUIREMENTSNOTMET", "SURFACEHUBPROFILENOTSUPPORTED", "HOLOLENSPROFILENOTSUPPORTED", "WINDOWSPCPROFILENOTSUPPORTED", "SURFACEHUB2SPROFILENOTSUPPORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWindowsAutopilotProfileAssignmentDetailedStatus(v string) (interface{}, error) {
    result := NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "HARDWAREREQUIREMENTSNOTMET":
            result = HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "SURFACEHUBPROFILENOTSUPPORTED":
            result = SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "HOLOLENSPROFILENOTSUPPORTED":
            result = HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "WINDOWSPCPROFILENOTSUPPORTED":
            result = WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "SURFACEHUB2SPROFILENOTSUPPORTED":
            result = SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentDetailedStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotProfileAssignmentDetailedStatus(values []WindowsAutopilotProfileAssignmentDetailedStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}