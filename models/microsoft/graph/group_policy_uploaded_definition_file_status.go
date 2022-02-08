package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyUploadedDefinitionFileStatus int

const (
    NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS GroupPolicyUploadedDefinitionFileStatus = iota
    UPLOADINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    AVAILABLE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    ASSIGNED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    REMOVALINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    UPLOADFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    REMOVALFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
)

func (i GroupPolicyUploadedDefinitionFileStatus) String() string {
    return []string{"NONE", "UPLOADINPROGRESS", "AVAILABLE", "ASSIGNED", "REMOVALINPROGRESS", "UPLOADFAILED", "REMOVALFAILED"}[i]
}
func ParseGroupPolicyUploadedDefinitionFileStatus(v string) (interface{}, error) {
    result := NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "UPLOADINPROGRESS":
            result = UPLOADINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "AVAILABLE":
            result = AVAILABLE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "ASSIGNED":
            result = ASSIGNED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "REMOVALINPROGRESS":
            result = REMOVALINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "UPLOADFAILED":
            result = UPLOADFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        case "REMOVALFAILED":
            result = REMOVALFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS
        default:
            return 0, errors.New("Unknown GroupPolicyUploadedDefinitionFileStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyUploadedDefinitionFileStatus(values []GroupPolicyUploadedDefinitionFileStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
