package decryptbuffer

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DecryptBufferRequestBody 
type DecryptBufferRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    encryptedBuffer []byte;
    // 
    publishingLicense []byte;
}
// NewDecryptBufferRequestBody instantiates a new decryptBufferRequestBody and sets the default values.
func NewDecryptBufferRequestBody()(*DecryptBufferRequestBody) {
    m := &DecryptBufferRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DecryptBufferRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEncryptedBuffer gets the encryptedBuffer property value. 
func (m *DecryptBufferRequestBody) GetEncryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptedBuffer
    }
}
// GetPublishingLicense gets the publishingLicense property value. 
func (m *DecryptBufferRequestBody) GetPublishingLicense()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.publishingLicense
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DecryptBufferRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["encryptedBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedBuffer(val)
        }
        return nil
    }
    res["publishingLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingLicense(val)
        }
        return nil
    }
    return res
}
func (m *DecryptBufferRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DecryptBufferRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("encryptedBuffer", m.GetEncryptedBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("publishingLicense", m.GetPublishingLicense())
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
func (m *DecryptBufferRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEncryptedBuffer sets the encryptedBuffer property value. 
func (m *DecryptBufferRequestBody) SetEncryptedBuffer(value []byte)() {
    if m != nil {
        m.encryptedBuffer = value
    }
}
// SetPublishingLicense sets the publishingLicense property value. 
func (m *DecryptBufferRequestBody) SetPublishingLicense(value []byte)() {
    if m != nil {
        m.publishingLicense = value
    }
}
