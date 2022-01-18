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
    switch strings.ToUpper(v) {
        case "READY":
            return READY_CLOUDPCSNAPSHOTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTSTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcSnapshotStatus value: " + v)
}
func SerializeCloudPcSnapshotStatus(values []CloudPcSnapshotStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
