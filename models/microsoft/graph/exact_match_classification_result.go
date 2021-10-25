package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExactMatchClassificationResult struct {
    additionalData map[string]interface{};
    classification []ExactMatchDetectedSensitiveContent;
    errors []ClassificationError;
}
func NewExactMatchClassificationResult()(*ExactMatchClassificationResult) {
    m := &ExactMatchClassificationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExactMatchClassificationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExactMatchClassificationResult) GetClassification()([]ExactMatchDetectedSensitiveContent) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
func (m *ExactMatchClassificationResult) GetErrors()([]ClassificationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
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
func (m *ExactMatchClassificationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExactMatchClassificationResult) SetClassification(value []ExactMatchDetectedSensitiveContent)() {
    m.classification = value
}
func (m *ExactMatchClassificationResult) SetErrors(value []ClassificationError)() {
    m.errors = value
}
