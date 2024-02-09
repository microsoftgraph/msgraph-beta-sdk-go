package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting struct {
    CloudPcDisasterRecoveryNetworkSetting
}
// NewCloudPcDisasterRecoveryMicrosoftHostedNetworkSetting instantiates a new CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting and sets the default values.
func NewCloudPcDisasterRecoveryMicrosoftHostedNetworkSetting()(*CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) {
    m := &CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting{
        CloudPcDisasterRecoveryNetworkSetting: *NewCloudPcDisasterRecoveryNetworkSetting(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcDisasterRecoveryMicrosoftHostedNetworkSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcDisasterRecoveryMicrosoftHostedNetworkSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDisasterRecoveryMicrosoftHostedNetworkSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDisasterRecoveryMicrosoftHostedNetworkSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcDisasterRecoveryNetworkSetting.GetFieldDeserializers()
    res["regionGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRegionGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionGroup(val.(*CloudPcRegionGroup))
        }
        return nil
    }
    res["regionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionName(val)
        }
        return nil
    }
    return res
}
// GetRegionGroup gets the regionGroup property value. The regionGroup property
// returns a *CloudPcRegionGroup when successful
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) GetRegionGroup()(*CloudPcRegionGroup) {
    val, err := m.GetBackingStore().Get("regionGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcRegionGroup)
    }
    return nil
}
// GetRegionName gets the regionName property value. The regionName property
// returns a *string when successful
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) GetRegionName()(*string) {
    val, err := m.GetBackingStore().Get("regionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcDisasterRecoveryNetworkSetting.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRegionGroup() != nil {
        cast := (*m.GetRegionGroup()).String()
        err = writer.WriteStringValue("regionGroup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("regionName", m.GetRegionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegionGroup sets the regionGroup property value. The regionGroup property
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) SetRegionGroup(value *CloudPcRegionGroup)() {
    err := m.GetBackingStore().Set("regionGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionName sets the regionName property value. The regionName property
func (m *CloudPcDisasterRecoveryMicrosoftHostedNetworkSetting) SetRegionName(value *string)() {
    err := m.GetBackingStore().Set("regionName", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcDisasterRecoveryMicrosoftHostedNetworkSettingable interface {
    CloudPcDisasterRecoveryNetworkSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRegionGroup()(*CloudPcRegionGroup)
    GetRegionName()(*string)
    SetRegionGroup(value *CloudPcRegionGroup)()
    SetRegionName(value *string)()
}
