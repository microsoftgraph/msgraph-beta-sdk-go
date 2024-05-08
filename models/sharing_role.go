package models
import (
    "errors"
)
type SharingRole int

const (
    NONE_SHARINGROLE SharingRole = iota
    VIEW_SHARINGROLE
    EDIT_SHARINGROLE
    MANAGELIST_SHARINGROLE
    REVIEW_SHARINGROLE
    RESTRICTEDVIEW_SHARINGROLE
    SUBMITONLY_SHARINGROLE
    UNKNOWNFUTUREVALUE_SHARINGROLE
)

func (i SharingRole) String() string {
    return []string{"none", "view", "edit", "manageList", "review", "restrictedView", "submitOnly", "unknownFutureValue"}[i]
}
func ParseSharingRole(v string) (any, error) {
    result := NONE_SHARINGROLE
    switch v {
        case "none":
            result = NONE_SHARINGROLE
        case "view":
            result = VIEW_SHARINGROLE
        case "edit":
            result = EDIT_SHARINGROLE
        case "manageList":
            result = MANAGELIST_SHARINGROLE
        case "review":
            result = REVIEW_SHARINGROLE
        case "restrictedView":
            result = RESTRICTEDVIEW_SHARINGROLE
        case "submitOnly":
            result = SUBMITONLY_SHARINGROLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SHARINGROLE
        default:
            return 0, errors.New("Unknown SharingRole value: " + v)
    }
    return &result, nil
}
func SerializeSharingRole(values []SharingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SharingRole) isMultiValue() bool {
    return false
}
