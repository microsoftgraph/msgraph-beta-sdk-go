package graph
import (
    "strings"
    "errors"
)
// 
type DriveItemSourceApplication int

const (
    TEAMS_DRIVEITEMSOURCEAPPLICATION DriveItemSourceApplication = iota
    YAMMER_DRIVEITEMSOURCEAPPLICATION
    SHAREPOINT_DRIVEITEMSOURCEAPPLICATION
    ONEDRIVE_DRIVEITEMSOURCEAPPLICATION
    STREAM_DRIVEITEMSOURCEAPPLICATION
    POWERPOINT_DRIVEITEMSOURCEAPPLICATION
    OFFICE_DRIVEITEMSOURCEAPPLICATION
    UNKNOWNFUTUREVALUE_DRIVEITEMSOURCEAPPLICATION
)

func (i DriveItemSourceApplication) String() string {
    return []string{"TEAMS", "YAMMER", "SHAREPOINT", "ONEDRIVE", "STREAM", "POWERPOINT", "OFFICE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDriveItemSourceApplication(v string) (interface{}, error) {
    result := TEAMS_DRIVEITEMSOURCEAPPLICATION
    switch strings.ToUpper(v) {
        case "TEAMS":
            result = TEAMS_DRIVEITEMSOURCEAPPLICATION
        case "YAMMER":
            result = YAMMER_DRIVEITEMSOURCEAPPLICATION
        case "SHAREPOINT":
            result = SHAREPOINT_DRIVEITEMSOURCEAPPLICATION
        case "ONEDRIVE":
            result = ONEDRIVE_DRIVEITEMSOURCEAPPLICATION
        case "STREAM":
            result = STREAM_DRIVEITEMSOURCEAPPLICATION
        case "POWERPOINT":
            result = POWERPOINT_DRIVEITEMSOURCEAPPLICATION
        case "OFFICE":
            result = OFFICE_DRIVEITEMSOURCEAPPLICATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DRIVEITEMSOURCEAPPLICATION
        default:
            return 0, errors.New("Unknown DriveItemSourceApplication value: " + v)
    }
    return &result, nil
}
func SerializeDriveItemSourceApplication(values []DriveItemSourceApplication) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
