package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidEnrollmentCompanyCode 
type AndroidEnrollmentCompanyCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Enrollment Token used by the User to enroll their device.
    enrollmentToken *string;
    // String used to generate a QR code for the token.
    qrCodeContent *string;
    // Generated QR code for the token.
    qrCodeImage *MimeContent;
}
// NewAndroidEnrollmentCompanyCode instantiates a new androidEnrollmentCompanyCode and sets the default values.
func NewAndroidEnrollmentCompanyCode()(*AndroidEnrollmentCompanyCode) {
    m := &AndroidEnrollmentCompanyCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidEnrollmentCompanyCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnrollmentToken gets the enrollmentToken property value. Enrollment Token used by the User to enroll their device.
func (m *AndroidEnrollmentCompanyCode) GetEnrollmentToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentToken
    }
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidEnrollmentCompanyCode) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
// GetQrCodeImage gets the qrCodeImage property value. Generated QR code for the token.
func (m *AndroidEnrollmentCompanyCode) GetQrCodeImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidEnrollmentCompanyCode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enrollmentToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentToken(val)
        }
        return nil
    }
    res["qrCodeContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeContent(val)
        }
        return nil
    }
    res["qrCodeImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeImage(val.(*MimeContent))
        }
        return nil
    }
    return res
}
func (m *AndroidEnrollmentCompanyCode) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidEnrollmentCompanyCode) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnrollmentToken sets the enrollmentToken property value. Enrollment Token used by the User to enroll their device.
func (m *AndroidEnrollmentCompanyCode) SetEnrollmentToken(value *string)() {
    if m != nil {
        m.enrollmentToken = value
    }
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidEnrollmentCompanyCode) SetQrCodeContent(value *string)() {
    if m != nil {
        m.qrCodeContent = value
    }
}
// SetQrCodeImage sets the qrCodeImage property value. Generated QR code for the token.
func (m *AndroidEnrollmentCompanyCode) SetQrCodeImage(value *MimeContent)() {
    if m != nil {
        m.qrCodeImage = value
    }
}
