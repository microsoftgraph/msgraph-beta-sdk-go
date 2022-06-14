package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityConfigurationTask 
type SecurityConfigurationTask struct {
    DeviceAppManagementTask
    // The applicable platform. Possible values are: unknown, macOS, windows10AndLater, windows10AndWindowsServer.
    applicablePlatform *EndpointSecurityConfigurationApplicablePlatform
    // The endpoint security policy type. Possible values are: unknown, antivirus, diskEncryption, firewall, endpointDetectionAndResponse, attackSurfaceReduction, accountProtection.
    endpointSecurityPolicy *EndpointSecurityConfigurationType
    // The endpoint security policy profile. Possible values are: unknown, antivirus, windowsSecurity, bitLocker, fileVault, firewall, firewallRules, endpointDetectionAndResponse, deviceControl, appAndBrowserIsolation, exploitProtection, webProtection, applicationControl, attackSurfaceReductionRules, accountProtection.
    endpointSecurityPolicyProfile *EndpointSecurityConfigurationProfileType
    // Information about the mitigation.
    insights *string
    // The intended settings and their values.
    intendedSettings []KeyValuePairable
    // The number of vulnerable devices.
    managedDeviceCount *int32
    // The vulnerable managed devices.
    managedDevices []VulnerableManagedDeviceable
}
// NewSecurityConfigurationTask instantiates a new SecurityConfigurationTask and sets the default values.
func NewSecurityConfigurationTask()(*SecurityConfigurationTask) {
    m := &SecurityConfigurationTask{
        DeviceAppManagementTask: *NewDeviceAppManagementTask(),
    }
    return m
}
// CreateSecurityConfigurationTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityConfigurationTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityConfigurationTask(), nil
}
// GetApplicablePlatform gets the applicablePlatform property value. The applicable platform. Possible values are: unknown, macOS, windows10AndLater, windows10AndWindowsServer.
func (m *SecurityConfigurationTask) GetApplicablePlatform()(*EndpointSecurityConfigurationApplicablePlatform) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatform
    }
}
// GetEndpointSecurityPolicy gets the endpointSecurityPolicy property value. The endpoint security policy type. Possible values are: unknown, antivirus, diskEncryption, firewall, endpointDetectionAndResponse, attackSurfaceReduction, accountProtection.
func (m *SecurityConfigurationTask) GetEndpointSecurityPolicy()(*EndpointSecurityConfigurationType) {
    if m == nil {
        return nil
    } else {
        return m.endpointSecurityPolicy
    }
}
// GetEndpointSecurityPolicyProfile gets the endpointSecurityPolicyProfile property value. The endpoint security policy profile. Possible values are: unknown, antivirus, windowsSecurity, bitLocker, fileVault, firewall, firewallRules, endpointDetectionAndResponse, deviceControl, appAndBrowserIsolation, exploitProtection, webProtection, applicationControl, attackSurfaceReductionRules, accountProtection.
func (m *SecurityConfigurationTask) GetEndpointSecurityPolicyProfile()(*EndpointSecurityConfigurationProfileType) {
    if m == nil {
        return nil
    } else {
        return m.endpointSecurityPolicyProfile
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityConfigurationTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAppManagementTask.GetFieldDeserializers()
    res["applicablePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointSecurityConfigurationApplicablePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicablePlatform(val.(*EndpointSecurityConfigurationApplicablePlatform))
        }
        return nil
    }
    res["endpointSecurityPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointSecurityConfigurationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointSecurityPolicy(val.(*EndpointSecurityConfigurationType))
        }
        return nil
    }
    res["endpointSecurityPolicyProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointSecurityConfigurationProfileType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointSecurityPolicyProfile(val.(*EndpointSecurityConfigurationProfileType))
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val)
        }
        return nil
    }
    res["intendedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetIntendedSettings(res)
        }
        return nil
    }
    res["managedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceCount(val)
        }
        return nil
    }
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerableManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VulnerableManagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(VulnerableManagedDeviceable)
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    return res
}
// GetInsights gets the insights property value. Information about the mitigation.
func (m *SecurityConfigurationTask) GetInsights()(*string) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetIntendedSettings gets the intendedSettings property value. The intended settings and their values.
func (m *SecurityConfigurationTask) GetIntendedSettings()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.intendedSettings
    }
}
// GetManagedDeviceCount gets the managedDeviceCount property value. The number of vulnerable devices.
func (m *SecurityConfigurationTask) GetManagedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCount
    }
}
// GetManagedDevices gets the managedDevices property value. The vulnerable managed devices.
func (m *SecurityConfigurationTask) GetManagedDevices()([]VulnerableManagedDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// Serialize serializes information the current object
func (m *SecurityConfigurationTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAppManagementTask.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicablePlatform() != nil {
        cast := (*m.GetApplicablePlatform()).String()
        err = writer.WriteStringValue("applicablePlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEndpointSecurityPolicy() != nil {
        cast := (*m.GetEndpointSecurityPolicy()).String()
        err = writer.WriteStringValue("endpointSecurityPolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEndpointSecurityPolicyProfile() != nil {
        cast := (*m.GetEndpointSecurityPolicyProfile()).String()
        err = writer.WriteStringValue("endpointSecurityPolicyProfile", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    if m.GetIntendedSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntendedSettings()))
        for i, v := range m.GetIntendedSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("intendedSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("managedDeviceCount", m.GetManagedDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicablePlatform sets the applicablePlatform property value. The applicable platform. Possible values are: unknown, macOS, windows10AndLater, windows10AndWindowsServer.
func (m *SecurityConfigurationTask) SetApplicablePlatform(value *EndpointSecurityConfigurationApplicablePlatform)() {
    if m != nil {
        m.applicablePlatform = value
    }
}
// SetEndpointSecurityPolicy sets the endpointSecurityPolicy property value. The endpoint security policy type. Possible values are: unknown, antivirus, diskEncryption, firewall, endpointDetectionAndResponse, attackSurfaceReduction, accountProtection.
func (m *SecurityConfigurationTask) SetEndpointSecurityPolicy(value *EndpointSecurityConfigurationType)() {
    if m != nil {
        m.endpointSecurityPolicy = value
    }
}
// SetEndpointSecurityPolicyProfile sets the endpointSecurityPolicyProfile property value. The endpoint security policy profile. Possible values are: unknown, antivirus, windowsSecurity, bitLocker, fileVault, firewall, firewallRules, endpointDetectionAndResponse, deviceControl, appAndBrowserIsolation, exploitProtection, webProtection, applicationControl, attackSurfaceReductionRules, accountProtection.
func (m *SecurityConfigurationTask) SetEndpointSecurityPolicyProfile(value *EndpointSecurityConfigurationProfileType)() {
    if m != nil {
        m.endpointSecurityPolicyProfile = value
    }
}
// SetInsights sets the insights property value. Information about the mitigation.
func (m *SecurityConfigurationTask) SetInsights(value *string)() {
    if m != nil {
        m.insights = value
    }
}
// SetIntendedSettings sets the intendedSettings property value. The intended settings and their values.
func (m *SecurityConfigurationTask) SetIntendedSettings(value []KeyValuePairable)() {
    if m != nil {
        m.intendedSettings = value
    }
}
// SetManagedDeviceCount sets the managedDeviceCount property value. The number of vulnerable devices.
func (m *SecurityConfigurationTask) SetManagedDeviceCount(value *int32)() {
    if m != nil {
        m.managedDeviceCount = value
    }
}
// SetManagedDevices sets the managedDevices property value. The vulnerable managed devices.
func (m *SecurityConfigurationTask) SetManagedDevices(value []VulnerableManagedDeviceable)() {
    if m != nil {
        m.managedDevices = value
    }
}
