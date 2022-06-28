package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientCertificateAuthentication 
type ClientCertificateAuthentication struct {
    ApiAuthenticationConfigurationBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of certificates uploaded for this API connector.
    certificateList []Pkcs12CertificateInformationable
}
// NewClientCertificateAuthentication instantiates a new ClientCertificateAuthentication and sets the default values.
func NewClientCertificateAuthentication()(*ClientCertificateAuthentication) {
    m := &ClientCertificateAuthentication{
        ApiAuthenticationConfigurationBase: *NewApiAuthenticationConfigurationBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClientCertificateAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientCertificateAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientCertificateAuthentication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClientCertificateAuthentication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateList gets the certificateList property value. The list of certificates uploaded for this API connector.
func (m *ClientCertificateAuthentication) GetCertificateList()([]Pkcs12CertificateInformationable) {
    if m == nil {
        return nil
    } else {
        return m.certificateList
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientCertificateAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApiAuthenticationConfigurationBase.GetFieldDeserializers()
    res["certificateList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePkcs12CertificateInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pkcs12CertificateInformationable, len(val))
            for i, v := range val {
                res[i] = v.(Pkcs12CertificateInformationable)
            }
            m.SetCertificateList(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ClientCertificateAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApiAuthenticationConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateList()))
        for i, v := range m.GetCertificateList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateList", cast)
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
func (m *ClientCertificateAuthentication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificateList sets the certificateList property value. The list of certificates uploaded for this API connector.
func (m *ClientCertificateAuthentication) SetCertificateList(value []Pkcs12CertificateInformationable)() {
    if m != nil {
        m.certificateList = value
    }
}
