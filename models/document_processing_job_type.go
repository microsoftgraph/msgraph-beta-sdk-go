package models
type DocumentProcessingJobType int

const (
    FILE_DOCUMENTPROCESSINGJOBTYPE DocumentProcessingJobType = iota
    FOLDER_DOCUMENTPROCESSINGJOBTYPE
    UNKNOWNFUTUREVALUE_DOCUMENTPROCESSINGJOBTYPE
)

func (i DocumentProcessingJobType) String() string {
    return []string{"file", "folder", "unknownFutureValue"}[i]
}
func ParseDocumentProcessingJobType(v string) (any, error) {
    result := FILE_DOCUMENTPROCESSINGJOBTYPE
    switch v {
        case "file":
            result = FILE_DOCUMENTPROCESSINGJOBTYPE
        case "folder":
            result = FOLDER_DOCUMENTPROCESSINGJOBTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DOCUMENTPROCESSINGJOBTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDocumentProcessingJobType(values []DocumentProcessingJobType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DocumentProcessingJobType) isMultiValue() bool {
    return false
}
