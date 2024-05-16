package security
type SubmissionSource int

const (
    USER_SUBMISSIONSOURCE SubmissionSource = iota
    ADMINISTRATOR_SUBMISSIONSOURCE
    UNKNOWNFUTUREVALUE_SUBMISSIONSOURCE
)

func (i SubmissionSource) String() string {
    return []string{"user", "administrator", "unknownFutureValue"}[i]
}
func ParseSubmissionSource(v string) (any, error) {
    result := USER_SUBMISSIONSOURCE
    switch v {
        case "user":
            result = USER_SUBMISSIONSOURCE
        case "administrator":
            result = ADMINISTRATOR_SUBMISSIONSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSubmissionSource(values []SubmissionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SubmissionSource) isMultiValue() bool {
    return false
}
