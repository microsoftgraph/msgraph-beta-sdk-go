package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcDisasterRecoveryAzureConnectionSetting struct {
    CloudPcDisasterRecoveryNetworkSetting
}
// NewCloudPcDisasterRecoveryAzureConnectionSetting instantiates a new CloudPcDisasterRecoveryAzureConnectionSetting and sets the default values.
func NewCloudPcDisasterRecoveryAzureConnectionSetting()(*CloudPcDisasterRecoveryAzureConnectionSetting) {
    m := &CloudPcDisasterRecoveryAzureConnectionSetting{
        CloudPcDisasterRecoveryNetworkSetting: *NewCloudPcDisasterRecoveryNetworkSetting(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcDisasterRecoveryAzureConnectionSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcDisasterRecoveryAzureConnectionSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDisasterRecoveryAzureConnectionSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDisasterRecoveryAzureConnectionSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDisasterRecoveryAzureConnectionSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcDisasterRecoveryNetworkSetting.GetFieldDeserializers()
    res["onPremisesConnectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    return res
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The onPremisesConnectionId property
// returns a *string when successful
func (m *CloudPcDisasterRecoveryAzureConnectionSetting) GetOnPremisesConnectionId()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesConnectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcDisasterRecoveryAzureConnectionSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcDisasterRecoveryNetworkSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The onPremisesConnectionId property
func (m *CloudPcDisasterRecoveryAzureConnectionSetting) SetOnPremisesConnectionId(value *string)() {
    err := m.GetBackingStore().Set("onPremisesConnectionId", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcDisasterRecoveryAzureConnectionSettingable interface {
    CloudPcDisasterRecoveryNetworkSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOnPremisesConnectionId()(*string)
    SetOnPremisesConnectionId(value *string)()
}
