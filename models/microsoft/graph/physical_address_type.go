package graph
import (
    "strings"
    "errors"
)
// 
type PhysicalAddressType int

const (
    UNKNOWN_PHYSICALADDRESSTYPE PhysicalAddressType = iota
    HOME_PHYSICALADDRESSTYPE
    BUSINESS_PHYSICALADDRESSTYPE
    OTHER_PHYSICALADDRESSTYPE
)

func (i PhysicalAddressType) String() string {
    return []string{"UNKNOWN", "HOME", "BUSINESS", "OTHER"}[i]
}
func ParsePhysicalAddressType(v string) (interface{}, error) {
    result := UNKNOWN_PHYSICALADDRESSTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PHYSICALADDRESSTYPE
        case "HOME":
            result = HOME_PHYSICALADDRESSTYPE
        case "BUSINESS":
            result = BUSINESS_PHYSICALADDRESSTYPE
        case "OTHER":
            result = OTHER_PHYSICALADDRESSTYPE
        default:
            return 0, errors.New("Unknown PhysicalAddressType value: " + v)
    }
    return &result, nil
}
func SerializePhysicalAddressType(values []PhysicalAddressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
