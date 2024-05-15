package models
type PrintPresentationDirection int

const (
    CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION PrintPresentationDirection = iota
    COUNTERCLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
    COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
    CLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
)

func (i PrintPresentationDirection) String() string {
    return []string{"clockwiseFromTopLeft", "counterClockwiseFromTopLeft", "counterClockwiseFromTopRight", "clockwiseFromTopRight", "counterClockwiseFromBottomLeft", "clockwiseFromBottomLeft", "counterClockwiseFromBottomRight", "clockwiseFromBottomRight"}[i]
}
func ParsePrintPresentationDirection(v string) (any, error) {
    result := CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
    switch v {
        case "clockwiseFromTopLeft":
            result = CLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
        case "counterClockwiseFromTopLeft":
            result = COUNTERCLOCKWISEFROMTOPLEFT_PRINTPRESENTATIONDIRECTION
        case "counterClockwiseFromTopRight":
            result = COUNTERCLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
        case "clockwiseFromTopRight":
            result = CLOCKWISEFROMTOPRIGHT_PRINTPRESENTATIONDIRECTION
        case "counterClockwiseFromBottomLeft":
            result = COUNTERCLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
        case "clockwiseFromBottomLeft":
            result = CLOCKWISEFROMBOTTOMLEFT_PRINTPRESENTATIONDIRECTION
        case "counterClockwiseFromBottomRight":
            result = COUNTERCLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
        case "clockwiseFromBottomRight":
            result = CLOCKWISEFROMBOTTOMRIGHT_PRINTPRESENTATIONDIRECTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePrintPresentationDirection(values []PrintPresentationDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrintPresentationDirection) isMultiValue() bool {
    return false
}
