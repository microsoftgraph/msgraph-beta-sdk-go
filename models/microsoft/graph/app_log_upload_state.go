package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_APPLOGUPLOADSTATE, nil
        case "COMPLETED":
            return COMPLETED_APPLOGUPLOADSTATE, nil
        case "FAILED":
            return FAILED_APPLOGUPLOADSTATE, nil
    }
    return 0, errors.New("Unknown AppLogUploadState value: " + v)
}
func SerializeAppLogUploadState(values []AppLogUploadState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
