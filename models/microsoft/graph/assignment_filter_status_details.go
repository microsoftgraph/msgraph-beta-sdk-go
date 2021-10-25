package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterStatusDetails struct {
    additionalData map[string]interface{};
    deviceProperties []KeyValuePair;
    evalutionSummaries []AssignmentFilterEvaluationSummary;
    managedDeviceId *string;
    payloadId *string;
    userId *string;
}
func NewAssignmentFilterStatusDetails()(*AssignmentFilterStatusDetails) {
    m := &AssignmentFilterStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentFilterStatusDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentFilterStatusDetails) GetDeviceProperties()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.deviceProperties
    }
}
func (m *AssignmentFilterStatusDetails) GetEvalutionSummaries()([]AssignmentFilterEvaluationSummary) {
    if m == nil {
        return nil
    } else {
        return m.evalutionSummaries
    }
}
func (m *AssignmentFilterStatusDetails) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
func (m *AssignmentFilterStatusDetails) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
func (m *AssignmentFilterStatusDetails) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *AssignmentFilterStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetDeviceProperties(res)
        return nil
    }
    res["evalutionSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentFilterEvaluationSummary() })
        if err != nil {
            return err
        }
        res := make([]AssignmentFilterEvaluationSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*AssignmentFilterEvaluationSummary))
        }
        m.SetEvalutionSummaries(res)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayloadId(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *AssignmentFilterStatusDetails) IsNil()(bool) {
    return m == nil
}
func (m *AssignmentFilterStatusDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceProperties()))
        for i, v := range m.GetDeviceProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("deviceProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvalutionSummaries()))
        for i, v := range m.GetEvalutionSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("evalutionSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *AssignmentFilterStatusDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentFilterStatusDetails) SetDeviceProperties(value []KeyValuePair)() {
    m.deviceProperties = value
}
func (m *AssignmentFilterStatusDetails) SetEvalutionSummaries(value []AssignmentFilterEvaluationSummary)() {
    m.evalutionSummaries = value
}
func (m *AssignmentFilterStatusDetails) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
func (m *AssignmentFilterStatusDetails) SetPayloadId(value *string)() {
    m.payloadId = value
}
func (m *AssignmentFilterStatusDetails) SetUserId(value *string)() {
    m.userId = value
}
