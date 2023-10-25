package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SuperAwsResourceFinding 
type SuperAwsResourceFinding struct {
    IdentityFinding
}
// NewSuperAwsResourceFinding instantiates a new superAwsResourceFinding and sets the default values.
func NewSuperAwsResourceFinding()(*SuperAwsResourceFinding) {
    m := &SuperAwsResourceFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateSuperAwsResourceFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSuperAwsResourceFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSuperAwsResourceFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SuperAwsResourceFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SuperAwsResourceFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SuperAwsResourceFindingable 
type SuperAwsResourceFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
