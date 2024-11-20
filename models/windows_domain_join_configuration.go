package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDomainJoinConfiguration windows Domain Join device configuration.
type WindowsDomainJoinConfiguration struct {
    DeviceConfiguration
}
// NewWindowsDomainJoinConfiguration instantiates a new WindowsDomainJoinConfiguration and sets the default values.
func NewWindowsDomainJoinConfiguration()(*WindowsDomainJoinConfiguration) {
    m := &WindowsDomainJoinConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDomainJoinConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsDomainJoinConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsDomainJoinConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDomainJoinConfiguration(), nil
}
// GetActiveDirectoryDomainName gets the activeDirectoryDomainName property value. Active Directory domain name to join.
// returns a *string when successful
func (m *WindowsDomainJoinConfiguration) GetActiveDirectoryDomainName()(*string) {
    val, err := m.GetBackingStore().Get("activeDirectoryDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComputerNameStaticPrefix gets the computerNameStaticPrefix property value. Fixed prefix to be used for computer name.
// returns a *string when successful
func (m *WindowsDomainJoinConfiguration) GetComputerNameStaticPrefix()(*string) {
    val, err := m.GetBackingStore().Get("computerNameStaticPrefix")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComputerNameSuffixRandomCharCount gets the computerNameSuffixRandomCharCount property value. Dynamically generated characters used as suffix for computer name. Valid values 3 to 14
// returns a *int32 when successful
func (m *WindowsDomainJoinConfiguration) GetComputerNameSuffixRandomCharCount()(*int32) {
    val, err := m.GetBackingStore().Get("computerNameSuffixRandomCharCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsDomainJoinConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["activeDirectoryDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDirectoryDomainName(val)
        }
        return nil
    }
    res["computerNameStaticPrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputerNameStaticPrefix(val)
        }
        return nil
    }
    res["computerNameSuffixRandomCharCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputerNameSuffixRandomCharCount(val)
        }
        return nil
    }
    res["networkAccessConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationable)
                }
            }
            m.SetNetworkAccessConfigurations(res)
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val)
        }
        return nil
    }
    return res
}
// GetNetworkAccessConfigurations gets the networkAccessConfigurations property value. Reference to device configurations required for network connectivity. This collection can contain a maximum of 2 elements.
// returns a []DeviceConfigurationable when successful
func (m *WindowsDomainJoinConfiguration) GetNetworkAccessConfigurations()([]DeviceConfigurationable) {
    val, err := m.GetBackingStore().Get("networkAccessConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceConfigurationable)
    }
    return nil
}
// GetOrganizationalUnit gets the organizationalUnit property value. Organizational unit (OU) where the computer account will be created. If this parameter is NULL, the well known computer object container will be used as published in the domain.
// returns a *string when successful
func (m *WindowsDomainJoinConfiguration) GetOrganizationalUnit()(*string) {
    val, err := m.GetBackingStore().Get("organizationalUnit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsDomainJoinConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activeDirectoryDomainName", m.GetActiveDirectoryDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("computerNameStaticPrefix", m.GetComputerNameStaticPrefix())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("computerNameSuffixRandomCharCount", m.GetComputerNameSuffixRandomCharCount())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkAccessConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkAccessConfigurations()))
        for i, v := range m.GetNetworkAccessConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("networkAccessConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDirectoryDomainName sets the activeDirectoryDomainName property value. Active Directory domain name to join.
func (m *WindowsDomainJoinConfiguration) SetActiveDirectoryDomainName(value *string)() {
    err := m.GetBackingStore().Set("activeDirectoryDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetComputerNameStaticPrefix sets the computerNameStaticPrefix property value. Fixed prefix to be used for computer name.
func (m *WindowsDomainJoinConfiguration) SetComputerNameStaticPrefix(value *string)() {
    err := m.GetBackingStore().Set("computerNameStaticPrefix", value)
    if err != nil {
        panic(err)
    }
}
// SetComputerNameSuffixRandomCharCount sets the computerNameSuffixRandomCharCount property value. Dynamically generated characters used as suffix for computer name. Valid values 3 to 14
func (m *WindowsDomainJoinConfiguration) SetComputerNameSuffixRandomCharCount(value *int32)() {
    err := m.GetBackingStore().Set("computerNameSuffixRandomCharCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkAccessConfigurations sets the networkAccessConfigurations property value. Reference to device configurations required for network connectivity. This collection can contain a maximum of 2 elements.
func (m *WindowsDomainJoinConfiguration) SetNetworkAccessConfigurations(value []DeviceConfigurationable)() {
    err := m.GetBackingStore().Set("networkAccessConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnit sets the organizationalUnit property value. Organizational unit (OU) where the computer account will be created. If this parameter is NULL, the well known computer object container will be used as published in the domain.
func (m *WindowsDomainJoinConfiguration) SetOrganizationalUnit(value *string)() {
    err := m.GetBackingStore().Set("organizationalUnit", value)
    if err != nil {
        panic(err)
    }
}
type WindowsDomainJoinConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDirectoryDomainName()(*string)
    GetComputerNameStaticPrefix()(*string)
    GetComputerNameSuffixRandomCharCount()(*int32)
    GetNetworkAccessConfigurations()([]DeviceConfigurationable)
    GetOrganizationalUnit()(*string)
    SetActiveDirectoryDomainName(value *string)()
    SetComputerNameStaticPrefix(value *string)()
    SetComputerNameSuffixRandomCharCount(value *int32)()
    SetNetworkAccessConfigurations(value []DeviceConfigurationable)()
    SetOrganizationalUnit(value *string)()
}
