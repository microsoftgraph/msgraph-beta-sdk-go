package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DlpEvaluatePoliciesJobResponse struct {
    JobResponseBase
    result *DlpPoliciesJobResult;
}
func NewDlpEvaluatePoliciesJobResponse()(*DlpEvaluatePoliciesJobResponse) {
    m := &DlpEvaluatePoliciesJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
func (m *DlpEvaluatePoliciesJobResponse) GetResult()(*DlpPoliciesJobResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *DlpEvaluatePoliciesJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDlpPoliciesJobResult() })
        if err != nil {
            return err
        }
        m.SetResult(val.(*DlpPoliciesJobResult))
        return nil
    }
    return res
}
func (m *DlpEvaluatePoliciesJobResponse) IsNil()(bool) {
    return m == nil
}
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
func (m *DlpEvaluatePoliciesJobResponse) SetResult(value *DlpPoliciesJobResult)() {
    m.result = value
}
