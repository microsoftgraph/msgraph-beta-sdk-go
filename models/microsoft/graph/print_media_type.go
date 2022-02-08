package graph
import (
    "strings"
    "errors"
)
// 
type PrintMediaType int

const (
    STATIONERY_PRINTMEDIATYPE PrintMediaType = iota
    TRANSPARENCY_PRINTMEDIATYPE
    ENVELOPE_PRINTMEDIATYPE
    ENVELOPEPLAIN_PRINTMEDIATYPE
    CONTINUOUS_PRINTMEDIATYPE
    SCREEN_PRINTMEDIATYPE
    SCREENPAGED_PRINTMEDIATYPE
    CONTINUOUSLONG_PRINTMEDIATYPE
    CONTINUOUSSHORT_PRINTMEDIATYPE
    ENVELOPEWINDOW_PRINTMEDIATYPE
    MULTIPARTFORM_PRINTMEDIATYPE
    MULTILAYER_PRINTMEDIATYPE
    LABELS_PRINTMEDIATYPE
)

func (i PrintMediaType) String() string {
    return []string{"STATIONERY", "TRANSPARENCY", "ENVELOPE", "ENVELOPEPLAIN", "CONTINUOUS", "SCREEN", "SCREENPAGED", "CONTINUOUSLONG", "CONTINUOUSSHORT", "ENVELOPEWINDOW", "MULTIPARTFORM", "MULTILAYER", "LABELS"}[i]
}
func ParsePrintMediaType(v string) (interface{}, error) {
    result := STATIONERY_PRINTMEDIATYPE
    switch strings.ToUpper(v) {
        case "STATIONERY":
            result = STATIONERY_PRINTMEDIATYPE
        case "TRANSPARENCY":
            result = TRANSPARENCY_PRINTMEDIATYPE
        case "ENVELOPE":
            result = ENVELOPE_PRINTMEDIATYPE
        case "ENVELOPEPLAIN":
            result = ENVELOPEPLAIN_PRINTMEDIATYPE
        case "CONTINUOUS":
            result = CONTINUOUS_PRINTMEDIATYPE
        case "SCREEN":
            result = SCREEN_PRINTMEDIATYPE
        case "SCREENPAGED":
            result = SCREENPAGED_PRINTMEDIATYPE
        case "CONTINUOUSLONG":
            result = CONTINUOUSLONG_PRINTMEDIATYPE
        case "CONTINUOUSSHORT":
            result = CONTINUOUSSHORT_PRINTMEDIATYPE
        case "ENVELOPEWINDOW":
            result = ENVELOPEWINDOW_PRINTMEDIATYPE
        case "MULTIPARTFORM":
            result = MULTIPARTFORM_PRINTMEDIATYPE
        case "MULTILAYER":
            result = MULTILAYER_PRINTMEDIATYPE
        case "LABELS":
            result = LABELS_PRINTMEDIATYPE
        default:
            return 0, errors.New("Unknown PrintMediaType value: " + v)
    }
    return &result, nil
}
func SerializePrintMediaType(values []PrintMediaType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
