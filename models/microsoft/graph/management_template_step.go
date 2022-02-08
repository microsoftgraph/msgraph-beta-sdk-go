package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementTemplateStep 
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
// NewManagementTemplateStep instantiates a new managementTemplateStep and sets the default values.
func NewManagementTemplateStep()(*ManagementTemplateStep) {
    m := &ManagementTemplateStep{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. 
func (m *ManagementTemplateStep) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDescription gets the description property value. 
func (m *ManagementTemplateStep) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *ManagementTemplateStep) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetManagementPortal gets the managementPortal property value. 
func (m *ManagementTemplateStep) GetManagementPortal()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementPortal
    }
}
// GetManagementTemplate gets the managementTemplate property value. 
func (m *ManagementTemplateStep) GetManagementTemplate()(*ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplate
    }
}
// GetPortalLink gets the portalLink property value. 
func (m *ManagementTemplateStep) GetPortalLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.portalLink
    }
}
// GetPriority gets the priority property value. 
func (m *ManagementTemplateStep) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetProvider gets the provider property value. 
func (m *ManagementTemplateStep) GetProvider()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// GetStepVersions gets the stepVersions property value. 
func (m *ManagementTemplateStep) GetStepVersions()([]ManagementTemplateStepVersion) {
    if m == nil {
        return nil
    } else {
        return m.stepVersions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
            m.SetProvider(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider))
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
// Serialize serializes information the current object
func (m *ManagementTemplateStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := (*m.GetProvider()).String()
        err = writer.WriteStringValue("provider", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStepVersions() != nil {
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
// SetCategory sets the category property value. 
func (m *ManagementTemplateStep) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDescription sets the description property value. 
func (m *ManagementTemplateStep) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *ManagementTemplateStep) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagementPortal sets the managementPortal property value. 
func (m *ManagementTemplateStep) SetManagementPortal(value *string)() {
    if m != nil {
        m.managementPortal = value
    }
}
// SetManagementTemplate sets the managementTemplate property value. 
func (m *ManagementTemplateStep) SetManagementTemplate(value *ManagementTemplate)() {
    if m != nil {
        m.managementTemplate = value
    }
}
// SetPortalLink sets the portalLink property value. 
func (m *ManagementTemplateStep) SetPortalLink(value *string)() {
    if m != nil {
        m.portalLink = value
    }
}
// SetPriority sets the priority property value. 
func (m *ManagementTemplateStep) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetProvider sets the provider property value. 
func (m *ManagementTemplateStep) SetProvider(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)() {
    if m != nil {
        m.provider = value
    }
}
// SetStepVersions sets the stepVersions property value. 
func (m *ManagementTemplateStep) SetStepVersions(value []ManagementTemplateStepVersion)() {
    if m != nil {
        m.stepVersions = value
    }
}
