package models
import (
    "errors"
)
// 
type CloudPcSnapshotStatus int

const (
    READY_CLOUDPCSNAPSHOTSTATUS CloudPcSnapshotStatus = iota
    UNKNOWNFUTUREVALUE_CLOUDPCSNAPSHOTSTATUS
)

func (i CloudPcSnapshotStatus) String() string {
    return []string{"ready", "unknownFutureValue"}[i]
}
func ParseCloudPcSnapshotStatus(v string) (any, error) {
    result := READY_CLOUDPCSNAPSHOTSTATUS
    switch v {
        case "ready":
            result = READY_CLOUDPCSNAPSHOTSTATUS
        case "unknownFutureValue":
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
func (i CloudPcSnapshotStatus) isMultiValue() bool {
    return false
}
