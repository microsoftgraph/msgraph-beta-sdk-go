package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
    Entity
    // Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
    deployedDeviceCount *int32;
    // Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
    failedDeviceCount *int32;
}
// Instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentSummary and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetDeployedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedDeviceCount
    }
}
// Gets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
// Parameters:
//  - value : Value to set for the deployedDeviceCount property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetDeployedDeviceCount(value *int32)() {
    m.deployedDeviceCount = value
}
// Sets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
// Parameters:
//  - value : Value to set for the failedDeviceCount property.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
