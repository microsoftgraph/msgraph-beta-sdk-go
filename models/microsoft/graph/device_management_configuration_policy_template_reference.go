package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationPolicyTemplateReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Template Display Name of the referenced template. This property is read-only.
    templateDisplayName *string;
    // Template Display Version of the referenced Template. This property is read-only.
    templateDisplayVersion *string;
    // Template Family of the referenced Template. This property is read-only. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
    templateFamily *DeviceManagementConfigurationTemplateFamily;
    // Template id
    templateId *string;
}
// Instantiates a new deviceManagementConfigurationPolicyTemplateReference and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateReference()(*DeviceManagementConfigurationPolicyTemplateReference) {
    m := &DeviceManagementConfigurationPolicyTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayName
    }
}
// Gets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayVersion
    }
}
// Gets the templateFamily property value. Template Family of the referenced Template. This property is read-only. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    if m == nil {
        return nil
    } else {
        return m.templateFamily
    }
}
// Gets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["templateDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayName(val)
        }
        return nil
    }
    res["templateDisplayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayVersion(val)
        }
        return nil
    }
    res["templateFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementConfigurationTemplateFamily)
            m.SetTemplateFamily(&cast)
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationPolicyTemplateReference) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationPolicyTemplateReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("templateDisplayName", m.GetTemplateDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateDisplayVersion", m.GetTemplateDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateFamily() != nil {
        cast := m.GetTemplateFamily().String()
        err := writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
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
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
// Parameters:
//  - value : Value to set for the templateDisplayName property.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayName(value *string)() {
    m.templateDisplayName = value
}
// Sets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
// Parameters:
//  - value : Value to set for the templateDisplayVersion property.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayVersion(value *string)() {
    m.templateDisplayVersion = value
}
// Sets the templateFamily property value. Template Family of the referenced Template. This property is read-only. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
// Parameters:
//  - value : Value to set for the templateFamily property.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    m.templateFamily = value
}
// Sets the templateId property value. Template id
// Parameters:
//  - value : Value to set for the templateId property.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateId(value *string)() {
    m.templateId = value
}
