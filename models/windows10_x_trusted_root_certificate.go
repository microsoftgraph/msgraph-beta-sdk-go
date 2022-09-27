package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XTrustedRootCertificate 
type Windows10XTrustedRootCertificate struct {
    DeviceManagementResourceAccessProfileBase
    // File name to display in UI.
    certFileName *string
    // Possible values for the Certificate Destination Store.
    destinationStore *CertificateDestinationStore
    // Trusted Root Certificate
    trustedRootCertificate []byte
}
// NewWindows10XTrustedRootCertificate instantiates a new Windows10XTrustedRootCertificate and sets the default values.
func NewWindows10XTrustedRootCertificate()(*Windows10XTrustedRootCertificate) {
    m := &Windows10XTrustedRootCertificate{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10XTrustedRootCertificate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XTrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XTrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XTrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *Windows10XTrustedRootCertificate) GetCertFileName()(*string) {
    return m.certFileName
}
// GetDestinationStore gets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows10XTrustedRootCertificate) GetDestinationStore()(*CertificateDestinationStore) {
    return m.destinationStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XTrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
    res["certFileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCertFileName)
    res["destinationStore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCertificateDestinationStore , m.SetDestinationStore)
    res["trustedRootCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetTrustedRootCertificate)
    return res
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows10XTrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
    return m.trustedRootCertificate
}
// Serialize serializes information the current object
func (m *Windows10XTrustedRootCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementResourceAccessProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certFileName", m.GetCertFileName())
        if err != nil {
            return err
        }
    }
    if m.GetDestinationStore() != nil {
        cast := (*m.GetDestinationStore()).String()
        err = writer.WriteStringValue("destinationStore", &cast)
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
func (m *Windows10XTrustedRootCertificate) SetCertFileName(value *string)() {
    m.certFileName = value
}
// SetDestinationStore sets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows10XTrustedRootCertificate) SetDestinationStore(value *CertificateDestinationStore)() {
    m.destinationStore = value
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows10XTrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    m.trustedRootCertificate = value
}
