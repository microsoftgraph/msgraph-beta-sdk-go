package classifyexactmatches

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new classifyExactMatchesRequestBody and sets the default values.
func NewClassifyExactMatchesRequestBody()(*ClassifyExactMatchesRequestBody) {
    m := &ClassifyExactMatchesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contentClassifications property value. 
func (m *ClassifyExactMatchesRequestBody) GetContentClassifications()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification) {
    if m == nil {
        return nil
    } else {
        return m.contentClassifications
    }
}
// Gets the sensitiveTypeIds property value. 
func (m *ClassifyExactMatchesRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// Gets the text property value. 
func (m *ClassifyExactMatchesRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Gets the timeoutInMs property value. 
func (m *ClassifyExactMatchesRequestBody) GetTimeoutInMs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeoutInMs
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassifyExactMatchesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ClassifyExactMatchesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contentClassifications property value. 
// Parameters:
//  - value : Value to set for the contentClassifications property.
func (m *ClassifyExactMatchesRequestBody) SetContentClassifications(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentClassification)() {
    m.contentClassifications = value
}
// Sets the sensitiveTypeIds property value. 
// Parameters:
//  - value : Value to set for the sensitiveTypeIds property.
func (m *ClassifyExactMatchesRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
// Sets the text property value. 
// Parameters:
//  - value : Value to set for the text property.
func (m *ClassifyExactMatchesRequestBody) SetText(value *string)() {
    m.text = value
}
// Sets the timeoutInMs property value. 
// Parameters:
//  - value : Value to set for the timeoutInMs property.
func (m *ClassifyExactMatchesRequestBody) SetTimeoutInMs(value *string)() {
    m.timeoutInMs = value
}
