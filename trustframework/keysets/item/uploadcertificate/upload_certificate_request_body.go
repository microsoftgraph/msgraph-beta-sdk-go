package uploadcertificate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UploadCertificateRequestBody struct {
    additionalData map[string]interface{};
    key *string;
}
func NewUploadCertificateRequestBody()(*UploadCertificateRequestBody) {
    m := &UploadCertificateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UploadCertificateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UploadCertificateRequestBody) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *UploadCertificateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    return res
}
func (m *UploadCertificateRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *UploadCertificateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *UploadCertificateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UploadCertificateRequestBody) SetKey(value *string)() {
    m.key = value
}
