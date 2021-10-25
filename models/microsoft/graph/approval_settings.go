package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApprovalSettings struct {
    additionalData map[string]interface{};
    approvalMode *string;
    approvalStages []ApprovalStage;
    isApprovalRequired *bool;
    isApprovalRequiredForExtension *bool;
    isRequestorJustificationRequired *bool;
}
func NewApprovalSettings()(*ApprovalSettings) {
    m := &ApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApprovalSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApprovalSettings) GetApprovalMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalMode
    }
}
func (m *ApprovalSettings) GetApprovalStages()([]ApprovalStage) {
    if m == nil {
        return nil
    } else {
        return m.approvalStages
    }
}
func (m *ApprovalSettings) GetIsApprovalRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequired
    }
}
func (m *ApprovalSettings) GetIsApprovalRequiredForExtension()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForExtension
    }
}
func (m *ApprovalSettings) GetIsRequestorJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequestorJustificationRequired
    }
}
func (m *ApprovalSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvalMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApprovalMode(val)
        return nil
    }
    res["approvalStages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalStage() })
        if err != nil {
            return err
        }
        res := make([]ApprovalStage, len(val))
        for i, v := range val {
            res[i] = *(v.(*ApprovalStage))
        }
        m.SetApprovalStages(res)
        return nil
    }
    res["isApprovalRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApprovalRequired(val)
        return nil
    }
    res["isApprovalRequiredForExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApprovalRequiredForExtension(val)
        return nil
    }
    res["isRequestorJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRequestorJustificationRequired(val)
        return nil
    }
    return res
}
func (m *ApprovalSettings) IsNil()(bool) {
    return m == nil
}
func (m *ApprovalSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("approvalMode", m.GetApprovalMode())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApprovalStages()))
        for i, v := range m.GetApprovalStages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("approvalStages", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequired", m.GetIsApprovalRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForExtension", m.GetIsApprovalRequiredForExtension())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequestorJustificationRequired", m.GetIsRequestorJustificationRequired())
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
func (m *ApprovalSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApprovalSettings) SetApprovalMode(value *string)() {
    m.approvalMode = value
}
func (m *ApprovalSettings) SetApprovalStages(value []ApprovalStage)() {
    m.approvalStages = value
}
func (m *ApprovalSettings) SetIsApprovalRequired(value *bool)() {
    m.isApprovalRequired = value
}
func (m *ApprovalSettings) SetIsApprovalRequiredForExtension(value *bool)() {
    m.isApprovalRequiredForExtension = value
}
func (m *ApprovalSettings) SetIsRequestorJustificationRequired(value *bool)() {
    m.isRequestorJustificationRequired = value
}
