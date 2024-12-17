package models
type ChangeItemState int

const (
    AVAILABLE_CHANGEITEMSTATE ChangeItemState = iota
    COMINGSOON_CHANGEITEMSTATE
    UNKNOWNFUTUREVALUE_CHANGEITEMSTATE
)

func (i ChangeItemState) String() string {
    return []string{"available", "comingSoon", "unknownFutureValue"}[i]
}
func ParseChangeItemState(v string) (any, error) {
    result := AVAILABLE_CHANGEITEMSTATE
    switch v {
        case "available":
            result = AVAILABLE_CHANGEITEMSTATE
        case "comingSoon":
            result = COMINGSOON_CHANGEITEMSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CHANGEITEMSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeChangeItemState(values []ChangeItemState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChangeItemState) isMultiValue() bool {
    return false
}
