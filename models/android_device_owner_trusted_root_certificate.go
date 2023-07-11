package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerTrustedRootCertificate android Device Owner Trusted Root Certificate configuration profile
type AndroidDeviceOwnerTrustedRootCertificate struct {
    DeviceConfiguration
    // The OdataType property
    OdataType *string
}
// NewAndroidDeviceOwnerTrustedRootCertificate instantiates a new androidDeviceOwnerTrustedRootCertificate and sets the default values.
func NewAndroidDeviceOwnerTrustedRootCertificate()(*AndroidDeviceOwnerTrustedRootCertificate) {
    m := &AndroidDeviceOwnerTrustedRootCertificate{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerTrustedRootCertificate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerTrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerTrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerTrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *AndroidDeviceOwnerTrustedRootCertificate) GetCertFileName()(*string) {
    val, err := m.GetBackingStore().Get("certFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerTrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["certFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertFileName(val)
        }
        return nil
    }
    res["trustedRootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedRootCertificate(val)
        }
        return nil
    }
    return res
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate
func (m *AndroidDeviceOwnerTrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
    val, err := m.GetBackingStore().Get("trustedRootCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerTrustedRootCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certFileName", m.GetCertFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("trustedRootCertificate", m.GetTrustedRootCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertFileName sets the certFileName property value. File name to display in UI.
func (m *AndroidDeviceOwnerTrustedRootCertificate) SetCertFileName(value *string)() {
    err := m.GetBackingStore().Set("certFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate
func (m *AndroidDeviceOwnerTrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    err := m.GetBackingStore().Set("trustedRootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerTrustedRootCertificateable 
type AndroidDeviceOwnerTrustedRootCertificateable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertFileName()(*string)
    GetTrustedRootCertificate()([]byte)
    SetCertFileName(value *string)()
    SetTrustedRootCertificate(value []byte)()
}
