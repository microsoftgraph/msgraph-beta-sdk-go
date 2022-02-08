package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseAction int

const (
    CONTENTEXPORT_CASEACTION CaseAction = iota
    APPLYTAGS_CASEACTION
    CONVERTTOPDF_CASEACTION
    INDEX_CASEACTION
    ESTIMATESTATISTICS_CASEACTION
    ADDTOREVIEWSET_CASEACTION
    HOLDUPDATE_CASEACTION
    UNKNOWNFUTUREVALUE_CASEACTION
)

func (i CaseAction) String() string {
    return []string{"CONTENTEXPORT", "APPLYTAGS", "CONVERTTOPDF", "INDEX", "ESTIMATESTATISTICS", "ADDTOREVIEWSET", "HOLDUPDATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCaseAction(v string) (interface{}, error) {
    result := CONTENTEXPORT_CASEACTION
    switch strings.ToUpper(v) {
        case "CONTENTEXPORT":
            result = CONTENTEXPORT_CASEACTION
        case "APPLYTAGS":
            result = APPLYTAGS_CASEACTION
        case "CONVERTTOPDF":
            result = CONVERTTOPDF_CASEACTION
        case "INDEX":
            result = INDEX_CASEACTION
        case "ESTIMATESTATISTICS":
            result = ESTIMATESTATISTICS_CASEACTION
        case "ADDTOREVIEWSET":
            result = ADDTOREVIEWSET_CASEACTION
        case "HOLDUPDATE":
            result = HOLDUPDATE_CASEACTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CASEACTION
        default:
            return 0, errors.New("Unknown CaseAction value: " + v)
    }
    return &result, nil
}
func SerializeCaseAction(values []CaseAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
