package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BufferDecryptionResult struct {
    additionalData map[string]interface{};
    decryptedBuffer []byte;
}
func NewBufferDecryptionResult()(*BufferDecryptionResult) {
    m := &BufferDecryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BufferDecryptionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BufferDecryptionResult) GetDecryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.decryptedBuffer
    }
}
func (m *BufferDecryptionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decryptedBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetDecryptedBuffer(val)
        return nil
    }
    return res
}
func (m *BufferDecryptionResult) IsNil()(bool) {
    return m == nil
}
func (m *BufferDecryptionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("decryptedBuffer", m.GetDecryptedBuffer())
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
func (m *BufferDecryptionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BufferDecryptionResult) SetDecryptedBuffer(value []byte)() {
    m.decryptedBuffer = value
}
