package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2CombinationConfiguration 
type Fido2CombinationConfiguration struct {
    AuthenticationCombinationConfiguration
    // The allowedAAGUIDs property
    allowedAAGUIDs []string
}
// NewFido2CombinationConfiguration instantiates a new Fido2CombinationConfiguration and sets the default values.
func NewFido2CombinationConfiguration()(*Fido2CombinationConfiguration) {
    m := &Fido2CombinationConfiguration{
        AuthenticationCombinationConfiguration: *NewAuthenticationCombinationConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.fido2CombinationConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFido2CombinationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2CombinationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFido2CombinationConfiguration(), nil
}
// GetAllowedAAGUIDs gets the allowedAAGUIDs property value. The allowedAAGUIDs property
func (m *Fido2CombinationConfiguration) GetAllowedAAGUIDs()([]string) {
    return m.allowedAAGUIDs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2CombinationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationCombinationConfiguration.GetFieldDeserializers()
    res["allowedAAGUIDs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedAAGUIDs)
    return res
}
// Serialize serializes information the current object
func (m *Fido2CombinationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationCombinationConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedAAGUIDs() != nil {
        err = writer.WriteCollectionOfStringValues("allowedAAGUIDs", m.GetAllowedAAGUIDs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAAGUIDs sets the allowedAAGUIDs property value. The allowedAAGUIDs property
func (m *Fido2CombinationConfiguration) SetAllowedAAGUIDs(value []string)() {
    m.allowedAAGUIDs = value
}
