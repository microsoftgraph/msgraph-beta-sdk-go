package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementCertificateWithThumbprint 
type ManagementCertificateWithThumbprint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Base 64 encoded management certificate
    certificate *string
    // The thumbprint of the management certificate
    thumbprint *string
}
// NewManagementCertificateWithThumbprint instantiates a new managementCertificateWithThumbprint and sets the default values.
func NewManagementCertificateWithThumbprint()(*ManagementCertificateWithThumbprint) {
    m := &ManagementCertificateWithThumbprint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementCertificateWithThumbprintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementCertificateWithThumbprintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementCertificateWithThumbprint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementCertificateWithThumbprint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificate gets the certificate property value. The Base 64 encoded management certificate
func (m *ManagementCertificateWithThumbprint) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementCertificateWithThumbprint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    return res
}
// GetThumbprint gets the thumbprint property value. The thumbprint of the management certificate
func (m *ManagementCertificateWithThumbprint) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// Serialize serializes information the current object
func (m *ManagementCertificateWithThumbprint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementCertificateWithThumbprint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificate sets the certificate property value. The Base 64 encoded management certificate
func (m *ManagementCertificateWithThumbprint) SetCertificate(value *string)() {
    if m != nil {
        m.certificate = value
    }
}
// SetThumbprint sets the thumbprint property value. The thumbprint of the management certificate
func (m *ManagementCertificateWithThumbprint) SetThumbprint(value *string)() {
    if m != nil {
        m.thumbprint = value
    }
}
