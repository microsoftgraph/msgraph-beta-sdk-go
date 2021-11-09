package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DlpEvaluatePoliciesJobResponse struct {
    JobResponseBase
    // 
    result *DlpPoliciesJobResult;
}
// Instantiates a new dlpEvaluatePoliciesJobResponse and sets the default values.
func NewDlpEvaluatePoliciesJobResponse()(*DlpEvaluatePoliciesJobResponse) {
    m := &DlpEvaluatePoliciesJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// Gets the result property value. 
func (m *DlpEvaluatePoliciesJobResponse) GetResult()(*DlpPoliciesJobResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// The deserialization information for the current model
func (m *DlpEvaluatePoliciesJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDlpPoliciesJobResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*DlpPoliciesJobResult))
        }
        return nil
    }
    return res
}
func (m *DlpEvaluatePoliciesJobResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the result property value. 
// Parameters:
//  - value : Value to set for the result property.
func (m *DlpEvaluatePoliciesJobResponse) SetResult(value *DlpPoliciesJobResult)() {
    m.result = value
}
