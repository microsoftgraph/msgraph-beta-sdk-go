package models
// Indicates whether the target of a relationship is the parent or the child in the relationship.
type MobileAppRelationshipType int

const (
    // Indicates that the target of a relationship is the child in the relationship.
    CHILD_MOBILEAPPRELATIONSHIPTYPE MobileAppRelationshipType = iota
    // Indicates that the target of a relationship is the parent in the relationship.
    PARENT_MOBILEAPPRELATIONSHIPTYPE
)

func (i MobileAppRelationshipType) String() string {
    return []string{"child", "parent"}[i]
}
func ParseMobileAppRelationshipType(v string) (any, error) {
    result := CHILD_MOBILEAPPRELATIONSHIPTYPE
    switch v {
        case "child":
            result = CHILD_MOBILEAPPRELATIONSHIPTYPE
        case "parent":
            result = PARENT_MOBILEAPPRELATIONSHIPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMobileAppRelationshipType(values []MobileAppRelationshipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MobileAppRelationshipType) isMultiValue() bool {
    return false
}
