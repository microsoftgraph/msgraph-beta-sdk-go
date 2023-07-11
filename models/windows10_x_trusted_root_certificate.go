package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XTrustedRootCertificate windows X Trusted Root Certificate configuration profile
type Windows10XTrustedRootCertificate struct {
    DeviceManagementResourceAccessProfileBase
}
// NewWindows10XTrustedRootCertificate instantiates a new windows10XTrustedRootCertificate and sets the default values.
func NewWindows10XTrustedRootCertificate()(*Windows10XTrustedRootCertificate) {
    m := &Windows10XTrustedRootCertificate{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10XTrustedRootCertificate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10XTrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XTrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XTrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *Windows10XTrustedRootCertificate) GetCertFileName()(*string) {
    val, err := m.GetBackingStore().Get("certFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationStore gets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows10XTrustedRootCertificate) GetDestinationStore()(*CertificateDestinationStore) {
    val, err := m.GetBackingStore().Get("destinationStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateDestinationStore)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XTrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
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
    res["destinationStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateDestinationStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationStore(val.(*CertificateDestinationStore))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Windows10XTrustedRootCertificate) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows10XTrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("certFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationStore sets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows10XTrustedRootCertificate) SetDestinationStore(value *CertificateDestinationStore)() {
    err := m.GetBackingStore().Set("destinationStore", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Windows10XTrustedRootCertificate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows10XTrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    err := m.GetBackingStore().Set("trustedRootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// Windows10XTrustedRootCertificateable 
type Windows10XTrustedRootCertificateable interface {
    DeviceManagementResourceAccessProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertFileName()(*string)
    GetDestinationStore()(*CertificateDestinationStore)
    GetOdataType()(*string)
    GetTrustedRootCertificate()([]byte)
    SetCertFileName(value *string)()
    SetDestinationStore(value *CertificateDestinationStore)()
    SetOdataType(value *string)()
    SetTrustedRootCertificate(value []byte)()
}
