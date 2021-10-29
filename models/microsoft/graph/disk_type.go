package graph
import (
    "strings"
    "errors"
)
// 
type DiskType int

const (
    UNKOWN_DISKTYPE DiskType = iota
    HDD_DISKTYPE
    SSD_DISKTYPE
)

func (i DiskType) String() string {
    return []string{"UNKOWN", "HDD", "SSD"}[i]
}
func ParseDiskType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKOWN":
            return UNKOWN_DISKTYPE, nil
        case "HDD":
            return HDD_DISKTYPE, nil
        case "SSD":
            return SSD_DISKTYPE, nil
    }
    return 0, errors.New("Unknown DiskType value: " + v)
}
func SerializeDiskType(values []DiskType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
