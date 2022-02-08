package graph
import (
    "strings"
    "errors"
)
// 
type DlpAction int

const (
    NOTIFYUSER_DLPACTION DlpAction = iota
    BLOCKACCESS_DLPACTION
    DEVICERESTRICTION_DLPACTION
)

func (i DlpAction) String() string {
    return []string{"NOTIFYUSER", "BLOCKACCESS", "DEVICERESTRICTION"}[i]
}
func ParseDlpAction(v string) (interface{}, error) {
    result := NOTIFYUSER_DLPACTION
    switch strings.ToUpper(v) {
        case "NOTIFYUSER":
            result = NOTIFYUSER_DLPACTION
        case "BLOCKACCESS":
            result = BLOCKACCESS_DLPACTION
        case "DEVICERESTRICTION":
            result = DEVICERESTRICTION_DLPACTION
        default:
            return 0, errors.New("Unknown DlpAction value: " + v)
    }
    return &result, nil
}
func SerializeDlpAction(values []DlpAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
