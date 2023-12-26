package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnAttributeCollectionStartCustomExtension 
type OnAttributeCollectionStartCustomExtension struct {
    CustomAuthenticationExtension
}
// NewOnAttributeCollectionStartCustomExtension instantiates a new onAttributeCollectionStartCustomExtension and sets the default values.
func NewOnAttributeCollectionStartCustomExtension()(*OnAttributeCollectionStartCustomExtension) {
    m := &OnAttributeCollectionStartCustomExtension{
        CustomAuthenticationExtension: *NewCustomAuthenticationExtension(),
    }
    odataTypeValue := "#microsoft.graph.onAttributeCollectionStartCustomExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnAttributeCollectionStartCustomExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnAttributeCollectionStartCustomExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnAttributeCollectionStartCustomExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnAttributeCollectionStartCustomExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomAuthenticationExtension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OnAttributeCollectionStartCustomExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomAuthenticationExtension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OnAttributeCollectionStartCustomExtensionable 
type OnAttributeCollectionStartCustomExtensionable interface {
    CustomAuthenticationExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
