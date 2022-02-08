package graph
import (
    "strings"
    "errors"
)
// 
type DiskType int

const (
    HDD_DISKTYPE DiskType = iota
    SSD_DISKTYPE
    UNKNOWN_DISKTYPE
)

func (i DiskType) String() string {
    return []string{"HDD", "SSD", "UNKNOWN"}[i]
}
func ParseDiskType(v string) (interface{}, error) {
    result := HDD_DISKTYPE
    switch strings.ToUpper(v) {
        case "HDD":
            result = HDD_DISKTYPE
        case "SSD":
            result = SSD_DISKTYPE
        case "UNKNOWN":
            result = UNKNOWN_DISKTYPE
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
