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
    switch strings.ToUpper(v) {
        case "STATIONERY":
            return STATIONERY_PRINTMEDIATYPE, nil
        case "TRANSPARENCY":
            return TRANSPARENCY_PRINTMEDIATYPE, nil
        case "ENVELOPE":
            return ENVELOPE_PRINTMEDIATYPE, nil
        case "ENVELOPEPLAIN":
            return ENVELOPEPLAIN_PRINTMEDIATYPE, nil
        case "CONTINUOUS":
            return CONTINUOUS_PRINTMEDIATYPE, nil
        case "SCREEN":
            return SCREEN_PRINTMEDIATYPE, nil
        case "SCREENPAGED":
            return SCREENPAGED_PRINTMEDIATYPE, nil
        case "CONTINUOUSLONG":
            return CONTINUOUSLONG_PRINTMEDIATYPE, nil
        case "CONTINUOUSSHORT":
            return CONTINUOUSSHORT_PRINTMEDIATYPE, nil
        case "ENVELOPEWINDOW":
            return ENVELOPEWINDOW_PRINTMEDIATYPE, nil
        case "MULTIPARTFORM":
            return MULTIPARTFORM_PRINTMEDIATYPE, nil
        case "MULTILAYER":
            return MULTILAYER_PRINTMEDIATYPE, nil
        case "LABELS":
            return LABELS_PRINTMEDIATYPE, nil
    }
    return 0, errors.New("Unknown PrintMediaType value: " + v)
}
func SerializePrintMediaType(values []PrintMediaType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
