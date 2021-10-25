package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_OBJECTFLOWTYPES, nil
        case "ADD":
            return ADD_OBJECTFLOWTYPES, nil
        case "UPDATE":
            return UPDATE_OBJECTFLOWTYPES, nil
        case "DELETE":
            return DELETE_OBJECTFLOWTYPES, nil
    }
    return 0, errors.New("Unknown ObjectFlowTypes value: " + v)
}
func SerializeObjectFlowTypes(values []ObjectFlowTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
