package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileFormatReferenceValue struct {
    ReferenceValue
}
// NewFileFormatReferenceValue instantiates a new FileFormatReferenceValue and sets the default values.
func NewFileFormatReferenceValue()(*FileFormatReferenceValue) {
    m := &FileFormatReferenceValue{
        ReferenceValue: *NewReferenceValue(),
    }
    odataTypeValue := "#microsoft.graph.industryData.fileFormatReferenceValue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileFormatReferenceValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileFormatReferenceValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileFormatReferenceValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileFormatReferenceValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ReferenceValue.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FileFormatReferenceValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ReferenceValue.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type FileFormatReferenceValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ReferenceValueable
}
