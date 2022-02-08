package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementTemplateStepDeployment 
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
// NewManagementTemplateStepDeployment instantiates a new managementTemplateStepDeployment and sets the default values.
func NewManagementTemplateStepDeployment()(*ManagementTemplateStepDeployment) {
    m := &ManagementTemplateStepDeployment{
        Entity: *NewEntity(),
    }
    return m
}
// GetError gets the error property value. 
func (m *ManagementTemplateStepDeployment) GetError()(*GraphAPIErrorDetails) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetSettings gets the settings property value. 
func (m *ManagementTemplateStepDeployment) GetSettings()([]Setting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetStatus gets the status property value. 
func (m *ManagementTemplateStepDeployment) GetStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTemplateStepVersion gets the templateStepVersion property value. 
func (m *ManagementTemplateStepDeployment) GetTemplateStepVersion()(*ManagementTemplateStepVersion) {
    if m == nil {
        return nil
    } else {
        return m.templateStepVersion
    }
}
// GetTenantId gets the tenantId property value. 
func (m *ManagementTemplateStepDeployment) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetStatus(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus))
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
// Serialize serializes information the current object
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
    if m.GetSettings() != nil {
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
        cast := (*m.GetStatus()).String()
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
// SetError sets the error property value. 
func (m *ManagementTemplateStepDeployment) SetError(value *GraphAPIErrorDetails)() {
    if m != nil {
        m.error = value
    }
}
// SetSettings sets the settings property value. 
func (m *ManagementTemplateStepDeployment) SetSettings(value []Setting)() {
    if m != nil {
        m.settings = value
    }
}
// SetStatus sets the status property value. 
func (m *ManagementTemplateStepDeployment) SetStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementTemplateDeploymentStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTemplateStepVersion sets the templateStepVersion property value. 
func (m *ManagementTemplateStepDeployment) SetTemplateStepVersion(value *ManagementTemplateStepVersion)() {
    if m != nil {
        m.templateStepVersion = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *ManagementTemplateStepDeployment) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
