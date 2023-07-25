package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// YearReferenceValue 
type YearReferenceValue struct {
    ReferenceValue
}
// NewYearReferenceValue instantiates a new yearReferenceValue and sets the default values.
func NewYearReferenceValue()(*YearReferenceValue) {
    m := &YearReferenceValue{
        ReferenceValue: *NewReferenceValue(),
    }
    odataTypeValue := "#microsoft.graph.industryData.yearReferenceValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateYearReferenceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateYearReferenceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewYearReferenceValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *YearReferenceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ReferenceValue.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *YearReferenceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ReferenceValue.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// YearReferenceValueable 
type YearReferenceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReferenceValueable
}
