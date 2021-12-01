package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassificationJobResponse 
type ClassificationJobResponse struct {
    JobResponseBase
    // 
    result *DetectedSensitiveContentWrapper;
}
// NewClassificationJobResponse instantiates a new classificationJobResponse and sets the default values.
func NewClassificationJobResponse()(*ClassificationJobResponse) {
    m := &ClassificationJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// GetResult gets the result property value. 
func (m *ClassificationJobResponse) GetResult()(*DetectedSensitiveContentWrapper) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedSensitiveContentWrapper() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*DetectedSensitiveContentWrapper))
        }
        return nil
    }
    return res
}
func (m *ClassificationJobResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassificationJobResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.JobResponseBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("result", m.GetResult())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResult sets the result property value. 
func (m *ClassificationJobResponse) SetResult(value *DetectedSensitiveContentWrapper)() {
    if m != nil {
        m.result = value
    }
}
