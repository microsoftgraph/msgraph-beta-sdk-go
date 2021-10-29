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
    switch strings.ToUpper(v) {
        case "CRITICAL":
            return CRITICAL_MACOSSOFTWAREUPDATECATEGORY, nil
        case "CONFIGURATIONDATAFILE":
            return CONFIGURATIONDATAFILE_MACOSSOFTWAREUPDATECATEGORY, nil
        case "FIRMWARE":
            return FIRMWARE_MACOSSOFTWAREUPDATECATEGORY, nil
        case "OTHER":
            return OTHER_MACOSSOFTWAREUPDATECATEGORY, nil
    }
    return 0, errors.New("Unknown MacOSSoftwareUpdateCategory value: " + v)
}
func SerializeMacOSSoftwareUpdateCategory(values []MacOSSoftwareUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
