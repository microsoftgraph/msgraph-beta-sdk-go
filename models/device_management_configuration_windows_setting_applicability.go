package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationWindowsSettingApplicability 
type DeviceManagementConfigurationWindowsSettingApplicability struct {
    DeviceManagementConfigurationSettingApplicability
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementConfigurationWindowsSettingApplicability instantiates a new deviceManagementConfigurationWindowsSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationWindowsSettingApplicability()(*DeviceManagementConfigurationWindowsSettingApplicability) {
    m := &DeviceManagementConfigurationWindowsSettingApplicability{
        DeviceManagementConfigurationSettingApplicability: *NewDeviceManagementConfigurationSettingApplicability(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationWindowsSettingApplicability"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationWindowsSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationWindowsSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationWindowsSettingApplicability(), nil
}
// GetConfigurationServiceProviderVersion gets the configurationServiceProviderVersion property value. Version of CSP setting is a part of
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetConfigurationServiceProviderVersion()(*string) {
    val, err := m.GetBackingStore().Get("configurationServiceProviderVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingApplicability.GetFieldDeserializers()
    res["configurationServiceProviderVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationServiceProviderVersion(val)
        }
        return nil
    }
    res["maximumSupportedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumSupportedVersion(val)
        }
        return nil
    }
    res["minimumSupportedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedVersion(val)
        }
        return nil
    }
    res["requiredAzureAdTrustType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationAzureAdTrustType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAzureAdTrustType(val.(*DeviceManagementConfigurationAzureAdTrustType))
        }
        return nil
    }
    res["requiresAzureAd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresAzureAd(val)
        }
        return nil
    }
    res["windowsSkus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDeviceManagementConfigurationWindowsSkus)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationWindowsSkus, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*DeviceManagementConfigurationWindowsSkus))
                }
            }
            m.SetWindowsSkus(res)
        }
        return nil
    }
    return res
}
// GetMaximumSupportedVersion gets the maximumSupportedVersion property value. Maximum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetMaximumSupportedVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumSupportedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumSupportedVersion gets the minimumSupportedVersion property value. Minimum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetMinimumSupportedVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumSupportedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequiredAzureAdTrustType gets the requiredAzureAdTrustType property value. Required AAD Trust Type
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetRequiredAzureAdTrustType()(*DeviceManagementConfigurationAzureAdTrustType) {
    val, err := m.GetBackingStore().Get("requiredAzureAdTrustType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationAzureAdTrustType)
    }
    return nil
}
// GetRequiresAzureAd gets the requiresAzureAd property value. AzureAD setting requirement
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetRequiresAzureAd()(*bool) {
    val, err := m.GetBackingStore().Get("requiresAzureAd")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWindowsSkus gets the windowsSkus property value. List of Windows SKUs that the setting is applicable for
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetWindowsSkus()([]DeviceManagementConfigurationWindowsSkus) {
    val, err := m.GetBackingStore().Get("windowsSkus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationWindowsSkus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationWindowsSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingApplicability.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationServiceProviderVersion", m.GetConfigurationServiceProviderVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumSupportedVersion", m.GetMaximumSupportedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumSupportedVersion", m.GetMinimumSupportedVersion())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAzureAdTrustType() != nil {
        cast := (*m.GetRequiredAzureAdTrustType()).String()
        err = writer.WriteStringValue("requiredAzureAdTrustType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresAzureAd", m.GetRequiresAzureAd())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsSkus() != nil {
        err = writer.WriteCollectionOfStringValues("windowsSkus", SerializeDeviceManagementConfigurationWindowsSkus(m.GetWindowsSkus()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationServiceProviderVersion sets the configurationServiceProviderVersion property value. Version of CSP setting is a part of
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetConfigurationServiceProviderVersion(value *string)() {
    err := m.GetBackingStore().Set("configurationServiceProviderVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumSupportedVersion sets the maximumSupportedVersion property value. Maximum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetMaximumSupportedVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumSupportedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedVersion sets the minimumSupportedVersion property value. Minimum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetMinimumSupportedVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumSupportedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredAzureAdTrustType sets the requiredAzureAdTrustType property value. Required AAD Trust Type
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetRequiredAzureAdTrustType(value *DeviceManagementConfigurationAzureAdTrustType)() {
    err := m.GetBackingStore().Set("requiredAzureAdTrustType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiresAzureAd sets the requiresAzureAd property value. AzureAD setting requirement
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetRequiresAzureAd(value *bool)() {
    err := m.GetBackingStore().Set("requiresAzureAd", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSkus sets the windowsSkus property value. List of Windows SKUs that the setting is applicable for
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetWindowsSkus(value []DeviceManagementConfigurationWindowsSkus)() {
    err := m.GetBackingStore().Set("windowsSkus", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationWindowsSettingApplicabilityable 
type DeviceManagementConfigurationWindowsSettingApplicabilityable interface {
    DeviceManagementConfigurationSettingApplicabilityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationServiceProviderVersion()(*string)
    GetMaximumSupportedVersion()(*string)
    GetMinimumSupportedVersion()(*string)
    GetRequiredAzureAdTrustType()(*DeviceManagementConfigurationAzureAdTrustType)
    GetRequiresAzureAd()(*bool)
    GetWindowsSkus()([]DeviceManagementConfigurationWindowsSkus)
    SetConfigurationServiceProviderVersion(value *string)()
    SetMaximumSupportedVersion(value *string)()
    SetMinimumSupportedVersion(value *string)()
    SetRequiredAzureAdTrustType(value *DeviceManagementConfigurationAzureAdTrustType)()
    SetRequiresAzureAd(value *bool)()
    SetWindowsSkus(value []DeviceManagementConfigurationWindowsSkus)()
}
