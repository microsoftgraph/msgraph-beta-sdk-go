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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "UPLOADINPROGRESS":
            return UPLOADINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "AVAILABLE":
            return AVAILABLE_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "ASSIGNED":
            return ASSIGNED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "REMOVALINPROGRESS":
            return REMOVALINPROGRESS_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "UPLOADFAILED":
            return UPLOADFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
        case "REMOVALFAILED":
            return REMOVALFAILED_GROUPPOLICYUPLOADEDDEFINITIONFILESTATUS, nil
    }
    return 0, errors.New("Unknown GroupPolicyUploadedDefinitionFileStatus value: " + v)
}
func SerializeGroupPolicyUploadedDefinitionFileStatus(values []GroupPolicyUploadedDefinitionFileStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
