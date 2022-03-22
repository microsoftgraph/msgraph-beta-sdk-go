package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateLabelJobResponse 
type EvaluateLabelJobResponse struct {
    JobResponseBase
    // 
    result EvaluateLabelJobResultGroupable;
}
// NewEvaluateLabelJobResponse instantiates a new evaluateLabelJobResponse and sets the default values.
func NewEvaluateLabelJobResponse()(*EvaluateLabelJobResponse) {
    m := &EvaluateLabelJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateEvaluateLabelJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateLabelJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(EvaluateLabelJobResultGroupable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. 
func (m *EvaluateLabelJobResponse) GetResult()(EvaluateLabelJobResultGroupable) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *EvaluateLabelJobResponse) SetResult(value EvaluateLabelJobResultGroupable)() {
    if m != nil {
        m.result = value
    }
}
