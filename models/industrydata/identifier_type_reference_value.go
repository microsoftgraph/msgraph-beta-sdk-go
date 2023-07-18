package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentifierTypeReferenceValue 
type IdentifierTypeReferenceValue struct {
    ReferenceValue
    // The OdataType property
    OdataType *string
}
// NewIdentifierTypeReferenceValue instantiates a new identifierTypeReferenceValue and sets the default values.
func NewIdentifierTypeReferenceValue()(*IdentifierTypeReferenceValue) {
    m := &IdentifierTypeReferenceValue{
        ReferenceValue: *NewReferenceValue(),
    }
    odataTypeValue := "#microsoft.graph.industryData.identifierTypeReferenceValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIdentifierTypeReferenceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentifierTypeReferenceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentifierTypeReferenceValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentifierTypeReferenceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ReferenceValue.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IdentifierTypeReferenceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ReferenceValue.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// IdentifierTypeReferenceValueable 
type IdentifierTypeReferenceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReferenceValueable
}
