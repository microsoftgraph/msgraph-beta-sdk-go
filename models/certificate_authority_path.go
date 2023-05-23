package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateAuthorityPath 
type CertificateAuthorityPath struct {
    Entity
}
// NewCertificateAuthorityPath instantiates a new certificateAuthorityPath and sets the default values.
func NewCertificateAuthorityPath()(*CertificateAuthorityPath) {
    m := &CertificateAuthorityPath{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCertificateAuthorityPathFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateAuthorityPathFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateAuthorityPath(), nil
}
// GetCertificateBasedApplicationConfigurations gets the certificateBasedApplicationConfigurations property value. The certificateBasedApplicationConfigurations property
func (m *CertificateAuthorityPath) GetCertificateBasedApplicationConfigurations()([]CertificateBasedApplicationConfigurationable) {
    val, err := m.GetBackingStore().Get("certificateBasedApplicationConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateBasedApplicationConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateAuthorityPath) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateBasedApplicationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateBasedApplicationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateBasedApplicationConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateBasedApplicationConfigurationable)
            }
            m.SetCertificateBasedApplicationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CertificateAuthorityPath) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateBasedApplicationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateBasedApplicationConfigurations()))
        for i, v := range m.GetCertificateBasedApplicationConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateBasedApplicationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateBasedApplicationConfigurations sets the certificateBasedApplicationConfigurations property value. The certificateBasedApplicationConfigurations property
func (m *CertificateAuthorityPath) SetCertificateBasedApplicationConfigurations(value []CertificateBasedApplicationConfigurationable)() {
    err := m.GetBackingStore().Set("certificateBasedApplicationConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// CertificateAuthorityPathable 
type CertificateAuthorityPathable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateBasedApplicationConfigurations()([]CertificateBasedApplicationConfigurationable)
    SetCertificateBasedApplicationConfigurations(value []CertificateBasedApplicationConfigurationable)()
}
