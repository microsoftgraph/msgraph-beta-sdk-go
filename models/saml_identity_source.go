package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SamlIdentitySource 
type SamlIdentitySource struct {
    PermissionsDefinitionIdentitySource
}
// NewSamlIdentitySource instantiates a new samlIdentitySource and sets the default values.
func NewSamlIdentitySource()(*SamlIdentitySource) {
    m := &SamlIdentitySource{
        PermissionsDefinitionIdentitySource: *NewPermissionsDefinitionIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.samlIdentitySource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSamlIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSamlIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSamlIdentitySource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionIdentitySource.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SamlIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SamlIdentitySourceable 
type SamlIdentitySourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionIdentitySourceable
}
