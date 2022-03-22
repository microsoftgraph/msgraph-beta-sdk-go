package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingDependency dependency information for a setting
type DeviceManagementSettingDependency struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of constraints for the dependency setting value
    constraints []DeviceManagementConstraintable;
    // The setting definition ID of the setting depended on
    definitionId *string;
}
// NewDeviceManagementSettingDependency instantiates a new deviceManagementSettingDependency and sets the default values.
func NewDeviceManagementSettingDependency()(*DeviceManagementSettingDependency) {
    m := &DeviceManagementSettingDependency{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingDependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingDependencyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementSettingDependency(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingDependency) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConstraints gets the constraints property value. Collection of constraints for the dependency setting value
func (m *DeviceManagementSettingDependency) GetConstraints()([]DeviceManagementConstraintable) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
// GetDefinitionId gets the definitionId property value. The setting definition ID of the setting depended on
func (m *DeviceManagementSettingDependency) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingDependency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["constraints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConstraintable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConstraintable)
            }
            m.SetConstraints(res)
        }
        return nil
    }
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingDependency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetConstraints() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingDependency) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConstraints sets the constraints property value. Collection of constraints for the dependency setting value
func (m *DeviceManagementSettingDependency) SetConstraints(value []DeviceManagementConstraintable)() {
    if m != nil {
        m.constraints = value
    }
}
// SetDefinitionId sets the definitionId property value. The setting definition ID of the setting depended on
func (m *DeviceManagementSettingDependency) SetDefinitionId(value *string)() {
    if m != nil {
        m.definitionId = value
    }
}
