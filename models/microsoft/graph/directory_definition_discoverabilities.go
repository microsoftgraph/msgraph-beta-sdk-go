package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
        case "ATTRIBUTENAMES":
            return ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
        case "ATTRIBUTEDATATYPES":
            return ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
        case "ATTRIBUTEREADONLY":
            return ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
        case "REFERENCEATTRIBUTES":
            return REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES, nil
    }
    return 0, errors.New("Unknown DirectoryDefinitionDiscoverabilities value: " + v)
}
func SerializeDirectoryDefinitionDiscoverabilities(values []DirectoryDefinitionDiscoverabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
