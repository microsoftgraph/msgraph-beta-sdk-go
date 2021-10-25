package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationJobApplicationParameters struct {
    additionalData map[string]interface{};
    ruleId *string;
    subjects []SynchronizationJobSubject;
}
func NewSynchronizationJobApplicationParameters()(*SynchronizationJobApplicationParameters) {
    m := &SynchronizationJobApplicationParameters{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationJobApplicationParameters) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationJobApplicationParameters) GetRuleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleId
    }
}
func (m *SynchronizationJobApplicationParameters) GetSubjects()([]SynchronizationJobSubject) {
    if m == nil {
        return nil
    } else {
        return m.subjects
    }
}
func (m *SynchronizationJobApplicationParameters) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ruleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleId(val)
        return nil
    }
    res["subjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationJobSubject() })
        if err != nil {
            return err
        }
        res := make([]SynchronizationJobSubject, len(val))
        for i, v := range val {
            res[i] = *(v.(*SynchronizationJobSubject))
        }
        m.SetSubjects(res)
        return nil
    }
    return res
}
func (m *SynchronizationJobApplicationParameters) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationJobApplicationParameters) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("subjects", cast)
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
func (m *SynchronizationJobApplicationParameters) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationJobApplicationParameters) SetRuleId(value *string)() {
    m.ruleId = value
}
func (m *SynchronizationJobApplicationParameters) SetSubjects(value []SynchronizationJobSubject)() {
    m.subjects = value
}
