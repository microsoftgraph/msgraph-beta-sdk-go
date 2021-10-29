package classifytext

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type ClassifyTextRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    fileExtension *string;
    // 
    matchTolerancesToInclude *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance;
    // 
    scopesToRun *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope;
    // 
    sensitiveTypeIds []string;
    // 
    text *string;
}
// Instantiates a new classifyTextRequestBody and sets the default values.
func NewClassifyTextRequestBody()(*ClassifyTextRequestBody) {
    m := &ClassifyTextRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the fileExtension property value. 
func (m *ClassifyTextRequestBody) GetFileExtension()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileExtension
    }
}
// Gets the matchTolerancesToInclude property value. 
func (m *ClassifyTextRequestBody) GetMatchTolerancesToInclude()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance) {
    if m == nil {
        return nil
    } else {
        return m.matchTolerancesToInclude
    }
}
// Gets the scopesToRun property value. 
func (m *ClassifyTextRequestBody) GetScopesToRun()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scopesToRun
    }
}
// Gets the sensitiveTypeIds property value. 
func (m *ClassifyTextRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// Gets the text property value. 
func (m *ClassifyTextRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// The deserialization information for the current model
func (m *ClassifyTextRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fileExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileExtension(val)
        return nil
    }
    res["matchTolerancesToInclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance)
        m.SetMatchTolerancesToInclude(&cast)
        return nil
    }
    res["scopesToRun"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope)
        m.SetScopesToRun(&cast)
        return nil
    }
    res["sensitiveTypeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSensitiveTypeIds(res)
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetText(val)
        return nil
    }
    return res
}
func (m *ClassifyTextRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassifyTextRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := m.GetMatchTolerancesToInclude().String()
        err := writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := m.GetScopesToRun().String()
        err := writer.WriteStringValue("scopesToRun", &cast)
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
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *ClassifyTextRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the fileExtension property value. 
// Parameters:
//  - value : Value to set for the fileExtension property.
func (m *ClassifyTextRequestBody) SetFileExtension(value *string)() {
    m.fileExtension = value
}
// Sets the matchTolerancesToInclude property value. 
// Parameters:
//  - value : Value to set for the matchTolerancesToInclude property.
func (m *ClassifyTextRequestBody) SetMatchTolerancesToInclude(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance)() {
    m.matchTolerancesToInclude = value
}
// Sets the scopesToRun property value. 
// Parameters:
//  - value : Value to set for the scopesToRun property.
func (m *ClassifyTextRequestBody) SetScopesToRun(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope)() {
    m.scopesToRun = value
}
// Sets the sensitiveTypeIds property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeIds property.
func (m *ClassifyTextRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
// Sets the text property value. 
// Parameters:
//  - value : Value to set for the text property.
func (m *ClassifyTextRequestBody) SetText(value *string)() {
    m.text = value
}
