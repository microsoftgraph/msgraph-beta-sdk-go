package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
    Entity
    // Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
    deployedDeviceCount *int32;
    // Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
    failedDeviceCount *int32;
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentSummary and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeployedDeviceCount gets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetDeployedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedDeviceCount
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deployedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedDeviceCount(val)
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
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deployedDeviceCount", m.GetDeployedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeployedDeviceCount sets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetDeployedDeviceCount(value *int32)() {
    if m != nil {
        m.deployedDeviceCount = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
