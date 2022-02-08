package graph
import (
    "strings"
    "errors"
)
// 
type MacOSSoftwareUpdateCategory int

const (
    CRITICAL_MACOSSOFTWAREUPDATECATEGORY MacOSSoftwareUpdateCategory = iota
    CONFIGURATIONDATAFILE_MACOSSOFTWAREUPDATECATEGORY
    FIRMWARE_MACOSSOFTWAREUPDATECATEGORY
    OTHER_MACOSSOFTWAREUPDATECATEGORY
)

func (i MacOSSoftwareUpdateCategory) String() string {
    return []string{"CRITICAL", "CONFIGURATIONDATAFILE", "FIRMWARE", "OTHER"}[i]
}
func ParseMacOSSoftwareUpdateCategory(v string) (interface{}, error) {
    result := CRITICAL_MACOSSOFTWAREUPDATECATEGORY
    switch strings.ToUpper(v) {
        case "CRITICAL":
            result = CRITICAL_MACOSSOFTWAREUPDATECATEGORY
        case "CONFIGURATIONDATAFILE":
            result = CONFIGURATIONDATAFILE_MACOSSOFTWAREUPDATECATEGORY
        case "FIRMWARE":
            result = FIRMWARE_MACOSSOFTWAREUPDATECATEGORY
        case "OTHER":
            result = OTHER_MACOSSOFTWAREUPDATECATEGORY
        default:
            return 0, errors.New("Unknown MacOSSoftwareUpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateCategory(values []MacOSSoftwareUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
