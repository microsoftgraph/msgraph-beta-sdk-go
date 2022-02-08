package ediscovery
import (
    "strings"
    "errors"
)
// 
type ExportOptions int

const (
    ORIGINALFILES_EXPORTOPTIONS ExportOptions = iota
    TEXT_EXPORTOPTIONS
    PDFREPLACEMENT_EXPORTOPTIONS
    FILEINFO_EXPORTOPTIONS
    TAGS_EXPORTOPTIONS
    UNKNOWNFUTUREVALUE_EXPORTOPTIONS
)

func (i ExportOptions) String() string {
    return []string{"ORIGINALFILES", "TEXT", "PDFREPLACEMENT", "FILEINFO", "TAGS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExportOptions(v string) (interface{}, error) {
    result := ORIGINALFILES_EXPORTOPTIONS
    switch strings.ToUpper(v) {
        case "ORIGINALFILES":
            result = ORIGINALFILES_EXPORTOPTIONS
        case "TEXT":
            result = TEXT_EXPORTOPTIONS
        case "PDFREPLACEMENT":
            result = PDFREPLACEMENT_EXPORTOPTIONS
        case "FILEINFO":
            result = FILEINFO_EXPORTOPTIONS
        case "TAGS":
            result = TAGS_EXPORTOPTIONS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EXPORTOPTIONS
        default:
            return 0, errors.New("Unknown ExportOptions value: " + v)
    }
    return &result, nil
}
func SerializeExportOptions(values []ExportOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
