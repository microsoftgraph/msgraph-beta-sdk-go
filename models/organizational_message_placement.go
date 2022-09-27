package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type OrganizationalMessagePlacement int

const (
    // Indicates the default area for text to be displayed. This is the only valid placement value for the actionCenter and softLanding surfaces
    DEFAULT_ESCAPED_ORGANIZATIONALMESSAGEPLACEMENT OrganizationalMessagePlacement = iota
    // Indicates the area where the first card is displayed. Only applies to the getStarted surface
    CARD0_ORGANIZATIONALMESSAGEPLACEMENT
    // Indicates the area where the second card is displayed. Only applies to the getStarted surface
    CARD1_ORGANIZATIONALMESSAGEPLACEMENT
    // Indicates the area where the third card is displayed. Only applies to the getStarted surface
    CARD2_ORGANIZATIONALMESSAGEPLACEMENT
    // Indicates the area where the fourth card is displayed. Only applies to the getStarted surface
    CARD3_ORGANIZATIONALMESSAGEPLACEMENT
    // UnknownFutureValue, Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGEPLACEMENT
)

func (i OrganizationalMessagePlacement) String() string {
    return []string{"default", "card0", "card1", "card2", "card3", "unknownFutureValue"}[i]
}
func ParseOrganizationalMessagePlacement(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_ORGANIZATIONALMESSAGEPLACEMENT
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_ORGANIZATIONALMESSAGEPLACEMENT
        case "card0":
            result = CARD0_ORGANIZATIONALMESSAGEPLACEMENT
        case "card1":
            result = CARD1_ORGANIZATIONALMESSAGEPLACEMENT
        case "card2":
            result = CARD2_ORGANIZATIONALMESSAGEPLACEMENT
        case "card3":
            result = CARD3_ORGANIZATIONALMESSAGEPLACEMENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGEPLACEMENT
        default:
            return 0, errors.New("Unknown OrganizationalMessagePlacement value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessagePlacement(values []OrganizationalMessagePlacement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
