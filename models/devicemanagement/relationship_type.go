package devicemanagement
type RelationshipType int

const (
    AND_RELATIONSHIPTYPE RelationshipType = iota
    OR_RELATIONSHIPTYPE
    UNKNOWNFUTUREVALUE_RELATIONSHIPTYPE
)

func (i RelationshipType) String() string {
    return []string{"and", "or", "unknownFutureValue"}[i]
}
func ParseRelationshipType(v string) (any, error) {
    result := AND_RELATIONSHIPTYPE
    switch v {
        case "and":
            result = AND_RELATIONSHIPTYPE
        case "or":
            result = OR_RELATIONSHIPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RELATIONSHIPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRelationshipType(values []RelationshipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RelationshipType) isMultiValue() bool {
    return false
}
