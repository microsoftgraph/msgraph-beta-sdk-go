package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BufferDecryptionResult provides operations to call the decryptBuffer method.
type BufferDecryptionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decryptedBuffer []byte;
}
// NewBufferDecryptionResult instantiates a new bufferDecryptionResult and sets the default values.
func NewBufferDecryptionResult()(*BufferDecryptionResult) {
    m := &BufferDecryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBufferDecryptionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBufferDecryptionResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBufferDecryptionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BufferDecryptionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecryptedBuffer gets the decryptedBuffer property value. 
func (m *BufferDecryptionResult) GetDecryptedBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.decryptedBuffer
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BufferDecryptionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decryptedBuffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.Get[]byteValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecryptedBuffer(val)
        }
        return nil
    }
    return res
}
func (m *BufferDecryptionResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BufferDecryptionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecryptedBuffer sets the decryptedBuffer property value. 
func (m *BufferDecryptionResult) SetDecryptedBuffer(value []byte)() {
    if m != nil {
        m.decryptedBuffer = value
    }
}
