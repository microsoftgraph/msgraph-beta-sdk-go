package models
// Windows quality update category
type WindowsQualityUpdateCategory int

const (
    // All update type
    ALL_WINDOWSQUALITYUPDATECATEGORY WindowsQualityUpdateCategory = iota
    // Security only update type
    SECURITY_WINDOWSQUALITYUPDATECATEGORY
    // Non security only update type
    NONSECURITY_WINDOWSQUALITYUPDATECATEGORY
)

func (i WindowsQualityUpdateCategory) String() string {
    return []string{"all", "security", "nonSecurity"}[i]
}
func ParseWindowsQualityUpdateCategory(v string) (any, error) {
    result := ALL_WINDOWSQUALITYUPDATECATEGORY
    switch v {
        case "all":
            result = ALL_WINDOWSQUALITYUPDATECATEGORY
        case "security":
            result = SECURITY_WINDOWSQUALITYUPDATECATEGORY
        case "nonSecurity":
            result = NONSECURITY_WINDOWSQUALITYUPDATECATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsQualityUpdateCategory(values []WindowsQualityUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsQualityUpdateCategory) isMultiValue() bool {
    return false
}
