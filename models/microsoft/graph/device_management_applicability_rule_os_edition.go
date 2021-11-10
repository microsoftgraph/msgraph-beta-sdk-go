package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementApplicabilityRuleOsEdition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name for object.
    name *string;
    // Applicability rule OS edition type.
    osEditionTypes []Windows10EditionType;
    // Applicability Rule type. Possible values are: include, exclude.
    ruleType *DeviceManagementApplicabilityRuleType;
}
// Instantiates a new deviceManagementApplicabilityRuleOsEdition and sets the default values.
func NewDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    m := &DeviceManagementApplicabilityRuleOsEdition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the osEditionTypes property value. Applicability rule OS edition type.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetOsEditionTypes()([]Windows10EditionType) {
    if m == nil {
        return nil
    } else {
        return m.osEditionTypes
    }
}
// Gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// The deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleOsEdition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["osEditionTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseWindows10EditionType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10EditionType, len(val))
            for i, v := range val {
                res[i] = *(v.(*Windows10EditionType))
            }
            m.SetOsEditionTypes(res)
        }
        return nil
    }
    res["ruleType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementApplicabilityRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementApplicabilityRuleType)
            m.SetRuleType(&cast)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementApplicabilityRuleOsEdition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementApplicabilityRuleOsEdition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("osEditionTypes", SerializeWindows10EditionType(m.GetOsEditionTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetRuleType() != nil {
        cast := m.GetRuleType().String()
        err := writer.WriteStringValue("ruleType", &cast)
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
func (m *DeviceManagementApplicabilityRuleOsEdition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the name property value. Name for object.
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetName(value *string)() {
    m.name = value
}
// Sets the osEditionTypes property value. Applicability rule OS edition type.
// Parameters:
//  - value : Value to set for the osEditionTypes property.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetOsEditionTypes(value []Windows10EditionType)() {
    m.osEditionTypes = value
}
// Sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
// Parameters:
//  - value : Value to set for the ruleType property.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}
