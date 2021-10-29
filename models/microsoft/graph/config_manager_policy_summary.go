package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConfigManagerPolicySummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of devices evaluated to be compliant by the policy.
    compliantDeviceCount *int32;
    // The number of devices that have have been remediated by the policy.
    enforcedDeviceCount *int32;
    // The number of devices that failed to be evaluated by the policy.
    failedDeviceCount *int32;
    // The number of devices evaluated to be noncompliant by the policy.
    nonCompliantDeviceCount *int32;
    // The number of devices that have acknowledged the policy but are pending evaluation.
    pendingDeviceCount *int32;
    // The number of devices targeted by the policy.
    targetedDeviceCount *int32;
}
// Instantiates a new configManagerPolicySummary and sets the default values.
func NewConfigManagerPolicySummary()(*ConfigManagerPolicySummary) {
    m := &ConfigManagerPolicySummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigManagerPolicySummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) GetEnforcedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enforcedDeviceCount
    }
}
// Gets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// Gets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// Gets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) GetPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingDeviceCount
    }
}
// Gets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) GetTargetedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceCount
    }
}
// The deserialization information for the current model
func (m *ConfigManagerPolicySummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
        return nil
    }
    res["enforcedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEnforcedDeviceCount(val)
        return nil
    }
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedDeviceCount(val)
        return nil
    }
    res["nonCompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantDeviceCount(val)
        return nil
    }
    res["pendingDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingDeviceCount(val)
        return nil
    }
    res["targetedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTargetedDeviceCount(val)
        return nil
    }
    return res
}
func (m *ConfigManagerPolicySummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConfigManagerPolicySummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("enforcedDeviceCount", m.GetEnforcedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingDeviceCount", m.GetPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("targetedDeviceCount", m.GetTargetedDeviceCount())
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
func (m *ConfigManagerPolicySummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *ConfigManagerPolicySummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
// Parameters:
//  - value : Value to set for the enforcedDeviceCount property.
func (m *ConfigManagerPolicySummary) SetEnforcedDeviceCount(value *int32)() {
    m.enforcedDeviceCount = value
}
// Sets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
// Parameters:
//  - value : Value to set for the failedDeviceCount property.
func (m *ConfigManagerPolicySummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// Sets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
// Parameters:
//  - value : Value to set for the nonCompliantDeviceCount property.
func (m *ConfigManagerPolicySummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// Sets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
// Parameters:
//  - value : Value to set for the pendingDeviceCount property.
func (m *ConfigManagerPolicySummary) SetPendingDeviceCount(value *int32)() {
    m.pendingDeviceCount = value
}
// Sets the targetedDeviceCount property value. The number of devices targeted by the policy.
// Parameters:
//  - value : Value to set for the targetedDeviceCount property.
func (m *ConfigManagerPolicySummary) SetTargetedDeviceCount(value *int32)() {
    m.targetedDeviceCount = value
}
