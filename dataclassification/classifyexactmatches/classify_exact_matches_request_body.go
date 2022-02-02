package classifyexactmatches

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ClassifyExactMatchesRequestBody 
type ClassifyExactMatchesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    contentClassifications []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification;
    // 
    sensitiveTypeIds []string;
    // 
    text *string;
    // 
    timeoutInMs *string;
}
// NewClassifyExactMatchesRequestBody instantiates a new classifyExactMatchesRequestBody and sets the default values.
func NewClassifyExactMatchesRequestBody()(*ClassifyExactMatchesRequestBody) {
    m := &ClassifyExactMatchesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentClassifications gets the contentClassifications property value. 
func (m *ClassifyExactMatchesRequestBody) GetContentClassifications()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification) {
    if m == nil {
        return nil
    } else {
        return m.contentClassifications
    }
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. 
func (m *ClassifyExactMatchesRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// GetText gets the text property value. 
func (m *ClassifyExactMatchesRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetTimeoutInMs gets the timeoutInMs property value. 
func (m *ClassifyExactMatchesRequestBody) GetTimeoutInMs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeoutInMs
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyExactMatchesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentClassifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentClassification() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification))
            }
            m.SetContentClassifications(res)
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
    res["timeoutInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeoutInMs(val)
        }
        return nil
    }
    return res
}
func (m *ClassifyExactMatchesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassifyExactMatchesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetContentClassifications() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContentClassifications()))
        for i, v := range m.GetContentClassifications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("contentClassifications", cast)
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
        err := writer.WriteStringValue("timeoutInMs", m.GetTimeoutInMs())
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
func (m *ClassifyExactMatchesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentClassifications sets the contentClassifications property value. 
func (m *ClassifyExactMatchesRequestBody) SetContentClassifications(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification)() {
    if m != nil {
        m.contentClassifications = value
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. 
func (m *ClassifyExactMatchesRequestBody) SetSensitiveTypeIds(value []string)() {
    if m != nil {
        m.sensitiveTypeIds = value
    }
}
// SetText sets the text property value. 
func (m *ClassifyExactMatchesRequestBody) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
// SetTimeoutInMs sets the timeoutInMs property value. 
func (m *ClassifyExactMatchesRequestBody) SetTimeoutInMs(value *string)() {
    if m != nil {
        m.timeoutInMs = value
    }
}
