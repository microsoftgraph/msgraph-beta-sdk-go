package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassificationJobResponse 
type ClassificationJobResponse struct {
    JobResponseBase
    // 
    result DetectedSensitiveContentWrapperable;
}
// NewClassificationJobResponse instantiates a new classificationJobResponse and sets the default values.
func NewClassificationJobResponse()(*ClassificationJobResponse) {
    m := &ClassificationJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateClassificationJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationJobResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewClassificationJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetectedSensitiveContentWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(DetectedSensitiveContentWrapperable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. 
func (m *ClassificationJobResponse) GetResult()(DetectedSensitiveContentWrapperable) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
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
func (m *ClassificationJobResponse) SetResult(value DetectedSensitiveContentWrapperable)() {
    if m != nil {
        m.result = value
    }
}
