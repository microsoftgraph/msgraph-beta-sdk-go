package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustedCertificateAuthorityAsEntityBase 
type TrustedCertificateAuthorityAsEntityBase struct {
    DirectoryObject
}
// NewTrustedCertificateAuthorityAsEntityBase instantiates a new trustedCertificateAuthorityAsEntityBase and sets the default values.
func NewTrustedCertificateAuthorityAsEntityBase()(*TrustedCertificateAuthorityAsEntityBase) {
    m := &TrustedCertificateAuthorityAsEntityBase{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.trustedCertificateAuthorityAsEntityBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTrustedCertificateAuthorityAsEntityBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustedCertificateAuthorityAsEntityBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.certificateBasedApplicationConfiguration":
                        return NewCertificateBasedApplicationConfiguration(), nil
                }
            }
        }
    }
    return NewTrustedCertificateAuthorityAsEntityBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustedCertificateAuthorityAsEntityBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["trustedCertificateAuthorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateAuthorityAsEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateAuthorityAsEntityable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateAuthorityAsEntityable)
            }
            m.SetTrustedCertificateAuthorities(res)
        }
        return nil
    }
    return res
}
// GetTrustedCertificateAuthorities gets the trustedCertificateAuthorities property value. The trustedCertificateAuthorities property
func (m *TrustedCertificateAuthorityAsEntityBase) GetTrustedCertificateAuthorities()([]CertificateAuthorityAsEntityable) {
    val, err := m.GetBackingStore().Get("trustedCertificateAuthorities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateAuthorityAsEntityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrustedCertificateAuthorityAsEntityBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTrustedCertificateAuthorities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrustedCertificateAuthorities()))
        for i, v := range m.GetTrustedCertificateAuthorities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("trustedCertificateAuthorities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTrustedCertificateAuthorities sets the trustedCertificateAuthorities property value. The trustedCertificateAuthorities property
func (m *TrustedCertificateAuthorityAsEntityBase) SetTrustedCertificateAuthorities(value []CertificateAuthorityAsEntityable)() {
    err := m.GetBackingStore().Set("trustedCertificateAuthorities", value)
    if err != nil {
        panic(err)
    }
}
// TrustedCertificateAuthorityAsEntityBaseable 
type TrustedCertificateAuthorityAsEntityBaseable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTrustedCertificateAuthorities()([]CertificateAuthorityAsEntityable)
    SetTrustedCertificateAuthorities(value []CertificateAuthorityAsEntityable)()
}
