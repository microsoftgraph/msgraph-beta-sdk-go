package graph
import (
    "strings"
    "errors"
)
// 
type UserPfxIntendedPurpose int

const (
    UNASSIGNED_USERPFXINTENDEDPURPOSE UserPfxIntendedPurpose = iota
    SMIMEENCRYPTION_USERPFXINTENDEDPURPOSE
    SMIMESIGNING_USERPFXINTENDEDPURPOSE
    VPN_USERPFXINTENDEDPURPOSE
    WIFI_USERPFXINTENDEDPURPOSE
)

func (i UserPfxIntendedPurpose) String() string {
    return []string{"UNASSIGNED", "SMIMEENCRYPTION", "SMIMESIGNING", "VPN", "WIFI"}[i]
}
func ParseUserPfxIntendedPurpose(v string) (interface{}, error) {
    result := UNASSIGNED_USERPFXINTENDEDPURPOSE
    switch strings.ToUpper(v) {
        case "UNASSIGNED":
            result = UNASSIGNED_USERPFXINTENDEDPURPOSE
        case "SMIMEENCRYPTION":
            result = SMIMEENCRYPTION_USERPFXINTENDEDPURPOSE
        case "SMIMESIGNING":
            result = SMIMESIGNING_USERPFXINTENDEDPURPOSE
        case "VPN":
            result = VPN_USERPFXINTENDEDPURPOSE
        case "WIFI":
            result = WIFI_USERPFXINTENDEDPURPOSE
        default:
            return 0, errors.New("Unknown UserPfxIntendedPurpose value: " + v)
    }
    return &result, nil
}
func SerializeUserPfxIntendedPurpose(values []UserPfxIntendedPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
