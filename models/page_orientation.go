package models
type PageOrientation int

const (
    HORIZONTAL_PAGEORIENTATION PageOrientation = iota
    DIAGONAL_PAGEORIENTATION
)

func (i PageOrientation) String() string {
    return []string{"horizontal", "diagonal"}[i]
}
func ParsePageOrientation(v string) (any, error) {
    result := HORIZONTAL_PAGEORIENTATION
    switch v {
        case "horizontal":
            result = HORIZONTAL_PAGEORIENTATION
        case "diagonal":
            result = DIAGONAL_PAGEORIENTATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePageOrientation(values []PageOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PageOrientation) isMultiValue() bool {
    return false
}
