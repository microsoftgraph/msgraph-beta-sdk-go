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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_GROUPPOLICYOPERATIONTYPE, nil
        case "UPLOAD":
            return UPLOAD_GROUPPOLICYOPERATIONTYPE, nil
        case "UPLOADNEWVERSION":
            return UPLOADNEWVERSION_GROUPPOLICYOPERATIONTYPE, nil
        case "ADDLANGUAGEFILES":
            return ADDLANGUAGEFILES_GROUPPOLICYOPERATIONTYPE, nil
        case "REMOVELANGUAGEFILES":
            return REMOVELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE, nil
        case "UPDATELANGUAGEFILES":
            return UPDATELANGUAGEFILES_GROUPPOLICYOPERATIONTYPE, nil
        case "REMOVE":
            return REMOVE_GROUPPOLICYOPERATIONTYPE, nil
    }
    return 0, errors.New("Unknown GroupPolicyOperationType value: " + v)
}
func SerializeGroupPolicyOperationType(values []GroupPolicyOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
