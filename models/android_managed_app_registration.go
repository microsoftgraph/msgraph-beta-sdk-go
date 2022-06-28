package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppRegistration 
type AndroidManagedAppRegistration struct {
    ManagedAppRegistration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The patch version for the current android app registration
    patchVersion *string
}
// NewAndroidManagedAppRegistration instantiates a new AndroidManagedAppRegistration and sets the default values.
func NewAndroidManagedAppRegistration()(*AndroidManagedAppRegistration) {
    m := &AndroidManagedAppRegistration{
        ManagedAppRegistration: *NewManagedAppRegistration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidManagedAppRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppRegistration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedAppRegistration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppRegistration.GetFieldDeserializers()
    res["patchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchVersion(val)
        }
        return nil
    }
    return res
}
// GetPatchVersion gets the patchVersion property value. The patch version for the current android app registration
func (m *AndroidManagedAppRegistration) GetPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.patchVersion
    }
}
// Serialize serializes information the current object
func (m *AndroidManagedAppRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppRegistration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("patchVersion", m.GetPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedAppRegistration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPatchVersion sets the patchVersion property value. The patch version for the current android app registration
func (m *AndroidManagedAppRegistration) SetPatchVersion(value *string)() {
    if m != nil {
        m.patchVersion = value
    }
}
