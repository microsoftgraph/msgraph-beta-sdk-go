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
    switch strings.ToUpper(v) {
        case "TEAMS":
            return TEAMS_DRIVEITEMSOURCEAPPLICATION, nil
        case "YAMMER":
            return YAMMER_DRIVEITEMSOURCEAPPLICATION, nil
        case "SHAREPOINT":
            return SHAREPOINT_DRIVEITEMSOURCEAPPLICATION, nil
        case "ONEDRIVE":
            return ONEDRIVE_DRIVEITEMSOURCEAPPLICATION, nil
        case "STREAM":
            return STREAM_DRIVEITEMSOURCEAPPLICATION, nil
        case "POWERPOINT":
            return POWERPOINT_DRIVEITEMSOURCEAPPLICATION, nil
        case "OFFICE":
            return OFFICE_DRIVEITEMSOURCEAPPLICATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DRIVEITEMSOURCEAPPLICATION, nil
    }
    return 0, errors.New("Unknown DriveItemSourceApplication value: " + v)
}
func SerializeDriveItemSourceApplication(values []DriveItemSourceApplication) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
