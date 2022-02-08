package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcSnapshotStatus int

const (
    READY_CLOUDPCSNAPSHOTSTATUS CloudPcSnapshotStatus = iota
    UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTSTATUS
)

func (i CloudPcSnapshotStatus) String() string {
    return []string{"READY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcSnapshotStatus(v string) (interface{}, error) {
    result := READY_CLOUDPCSNAPSHOTSTATUS
    switch strings.ToUpper(v) {
        case "READY":
            result = READY_CLOUDPCSNAPSHOTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTSTATUS
        default:
            return 0, errors.New("Unknown CloudPcSnapshotStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcSnapshotStatus(values []CloudPcSnapshotStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
