package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateReference 
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
// NewDeviceManagementConfigurationPolicyTemplateReference instantiates a new deviceManagementConfigurationPolicyTemplateReference and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateReference()(*DeviceManagementConfigurationPolicyTemplateReference) {
    m := &DeviceManagementConfigurationPolicyTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTemplateDisplayName gets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayName
    }
}
// GetTemplateDisplayVersion gets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayVersion
    }
}
// GetTemplateFamily gets the templateFamily property value. Template Family of the referenced Template. This property is read-only. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    if m == nil {
        return nil
    } else {
        return m.templateFamily
    }
}
// GetTemplateId gets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
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
// Serialize serializes information the current object
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
        cast := (*m.GetTemplateFamily()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTemplateDisplayName sets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayName(value *string)() {
    if m != nil {
        m.templateDisplayName = value
    }
}
// SetTemplateDisplayVersion sets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayVersion(value *string)() {
    if m != nil {
        m.templateDisplayVersion = value
    }
}
// SetTemplateFamily sets the templateFamily property value. Template Family of the referenced Template. This property is read-only. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    if m != nil {
        m.templateFamily = value
    }
}
// SetTemplateId sets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
