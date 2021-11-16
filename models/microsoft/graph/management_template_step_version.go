package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new managementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configurationAction property value. 
func (m *ManagementTemplateStepVersion) GetConfigurationAction()(*TemplateAction) {
    if m == nil {
        return nil
    } else {
        return m.configurationAction
    }
}
// Gets the deployments property value. 
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeployment) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// Gets the templateStep property value. 
func (m *ManagementTemplateStepVersion) GetTemplateStep()(*ManagementTemplateStep) {
    if m == nil {
        return nil
    } else {
        return m.templateStep
    }
}
// Gets the validationAction property value. 
func (m *ManagementTemplateStepVersion) GetValidationAction()(*TemplateAction) {
    if m == nil {
        return nil
    } else {
        return m.validationAction
    }
}
// Gets the version property value. 
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
// Sets the configurationAction property value. 
// Parameters:
//  - value : Value to set for the configurationAction property.
func (m *ManagementTemplateStepVersion) SetConfigurationAction(value *TemplateAction)() {
    m.configurationAction = value
}
// Sets the deployments property value. 
// Parameters:
//  - value : Value to set for the deployments property.
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeployment)() {
    m.deployments = value
}
// Sets the templateStep property value. 
// Parameters:
//  - value : Value to set for the templateStep property.
func (m *ManagementTemplateStepVersion) SetTemplateStep(value *ManagementTemplateStep)() {
    m.templateStep = value
}
// Sets the validationAction property value. 
// Parameters:
//  - value : Value to set for the validationAction property.
func (m *ManagementTemplateStepVersion) SetValidationAction(value *TemplateAction)() {
    m.validationAction = value
}
// Sets the version property value. 
// Parameters:
//  - value : Value to set for the version property.
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    m.version = value
}
