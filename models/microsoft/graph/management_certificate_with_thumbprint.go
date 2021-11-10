package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementCertificateWithThumbprint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Base 64 encoded management certificate
    certificate *string;
    // The thumbprint of the management certificate
    thumbprint *string;
}
// Instantiates a new managementCertificateWithThumbprint and sets the default values.
func NewManagementCertificateWithThumbprint()(*ManagementCertificateWithThumbprint) {
    m := &ManagementCertificateWithThumbprint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementCertificateWithThumbprint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certificate property value. The Base 64 encoded management certificate
func (m *ManagementCertificateWithThumbprint) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// Gets the thumbprint property value. The thumbprint of the management certificate
func (m *ManagementCertificateWithThumbprint) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// The deserialization information for the current model
func (m *ManagementCertificateWithThumbprint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ManagementCertificateWithThumbprint) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementCertificateWithThumbprint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagementCertificateWithThumbprint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certificate property value. The Base 64 encoded management certificate
// Parameters:
//  - value : Value to set for the certificate property.
func (m *ManagementCertificateWithThumbprint) SetCertificate(value *string)() {
    m.certificate = value
}
// Sets the thumbprint property value. The thumbprint of the management certificate
// Parameters:
//  - value : Value to set for the thumbprint property.
func (m *ManagementCertificateWithThumbprint) SetThumbprint(value *string)() {
    m.thumbprint = value
}
