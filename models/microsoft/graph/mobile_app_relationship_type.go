package graph
import (
    "strings"
    "errors"
)
// 
type MobileAppRelationshipType int

const (
    CHILD_MOBILEAPPRELATIONSHIPTYPE MobileAppRelationshipType = iota
    PARENT_MOBILEAPPRELATIONSHIPTYPE
)

func (i MobileAppRelationshipType) String() string {
    return []string{"CHILD", "PARENT"}[i]
}
func ParseMobileAppRelationshipType(v string) (interface{}, error) {
    result := CHILD_MOBILEAPPRELATIONSHIPTYPE
    switch strings.ToUpper(v) {
        case "CHILD":
            result = CHILD_MOBILEAPPRELATIONSHIPTYPE
        case "PARENT":
            result = PARENT_MOBILEAPPRELATIONSHIPTYPE
        default:
            return 0, errors.New("Unknown MobileAppRelationshipType value: " + v)
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
