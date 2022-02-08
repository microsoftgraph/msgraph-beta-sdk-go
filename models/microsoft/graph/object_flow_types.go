package graph
import (
    "strings"
    "errors"
)
// 
type ObjectFlowTypes int

const (
    NONE_OBJECTFLOWTYPES ObjectFlowTypes = iota
    ADD_OBJECTFLOWTYPES
    UPDATE_OBJECTFLOWTYPES
    DELETE_OBJECTFLOWTYPES
)

func (i ObjectFlowTypes) String() string {
    return []string{"NONE", "ADD", "UPDATE", "DELETE"}[i]
}
func ParseObjectFlowTypes(v string) (interface{}, error) {
    result := NONE_OBJECTFLOWTYPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_OBJECTFLOWTYPES
        case "ADD":
            result = ADD_OBJECTFLOWTYPES
        case "UPDATE":
            result = UPDATE_OBJECTFLOWTYPES
        case "DELETE":
            result = DELETE_OBJECTFLOWTYPES
        default:
            return 0, errors.New("Unknown ObjectFlowTypes value: " + v)
    }
    return &result, nil
}
func SerializeObjectFlowTypes(values []ObjectFlowTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
