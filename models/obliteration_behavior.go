package models
import (
    "errors"
)
// In macOS 12 and later, this command uses Erase All Content and Settings (EACS) on Mac computers with the Apple M1 chip or the Apple T2 Security Chip. On those devices, if EACS can’t run, the device can use obliteration (macOS 11.x behavior). This key has no effect on machines prior to the T2 chip. Upon receiving this command, the device performs preflight checks to determine if the device is in a state that allows EACS. The ObliterationBehavior value defines the device's fallback behavior.
type ObliterationBehavior int

const (
    // Default. If Erase All Content and Settings (EACS) preflight fails, the device responds to the server with an Error status and then attempts to erase itself. If EACS preflight succeeds but EACS fails, then the device attempts to erase itself.
    DEFAULTESCAPED_OBLITERATIONBEHAVIOR ObliterationBehavior = iota
    // If Erase All Content and Settings (EACS) preflight fails, the device responds to the server with an Error status and doesn’t attempt to erase itself. If EACS preflight succeeds but EACS fails, then the device doesn’t attempt to erase itself.
    DONOTOBLITERATE_OBLITERATIONBEHAVIOR
    // If Erase All Content and Settings (EACS) preflight fails, the device responds with an Acknowledged status and then attempts to erase itself. If EACS preflight succeeds but EACS fails, then the device attempts to erase itself.
    OBLITERATEWITHWARNING_OBLITERATIONBEHAVIOR
    // The system doesn’t attempt Erase All Content and Settings (EACS). T2 and later devices always obliterate.
    ALWAYS_OBLITERATIONBEHAVIOR
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OBLITERATIONBEHAVIOR
)

func (i ObliterationBehavior) String() string {
    return []string{"default", "doNotObliterate", "obliterateWithWarning", "always", "unknownFutureValue"}[i]
}
func ParseObliterationBehavior(v string) (any, error) {
    result := DEFAULTESCAPED_OBLITERATIONBEHAVIOR
    switch v {
        case "default":
            result = DEFAULTESCAPED_OBLITERATIONBEHAVIOR
        case "doNotObliterate":
            result = DONOTOBLITERATE_OBLITERATIONBEHAVIOR
        case "obliterateWithWarning":
            result = OBLITERATEWITHWARNING_OBLITERATIONBEHAVIOR
        case "always":
            result = ALWAYS_OBLITERATIONBEHAVIOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OBLITERATIONBEHAVIOR
        default:
            return 0, errors.New("Unknown ObliterationBehavior value: " + v)
    }
    return &result, nil
}
func SerializeObliterationBehavior(values []ObliterationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
