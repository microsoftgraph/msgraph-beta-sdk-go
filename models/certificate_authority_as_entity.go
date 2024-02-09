package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CertificateAuthorityAsEntity struct {
    Entity
}
// NewCertificateAuthorityAsEntity instantiates a new CertificateAuthorityAsEntity and sets the default values.
func NewCertificateAuthorityAsEntity()(*CertificateAuthorityAsEntity) {
    m := &CertificateAuthorityAsEntity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCertificateAuthorityAsEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCertificateAuthorityAsEntityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateAuthorityAsEntity(), nil
}
// GetCertificate gets the certificate property value. The trusted certificate.
// returns a []byte when successful
func (m *CertificateAuthorityAsEntity) GetCertificate()([]byte) {
    val, err := m.GetBackingStore().Get("certificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CertificateAuthorityAsEntity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["isRootAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRootAuthority(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerSubjectKeyIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerSubjectKeyIdentifier(val)
        }
        return nil
    }
    return res
}
// GetIsRootAuthority gets the isRootAuthority property value. Indicates if the certificate is a root authority. In a certificateBasedApplicationConfiguration object, at least one object in the trustedCertificateAuthorities collection must be a root authority.
// returns a *bool when successful
func (m *CertificateAuthorityAsEntity) GetIsRootAuthority()(*bool) {
    val, err := m.GetBackingStore().Get("isRootAuthority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIssuer gets the issuer property value. The issuer of the trusted certificate.
// returns a *string when successful
func (m *CertificateAuthorityAsEntity) GetIssuer()(*string) {
    val, err := m.GetBackingStore().Get("issuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIssuerSubjectKeyIdentifier gets the issuerSubjectKeyIdentifier property value. The subject key identifier of the trusted certificate.
// returns a *string when successful
func (m *CertificateAuthorityAsEntity) GetIssuerSubjectKeyIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("issuerSubjectKeyIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CertificateAuthorityAsEntity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRootAuthority", m.GetIsRootAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerSubjectKeyIdentifier", m.GetIssuerSubjectKeyIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificate sets the certificate property value. The trusted certificate.
func (m *CertificateAuthorityAsEntity) SetCertificate(value []byte)() {
    err := m.GetBackingStore().Set("certificate", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRootAuthority sets the isRootAuthority property value. Indicates if the certificate is a root authority. In a certificateBasedApplicationConfiguration object, at least one object in the trustedCertificateAuthorities collection must be a root authority.
func (m *CertificateAuthorityAsEntity) SetIsRootAuthority(value *bool)() {
    err := m.GetBackingStore().Set("isRootAuthority", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuer sets the issuer property value. The issuer of the trusted certificate.
func (m *CertificateAuthorityAsEntity) SetIssuer(value *string)() {
    err := m.GetBackingStore().Set("issuer", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerSubjectKeyIdentifier sets the issuerSubjectKeyIdentifier property value. The subject key identifier of the trusted certificate.
func (m *CertificateAuthorityAsEntity) SetIssuerSubjectKeyIdentifier(value *string)() {
    err := m.GetBackingStore().Set("issuerSubjectKeyIdentifier", value)
    if err != nil {
        panic(err)
    }
}
type CertificateAuthorityAsEntityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()([]byte)
    GetIsRootAuthority()(*bool)
    GetIssuer()(*string)
    GetIssuerSubjectKeyIdentifier()(*string)
    SetCertificate(value []byte)()
    SetIsRootAuthority(value *bool)()
    SetIssuer(value *string)()
    SetIssuerSubjectKeyIdentifier(value *string)()
}
