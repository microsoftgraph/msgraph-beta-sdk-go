package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnAttributeCollectionSubmitCustomExtension 
type OnAttributeCollectionSubmitCustomExtension struct {
    CustomAuthenticationExtension
}
// NewOnAttributeCollectionSubmitCustomExtension instantiates a new onAttributeCollectionSubmitCustomExtension and sets the default values.
func NewOnAttributeCollectionSubmitCustomExtension()(*OnAttributeCollectionSubmitCustomExtension) {
    m := &OnAttributeCollectionSubmitCustomExtension{
        CustomAuthenticationExtension: *NewCustomAuthenticationExtension(),
    }
    odataTypeValue := "#microsoft.graph.onAttributeCollectionSubmitCustomExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnAttributeCollectionSubmitCustomExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnAttributeCollectionSubmitCustomExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnAttributeCollectionSubmitCustomExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnAttributeCollectionSubmitCustomExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomAuthenticationExtension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OnAttributeCollectionSubmitCustomExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomAuthenticationExtension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OnAttributeCollectionSubmitCustomExtensionable 
type OnAttributeCollectionSubmitCustomExtensionable interface {
    CustomAuthenticationExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
