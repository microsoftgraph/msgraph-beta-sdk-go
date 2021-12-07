package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigManagerPolicySummary 
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
// NewConfigManagerPolicySummary instantiates a new configManagerPolicySummary and sets the default values.
func NewConfigManagerPolicySummary()(*ConfigManagerPolicySummary) {
    m := &ConfigManagerPolicySummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigManagerPolicySummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetEnforcedDeviceCount gets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) GetEnforcedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enforcedDeviceCount
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) GetPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingDeviceCount
    }
}
// GetTargetedDeviceCount gets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) GetTargetedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigManagerPolicySummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["enforcedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcedDeviceCount(val)
        }
        return nil
    }
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["nonCompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantDeviceCount(val)
        }
        return nil
    }
    res["pendingDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingDeviceCount(val)
        }
        return nil
    }
    res["targetedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *ConfigManagerPolicySummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigManagerPolicySummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetEnforcedDeviceCount sets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) SetEnforcedDeviceCount(value *int32)() {
    if m != nil {
        m.enforcedDeviceCount = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) SetNonCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.nonCompliantDeviceCount = value
    }
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) SetPendingDeviceCount(value *int32)() {
    if m != nil {
        m.pendingDeviceCount = value
    }
}
// SetTargetedDeviceCount sets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) SetTargetedDeviceCount(value *int32)() {
    if m != nil {
        m.targetedDeviceCount = value
    }
}
