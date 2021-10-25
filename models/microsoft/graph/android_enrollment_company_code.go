package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AndroidEnrollmentCompanyCode struct {
    additionalData map[string]interface{};
    enrollmentToken *string;
    qrCodeContent *string;
    qrCodeImage *MimeContent;
}
func NewAndroidEnrollmentCompanyCode()(*AndroidEnrollmentCompanyCode) {
    m := &AndroidEnrollmentCompanyCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AndroidEnrollmentCompanyCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AndroidEnrollmentCompanyCode) GetEnrollmentToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentToken
    }
}
func (m *AndroidEnrollmentCompanyCode) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
func (m *AndroidEnrollmentCompanyCode) GetQrCodeImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
func (m *AndroidEnrollmentCompanyCode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enrollmentToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentToken(val)
        return nil
    }
    res["qrCodeContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQrCodeContent(val)
        return nil
    }
    res["qrCodeImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        m.SetQrCodeImage(val.(*MimeContent))
        return nil
    }
    return res
}
func (m *AndroidEnrollmentCompanyCode) IsNil()(bool) {
    return m == nil
}
func (m *AndroidEnrollmentCompanyCode) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enrollmentToken", m.GetEnrollmentToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("qrCodeContent", m.GetQrCodeContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("qrCodeImage", m.GetQrCodeImage())
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
func (m *AndroidEnrollmentCompanyCode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AndroidEnrollmentCompanyCode) SetEnrollmentToken(value *string)() {
    m.enrollmentToken = value
}
func (m *AndroidEnrollmentCompanyCode) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
func (m *AndroidEnrollmentCompanyCode) SetQrCodeImage(value *MimeContent)() {
    m.qrCodeImage = value
}
