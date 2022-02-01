package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalSettings 
type ApprovalSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // One of NoApproval, SingleStage or Serial. The NoApproval is used when isApprovalRequired is false.
    approvalMode *string;
    // If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
    approvalStages []ApprovalStage;
    // If false, then approval is not required for requests in this policy.
    isApprovalRequired *bool;
    // If false, then approval is not required for a user who already has an assignment to extend their assignment.
    isApprovalRequiredForExtension *bool;
    // Indicates whether the requestor is required to supply a justification in their request.
    isRequestorJustificationRequired *bool;
}
// NewApprovalSettings instantiates a new approvalSettings and sets the default values.
func NewApprovalSettings()(*ApprovalSettings) {
    m := &ApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApprovalMode gets the approvalMode property value. One of NoApproval, SingleStage or Serial. The NoApproval is used when isApprovalRequired is false.
func (m *ApprovalSettings) GetApprovalMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalMode
    }
}
// GetApprovalStages gets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
func (m *ApprovalSettings) GetApprovalStages()([]ApprovalStage) {
    if m == nil {
        return nil
    } else {
        return m.approvalStages
    }
}
// GetIsApprovalRequired gets the isApprovalRequired property value. If false, then approval is not required for requests in this policy.
func (m *ApprovalSettings) GetIsApprovalRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequired
    }
}
// GetIsApprovalRequiredForExtension gets the isApprovalRequiredForExtension property value. If false, then approval is not required for a user who already has an assignment to extend their assignment.
func (m *ApprovalSettings) GetIsApprovalRequiredForExtension()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForExtension
    }
}
// GetIsRequestorJustificationRequired gets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
func (m *ApprovalSettings) GetIsRequestorJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequestorJustificationRequired
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvalMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalMode(val)
        }
        return nil
    }
    res["approvalStages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalStage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalStage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ApprovalStage))
            }
            m.SetApprovalStages(res)
        }
        return nil
    }
    res["isApprovalRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequired(val)
        }
        return nil
    }
    res["isApprovalRequiredForExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForExtension(val)
        }
        return nil
    }
    res["isRequestorJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequestorJustificationRequired(val)
        }
        return nil
    }
    return res
}
func (m *ApprovalSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApprovalSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("approvalMode", m.GetApprovalMode())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalStages() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApprovalMode sets the approvalMode property value. One of NoApproval, SingleStage or Serial. The NoApproval is used when isApprovalRequired is false.
func (m *ApprovalSettings) SetApprovalMode(value *string)() {
    if m != nil {
        m.approvalMode = value
    }
}
// SetApprovalStages sets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
func (m *ApprovalSettings) SetApprovalStages(value []ApprovalStage)() {
    if m != nil {
        m.approvalStages = value
    }
}
// SetIsApprovalRequired sets the isApprovalRequired property value. If false, then approval is not required for requests in this policy.
func (m *ApprovalSettings) SetIsApprovalRequired(value *bool)() {
    if m != nil {
        m.isApprovalRequired = value
    }
}
// SetIsApprovalRequiredForExtension sets the isApprovalRequiredForExtension property value. If false, then approval is not required for a user who already has an assignment to extend their assignment.
func (m *ApprovalSettings) SetIsApprovalRequiredForExtension(value *bool)() {
    if m != nil {
        m.isApprovalRequiredForExtension = value
    }
}
// SetIsRequestorJustificationRequired sets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
func (m *ApprovalSettings) SetIsRequestorJustificationRequired(value *bool)() {
    if m != nil {
        m.isRequestorJustificationRequired = value
    }
}
