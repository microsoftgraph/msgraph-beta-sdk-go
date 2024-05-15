package models
// Describes the OfficeSuiteApp file format types that can be selected.
type OfficeSuiteDefaultFileFormatType int

const (
    // No file format selected
    NOTCONFIGURED_OFFICESUITEDEFAULTFILEFORMATTYPE OfficeSuiteDefaultFileFormatType = iota
    // Office Open XML Format selected
    OFFICEOPENXMLFORMAT_OFFICESUITEDEFAULTFILEFORMATTYPE
    // Office Open Document Format selected
    OFFICEOPENDOCUMENTFORMAT_OFFICESUITEDEFAULTFILEFORMATTYPE
    // Placeholder for evolvable enum.
    UNKNOWNFUTUREVALUE_OFFICESUITEDEFAULTFILEFORMATTYPE
)

func (i OfficeSuiteDefaultFileFormatType) String() string {
    return []string{"notConfigured", "officeOpenXMLFormat", "officeOpenDocumentFormat", "unknownFutureValue"}[i]
}
func ParseOfficeSuiteDefaultFileFormatType(v string) (any, error) {
    result := NOTCONFIGURED_OFFICESUITEDEFAULTFILEFORMATTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_OFFICESUITEDEFAULTFILEFORMATTYPE
        case "officeOpenXMLFormat":
            result = OFFICEOPENXMLFORMAT_OFFICESUITEDEFAULTFILEFORMATTYPE
        case "officeOpenDocumentFormat":
            result = OFFICEOPENDOCUMENTFORMAT_OFFICESUITEDEFAULTFILEFORMATTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OFFICESUITEDEFAULTFILEFORMATTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOfficeSuiteDefaultFileFormatType(values []OfficeSuiteDefaultFileFormatType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OfficeSuiteDefaultFileFormatType) isMultiValue() bool {
    return false
}
