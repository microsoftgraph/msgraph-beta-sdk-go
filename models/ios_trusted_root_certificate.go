package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosTrustedRootCertificate iOS Trusted Root Certificate configuration profile.
type IosTrustedRootCertificate struct {
    DeviceConfiguration
}
// NewIosTrustedRootCertificate instantiates a new IosTrustedRootCertificate and sets the default values.
func NewIosTrustedRootCertificate()(*IosTrustedRootCertificate) {
    m := &IosTrustedRootCertificate{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosTrustedRootCertificate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosTrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosTrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosTrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
// returns a *string when successful
func (m *IosTrustedRootCertificate) GetCertFileName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IosTrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate.
// returns a []byte when successful
func (m *IosTrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
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
func (m *IosTrustedRootCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IosTrustedRootCertificate) SetCertFileName(value *string)() {
    err := m.GetBackingStore().Set("certFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *IosTrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    err := m.GetBackingStore().Set("trustedRootCertificate", value)
    if err != nil {
        panic(err)
    }
}
type IosTrustedRootCertificateable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertFileName()(*string)
    GetTrustedRootCertificate()([]byte)
    SetCertFileName(value *string)()
    SetTrustedRootCertificate(value []byte)()
}
