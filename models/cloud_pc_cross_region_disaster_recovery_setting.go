package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudPcCrossRegionDisasterRecoverySetting struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcCrossRegionDisasterRecoverySetting instantiates a new CloudPcCrossRegionDisasterRecoverySetting and sets the default values.
func NewCloudPcCrossRegionDisasterRecoverySetting()(*CloudPcCrossRegionDisasterRecoverySetting) {
    m := &CloudPcCrossRegionDisasterRecoverySetting{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcCrossRegionDisasterRecoverySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcCrossRegionDisasterRecoverySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcCrossRegionDisasterRecoverySetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCrossRegionDisasterRecoveryEnabled gets the crossRegionDisasterRecoveryEnabled property value. True if an end user is allowed to set up cross-region disaster recovery for Cloud PC; otherwise, false. The default value is false. This property is deprecated and will no longer be supported effective February 11, 2025. For scenarios where crossRegionDisasterRecoveryEnabled is true, set disasterRecoveryType to crossRegion. For scenarios where crossRegionDisasterRecoveryEnabled is false,  set disasterRecoveryType to notconfigured.
// returns a *bool when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetCrossRegionDisasterRecoveryEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("crossRegionDisasterRecoveryEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisasterRecoveryNetworkSetting gets the disasterRecoveryNetworkSetting property value. Indicates the network settings of the Cloud PC during a cross-region disaster recovery operation.
// returns a CloudPcDisasterRecoveryNetworkSettingable when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetDisasterRecoveryNetworkSetting()(CloudPcDisasterRecoveryNetworkSettingable) {
    val, err := m.GetBackingStore().Get("disasterRecoveryNetworkSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcDisasterRecoveryNetworkSettingable)
    }
    return nil
}
// GetDisasterRecoveryType gets the disasterRecoveryType property value. Indicates the type of disaster recovery to perform when a disaster occurs on the user's Cloud PC. The possible values are: notConfigured, crossRegion, premium, unknownFutureValue. The default value is notConfigured.
// returns a *CloudPcDisasterRecoveryType when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetDisasterRecoveryType()(*CloudPcDisasterRecoveryType) {
    val, err := m.GetBackingStore().Get("disasterRecoveryType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDisasterRecoveryType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["crossRegionDisasterRecoveryEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossRegionDisasterRecoveryEnabled(val)
        }
        return nil
    }
    res["disasterRecoveryNetworkSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcDisasterRecoveryNetworkSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisasterRecoveryNetworkSetting(val.(CloudPcDisasterRecoveryNetworkSettingable))
        }
        return nil
    }
    res["disasterRecoveryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDisasterRecoveryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisasterRecoveryType(val.(*CloudPcDisasterRecoveryType))
        }
        return nil
    }
    res["maintainCrossRegionRestorePointEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintainCrossRegionRestorePointEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetMaintainCrossRegionRestorePointEnabled gets the maintainCrossRegionRestorePointEnabled property value. Indicates whether Windows 365 maintain the cross-region disaster recovery function generated restore points. If true, the Windows 365 stored restore points; false indicates that Windows 365 doesn't generate or keep the restore point from the original Cloud PC. If a disaster occurs, the new Cloud PC can only be provisioned using the initial image. This limitation can result in the loss of some user data on the original Cloud PC. The default value is false.
// returns a *bool when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetMaintainCrossRegionRestorePointEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("maintainCrossRegionRestorePointEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcCrossRegionDisasterRecoverySetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcCrossRegionDisasterRecoverySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("crossRegionDisasterRecoveryEnabled", m.GetCrossRegionDisasterRecoveryEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("disasterRecoveryNetworkSetting", m.GetDisasterRecoveryNetworkSetting())
        if err != nil {
            return err
        }
    }
    if m.GetDisasterRecoveryType() != nil {
        cast := (*m.GetDisasterRecoveryType()).String()
        err := writer.WriteStringValue("disasterRecoveryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maintainCrossRegionRestorePointEnabled", m.GetMaintainCrossRegionRestorePointEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCrossRegionDisasterRecoveryEnabled sets the crossRegionDisasterRecoveryEnabled property value. True if an end user is allowed to set up cross-region disaster recovery for Cloud PC; otherwise, false. The default value is false. This property is deprecated and will no longer be supported effective February 11, 2025. For scenarios where crossRegionDisasterRecoveryEnabled is true, set disasterRecoveryType to crossRegion. For scenarios where crossRegionDisasterRecoveryEnabled is false,  set disasterRecoveryType to notconfigured.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetCrossRegionDisasterRecoveryEnabled(value *bool)() {
    err := m.GetBackingStore().Set("crossRegionDisasterRecoveryEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDisasterRecoveryNetworkSetting sets the disasterRecoveryNetworkSetting property value. Indicates the network settings of the Cloud PC during a cross-region disaster recovery operation.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetDisasterRecoveryNetworkSetting(value CloudPcDisasterRecoveryNetworkSettingable)() {
    err := m.GetBackingStore().Set("disasterRecoveryNetworkSetting", value)
    if err != nil {
        panic(err)
    }
}
// SetDisasterRecoveryType sets the disasterRecoveryType property value. Indicates the type of disaster recovery to perform when a disaster occurs on the user's Cloud PC. The possible values are: notConfigured, crossRegion, premium, unknownFutureValue. The default value is notConfigured.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetDisasterRecoveryType(value *CloudPcDisasterRecoveryType)() {
    err := m.GetBackingStore().Set("disasterRecoveryType", value)
    if err != nil {
        panic(err)
    }
}
// SetMaintainCrossRegionRestorePointEnabled sets the maintainCrossRegionRestorePointEnabled property value. Indicates whether Windows 365 maintain the cross-region disaster recovery function generated restore points. If true, the Windows 365 stored restore points; false indicates that Windows 365 doesn't generate or keep the restore point from the original Cloud PC. If a disaster occurs, the new Cloud PC can only be provisioned using the initial image. This limitation can result in the loss of some user data on the original Cloud PC. The default value is false.
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetMaintainCrossRegionRestorePointEnabled(value *bool)() {
    err := m.GetBackingStore().Set("maintainCrossRegionRestorePointEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcCrossRegionDisasterRecoverySetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcCrossRegionDisasterRecoverySettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCrossRegionDisasterRecoveryEnabled()(*bool)
    GetDisasterRecoveryNetworkSetting()(CloudPcDisasterRecoveryNetworkSettingable)
    GetDisasterRecoveryType()(*CloudPcDisasterRecoveryType)
    GetMaintainCrossRegionRestorePointEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCrossRegionDisasterRecoveryEnabled(value *bool)()
    SetDisasterRecoveryNetworkSetting(value CloudPcDisasterRecoveryNetworkSettingable)()
    SetDisasterRecoveryType(value *CloudPcDisasterRecoveryType)()
    SetMaintainCrossRegionRestorePointEnabled(value *bool)()
    SetOdataType(value *string)()
}
