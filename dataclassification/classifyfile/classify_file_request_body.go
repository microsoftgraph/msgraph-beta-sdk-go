package classifyfile

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ClassifyFileRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    file []byte;
    // 
    sensitiveTypeIds []string;
}
// Instantiates a new classifyFileRequestBody and sets the default values.
func NewClassifyFileRequestBody()(*ClassifyFileRequestBody) {
    m := &ClassifyFileRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyFileRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the file property value. 
func (m *ClassifyFileRequestBody) GetFile()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// Gets the sensitiveTypeIds property value. 
func (m *ClassifyFileRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// The deserialization information for the current model
func (m *ClassifyFileRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetFile(val)
        return nil
    }
    res["sensitiveTypeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSensitiveTypeIds(res)
        return nil
    }
    return res
}
func (m *ClassifyFileRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassifyFileRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ClassifyFileRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the file property value. 
// Parameters:
//  - value : Value to set for the file property.
func (m *ClassifyFileRequestBody) SetFile(value []byte)() {
    m.file = value
}
// Sets the sensitiveTypeIds property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeIds property.
func (m *ClassifyFileRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
