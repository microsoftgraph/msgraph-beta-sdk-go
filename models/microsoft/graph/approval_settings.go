package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new approvalSettings and sets the default values.
func NewApprovalSettings()(*ApprovalSettings) {
    m := &ApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the approvalMode property value. One of NoApproval, SingleStage or Serial. The NoApproval is used when isApprovalRequired is false.
func (m *ApprovalSettings) GetApprovalMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalMode
    }
}
// Gets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
func (m *ApprovalSettings) GetApprovalStages()([]ApprovalStage) {
    if m == nil {
        return nil
    } else {
        return m.approvalStages
    }
}
// Gets the isApprovalRequired property value. If false, then approval is not required for requests in this policy.
func (m *ApprovalSettings) GetIsApprovalRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequired
    }
}
// Gets the isApprovalRequiredForExtension property value. If false, then approval is not required for a user who already has an assignment to extend their assignment.
func (m *ApprovalSettings) GetIsApprovalRequiredForExtension()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequiredForExtension
    }
}
// Gets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
func (m *ApprovalSettings) GetIsRequestorJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequestorJustificationRequired
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ApprovalSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the approvalMode property value. One of NoApproval, SingleStage or Serial. The NoApproval is used when isApprovalRequired is false.
// Parameters:
//  - value : Value to set for the approvalMode property.
func (m *ApprovalSettings) SetApprovalMode(value *string)() {
    m.approvalMode = value
}
// Sets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
// Parameters:
//  - value : Value to set for the approvalStages property.
func (m *ApprovalSettings) SetApprovalStages(value []ApprovalStage)() {
    m.approvalStages = value
}
// Sets the isApprovalRequired property value. If false, then approval is not required for requests in this policy.
// Parameters:
//  - value : Value to set for the isApprovalRequired property.
func (m *ApprovalSettings) SetIsApprovalRequired(value *bool)() {
    m.isApprovalRequired = value
}
// Sets the isApprovalRequiredForExtension property value. If false, then approval is not required for a user who already has an assignment to extend their assignment.
// Parameters:
//  - value : Value to set for the isApprovalRequiredForExtension property.
func (m *ApprovalSettings) SetIsApprovalRequiredForExtension(value *bool)() {
    m.isApprovalRequiredForExtension = value
}
// Sets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
// Parameters:
//  - value : Value to set for the isRequestorJustificationRequired property.
func (m *ApprovalSettings) SetIsRequestorJustificationRequired(value *bool)() {
    m.isRequestorJustificationRequired = value
}
