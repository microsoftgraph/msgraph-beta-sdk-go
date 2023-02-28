package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityConfigurationTask 
type SecurityConfigurationTask struct {
    DeviceAppManagementTask
}
// NewSecurityConfigurationTask instantiates a new SecurityConfigurationTask and sets the default values.
func NewSecurityConfigurationTask()(*SecurityConfigurationTask) {
    m := &SecurityConfigurationTask{
        DeviceAppManagementTask: *NewDeviceAppManagementTask(),
    }
    odataTypeValue := "#microsoft.graph.securityConfigurationTask"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSecurityConfigurationTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityConfigurationTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityConfigurationTask(), nil
}
// GetApplicablePlatform gets the applicablePlatform property value. The endpoint security configuration applicable platform.
func (m *SecurityConfigurationTask) GetApplicablePlatform()(*EndpointSecurityConfigurationApplicablePlatform) {
    val, err := m.GetBackingStore().Get("applicablePlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndpointSecurityConfigurationApplicablePlatform)
    }
    return nil
}
// GetEndpointSecurityPolicy gets the endpointSecurityPolicy property value. The endpoint security policy type.
func (m *SecurityConfigurationTask) GetEndpointSecurityPolicy()(*EndpointSecurityConfigurationType) {
    val, err := m.GetBackingStore().Get("endpointSecurityPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndpointSecurityConfigurationType)
    }
    return nil
}
// GetEndpointSecurityPolicyProfile gets the endpointSecurityPolicyProfile property value. The endpoint security policy profile type.
func (m *SecurityConfigurationTask) GetEndpointSecurityPolicyProfile()(*EndpointSecurityConfigurationProfileType) {
    val, err := m.GetBackingStore().Get("endpointSecurityPolicyProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndpointSecurityConfigurationProfileType)
    }
    return nil
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
    val, err := m.GetBackingStore().Get("insights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIntendedSettings gets the intendedSettings property value. The intended settings and their values.
func (m *SecurityConfigurationTask) GetIntendedSettings()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("intendedSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetManagedDeviceCount gets the managedDeviceCount property value. The number of vulnerable devices. Valid values 0 to 65536
func (m *SecurityConfigurationTask) GetManagedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("managedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetManagedDevices gets the managedDevices property value. The vulnerable managed devices.
func (m *SecurityConfigurationTask) GetManagedDevices()([]VulnerableManagedDeviceable) {
    val, err := m.GetBackingStore().Get("managedDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VulnerableManagedDeviceable)
    }
    return nil
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
// SetApplicablePlatform sets the applicablePlatform property value. The endpoint security configuration applicable platform.
func (m *SecurityConfigurationTask) SetApplicablePlatform(value *EndpointSecurityConfigurationApplicablePlatform)() {
    err := m.GetBackingStore().Set("applicablePlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointSecurityPolicy sets the endpointSecurityPolicy property value. The endpoint security policy type.
func (m *SecurityConfigurationTask) SetEndpointSecurityPolicy(value *EndpointSecurityConfigurationType)() {
    err := m.GetBackingStore().Set("endpointSecurityPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointSecurityPolicyProfile sets the endpointSecurityPolicyProfile property value. The endpoint security policy profile type.
func (m *SecurityConfigurationTask) SetEndpointSecurityPolicyProfile(value *EndpointSecurityConfigurationProfileType)() {
    err := m.GetBackingStore().Set("endpointSecurityPolicyProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetInsights sets the insights property value. Information about the mitigation.
func (m *SecurityConfigurationTask) SetInsights(value *string)() {
    err := m.GetBackingStore().Set("insights", value)
    if err != nil {
        panic(err)
    }
}
// SetIntendedSettings sets the intendedSettings property value. The intended settings and their values.
func (m *SecurityConfigurationTask) SetIntendedSettings(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("intendedSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCount sets the managedDeviceCount property value. The number of vulnerable devices. Valid values 0 to 65536
func (m *SecurityConfigurationTask) SetManagedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("managedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDevices sets the managedDevices property value. The vulnerable managed devices.
func (m *SecurityConfigurationTask) SetManagedDevices(value []VulnerableManagedDeviceable)() {
    err := m.GetBackingStore().Set("managedDevices", value)
    if err != nil {
        panic(err)
    }
}
// SecurityConfigurationTaskable 
type SecurityConfigurationTaskable interface {
    DeviceAppManagementTaskable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicablePlatform()(*EndpointSecurityConfigurationApplicablePlatform)
    GetEndpointSecurityPolicy()(*EndpointSecurityConfigurationType)
    GetEndpointSecurityPolicyProfile()(*EndpointSecurityConfigurationProfileType)
    GetInsights()(*string)
    GetIntendedSettings()([]KeyValuePairable)
    GetManagedDeviceCount()(*int32)
    GetManagedDevices()([]VulnerableManagedDeviceable)
    SetApplicablePlatform(value *EndpointSecurityConfigurationApplicablePlatform)()
    SetEndpointSecurityPolicy(value *EndpointSecurityConfigurationType)()
    SetEndpointSecurityPolicyProfile(value *EndpointSecurityConfigurationProfileType)()
    SetInsights(value *string)()
    SetIntendedSettings(value []KeyValuePairable)()
    SetManagedDeviceCount(value *int32)()
    SetManagedDevices(value []VulnerableManagedDeviceable)()
}
