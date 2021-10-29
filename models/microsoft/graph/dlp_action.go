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
    switch strings.ToUpper(v) {
        case "NOTIFYUSER":
            return NOTIFYUSER_DLPACTION, nil
        case "BLOCKACCESS":
            return BLOCKACCESS_DLPACTION, nil
        case "DEVICERESTRICTION":
            return DEVICERESTRICTION_DLPACTION, nil
    }
    return 0, errors.New("Unknown DlpAction value: " + v)
}
func SerializeDlpAction(values []DlpAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
