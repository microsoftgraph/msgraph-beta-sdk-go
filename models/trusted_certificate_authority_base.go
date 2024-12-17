package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrustedCertificateAuthorityBase struct {
    DirectoryObject
}
// NewTrustedCertificateAuthorityBase instantiates a new TrustedCertificateAuthorityBase and sets the default values.
func NewTrustedCertificateAuthorityBase()(*TrustedCertificateAuthorityBase) {
    m := &TrustedCertificateAuthorityBase{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.trustedCertificateAuthorityBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTrustedCertificateAuthorityBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrustedCertificateAuthorityBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.mutualTlsOauthConfiguration":
                        return NewMutualTlsOauthConfiguration(), nil
                }
            }
        }
    }
    return NewTrustedCertificateAuthorityBase(), nil
}
// GetCertificateAuthorities gets the certificateAuthorities property value. The certificateAuthorities property
// returns a []CertificateAuthorityable when successful
func (m *TrustedCertificateAuthorityBase) GetCertificateAuthorities()([]CertificateAuthorityable) {
    val, err := m.GetBackingStore().Get("certificateAuthorities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateAuthorityable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrustedCertificateAuthorityBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["certificateAuthorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateAuthorityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateAuthorityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CertificateAuthorityable)
                }
            }
            m.SetCertificateAuthorities(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TrustedCertificateAuthorityBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateAuthorities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateAuthorities()))
        for i, v := range m.GetCertificateAuthorities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("certificateAuthorities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAuthorities sets the certificateAuthorities property value. The certificateAuthorities property
func (m *TrustedCertificateAuthorityBase) SetCertificateAuthorities(value []CertificateAuthorityable)() {
    err := m.GetBackingStore().Set("certificateAuthorities", value)
    if err != nil {
        panic(err)
    }
}
type TrustedCertificateAuthorityBaseable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAuthorities()([]CertificateAuthorityable)
    SetCertificateAuthorities(value []CertificateAuthorityable)()
}
