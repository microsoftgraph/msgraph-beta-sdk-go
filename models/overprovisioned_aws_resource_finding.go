package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedAwsResourceFinding 
type OverprovisionedAwsResourceFinding struct {
    IdentityFinding
}
// NewOverprovisionedAwsResourceFinding instantiates a new overprovisionedAwsResourceFinding and sets the default values.
func NewOverprovisionedAwsResourceFinding()(*OverprovisionedAwsResourceFinding) {
    m := &OverprovisionedAwsResourceFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedAwsResourceFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedAwsResourceFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedAwsResourceFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedAwsResourceFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedAwsResourceFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedAwsResourceFindingable 
type OverprovisionedAwsResourceFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
