package models
type PhysicalAddressType int

const (
    UNKNOWN_PHYSICALADDRESSTYPE PhysicalAddressType = iota
    HOME_PHYSICALADDRESSTYPE
    BUSINESS_PHYSICALADDRESSTYPE
    OTHER_PHYSICALADDRESSTYPE
)

func (i PhysicalAddressType) String() string {
    return []string{"unknown", "home", "business", "other"}[i]
}
func ParsePhysicalAddressType(v string) (any, error) {
    result := UNKNOWN_PHYSICALADDRESSTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_PHYSICALADDRESSTYPE
        case "home":
            result = HOME_PHYSICALADDRESSTYPE
        case "business":
            result = BUSINESS_PHYSICALADDRESSTYPE
        case "other":
            result = OTHER_PHYSICALADDRESSTYPE
        default:
            return nil, nil
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
func (i PhysicalAddressType) isMultiValue() bool {
    return false
}
