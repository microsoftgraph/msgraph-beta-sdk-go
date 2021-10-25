package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EvaluateLabelJobResponse struct {
    JobResponseBase
    result *EvaluateLabelJobResultGroup;
}
func NewEvaluateLabelJobResponse()(*EvaluateLabelJobResponse) {
    m := &EvaluateLabelJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
func (m *EvaluateLabelJobResponse) GetResult()(*EvaluateLabelJobResultGroup) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *EvaluateLabelJobResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateLabelJobResultGroup() })
        if err != nil {
            return err
        }
        m.SetResult(val.(*EvaluateLabelJobResultGroup))
        return nil
    }
    return res
}
func (m *EvaluateLabelJobResponse) IsNil()(bool) {
    return m == nil
}
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
func (m *EvaluateLabelJobResponse) SetResult(value *EvaluateLabelJobResultGroup)() {
    m.result = value
}
