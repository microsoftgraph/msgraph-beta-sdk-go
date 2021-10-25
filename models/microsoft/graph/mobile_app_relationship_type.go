package graph
import (
    "strings"
    "errors"
)
type MobileAppRelationshipType int

const (
    CHILD_MOBILEAPPRELATIONSHIPTYPE MobileAppRelationshipType = iota
    PARENT_MOBILEAPPRELATIONSHIPTYPE
)

func (i MobileAppRelationshipType) String() string {
    return []string{"CHILD", "PARENT"}[i]
}
func ParseMobileAppRelationshipType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CHILD":
            return CHILD_MOBILEAPPRELATIONSHIPTYPE, nil
        case "PARENT":
            return PARENT_MOBILEAPPRELATIONSHIPTYPE, nil
    }
    return 0, errors.New("Unknown MobileAppRelationshipType value: " + v)
}
func SerializeMobileAppRelationshipType(values []MobileAppRelationshipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
