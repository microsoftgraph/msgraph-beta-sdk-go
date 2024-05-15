package models
// Owner type of device.
type OwnerType int

const (
    // Unknown.
    UNKNOWN_OWNERTYPE OwnerType = iota
    // Owned by company.
    COMPANY_OWNERTYPE
    // Owned by person.
    PERSONAL_OWNERTYPE
)

func (i OwnerType) String() string {
    return []string{"unknown", "company", "personal"}[i]
}
func ParseOwnerType(v string) (any, error) {
    result := UNKNOWN_OWNERTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_OWNERTYPE
        case "company":
            result = COMPANY_OWNERTYPE
        case "personal":
            result = PERSONAL_OWNERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOwnerType(values []OwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnerType) isMultiValue() bool {
    return false
}
