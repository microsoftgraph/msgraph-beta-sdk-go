package models
type LabelKind int

const (
    ALL_LABELKIND LabelKind = iota
    ENUMERATED_LABELKIND
    UNKNOWNFUTUREVALUE_LABELKIND
)

func (i LabelKind) String() string {
    return []string{"all", "enumerated", "unknownFutureValue"}[i]
}
func ParseLabelKind(v string) (any, error) {
    result := ALL_LABELKIND
    switch v {
        case "all":
            result = ALL_LABELKIND
        case "enumerated":
            result = ENUMERATED_LABELKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LABELKIND
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLabelKind(values []LabelKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LabelKind) isMultiValue() bool {
    return false
}
