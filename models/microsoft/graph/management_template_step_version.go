package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementTemplateStepVersion 
type ManagementTemplateStepVersion struct {
    Entity
    // 
    configurationAction *TemplateAction;
    // 
    deployments []ManagementTemplateStepDeployment;
    // 
    templateStep *ManagementTemplateStep;
    // 
    validationAction *TemplateAction;
    // 
    version *int32;
}
// NewManagementTemplateStepVersion instantiates a new managementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfigurationAction gets the configurationAction property value. 
func (m *ManagementTemplateStepVersion) GetConfigurationAction()(*TemplateAction) {
    if m == nil {
        return nil
    } else {
        return m.configurationAction
    }
}
// GetDeployments gets the deployments property value. 
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeployment) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// GetTemplateStep gets the templateStep property value. 
func (m *ManagementTemplateStepVersion) GetTemplateStep()(*ManagementTemplateStep) {
    if m == nil {
        return nil
    } else {
        return m.templateStep
    }
}
// GetValidationAction gets the validationAction property value. 
func (m *ManagementTemplateStepVersion) GetValidationAction()(*TemplateAction) {
    if m == nil {
        return nil
    } else {
        return m.validationAction
    }
}
// GetVersion gets the version property value. 
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemplateAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationAction(val.(*TemplateAction))
        }
        return nil
    }
    res["deployments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStepDeployment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepDeployment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStepDeployment))
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["templateStep"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStep() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateStep(val.(*ManagementTemplateStep))
        }
        return nil
    }
    res["validationAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemplateAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationAction(val.(*TemplateAction))
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
    return res
}
func (m *ManagementTemplateStepVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configurationAction", m.GetConfigurationAction())
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStep", m.GetTemplateStep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("validationAction", m.GetValidationAction())
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
    return nil
}
// SetConfigurationAction sets the configurationAction property value. 
func (m *ManagementTemplateStepVersion) SetConfigurationAction(value *TemplateAction)() {
    if m != nil {
        m.configurationAction = value
    }
}
// SetDeployments sets the deployments property value. 
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeployment)() {
    if m != nil {
        m.deployments = value
    }
}
// SetTemplateStep sets the templateStep property value. 
func (m *ManagementTemplateStepVersion) SetTemplateStep(value *ManagementTemplateStep)() {
    if m != nil {
        m.templateStep = value
    }
}
// SetValidationAction sets the validationAction property value. 
func (m *ManagementTemplateStepVersion) SetValidationAction(value *TemplateAction)() {
    if m != nil {
        m.validationAction = value
    }
}
// SetVersion sets the version property value. 
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
