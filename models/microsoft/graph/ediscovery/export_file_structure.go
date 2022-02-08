package ediscovery
import (
    "strings"
    "errors"
)
// 
type ExportFileStructure int

const (
    NONE_EXPORTFILESTRUCTURE ExportFileStructure = iota
    DIRECTORY_EXPORTFILESTRUCTURE
    PST_EXPORTFILESTRUCTURE
    UNKNOWNFUTUREVALUE_EXPORTFILESTRUCTURE
)

func (i ExportFileStructure) String() string {
    return []string{"NONE", "DIRECTORY", "PST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExportFileStructure(v string) (interface{}, error) {
    result := NONE_EXPORTFILESTRUCTURE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_EXPORTFILESTRUCTURE
        case "DIRECTORY":
            result = DIRECTORY_EXPORTFILESTRUCTURE
        case "PST":
            result = PST_EXPORTFILESTRUCTURE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EXPORTFILESTRUCTURE
        default:
            return 0, errors.New("Unknown ExportFileStructure value: " + v)
    }
    return &result, nil
}
func SerializeExportFileStructure(values []ExportFileStructure) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
