package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchClassificationResult 
type ExactMatchClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    classification []ExactMatchDetectedSensitiveContent;
    // 
    errors []ClassificationError;
}
// NewExactMatchClassificationResult instantiates a new exactMatchClassificationResult and sets the default values.
func NewExactMatchClassificationResult()(*ExactMatchClassificationResult) {
    m := &ExactMatchClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactMatchClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClassification gets the classification property value. 
func (m *ExactMatchClassificationResult) GetClassification()([]ExactMatchDetectedSensitiveContent) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetErrors gets the errors property value. 
func (m *ExactMatchClassificationResult) GetErrors()([]ClassificationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchClassificationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchDetectedSensitiveContent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchDetectedSensitiveContent, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExactMatchDetectedSensitiveContent))
            }
            m.SetClassification(res)
        }
        return nil
    }
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationError, len(val))
            for i, v := range val {
                res[i] = *(v.(*ClassificationError))
            }
            m.SetErrors(res)
        }
        return nil
    }
    return res
}
func (m *ExactMatchClassificationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExactMatchClassificationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetClassification() != nil {
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
    if m.GetErrors() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactMatchClassificationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClassification sets the classification property value. 
func (m *ExactMatchClassificationResult) SetClassification(value []ExactMatchDetectedSensitiveContent)() {
    if m != nil {
        m.classification = value
    }
}
// SetErrors sets the errors property value. 
func (m *ExactMatchClassificationResult) SetErrors(value []ClassificationError)() {
    if m != nil {
        m.errors = value
    }
}
