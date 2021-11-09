package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BufferEncryptionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    encryptedBuffer []byte;
    // 
    publishingLicense []byte;
}
// Instantiates a new bufferEncryptionResult and sets the default values.
func NewBufferEncryptionResult()(*BufferEncryptionResult) {
    m := &BufferEncryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BufferEncryptionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the encryptedBuffer property value. 
func (m *BufferEncryptionResult) GetEncryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptedBuffer
    }
}
// Gets the publishingLicense property value. 
func (m *BufferEncryptionResult) GetPublishingLicense()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.publishingLicense
    }
}
// The deserialization information for the current model
func (m *BufferEncryptionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *BufferEncryptionResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BufferEncryptionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *BufferEncryptionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the encryptedBuffer property value. 
// Parameters:
//  - value : Value to set for the encryptedBuffer property.
func (m *BufferEncryptionResult) SetEncryptedBuffer(value []byte)() {
    m.encryptedBuffer = value
}
// Sets the publishingLicense property value. 
// Parameters:
//  - value : Value to set for the publishingLicense property.
func (m *BufferEncryptionResult) SetPublishingLicense(value []byte)() {
    m.publishingLicense = value
}
