package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type ManagementTemplate struct {
    Entity
    // The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory;
    // The description for the management template. Optional. Read-only.
    description *string;
    // The display name for the management template. Required. Read-only.
    displayName *string;
    // 
    managementTemplateCollections []ManagementTemplateCollection;
    // 
    managementTemplateSteps []ManagementTemplateStep;
    // The collection of parameters used by the management template. Optional. Read-only.
    parameters []TemplateParameter;
    // 
    version *int32;
    // The collection of workload actions associated with the management template. Optional. Read-only.
    workloadActions []WorkloadAction;
}
// Instantiates a new managementTemplate and sets the default values.
func NewManagementTemplate()(*ManagementTemplate) {
    m := &ManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the managementTemplateCollections property value. 
func (m *ManagementTemplate) GetManagementTemplateCollections()([]ManagementTemplateCollection) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateCollections
    }
}
// Gets the managementTemplateSteps property value. 
func (m *ManagementTemplate) GetManagementTemplateSteps()([]ManagementTemplateStep) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateSteps
    }
}
// Gets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) GetParameters()([]TemplateParameter) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// Gets the version property value. 
func (m *ManagementTemplate) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Gets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) GetWorkloadActions()([]WorkloadAction) {
    if m == nil {
        return nil
    } else {
        return m.workloadActions
    }
}
// The deserialization information for the current model
func (m *ManagementTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)
            m.SetCategory(&cast)
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
    res["managementTemplateCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollection, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateCollection))
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStep))
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemplateParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemplateParameter, len(val))
            for i, v := range val {
                res[i] = *(v.(*TemplateParameter))
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
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
func (m *ManagementTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
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
// Sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
// Parameters:
//  - value : Value to set for the category property.
func (m *ManagementTemplate) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    m.category = value
}
// Sets the description property value. The description for the management template. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the management template. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the managementTemplateCollections property value. 
// Parameters:
//  - value : Value to set for the managementTemplateCollections property.
func (m *ManagementTemplate) SetManagementTemplateCollections(value []ManagementTemplateCollection)() {
    m.managementTemplateCollections = value
}
// Sets the managementTemplateSteps property value. 
// Parameters:
//  - value : Value to set for the managementTemplateSteps property.
func (m *ManagementTemplate) SetManagementTemplateSteps(value []ManagementTemplateStep)() {
    m.managementTemplateSteps = value
}
// Sets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
// Parameters:
//  - value : Value to set for the parameters property.
func (m *ManagementTemplate) SetParameters(value []TemplateParameter)() {
    m.parameters = value
}
// Sets the version property value. 
// Parameters:
//  - value : Value to set for the version property.
func (m *ManagementTemplate) SetVersion(value *int32)() {
    m.version = value
}
// Sets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
// Parameters:
//  - value : Value to set for the workloadActions property.
func (m *ManagementTemplate) SetWorkloadActions(value []WorkloadAction)() {
    m.workloadActions = value
}
