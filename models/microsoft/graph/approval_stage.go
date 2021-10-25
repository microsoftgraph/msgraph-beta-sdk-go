package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApprovalStage struct {
    additionalData map[string]interface{};
    approvalStageTimeOutInDays *int32;
    escalationApprovers []UserSet;
    escalationTimeInMinutes *int32;
    isApproverJustificationRequired *bool;
    isEscalationEnabled *bool;
    primaryApprovers []UserSet;
}
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApprovalStage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApprovalStage) GetApprovalStageTimeOutInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.approvalStageTimeOutInDays
    }
}
func (m *ApprovalStage) GetEscalationApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.escalationApprovers
    }
}
func (m *ApprovalStage) GetEscalationTimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.escalationTimeInMinutes
    }
}
func (m *ApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApproverJustificationRequired
    }
}
func (m *ApprovalStage) GetIsEscalationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEscalationEnabled
    }
}
func (m *ApprovalStage) GetPrimaryApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.primaryApprovers
    }
}
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvalStageTimeOutInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetApprovalStageTimeOutInDays(val)
        return nil
    }
    res["escalationApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetEscalationApprovers(res)
        return nil
    }
    res["escalationTimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEscalationTimeInMinutes(val)
        return nil
    }
    res["isApproverJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApproverJustificationRequired(val)
        return nil
    }
    res["isEscalationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEscalationEnabled(val)
        return nil
    }
    res["primaryApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetPrimaryApprovers(res)
        return nil
    }
    return res
}
func (m *ApprovalStage) IsNil()(bool) {
    return m == nil
}
func (m *ApprovalStage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("approvalStageTimeOutInDays", m.GetApprovalStageTimeOutInDays())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEscalationApprovers()))
        for i, v := range m.GetEscalationApprovers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("escalationApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("escalationTimeInMinutes", m.GetEscalationTimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApproverJustificationRequired", m.GetIsApproverJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEscalationEnabled", m.GetIsEscalationEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrimaryApprovers()))
        for i, v := range m.GetPrimaryApprovers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("primaryApprovers", cast)
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
func (m *ApprovalStage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApprovalStage) SetApprovalStageTimeOutInDays(value *int32)() {
    m.approvalStageTimeOutInDays = value
}
func (m *ApprovalStage) SetEscalationApprovers(value []UserSet)() {
    m.escalationApprovers = value
}
func (m *ApprovalStage) SetEscalationTimeInMinutes(value *int32)() {
    m.escalationTimeInMinutes = value
}
func (m *ApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    m.isApproverJustificationRequired = value
}
func (m *ApprovalStage) SetIsEscalationEnabled(value *bool)() {
    m.isEscalationEnabled = value
}
func (m *ApprovalStage) SetPrimaryApprovers(value []UserSet)() {
    m.primaryApprovers = value
}
