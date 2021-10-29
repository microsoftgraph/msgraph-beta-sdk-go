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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_USERPFXPADDINGSCHEME, nil
        case "PKCS1":
            return PKCS1_USERPFXPADDINGSCHEME, nil
        case "OAEPSHA1":
            return OAEPSHA1_USERPFXPADDINGSCHEME, nil
        case "OAEPSHA256":
            return OAEPSHA256_USERPFXPADDINGSCHEME, nil
        case "OAEPSHA384":
            return OAEPSHA384_USERPFXPADDINGSCHEME, nil
        case "OAEPSHA512":
            return OAEPSHA512_USERPFXPADDINGSCHEME, nil
    }
    return 0, errors.New("Unknown UserPfxPaddingScheme value: " + v)
}
func SerializeUserPfxPaddingScheme(values []UserPfxPaddingScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
