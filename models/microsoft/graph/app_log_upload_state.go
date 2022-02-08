package graph
import (
    "strings"
    "errors"
)
// 
type AppLogUploadState int

const (
    PENDING_APPLOGUPLOADSTATE AppLogUploadState = iota
    COMPLETED_APPLOGUPLOADSTATE
    FAILED_APPLOGUPLOADSTATE
)

func (i AppLogUploadState) String() string {
    return []string{"PENDING", "COMPLETED", "FAILED"}[i]
}
func ParseAppLogUploadState(v string) (interface{}, error) {
    result := PENDING_APPLOGUPLOADSTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_APPLOGUPLOADSTATE
        case "COMPLETED":
            result = COMPLETED_APPLOGUPLOADSTATE
        case "FAILED":
            result = FAILED_APPLOGUPLOADSTATE
        default:
            return 0, errors.New("Unknown AppLogUploadState value: " + v)
    }
    return &result, nil
}
func SerializeAppLogUploadState(values []AppLogUploadState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
