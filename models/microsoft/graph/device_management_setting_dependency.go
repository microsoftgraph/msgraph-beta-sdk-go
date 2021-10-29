package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementSettingDependency struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of constraints for the dependency setting value
    constraints []DeviceManagementConstraint;
    // The setting definition ID of the setting depended on
    definitionId *string;
}
// Instantiates a new deviceManagementSettingDependency and sets the default values.
func NewDeviceManagementSettingDependency()(*DeviceManagementSettingDependency) {
    m := &DeviceManagementSettingDependency{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingDependency) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the constraints property value. Collection of constraints for the dependency setting value
func (m *DeviceManagementSettingDependency) GetConstraints()([]DeviceManagementConstraint) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
// Gets the definitionId property value. The setting definition ID of the setting depended on
func (m *DeviceManagementSettingDependency) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementSettingDependency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["constraints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConstraint() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConstraint, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConstraint))
        }
        m.SetConstraints(res)
        return nil
    }
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefinitionId(val)
        return nil
    }
    return res
}
func (m *DeviceManagementSettingDependency) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementSettingDependency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("definitionId", m.GetDefinitionId())
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
func (m *DeviceManagementSettingDependency) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the constraints property value. Collection of constraints for the dependency setting value
// Parameters:
//  - value : Value to set for the constraints property.
func (m *DeviceManagementSettingDependency) SetConstraints(value []DeviceManagementConstraint)() {
    m.constraints = value
}
// Sets the definitionId property value. The setting definition ID of the setting depended on
// Parameters:
//  - value : Value to set for the definitionId property.
func (m *DeviceManagementSettingDependency) SetDefinitionId(value *string)() {
    m.definitionId = value
}
