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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PHYSICALADDRESSTYPE, nil
        case "HOME":
            return HOME_PHYSICALADDRESSTYPE, nil
        case "BUSINESS":
            return BUSINESS_PHYSICALADDRESSTYPE, nil
        case "OTHER":
            return OTHER_PHYSICALADDRESSTYPE, nil
    }
    return 0, errors.New("Unknown PhysicalAddressType value: " + v)
}
func SerializePhysicalAddressType(values []PhysicalAddressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
