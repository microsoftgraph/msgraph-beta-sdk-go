package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementAction 
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
    // 
    referenceTemplateVersion *int32;
    // The collection of workload actions associated with the management action. Required. Read-only.
    workloadActions []WorkloadAction;
}
// NewManagementAction instantiates a new managementAction and sets the default values.
func NewManagementAction()(*ManagementAction) {
    m := &ManagementAction{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. The category for the management action. Possible values are: custom, devices, identity, unknownFutureValue. Optional. Read-only.
func (m *ManagementAction) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDescription gets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetReferenceTemplateId gets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) GetReferenceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceTemplateId
    }
}
// GetReferenceTemplateVersion gets the referenceTemplateVersion property value. 
func (m *ManagementAction) GetReferenceTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referenceTemplateVersion
    }
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) GetWorkloadActions()([]WorkloadAction) {
    if m == nil {
        return nil
    } else {
        return m.workloadActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["referenceTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateId(val)
        }
        return nil
    }
    res["referenceTemplateVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateVersion(val)
        }
        return nil
    }
    res["workloadActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkloadAction))
            }
            m.SetWorkloadActions(res)
        }
        return nil
    }
    return res
}
func (m *ManagementAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
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
        err = writer.WriteInt32Value("referenceTemplateVersion", m.GetReferenceTemplateVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadActions() != nil {
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
// SetCategory sets the category property value. The category for the management action. Possible values are: custom, devices, identity, unknownFutureValue. Optional. Read-only.
func (m *ManagementAction) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDescription sets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetReferenceTemplateId sets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) SetReferenceTemplateId(value *string)() {
    if m != nil {
        m.referenceTemplateId = value
    }
}
// SetReferenceTemplateVersion sets the referenceTemplateVersion property value. 
func (m *ManagementAction) SetReferenceTemplateVersion(value *int32)() {
    if m != nil {
        m.referenceTemplateVersion = value
    }
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) SetWorkloadActions(value []WorkloadAction)() {
    if m != nil {
        m.workloadActions = value
    }
}
