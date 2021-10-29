package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ExactMatchClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    classification []ExactMatchDetectedSensitiveContent;
    // 
    errors []ClassificationError;
}
// Instantiates a new exactMatchClassificationResult and sets the default values.
func NewExactMatchClassificationResult()(*ExactMatchClassificationResult) {
    m := &ExactMatchClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactMatchClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the classification property value. 
func (m *ExactMatchClassificationResult) GetClassification()([]ExactMatchDetectedSensitiveContent) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// Gets the errors property value. 
func (m *ExactMatchClassificationResult) GetErrors()([]ClassificationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// The deserialization information for the current model
func (m *ExactMatchClassificationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchDetectedSensitiveContent() })
        if err != nil {
            return err
        }
        res := make([]ExactMatchDetectedSensitiveContent, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExactMatchDetectedSensitiveContent))
        }
        m.SetClassification(res)
        return nil
    }
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationError() })
        if err != nil {
            return err
        }
        res := make([]ClassificationError, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClassificationError))
        }
        m.SetErrors(res)
        return nil
    }
    return res
}
func (m *ExactMatchClassificationResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExactMatchClassificationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClassification()))
        for i, v := range m.GetClassification() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("classification", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("errors", cast)
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
func (m *ExactMatchClassificationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the classification property value. 
// Parameters:
//  - value : Value to set for the classification property.
func (m *ExactMatchClassificationResult) SetClassification(value []ExactMatchDetectedSensitiveContent)() {
    m.classification = value
}
// Sets the errors property value. 
// Parameters:
//  - value : Value to set for the errors property.
func (m *ExactMatchClassificationResult) SetErrors(value []ClassificationError)() {
    m.errors = value
}
