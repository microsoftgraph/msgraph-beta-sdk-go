package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type ManagementAction struct {
    Entity
    // The category for the management action. Possible values are: custom, devices, identity, unknownFutureValue. Optional. Read-only.
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory;
    // The description for the management action. Optional. Read-only.
    description *string;
    // The display name for the management action. Optional. Read-only.
    displayName *string;
    // The reference for the management template used to generate the management action. Required. Read-only.
    referenceTemplateId *string;
    // The collection of workload actions associated with the management action. Required. Read-only.
    workloadActions []WorkloadAction;
}
// Instantiates a new managementAction and sets the default values.
func NewManagementAction()(*ManagementAction) {
    m := &ManagementAction{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the category property value. The category for the management action. Possible values are: custom, devices, identity, unknownFutureValue. Optional. Read-only.
func (m *ManagementAction) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) GetReferenceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceTemplateId
    }
}
// Gets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) GetWorkloadActions()([]WorkloadAction) {
    if m == nil {
        return nil
    } else {
        return m.workloadActions
    }
}
// The deserialization information for the current model
func (m *ManagementAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementCategory)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)
        m.SetCategory(&cast)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["referenceTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferenceTemplateId(val)
        return nil
    }
    res["workloadActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadAction() })
        if err != nil {
            return err
        }
        res := make([]WorkloadAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkloadAction))
        }
        m.SetWorkloadActions(res)
        return nil
    }
    return res
}
func (m *ManagementAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := m.GetCategory().String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceTemplateId", m.GetReferenceTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkloadActions()))
        for i, v := range m.GetWorkloadActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("workloadActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the category property value. The category for the management action. Possible values are: custom, devices, identity, unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the category property.
func (m *ManagementAction) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    m.category = value
}
// Sets the description property value. The description for the management action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementAction) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the management action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementAction) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
// Parameters:
//  - value : Value to set for the referenceTemplateId property.
func (m *ManagementAction) SetReferenceTemplateId(value *string)() {
    m.referenceTemplateId = value
}
// Sets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
// Parameters:
//  - value : Value to set for the workloadActions property.
func (m *ManagementAction) SetWorkloadActions(value []WorkloadAction)() {
    m.workloadActions = value
}
