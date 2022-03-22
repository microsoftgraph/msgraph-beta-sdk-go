package classifyfile

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassifyFileRequestBody provides operations to call the classifyFile method.
type ClassifyFileRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    file []byte;
    // 
    sensitiveTypeIds []string;
}
// NewClassifyFileRequestBody instantiates a new classifyFileRequestBody and sets the default values.
func NewClassifyFileRequestBody()(*ClassifyFileRequestBody) {
    m := &ClassifyFileRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassifyFileRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassifyFileRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewClassifyFileRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyFileRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyFileRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. 
func (m *ClassifyFileRequestBody) GetFile()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. 
func (m *ClassifyFileRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// Serialize serializes information the current object
func (m *ClassifyFileRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
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
func (m *ClassifyFileRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFile sets the file property value. 
func (m *ClassifyFileRequestBody) SetFile(value []byte)() {
    if m != nil {
        m.file = value
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. 
func (m *ClassifyFileRequestBody) SetSensitiveTypeIds(value []string)() {
    if m != nil {
        m.sensitiveTypeIds = value
    }
}
