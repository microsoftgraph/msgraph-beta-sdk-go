package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DlpEvaluatePoliciesJobResponse provides operations to call the evaluate method.
type DlpEvaluatePoliciesJobResponse struct {
    JobResponseBase
    // 
    result DlpPoliciesJobResultable;
}
// NewDlpEvaluatePoliciesJobResponse instantiates a new dlpEvaluatePoliciesJobResponse and sets the default values.
func NewDlpEvaluatePoliciesJobResponse()(*DlpEvaluatePoliciesJobResponse) {
    m := &DlpEvaluatePoliciesJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateDlpEvaluatePoliciesJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpEvaluatePoliciesJobResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDlpEvaluatePoliciesJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluatePoliciesJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDlpPoliciesJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(DlpPoliciesJobResultable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. 
func (m *DlpEvaluatePoliciesJobResponse) GetResult()(DlpPoliciesJobResultable) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *DlpEvaluatePoliciesJobResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DlpEvaluatePoliciesJobResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *DlpEvaluatePoliciesJobResponse) SetResult(value DlpPoliciesJobResultable)() {
    if m != nil {
        m.result = value
    }
}
