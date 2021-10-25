package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BufferEncryptionResult struct {
    additionalData map[string]interface{};
    encryptedBuffer []byte;
    publishingLicense []byte;
}
func NewBufferEncryptionResult()(*BufferEncryptionResult) {
    m := &BufferEncryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BufferEncryptionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BufferEncryptionResult) GetEncryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptedBuffer
    }
}
func (m *BufferEncryptionResult) GetPublishingLicense()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.publishingLicense
    }
}
func (m *BufferEncryptionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["encryptedBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetEncryptedBuffer(val)
        return nil
    }
    res["publishingLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetPublishingLicense(val)
        return nil
    }
    return res
}
func (m *BufferEncryptionResult) IsNil()(bool) {
    return m == nil
}
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
func (m *BufferEncryptionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BufferEncryptionResult) SetEncryptedBuffer(value []byte)() {
    m.encryptedBuffer = value
}
func (m *BufferEncryptionResult) SetPublishingLicense(value []byte)() {
    m.publishingLicense = value
}
