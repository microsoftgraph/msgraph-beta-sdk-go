package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagementCertificateWithThumbprint struct {
    additionalData map[string]interface{};
    certificate *string;
    thumbprint *string;
}
func NewManagementCertificateWithThumbprint()(*ManagementCertificateWithThumbprint) {
    m := &ManagementCertificateWithThumbprint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagementCertificateWithThumbprint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagementCertificateWithThumbprint) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
func (m *ManagementCertificateWithThumbprint) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
func (m *ManagementCertificateWithThumbprint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificate(val)
        return nil
    }
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbprint(val)
        return nil
    }
    return res
}
func (m *ManagementCertificateWithThumbprint) IsNil()(bool) {
    return m == nil
}
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
func (m *ManagementCertificateWithThumbprint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagementCertificateWithThumbprint) SetCertificate(value *string)() {
    m.certificate = value
}
func (m *ManagementCertificateWithThumbprint) SetThumbprint(value *string)() {
    m.thumbprint = value
}
