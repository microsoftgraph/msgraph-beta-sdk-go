package models
import (
    "errors"
)
// 
type EncryptWith int

const (
    TEMPLATE_ENCRYPTWITH EncryptWith = iota
    USERDEFINEDRIGHTS_ENCRYPTWITH
)

func (i EncryptWith) String() string {
    return []string{"template", "userDefinedRights"}[i]
}
func ParseEncryptWith(v string) (any, error) {
    result := TEMPLATE_ENCRYPTWITH
    switch v {
        case "template":
            result = TEMPLATE_ENCRYPTWITH
        case "userDefinedRights":
            result = USERDEFINEDRIGHTS_ENCRYPTWITH
        default:
            return 0, errors.New("Unknown EncryptWith value: " + v)
    }
    return &result, nil
}
func SerializeEncryptWith(values []EncryptWith) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EncryptWith) isMultiValue() bool {
    return false
}
