package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DiskType int

const (
    UNKNOWN_DISKTYPE DiskType = iota
    HDD_DISKTYPE
    SSD_DISKTYPE
    UNKNOWNFUTUREVALUE_DISKTYPE
)

func (i DiskType) String() string {
    return []string{"UNKNOWN", "HDD", "SSD", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDiskType(v string) (interface{}, error) {
    result := UNKNOWN_DISKTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DISKTYPE
        case "HDD":
            result = HDD_DISKTYPE
        case "SSD":
            result = SSD_DISKTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DISKTYPE
        default:
            return 0, errors.New("Unknown DiskType value: " + v)
    }
    return &result, nil
}
func SerializeDiskType(values []DiskType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
