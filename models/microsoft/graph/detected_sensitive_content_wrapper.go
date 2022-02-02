package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedSensitiveContentWrapper 
type DetectedSensitiveContentWrapper struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    classification []DetectedSensitiveContent;
}
// NewDetectedSensitiveContentWrapper instantiates a new detectedSensitiveContentWrapper and sets the default values.
func NewDetectedSensitiveContentWrapper()(*DetectedSensitiveContentWrapper) {
    m := &DetectedSensitiveContentWrapper{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentWrapper) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClassification gets the classification property value. 
func (m *DetectedSensitiveContentWrapper) GetClassification()([]DetectedSensitiveContent) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContentWrapper) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedSensitiveContent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedSensitiveContent, len(val))
            for i, v := range val {
                res[i] = *(v.(*DetectedSensitiveContent))
            }
            m.SetClassification(res)
        }
        return nil
    }
    return res
}
func (m *DetectedSensitiveContentWrapper) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContentWrapper) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetectedSensitiveContentWrapper) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClassification sets the classification property value. 
func (m *DetectedSensitiveContentWrapper) SetClassification(value []DetectedSensitiveContent)() {
    if m != nil {
        m.classification = value
    }
}
