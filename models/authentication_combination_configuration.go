package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationCombinationConfiguration provides operations to manage the collection of accessReviewDecision entities.
type AuthenticationCombinationConfiguration struct {
    Entity
    // The appliesToCombinations property
    appliesToCombinations []AuthenticationMethodModes
}
// NewAuthenticationCombinationConfiguration instantiates a new authenticationCombinationConfiguration and sets the default values.
func NewAuthenticationCombinationConfiguration()(*AuthenticationCombinationConfiguration) {
    m := &AuthenticationCombinationConfiguration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationCombinationConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationCombinationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationCombinationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.fido2CombinationConfiguration":
                        return NewFido2CombinationConfiguration(), nil
                }
            }
        }
    }
    return NewAuthenticationCombinationConfiguration(), nil
}
// GetAppliesToCombinations gets the appliesToCombinations property value. The appliesToCombinations property
func (m *AuthenticationCombinationConfiguration) GetAppliesToCombinations()([]AuthenticationMethodModes) {
    return m.appliesToCombinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationCombinationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesToCombinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseAuthenticationMethodModes , m.SetAppliesToCombinations)
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationCombinationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesToCombinations() != nil {
        err = writer.WriteCollectionOfStringValues("appliesToCombinations", SerializeAuthenticationMethodModes(m.GetAppliesToCombinations()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesToCombinations sets the appliesToCombinations property value. The appliesToCombinations property
func (m *AuthenticationCombinationConfiguration) SetAppliesToCombinations(value []AuthenticationMethodModes)() {
    m.appliesToCombinations = value
}
