package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementApplicabilityRuleOsEdition 
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
// NewDeviceManagementApplicabilityRuleOsEdition instantiates a new deviceManagementApplicabilityRuleOsEdition and sets the default values.
func NewDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    m := &DeviceManagementApplicabilityRuleOsEdition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOsEditionTypes gets the osEditionTypes property value. Applicability rule OS edition type.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetOsEditionTypes()([]Windows10EditionType) {
    if m == nil {
        return nil
    } else {
        return m.osEditionTypes
    }
}
// GetRuleType gets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    if m == nil {
        return nil
    } else {
        return m.ruleType
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetRuleType(val.(*DeviceManagementApplicabilityRuleType))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementApplicabilityRuleOsEdition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementApplicabilityRuleOsEdition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetOsEditionTypes() != nil {
        err := writer.WriteCollectionOfStringValues("osEditionTypes", SerializeWindows10EditionType(m.GetOsEditionTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetRuleType() != nil {
        cast := (*m.GetRuleType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOsEditionTypes sets the osEditionTypes property value. Applicability rule OS edition type.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetOsEditionTypes(value []Windows10EditionType)() {
    if m != nil {
        m.osEditionTypes = value
    }
}
// SetRuleType sets the ruleType property value. Applicability Rule type. Possible values are: include, exclude.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    if m != nil {
        m.ruleType = value
    }
}
