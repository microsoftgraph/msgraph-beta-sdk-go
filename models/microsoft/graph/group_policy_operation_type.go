package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyOperationType int

const (
    NONE_GROUPPOLICYOPERATIONTYPE GroupPolicyOperationType = iota
    UPLOAD_GROUPPOLICYOPERATIONTYPE
    UPLOADNEWVERSION_GROUPPOLICYOPERATIONTYPE
    ADDLANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    REMOVELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    UPDATELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
    REMOVE_GROUPPOLICYOPERATIONTYPE
)

func (i GroupPolicyOperationType) String() string {
    return []string{"NONE", "UPLOAD", "UPLOADNEWVERSION", "ADDLANGUAGEFILES", "REMOVELANGUAGEFILES", "UPDATELANGUAGEFILES", "REMOVE"}[i]
}
func ParseGroupPolicyOperationType(v string) (interface{}, error) {
    result := NONE_GROUPPOLICYOPERATIONTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_GROUPPOLICYOPERATIONTYPE
        case "UPLOAD":
            result = UPLOAD_GROUPPOLICYOPERATIONTYPE
        case "UPLOADNEWVERSION":
            result = UPLOADNEWVERSION_GROUPPOLICYOPERATIONTYPE
        case "ADDLANGUAGEFILES":
            result = ADDLANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "REMOVELANGUAGEFILES":
            result = REMOVELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "UPDATELANGUAGEFILES":
            result = UPDATELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE
        case "REMOVE":
            result = REMOVE_GROUPPOLICYOPERATIONTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyOperationType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyOperationType(values []GroupPolicyOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
