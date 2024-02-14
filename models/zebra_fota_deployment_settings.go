package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ZebraFotaDeploymentSettings the Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
type ZebraFotaDeploymentSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewZebraFotaDeploymentSettings instantiates a new ZebraFotaDeploymentSettings and sets the default values.
func NewZebraFotaDeploymentSettings()(*ZebraFotaDeploymentSettings) {
    m := &ZebraFotaDeploymentSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateZebraFotaDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateZebraFotaDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeploymentSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ZebraFotaDeploymentSettings) GetAdditionalData()(map[string]any) {
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
func (m *ZebraFotaDeploymentSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBatteryRuleMinimumBatteryLevelPercentage gets the batteryRuleMinimumBatteryLevelPercentage property value. Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
// returns a *int32 when successful
func (m *ZebraFotaDeploymentSettings) GetBatteryRuleMinimumBatteryLevelPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("batteryRuleMinimumBatteryLevelPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBatteryRuleRequireCharger gets the batteryRuleRequireCharger property value. Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or out of the charger. Applied only for installation. Defaults to false.
// returns a *bool when successful
func (m *ZebraFotaDeploymentSettings) GetBatteryRuleRequireCharger()(*bool) {
    val, err := m.GetBackingStore().Get("batteryRuleRequireCharger")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. Deploy update for devices with this model only.
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDownloadRuleNetworkType gets the downloadRuleNetworkType property value. Represents various network types for Zebra FOTA deployment.
// returns a *ZebraFotaNetworkType when successful
func (m *ZebraFotaDeploymentSettings) GetDownloadRuleNetworkType()(*ZebraFotaNetworkType) {
    val, err := m.GetBackingStore().Get("downloadRuleNetworkType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ZebraFotaNetworkType)
    }
    return nil
}
// GetDownloadRuleStartDateTime gets the downloadRuleStartDateTime property value. Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is UTC now and the maximum is 10 days from deployment creation.
// returns a *Time when successful
func (m *ZebraFotaDeploymentSettings) GetDownloadRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("downloadRuleStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ZebraFotaDeploymentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batteryRuleMinimumBatteryLevelPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuleMinimumBatteryLevelPercentage(val)
        }
        return nil
    }
    res["batteryRuleRequireCharger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuleRequireCharger(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["downloadRuleNetworkType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaNetworkType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadRuleNetworkType(val.(*ZebraFotaNetworkType))
        }
        return nil
    }
    res["downloadRuleStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadRuleStartDateTime(val)
        }
        return nil
    }
    res["firmwareTargetArtifactDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetArtifactDescription(val)
        }
        return nil
    }
    res["firmwareTargetBoardSupportPackageVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetBoardSupportPackageVersion(val)
        }
        return nil
    }
    res["firmwareTargetOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetOsVersion(val)
        }
        return nil
    }
    res["firmwareTargetPatch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareTargetPatch(val)
        }
        return nil
    }
    res["installRuleStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleStartDateTime(val)
        }
        return nil
    }
    res["installRuleWindowEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleWindowEndTime(val)
        }
        return nil
    }
    res["installRuleWindowStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallRuleWindowStartTime(val)
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
    res["scheduleDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleDurationInDays(val)
        }
        return nil
    }
    res["scheduleMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaScheduleMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleMode(val.(*ZebraFotaScheduleMode))
        }
        return nil
    }
    res["timeZoneOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZoneOffsetInMinutes(val)
        }
        return nil
    }
    res["updateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaUpdateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateType(val.(*ZebraFotaUpdateType))
        }
        return nil
    }
    return res
}
// GetFirmwareTargetArtifactDescription gets the firmwareTargetArtifactDescription property value. A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120 (released 29-June-2022).
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetArtifactDescription()(*string) {
    val, err := m.GetBackingStore().Get("firmwareTargetArtifactDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFirmwareTargetBoardSupportPackageVersion gets the firmwareTargetBoardSupportPackageVersion property value. Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetBoardSupportPackageVersion()(*string) {
    val, err := m.GetBackingStore().Get("firmwareTargetBoardSupportPackageVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFirmwareTargetOsVersion gets the firmwareTargetOsVersion property value. Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("firmwareTargetOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFirmwareTargetPatch gets the firmwareTargetPatch property value. Target patch name (e.g.: 'U06'). Required only for custom update type.
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetFirmwareTargetPatch()(*string) {
    val, err := m.GetBackingStore().Get("firmwareTargetPatch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInstallRuleStartDateTime gets the installRuleStartDateTime property value. Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise defaults to NOW. Ignored when deployment update type was set to auto.
// returns a *Time when successful
func (m *ZebraFotaDeploymentSettings) GetInstallRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("installRuleStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetInstallRuleWindowEndTime gets the installRuleWindowEndTime property value. Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
// returns a *TimeOnly when successful
func (m *ZebraFotaDeploymentSettings) GetInstallRuleWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("installRuleWindowEndTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetInstallRuleWindowStartTime gets the installRuleWindowStartTime property value. Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
// returns a *TimeOnly when successful
func (m *ZebraFotaDeploymentSettings) GetInstallRuleWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("installRuleWindowStartTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ZebraFotaDeploymentSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScheduleDurationInDays gets the scheduleDurationInDays property value. Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3) Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is used. If no values are provided, the string value of the current UTC is used.
// returns a *int32 when successful
func (m *ZebraFotaDeploymentSettings) GetScheduleDurationInDays()(*int32) {
    val, err := m.GetBackingStore().Get("scheduleDurationInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetScheduleMode gets the scheduleMode property value. Represents various schedule modes for Zebra FOTA deployment.
// returns a *ZebraFotaScheduleMode when successful
func (m *ZebraFotaDeploymentSettings) GetScheduleMode()(*ZebraFotaScheduleMode) {
    val, err := m.GetBackingStore().Get("scheduleMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ZebraFotaScheduleMode)
    }
    return nil
}
// GetTimeZoneOffsetInMinutes gets the timeZoneOffsetInMinutes property value. This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end data uses this timezone
// returns a *int32 when successful
func (m *ZebraFotaDeploymentSettings) GetTimeZoneOffsetInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("timeZoneOffsetInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUpdateType gets the updateType property value. Represents various update types for Zebra FOTA deployment.
// returns a *ZebraFotaUpdateType when successful
func (m *ZebraFotaDeploymentSettings) GetUpdateType()(*ZebraFotaUpdateType) {
    val, err := m.GetBackingStore().Get("updateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ZebraFotaUpdateType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ZebraFotaDeploymentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("batteryRuleMinimumBatteryLevelPercentage", m.GetBatteryRuleMinimumBatteryLevelPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("batteryRuleRequireCharger", m.GetBatteryRuleRequireCharger())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    if m.GetDownloadRuleNetworkType() != nil {
        cast := (*m.GetDownloadRuleNetworkType()).String()
        err := writer.WriteStringValue("downloadRuleNetworkType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("downloadRuleStartDateTime", m.GetDownloadRuleStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetArtifactDescription", m.GetFirmwareTargetArtifactDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetBoardSupportPackageVersion", m.GetFirmwareTargetBoardSupportPackageVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetOsVersion", m.GetFirmwareTargetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareTargetPatch", m.GetFirmwareTargetPatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("installRuleStartDateTime", m.GetInstallRuleStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("installRuleWindowEndTime", m.GetInstallRuleWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("installRuleWindowStartTime", m.GetInstallRuleWindowStartTime())
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
        err := writer.WriteInt32Value("scheduleDurationInDays", m.GetScheduleDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetScheduleMode() != nil {
        cast := (*m.GetScheduleMode()).String()
        err := writer.WriteStringValue("scheduleMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeZoneOffsetInMinutes", m.GetTimeZoneOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateType() != nil {
        cast := (*m.GetUpdateType()).String()
        err := writer.WriteStringValue("updateType", &cast)
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
func (m *ZebraFotaDeploymentSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ZebraFotaDeploymentSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBatteryRuleMinimumBatteryLevelPercentage sets the batteryRuleMinimumBatteryLevelPercentage property value. Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
func (m *ZebraFotaDeploymentSettings) SetBatteryRuleMinimumBatteryLevelPercentage(value *int32)() {
    err := m.GetBackingStore().Set("batteryRuleMinimumBatteryLevelPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryRuleRequireCharger sets the batteryRuleRequireCharger property value. Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or out of the charger. Applied only for installation. Defaults to false.
func (m *ZebraFotaDeploymentSettings) SetBatteryRuleRequireCharger(value *bool)() {
    err := m.GetBackingStore().Set("batteryRuleRequireCharger", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. Deploy update for devices with this model only.
func (m *ZebraFotaDeploymentSettings) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDownloadRuleNetworkType sets the downloadRuleNetworkType property value. Represents various network types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetDownloadRuleNetworkType(value *ZebraFotaNetworkType)() {
    err := m.GetBackingStore().Set("downloadRuleNetworkType", value)
    if err != nil {
        panic(err)
    }
}
// SetDownloadRuleStartDateTime sets the downloadRuleStartDateTime property value. Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is UTC now and the maximum is 10 days from deployment creation.
func (m *ZebraFotaDeploymentSettings) SetDownloadRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("downloadRuleStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareTargetArtifactDescription sets the firmwareTargetArtifactDescription property value. A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120 (released 29-June-2022).
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetArtifactDescription(value *string)() {
    err := m.GetBackingStore().Set("firmwareTargetArtifactDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareTargetBoardSupportPackageVersion sets the firmwareTargetBoardSupportPackageVersion property value. Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetBoardSupportPackageVersion(value *string)() {
    err := m.GetBackingStore().Set("firmwareTargetBoardSupportPackageVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareTargetOsVersion sets the firmwareTargetOsVersion property value. Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("firmwareTargetOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareTargetPatch sets the firmwareTargetPatch property value. Target patch name (e.g.: 'U06'). Required only for custom update type.
func (m *ZebraFotaDeploymentSettings) SetFirmwareTargetPatch(value *string)() {
    err := m.GetBackingStore().Set("firmwareTargetPatch", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallRuleStartDateTime sets the installRuleStartDateTime property value. Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise defaults to NOW. Ignored when deployment update type was set to auto.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("installRuleStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallRuleWindowEndTime sets the installRuleWindowEndTime property value. Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("installRuleWindowEndTime", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallRuleWindowStartTime sets the installRuleWindowStartTime property value. Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
func (m *ZebraFotaDeploymentSettings) SetInstallRuleWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("installRuleWindowStartTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduleDurationInDays sets the scheduleDurationInDays property value. Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3) Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is used. If no values are provided, the string value of the current UTC is used.
func (m *ZebraFotaDeploymentSettings) SetScheduleDurationInDays(value *int32)() {
    err := m.GetBackingStore().Set("scheduleDurationInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduleMode sets the scheduleMode property value. Represents various schedule modes for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetScheduleMode(value *ZebraFotaScheduleMode)() {
    err := m.GetBackingStore().Set("scheduleMode", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeZoneOffsetInMinutes sets the timeZoneOffsetInMinutes property value. This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end data uses this timezone
func (m *ZebraFotaDeploymentSettings) SetTimeZoneOffsetInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("timeZoneOffsetInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateType sets the updateType property value. Represents various update types for Zebra FOTA deployment.
func (m *ZebraFotaDeploymentSettings) SetUpdateType(value *ZebraFotaUpdateType)() {
    err := m.GetBackingStore().Set("updateType", value)
    if err != nil {
        panic(err)
    }
}
type ZebraFotaDeploymentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBatteryRuleMinimumBatteryLevelPercentage()(*int32)
    GetBatteryRuleRequireCharger()(*bool)
    GetDeviceModel()(*string)
    GetDownloadRuleNetworkType()(*ZebraFotaNetworkType)
    GetDownloadRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFirmwareTargetArtifactDescription()(*string)
    GetFirmwareTargetBoardSupportPackageVersion()(*string)
    GetFirmwareTargetOsVersion()(*string)
    GetFirmwareTargetPatch()(*string)
    GetInstallRuleStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstallRuleWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetInstallRuleWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetOdataType()(*string)
    GetScheduleDurationInDays()(*int32)
    GetScheduleMode()(*ZebraFotaScheduleMode)
    GetTimeZoneOffsetInMinutes()(*int32)
    GetUpdateType()(*ZebraFotaUpdateType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBatteryRuleMinimumBatteryLevelPercentage(value *int32)()
    SetBatteryRuleRequireCharger(value *bool)()
    SetDeviceModel(value *string)()
    SetDownloadRuleNetworkType(value *ZebraFotaNetworkType)()
    SetDownloadRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFirmwareTargetArtifactDescription(value *string)()
    SetFirmwareTargetBoardSupportPackageVersion(value *string)()
    SetFirmwareTargetOsVersion(value *string)()
    SetFirmwareTargetPatch(value *string)()
    SetInstallRuleStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstallRuleWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetInstallRuleWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetOdataType(value *string)()
    SetScheduleDurationInDays(value *int32)()
    SetScheduleMode(value *ZebraFotaScheduleMode)()
    SetTimeZoneOffsetInMinutes(value *int32)()
    SetUpdateType(value *ZebraFotaUpdateType)()
}
