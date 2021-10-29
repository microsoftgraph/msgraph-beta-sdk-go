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
    switch strings.ToUpper(v) {
        case "UNASSIGNED":
            return UNASSIGNED_USERPFXINTENDEDPURPOSE, nil
        case "SMIMEENCRYPTION":
            return SMIMEENCRYPTION_USERPFXINTENDEDPURPOSE, nil
        case "SMIMESIGNING":
            return SMIMESIGNING_USERPFXINTENDEDPURPOSE, nil
        case "VPN":
            return VPN_USERPFXINTENDEDPURPOSE, nil
        case "WIFI":
            return WIFI_USERPFXINTENDEDPURPOSE, nil
    }
    return 0, errors.New("Unknown UserPfxIntendedPurpose value: " + v)
}
func SerializeUserPfxIntendedPurpose(values []UserPfxIntendedPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
