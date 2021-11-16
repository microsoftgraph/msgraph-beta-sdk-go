package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type ManagementTemplateStepDeployment struct {
    Entity
    // 
    error *GraphAPIErrorDetails;
    // 
    settings []Setting;
    // 
    status *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus;
    // 
    templateStepVersion *ManagementTemplateStepVersion;
    // 
    tenantId *string;
}
// Instantiates a new managementTemplateStepDeployment and sets the default values.
func NewManagementTemplateStepDeployment()(*ManagementTemplateStepDeployment) {
    m := &ManagementTemplateStepDeployment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the error property value. 
func (m *ManagementTemplateStepDeployment) GetError()(*GraphAPIErrorDetails) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the settings property value. 
func (m *ManagementTemplateStepDeployment) GetSettings()([]Setting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the status property value. 
func (m *ManagementTemplateStepDeployment) GetStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the templateStepVersion property value. 
func (m *ManagementTemplateStepDeployment) GetTemplateStepVersion()(*ManagementTemplateStepVersion) {
    if m == nil {
        return nil
    } else {
        return m.templateStepVersion
    }
}
// Gets the tenantId property value. 
func (m *ManagementTemplateStepDeployment) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// The deserialization information for the current model
func (m *ManagementTemplateStepDeployment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGraphAPIErrorDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*GraphAPIErrorDetails))
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setting, len(val))
            for i, v := range val {
                res[i] = *(v.(*Setting))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementTemplateDeploymentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["templateStepVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStepVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateStepVersion(val.(*ManagementTemplateStepVersion))
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
func (m *ManagementTemplateStepDeployment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementTemplateStepDeployment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStepVersion", m.GetTemplateStepVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the error property value. 
// Parameters:
//  - value : Value to set for the error property.
func (m *ManagementTemplateStepDeployment) SetError(value *GraphAPIErrorDetails)() {
    m.error = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *ManagementTemplateStepDeployment) SetSettings(value []Setting)() {
    m.settings = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *ManagementTemplateStepDeployment) SetStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus)() {
    m.status = value
}
// Sets the templateStepVersion property value. 
// Parameters:
//  - value : Value to set for the templateStepVersion property.
func (m *ManagementTemplateStepDeployment) SetTemplateStepVersion(value *ManagementTemplateStepVersion)() {
    m.templateStepVersion = value
}
// Sets the tenantId property value. 
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *ManagementTemplateStepDeployment) SetTenantId(value *string)() {
    m.tenantId = value
}
