// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable instead.
type ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse struct {
    ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponse
}
// NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse instantiates a new ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse()(*ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse) {
    m := &ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse{
        ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponse: *NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponse(),
    }
    return m
}
// CreateConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponse(), nil
}
// Deprecated: This class is obsolete. Use ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable instead.
type ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable interface {
    ConditionalAccessAuthenticationStrengthsPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
