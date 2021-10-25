package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CHASSISTYPE, nil
        case "DESKTOP":
            return DESKTOP_CHASSISTYPE, nil
        case "LAPTOP":
            return LAPTOP_CHASSISTYPE, nil
        case "WORKSWORKSTATION":
            return WORKSWORKSTATION_CHASSISTYPE, nil
        case "ENTERPRISESERVER":
            return ENTERPRISESERVER_CHASSISTYPE, nil
        case "PHONE":
            return PHONE_CHASSISTYPE, nil
        case "TABLET":
            return TABLET_CHASSISTYPE, nil
        case "MOBILEOTHER":
            return MOBILEOTHER_CHASSISTYPE, nil
        case "MOBILEUNKNOWN":
            return MOBILEUNKNOWN_CHASSISTYPE, nil
    }
    return 0, errors.New("Unknown ChassisType value: " + v)
}
func SerializeChassisType(values []ChassisType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
