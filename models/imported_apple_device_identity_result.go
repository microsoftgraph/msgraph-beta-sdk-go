package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedAppleDeviceIdentityResult the importedAppleDeviceIdentityResult resource represents the result of attempting to import Apple devices identities.
type ImportedAppleDeviceIdentityResult struct {
    ImportedAppleDeviceIdentity
}
// NewImportedAppleDeviceIdentityResult instantiates a new importedAppleDeviceIdentityResult and sets the default values.
func NewImportedAppleDeviceIdentityResult()(*ImportedAppleDeviceIdentityResult) {
    m := &ImportedAppleDeviceIdentityResult{
        ImportedAppleDeviceIdentity: *NewImportedAppleDeviceIdentity(),
    }
    return m
}
// CreateImportedAppleDeviceIdentityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedAppleDeviceIdentityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedAppleDeviceIdentityResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedAppleDeviceIdentityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ImportedAppleDeviceIdentity.GetFieldDeserializers()
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Status of imported device identity
func (m *ImportedAppleDeviceIdentityResult) GetStatus()(*bool) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImportedAppleDeviceIdentityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ImportedAppleDeviceIdentity.Serialize(writer)
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
func (m *ImportedAppleDeviceIdentityResult) SetStatus(value *bool)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// ImportedAppleDeviceIdentityResultable 
type ImportedAppleDeviceIdentityResultable interface {
    ImportedAppleDeviceIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*bool)
    SetStatus(value *bool)()
}
