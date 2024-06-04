package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllScopeSensitivityLabels struct {
    ScopeSensitivityLabels
}
// NewAllScopeSensitivityLabels instantiates a new AllScopeSensitivityLabels and sets the default values.
func NewAllScopeSensitivityLabels()(*AllScopeSensitivityLabels) {
    m := &AllScopeSensitivityLabels{
        ScopeSensitivityLabels: *NewScopeSensitivityLabels(),
    }
    odataTypeValue := "#microsoft.graph.allScopeSensitivityLabels"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAllScopeSensitivityLabelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllScopeSensitivityLabelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllScopeSensitivityLabels(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllScopeSensitivityLabels) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ScopeSensitivityLabels.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllScopeSensitivityLabels) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ScopeSensitivityLabels.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AllScopeSensitivityLabelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ScopeSensitivityLabelsable
}
