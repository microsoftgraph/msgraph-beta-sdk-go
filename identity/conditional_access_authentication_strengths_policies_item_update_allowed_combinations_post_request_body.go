package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody 
type ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowedCombinations property
    allowedCombinations []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodModes
}
// NewConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody instantiates a new ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody and sets the default values.
func NewConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody()(*ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) {
    m := &ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedCombinations gets the allowedCombinations property value. The allowedCombinations property
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) GetAllowedCombinations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodModes) {
    return m.allowedCombinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedCombinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseAuthenticationMethodModes)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodModes, len(val))
            for i, v := range val {
                res[i] = *(v.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodModes))
            }
            m.SetAllowedCombinations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedCombinations() != nil {
        err := writer.WriteCollectionOfStringValues("allowedCombinations", ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SerializeAuthenticationMethodModes(m.GetAllowedCombinations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedCombinations sets the allowedCombinations property value. The allowedCombinations property
func (m *ConditionalAccessAuthenticationStrengthsPoliciesItemUpdateAllowedCombinationsPostRequestBody) SetAllowedCombinations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodModes)() {
    m.allowedCombinations = value
}
