package graph
import (
    "strings"
    "errors"
)
// 
type UserPfxPaddingScheme int

const (
    NONE_USERPFXPADDINGSCHEME UserPfxPaddingScheme = iota
    PKCS1_USERPFXPADDINGSCHEME
    OAEPSHA1_USERPFXPADDINGSCHEME
    OAEPSHA256_USERPFXPADDINGSCHEME
    OAEPSHA384_USERPFXPADDINGSCHEME
    OAEPSHA512_USERPFXPADDINGSCHEME
)

func (i UserPfxPaddingScheme) String() string {
    return []string{"NONE", "PKCS1", "OAEPSHA1", "OAEPSHA256", "OAEPSHA384", "OAEPSHA512"}[i]
}
func ParseUserPfxPaddingScheme(v string) (interface{}, error) {
    result := NONE_USERPFXPADDINGSCHEME
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_USERPFXPADDINGSCHEME
        case "PKCS1":
            result = PKCS1_USERPFXPADDINGSCHEME
        case "OAEPSHA1":
            result = OAEPSHA1_USERPFXPADDINGSCHEME
        case "OAEPSHA256":
            result = OAEPSHA256_USERPFXPADDINGSCHEME
        case "OAEPSHA384":
            result = OAEPSHA384_USERPFXPADDINGSCHEME
        case "OAEPSHA512":
            result = OAEPSHA512_USERPFXPADDINGSCHEME
        default:
            return 0, errors.New("Unknown UserPfxPaddingScheme value: " + v)
    }
    return &result, nil
}
func SerializeUserPfxPaddingScheme(values []UserPfxPaddingScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
