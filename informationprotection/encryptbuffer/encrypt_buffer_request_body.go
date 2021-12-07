package encryptbuffer

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EncryptBufferRequestBody 
type EncryptBufferRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    buffer []byte;
    // 
    labelId *string;
}
// NewEncryptBufferRequestBody instantiates a new encryptBufferRequestBody and sets the default values.
func NewEncryptBufferRequestBody()(*EncryptBufferRequestBody) {
    m := &EncryptBufferRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptBufferRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBuffer gets the buffer property value. 
func (m *EncryptBufferRequestBody) GetBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.buffer
    }
}
// GetLabelId gets the labelId property value. 
func (m *EncryptBufferRequestBody) GetLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.labelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptBufferRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["buffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuffer(val)
        }
        return nil
    }
    res["labelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
        }
        return nil
    }
    return res
}
func (m *EncryptBufferRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EncryptBufferRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *EncryptBufferRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBuffer sets the buffer property value. 
func (m *EncryptBufferRequestBody) SetBuffer(value []byte)() {
    if m != nil {
        m.buffer = value
    }
}
// SetLabelId sets the labelId property value. 
func (m *EncryptBufferRequestBody) SetLabelId(value *string)() {
    if m != nil {
        m.labelId = value
    }
}
