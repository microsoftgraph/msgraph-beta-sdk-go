package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type ManagementTemplateStep struct {
    Entity
    // 
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory;
    // 
    description *string;
    // 
    displayName *string;
    // 
    managementPortal *string;
    // 
    managementTemplate *ManagementTemplate;
    // 
    portalLink *string;
    // 
    priority *int32;
    // 
    provider *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider;
    // 
    stepVersions []ManagementTemplateStepVersion;
}
// Instantiates a new managementTemplateStep and sets the default values.
func NewManagementTemplateStep()(*ManagementTemplateStep) {
    m := &ManagementTemplateStep{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the category property value. 
func (m *ManagementTemplateStep) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the description property value. 
func (m *ManagementTemplateStep) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. 
func (m *ManagementTemplateStep) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the managementPortal property value. 
func (m *ManagementTemplateStep) GetManagementPortal()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementPortal
    }
}
// Gets the managementTemplate property value. 
func (m *ManagementTemplateStep) GetManagementTemplate()(*ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplate
    }
}
// Gets the portalLink property value. 
func (m *ManagementTemplateStep) GetPortalLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.portalLink
    }
}
// Gets the priority property value. 
func (m *ManagementTemplateStep) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the provider property value. 
func (m *ManagementTemplateStep) GetProvider()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// Gets the stepVersions property value. 
func (m *ManagementTemplateStep) GetStepVersions()([]ManagementTemplateStepVersion) {
    if m == nil {
        return nil
    } else {
        return m.stepVersions
    }
}
// The deserialization information for the current model
func (m *ManagementTemplateStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["managementPortal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementPortal(val)
        }
        return nil
    }
    res["managementTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplate(val.(*ManagementTemplate))
        }
        return nil
    }
    res["portalLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalLink(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["provider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementProvider)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)
            m.SetProvider(&cast)
        }
        return nil
    }
    res["stepVersions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStepVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersion, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStepVersion))
            }
            m.SetStepVersions(res)
        }
        return nil
    }
    return res
}
func (m *ManagementTemplateStep) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementTemplateStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("managementPortal", m.GetManagementPortal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managementTemplate", m.GetManagementTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("portalLink", m.GetPortalLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetProvider() != nil {
        cast := m.GetProvider().String()
        err = writer.WriteStringValue("provider", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStepVersions()))
        for i, v := range m.GetStepVersions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("stepVersions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the category property value. 
// Parameters:
//  - value : Value to set for the category property.
func (m *ManagementTemplateStep) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    m.category = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementTemplateStep) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementTemplateStep) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the managementPortal property value. 
// Parameters:
//  - value : Value to set for the managementPortal property.
func (m *ManagementTemplateStep) SetManagementPortal(value *string)() {
    m.managementPortal = value
}
// Sets the managementTemplate property value. 
// Parameters:
//  - value : Value to set for the managementTemplate property.
func (m *ManagementTemplateStep) SetManagementTemplate(value *ManagementTemplate)() {
    m.managementTemplate = value
}
// Sets the portalLink property value. 
// Parameters:
//  - value : Value to set for the portalLink property.
func (m *ManagementTemplateStep) SetPortalLink(value *string)() {
    m.portalLink = value
}
// Sets the priority property value. 
// Parameters:
//  - value : Value to set for the priority property.
func (m *ManagementTemplateStep) SetPriority(value *int32)() {
    m.priority = value
}
// Sets the provider property value. 
// Parameters:
//  - value : Value to set for the provider property.
func (m *ManagementTemplateStep) SetProvider(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)() {
    m.provider = value
}
// Sets the stepVersions property value. 
// Parameters:
//  - value : Value to set for the stepVersions property.
func (m *ManagementTemplateStep) SetStepVersions(value []ManagementTemplateStepVersion)() {
    m.stepVersions = value
}
