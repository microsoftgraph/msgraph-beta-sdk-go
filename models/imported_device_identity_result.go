package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedDeviceIdentityResult 
type ImportedDeviceIdentityResult struct {
    ImportedDeviceIdentity
    // Status of imported device identity
    status *bool
}
// NewImportedDeviceIdentityResult instantiates a new ImportedDeviceIdentityResult and sets the default values.
func NewImportedDeviceIdentityResult()(*ImportedDeviceIdentityResult) {
    m := &ImportedDeviceIdentityResult{
        ImportedDeviceIdentity: *NewImportedDeviceIdentity(),
    }
    return m
}
// CreateImportedDeviceIdentityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedDeviceIdentityResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ImportedDeviceIdentity.GetFieldDeserializers()
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStatus)
    return res
}
// GetStatus gets the status property value. Status of imported device identity
func (m *ImportedDeviceIdentityResult) GetStatus()(*bool) {
    return m.status
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ImportedDeviceIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStatus sets the status property value. Status of imported device identity
func (m *ImportedDeviceIdentityResult) SetStatus(value *bool)() {
    m.status = value
}
