package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateBasedAuthConfiguration provides operations to manage the collection of accessReview entities.
type CertificateBasedAuthConfiguration struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of certificate authorities which creates a trusted certificate chain.
    certificateAuthorities []CertificateAuthorityable
}
// NewCertificateBasedAuthConfiguration instantiates a new certificateBasedAuthConfiguration and sets the default values.
func NewCertificateBasedAuthConfiguration()(*CertificateBasedAuthConfiguration) {
    m := &CertificateBasedAuthConfiguration{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCertificateBasedAuthConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateBasedAuthConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateBasedAuthConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateBasedAuthConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateAuthorities gets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) GetCertificateAuthorities()([]CertificateAuthorityable) {
    if m == nil {
        return nil
    } else {
        return m.certificateAuthorities
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateBasedAuthConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateAuthorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateAuthorityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateAuthorityable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateAuthorityable)
            }
            m.SetCertificateAuthorities(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CertificateBasedAuthConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateAuthorities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateAuthorities()))
        for i, v := range m.GetCertificateAuthorities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateAuthorities", cast)
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
func (m *CertificateBasedAuthConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificateAuthorities sets the certificateAuthorities property value. Collection of certificate authorities which creates a trusted certificate chain.
func (m *CertificateBasedAuthConfiguration) SetCertificateAuthorities(value []CertificateAuthorityable)() {
    if m != nil {
        m.certificateAuthorities = value
    }
}
