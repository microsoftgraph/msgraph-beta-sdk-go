package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateBasedApplicationConfiguration 
type CertificateBasedApplicationConfiguration struct {
    TrustedCertificateAuthorityAsEntityBase
}
// NewCertificateBasedApplicationConfiguration instantiates a new certificateBasedApplicationConfiguration and sets the default values.
func NewCertificateBasedApplicationConfiguration()(*CertificateBasedApplicationConfiguration) {
    m := &CertificateBasedApplicationConfiguration{
        TrustedCertificateAuthorityAsEntityBase: *NewTrustedCertificateAuthorityAsEntityBase(),
    }
    odataTypeValue := "#microsoft.graph.certificateBasedApplicationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCertificateBasedApplicationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateBasedApplicationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateBasedApplicationConfiguration(), nil
}
// GetDescription gets the description property value. The description property
func (m *CertificateBasedApplicationConfiguration) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CertificateBasedApplicationConfiguration) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateBasedApplicationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TrustedCertificateAuthorityAsEntityBase.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CertificateBasedApplicationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TrustedCertificateAuthorityAsEntityBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description property
func (m *CertificateBasedApplicationConfiguration) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CertificateBasedApplicationConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// CertificateBasedApplicationConfigurationable 
type CertificateBasedApplicationConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrustedCertificateAuthorityAsEntityBaseable
    GetDescription()(*string)
    GetDisplayName()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
}
