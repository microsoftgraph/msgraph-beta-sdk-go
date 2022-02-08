package graph
import (
    "strings"
    "errors"
)
// 
type ChassisType int

const (
    UNKNOWN_CHASSISTYPE ChassisType = iota
    DESKTOP_CHASSISTYPE
    LAPTOP_CHASSISTYPE
    WORKSWORKSTATION_CHASSISTYPE
    ENTERPRISESERVER_CHASSISTYPE
    PHONE_CHASSISTYPE
    TABLET_CHASSISTYPE
    MOBILEOTHER_CHASSISTYPE
    MOBILEUNKNOWN_CHASSISTYPE
)

func (i ChassisType) String() string {
    return []string{"UNKNOWN", "DESKTOP", "LAPTOP", "WORKSWORKSTATION", "ENTERPRISESERVER", "PHONE", "TABLET", "MOBILEOTHER", "MOBILEUNKNOWN"}[i]
}
func ParseChassisType(v string) (interface{}, error) {
    result := UNKNOWN_CHASSISTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CHASSISTYPE
        case "DESKTOP":
            result = DESKTOP_CHASSISTYPE
        case "LAPTOP":
            result = LAPTOP_CHASSISTYPE
        case "WORKSWORKSTATION":
            result = WORKSWORKSTATION_CHASSISTYPE
        case "ENTERPRISESERVER":
            result = ENTERPRISESERVER_CHASSISTYPE
        case "PHONE":
            result = PHONE_CHASSISTYPE
        case "TABLET":
            result = TABLET_CHASSISTYPE
        case "MOBILEOTHER":
            result = MOBILEOTHER_CHASSISTYPE
        case "MOBILEUNKNOWN":
            result = MOBILEUNKNOWN_CHASSISTYPE
        default:
            return 0, errors.New("Unknown ChassisType value: " + v)
    }
    return &result, nil
}
func SerializeChassisType(values []ChassisType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
