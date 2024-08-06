package models
type DocumentProcessingJobStatus int

const (
    INPROGRESS_DOCUMENTPROCESSINGJOBSTATUS DocumentProcessingJobStatus = iota
    COMPLETED_DOCUMENTPROCESSINGJOBSTATUS
    FAILED_DOCUMENTPROCESSINGJOBSTATUS
    UNKNOWNFUTUREVALUE_DOCUMENTPROCESSINGJOBSTATUS
)

func (i DocumentProcessingJobStatus) String() string {
    return []string{"inProgress", "completed", "failed", "unknownFutureValue"}[i]
}
func ParseDocumentProcessingJobStatus(v string) (any, error) {
    result := INPROGRESS_DOCUMENTPROCESSINGJOBSTATUS
    switch v {
        case "inProgress":
            result = INPROGRESS_DOCUMENTPROCESSINGJOBSTATUS
        case "completed":
            result = COMPLETED_DOCUMENTPROCESSINGJOBSTATUS
        case "failed":
            result = FAILED_DOCUMENTPROCESSINGJOBSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DOCUMENTPROCESSINGJOBSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDocumentProcessingJobStatus(values []DocumentProcessingJobStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DocumentProcessingJobStatus) isMultiValue() bool {
    return false
}
