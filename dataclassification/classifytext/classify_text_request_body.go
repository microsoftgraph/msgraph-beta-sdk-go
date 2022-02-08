package classifytext

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ClassifyTextRequestBody 
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
// NewClassifyTextRequestBody instantiates a new classifyTextRequestBody and sets the default values.
func NewClassifyTextRequestBody()(*ClassifyTextRequestBody) {
    m := &ClassifyTextRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFileExtension gets the fileExtension property value. 
func (m *ClassifyTextRequestBody) GetFileExtension()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileExtension
    }
}
// GetMatchTolerancesToInclude gets the matchTolerancesToInclude property value. 
func (m *ClassifyTextRequestBody) GetMatchTolerancesToInclude()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance) {
    if m == nil {
        return nil
    } else {
        return m.matchTolerancesToInclude
    }
}
// GetScopesToRun gets the scopesToRun property value. 
func (m *ClassifyTextRequestBody) GetScopesToRun()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scopesToRun
    }
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. 
func (m *ClassifyTextRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// GetText gets the text property value. 
func (m *ClassifyTextRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyTextRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fileExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileExtension(val)
        }
        return nil
    }
    res["matchTolerancesToInclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTolerancesToInclude(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance))
        }
        return nil
    }
    res["scopesToRun"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopesToRun(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope))
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
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
func (m *ClassifyTextRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassifyTextRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := (*m.GetMatchTolerancesToInclude()).String()
        err := writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := (*m.GetScopesToRun()).String()
        err := writer.WriteStringValue("scopesToRun", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFileExtension sets the fileExtension property value. 
func (m *ClassifyTextRequestBody) SetFileExtension(value *string)() {
    if m != nil {
        m.fileExtension = value
    }
}
// SetMatchTolerancesToInclude sets the matchTolerancesToInclude property value. 
func (m *ClassifyTextRequestBody) SetMatchTolerancesToInclude(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MlClassificationMatchTolerance)() {
    if m != nil {
        m.matchTolerancesToInclude = value
    }
}
// SetScopesToRun sets the scopesToRun property value. 
func (m *ClassifyTextRequestBody) SetScopesToRun(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitiveTypeScope)() {
    if m != nil {
        m.scopesToRun = value
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. 
func (m *ClassifyTextRequestBody) SetSensitiveTypeIds(value []string)() {
    if m != nil {
        m.sensitiveTypeIds = value
    }
}
// SetText sets the text property value. 
func (m *ClassifyTextRequestBody) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
