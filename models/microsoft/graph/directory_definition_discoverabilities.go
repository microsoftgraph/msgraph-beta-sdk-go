package graph
import (
    "strings"
    "errors"
)
// 
type DirectoryDefinitionDiscoverabilities int

const (
    NONE_DIRECTORYDEFINITIONDISCOVERABILITIES DirectoryDefinitionDiscoverabilities = iota
    ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES
    ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES
    ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES
    REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES
    UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES
)

func (i DirectoryDefinitionDiscoverabilities) String() string {
    return []string{"NONE", "ATTRIBUTENAMES", "ATTRIBUTEDATATYPES", "ATTRIBUTEREADONLY", "REFERENCEATTRIBUTES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDirectoryDefinitionDiscoverabilities(v string) (interface{}, error) {
    result := NONE_DIRECTORYDEFINITIONDISCOVERABILITIES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "ATTRIBUTENAMES":
            result = ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "ATTRIBUTEDATATYPES":
            result = ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "ATTRIBUTEREADONLY":
            result = ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "REFERENCEATTRIBUTES":
            result = REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES
        default:
            return 0, errors.New("Unknown DirectoryDefinitionDiscoverabilities value: " + v)
    }
    return &result, nil
}
func SerializeDirectoryDefinitionDiscoverabilities(values []DirectoryDefinitionDiscoverabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
