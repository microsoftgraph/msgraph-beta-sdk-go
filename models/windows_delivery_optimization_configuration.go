package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDeliveryOptimizationConfiguration windows Delivery Optimization configuration
type WindowsDeliveryOptimizationConfiguration struct {
    DeviceConfiguration
}
// NewWindowsDeliveryOptimizationConfiguration instantiates a new windowsDeliveryOptimizationConfiguration and sets the default values.
func NewWindowsDeliveryOptimizationConfiguration()(*WindowsDeliveryOptimizationConfiguration) {
    m := &WindowsDeliveryOptimizationConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDeliveryOptimizationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsDeliveryOptimizationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDeliveryOptimizationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDeliveryOptimizationConfiguration(), nil
}
// GetBackgroundDownloadFromHttpDelayInSeconds gets the backgroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer. Valid values 0 to 4294967295
func (m *WindowsDeliveryOptimizationConfiguration) GetBackgroundDownloadFromHttpDelayInSeconds()(*int64) {
    val, err := m.GetBackingStore().Get("backgroundDownloadFromHttpDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetBandwidthMode gets the bandwidthMode property value. Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
func (m *WindowsDeliveryOptimizationConfiguration) GetBandwidthMode()(DeliveryOptimizationBandwidthable) {
    val, err := m.GetBackingStore().Get("bandwidthMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeliveryOptimizationBandwidthable)
    }
    return nil
}
// GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds gets the cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download. Valid values 0 to 2592000.
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds gets the cacheServerForegroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download. Valid values 0 to 2592000.​
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("cacheServerForegroundDownloadFallbackToHttpDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCacheServerHostNames gets the cacheServerHostNames property value. Specifies cache servers host names.
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerHostNames()([]string) {
    val, err := m.GetBackingStore().Get("cacheServerHostNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDeliveryOptimizationMode gets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsDeliveryOptimizationConfiguration) GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode) {
    val, err := m.GetBackingStore().Get("deliveryOptimizationMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsDeliveryOptimizationMode)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDeliveryOptimizationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["backgroundDownloadFromHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundDownloadFromHttpDelayInSeconds(val)
        }
        return nil
    }
    res["bandwidthMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationBandwidthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthMode(val.(DeliveryOptimizationBandwidthable))
        }
        return nil
    }
    res["cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(val)
        }
        return nil
    }
    res["cacheServerForegroundDownloadFallbackToHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(val)
        }
        return nil
    }
    res["cacheServerHostNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCacheServerHostNames(res)
        }
        return nil
    }
    res["deliveryOptimizationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeliveryOptimizationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryOptimizationMode(val.(*WindowsDeliveryOptimizationMode))
        }
        return nil
    }
    res["foregroundDownloadFromHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForegroundDownloadFromHttpDelayInSeconds(val)
        }
        return nil
    }
    res["groupIdSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationGroupIdSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupIdSource(val.(DeliveryOptimizationGroupIdSourceable))
        }
        return nil
    }
    res["maximumCacheAgeInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheAgeInDays(val)
        }
        return nil
    }
    res["maximumCacheSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationMaxCacheSizeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheSize(val.(DeliveryOptimizationMaxCacheSizeable))
        }
        return nil
    }
    res["minimumBatteryPercentageAllowedToUpload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumBatteryPercentageAllowedToUpload(val)
        }
        return nil
    }
    res["minimumDiskSizeAllowedToPeerInGigabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumDiskSizeAllowedToPeerInGigabytes(val)
        }
        return nil
    }
    res["minimumFileSizeToCacheInMegabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumFileSizeToCacheInMegabytes(val)
        }
        return nil
    }
    res["minimumRamAllowedToPeerInGigabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRamAllowedToPeerInGigabytes(val)
        }
        return nil
    }
    res["modifyCacheLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifyCacheLocation(val)
        }
        return nil
    }
    res["restrictPeerSelectionBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeliveryOptimizationRestrictPeerSelectionByOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictPeerSelectionBy(val.(*DeliveryOptimizationRestrictPeerSelectionByOptions))
        }
        return nil
    }
    res["vpnPeerCaching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnPeerCaching(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetForegroundDownloadFromHttpDelayInSeconds gets the foregroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer (0-86400). Valid values 0 to 86400
func (m *WindowsDeliveryOptimizationConfiguration) GetForegroundDownloadFromHttpDelayInSeconds()(*int64) {
    val, err := m.GetBackingStore().Get("foregroundDownloadFromHttpDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetGroupIdSource gets the groupIdSource property value. Specifies to restrict peer selection to a specfic source.
func (m *WindowsDeliveryOptimizationConfiguration) GetGroupIdSource()(DeliveryOptimizationGroupIdSourceable) {
    val, err := m.GetBackingStore().Get("groupIdSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeliveryOptimizationGroupIdSourceable)
    }
    return nil
}
// GetMaximumCacheAgeInDays gets the maximumCacheAgeInDays property value. Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading successfully (0-3650). Valid values 0 to 3650
func (m *WindowsDeliveryOptimizationConfiguration) GetMaximumCacheAgeInDays()(*int32) {
    val, err := m.GetBackingStore().Get("maximumCacheAgeInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumCacheSize gets the maximumCacheSize property value. Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
func (m *WindowsDeliveryOptimizationConfiguration) GetMaximumCacheSize()(DeliveryOptimizationMaxCacheSizeable) {
    val, err := m.GetBackingStore().Get("maximumCacheSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeliveryOptimizationMaxCacheSizeable)
    }
    return nil
}
// GetMinimumBatteryPercentageAllowedToUpload gets the minimumBatteryPercentageAllowedToUpload property value. Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumBatteryPercentageAllowedToUpload()(*int32) {
    val, err := m.GetBackingStore().Get("minimumBatteryPercentageAllowedToUpload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumDiskSizeAllowedToPeerInGigabytes gets the minimumDiskSizeAllowedToPeerInGigabytes property value. Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumDiskSizeAllowedToPeerInGigabytes()(*int32) {
    val, err := m.GetBackingStore().Get("minimumDiskSizeAllowedToPeerInGigabytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumFileSizeToCacheInMegabytes gets the minimumFileSizeToCacheInMegabytes property value. Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumFileSizeToCacheInMegabytes()(*int32) {
    val, err := m.GetBackingStore().Get("minimumFileSizeToCacheInMegabytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumRamAllowedToPeerInGigabytes gets the minimumRamAllowedToPeerInGigabytes property value. Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumRamAllowedToPeerInGigabytes()(*int32) {
    val, err := m.GetBackingStore().Get("minimumRamAllowedToPeerInGigabytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetModifyCacheLocation gets the modifyCacheLocation property value. Specifies the drive that Delivery Optimization should use for its cache.
func (m *WindowsDeliveryOptimizationConfiguration) GetModifyCacheLocation()(*string) {
    val, err := m.GetBackingStore().Get("modifyCacheLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestrictPeerSelectionBy gets the restrictPeerSelectionBy property value. Values to restrict peer selection by.
func (m *WindowsDeliveryOptimizationConfiguration) GetRestrictPeerSelectionBy()(*DeliveryOptimizationRestrictPeerSelectionByOptions) {
    val, err := m.GetBackingStore().Get("restrictPeerSelectionBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeliveryOptimizationRestrictPeerSelectionByOptions)
    }
    return nil
}
// GetVpnPeerCaching gets the vpnPeerCaching property value. Possible values of a property
func (m *WindowsDeliveryOptimizationConfiguration) GetVpnPeerCaching()(*Enablement) {
    val, err := m.GetBackingStore().Get("vpnPeerCaching")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsDeliveryOptimizationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("backgroundDownloadFromHttpDelayInSeconds", m.GetBackgroundDownloadFromHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bandwidthMode", m.GetBandwidthMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds", m.GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cacheServerForegroundDownloadFallbackToHttpDelayInSeconds", m.GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetCacheServerHostNames() != nil {
        err = writer.WriteCollectionOfStringValues("cacheServerHostNames", m.GetCacheServerHostNames())
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryOptimizationMode() != nil {
        cast := (*m.GetDeliveryOptimizationMode()).String()
        err = writer.WriteStringValue("deliveryOptimizationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("foregroundDownloadFromHttpDelayInSeconds", m.GetForegroundDownloadFromHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("groupIdSource", m.GetGroupIdSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumCacheAgeInDays", m.GetMaximumCacheAgeInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximumCacheSize", m.GetMaximumCacheSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumBatteryPercentageAllowedToUpload", m.GetMinimumBatteryPercentageAllowedToUpload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumDiskSizeAllowedToPeerInGigabytes", m.GetMinimumDiskSizeAllowedToPeerInGigabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumFileSizeToCacheInMegabytes", m.GetMinimumFileSizeToCacheInMegabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumRamAllowedToPeerInGigabytes", m.GetMinimumRamAllowedToPeerInGigabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifyCacheLocation", m.GetModifyCacheLocation())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictPeerSelectionBy() != nil {
        cast := (*m.GetRestrictPeerSelectionBy()).String()
        err = writer.WriteStringValue("restrictPeerSelectionBy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVpnPeerCaching() != nil {
        cast := (*m.GetVpnPeerCaching()).String()
        err = writer.WriteStringValue("vpnPeerCaching", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackgroundDownloadFromHttpDelayInSeconds sets the backgroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer. Valid values 0 to 4294967295
func (m *WindowsDeliveryOptimizationConfiguration) SetBackgroundDownloadFromHttpDelayInSeconds(value *int64)() {
    err := m.GetBackingStore().Set("backgroundDownloadFromHttpDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetBandwidthMode sets the bandwidthMode property value. Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
func (m *WindowsDeliveryOptimizationConfiguration) SetBandwidthMode(value DeliveryOptimizationBandwidthable)() {
    err := m.GetBackingStore().Set("bandwidthMode", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds sets the cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download. Valid values 0 to 2592000.
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds sets the cacheServerForegroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download. Valid values 0 to 2592000.​
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("cacheServerForegroundDownloadFallbackToHttpDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetCacheServerHostNames sets the cacheServerHostNames property value. Specifies cache servers host names.
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerHostNames(value []string)() {
    err := m.GetBackingStore().Set("cacheServerHostNames", value)
    if err != nil {
        panic(err)
    }
}
// SetDeliveryOptimizationMode sets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsDeliveryOptimizationConfiguration) SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)() {
    err := m.GetBackingStore().Set("deliveryOptimizationMode", value)
    if err != nil {
        panic(err)
    }
}
// SetForegroundDownloadFromHttpDelayInSeconds sets the foregroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer (0-86400). Valid values 0 to 86400
func (m *WindowsDeliveryOptimizationConfiguration) SetForegroundDownloadFromHttpDelayInSeconds(value *int64)() {
    err := m.GetBackingStore().Set("foregroundDownloadFromHttpDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupIdSource sets the groupIdSource property value. Specifies to restrict peer selection to a specfic source.
func (m *WindowsDeliveryOptimizationConfiguration) SetGroupIdSource(value DeliveryOptimizationGroupIdSourceable)() {
    err := m.GetBackingStore().Set("groupIdSource", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumCacheAgeInDays sets the maximumCacheAgeInDays property value. Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading successfully (0-3650). Valid values 0 to 3650
func (m *WindowsDeliveryOptimizationConfiguration) SetMaximumCacheAgeInDays(value *int32)() {
    err := m.GetBackingStore().Set("maximumCacheAgeInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumCacheSize sets the maximumCacheSize property value. Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
func (m *WindowsDeliveryOptimizationConfiguration) SetMaximumCacheSize(value DeliveryOptimizationMaxCacheSizeable)() {
    err := m.GetBackingStore().Set("maximumCacheSize", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumBatteryPercentageAllowedToUpload sets the minimumBatteryPercentageAllowedToUpload property value. Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumBatteryPercentageAllowedToUpload(value *int32)() {
    err := m.GetBackingStore().Set("minimumBatteryPercentageAllowedToUpload", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumDiskSizeAllowedToPeerInGigabytes sets the minimumDiskSizeAllowedToPeerInGigabytes property value. Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumDiskSizeAllowedToPeerInGigabytes(value *int32)() {
    err := m.GetBackingStore().Set("minimumDiskSizeAllowedToPeerInGigabytes", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumFileSizeToCacheInMegabytes sets the minimumFileSizeToCacheInMegabytes property value. Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumFileSizeToCacheInMegabytes(value *int32)() {
    err := m.GetBackingStore().Set("minimumFileSizeToCacheInMegabytes", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRamAllowedToPeerInGigabytes sets the minimumRamAllowedToPeerInGigabytes property value. Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumRamAllowedToPeerInGigabytes(value *int32)() {
    err := m.GetBackingStore().Set("minimumRamAllowedToPeerInGigabytes", value)
    if err != nil {
        panic(err)
    }
}
// SetModifyCacheLocation sets the modifyCacheLocation property value. Specifies the drive that Delivery Optimization should use for its cache.
func (m *WindowsDeliveryOptimizationConfiguration) SetModifyCacheLocation(value *string)() {
    err := m.GetBackingStore().Set("modifyCacheLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetRestrictPeerSelectionBy sets the restrictPeerSelectionBy property value. Values to restrict peer selection by.
func (m *WindowsDeliveryOptimizationConfiguration) SetRestrictPeerSelectionBy(value *DeliveryOptimizationRestrictPeerSelectionByOptions)() {
    err := m.GetBackingStore().Set("restrictPeerSelectionBy", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnPeerCaching sets the vpnPeerCaching property value. Possible values of a property
func (m *WindowsDeliveryOptimizationConfiguration) SetVpnPeerCaching(value *Enablement)() {
    err := m.GetBackingStore().Set("vpnPeerCaching", value)
    if err != nil {
        panic(err)
    }
}
// WindowsDeliveryOptimizationConfigurationable 
type WindowsDeliveryOptimizationConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackgroundDownloadFromHttpDelayInSeconds()(*int64)
    GetBandwidthMode()(DeliveryOptimizationBandwidthable)
    GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds()(*int32)
    GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds()(*int32)
    GetCacheServerHostNames()([]string)
    GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode)
    GetForegroundDownloadFromHttpDelayInSeconds()(*int64)
    GetGroupIdSource()(DeliveryOptimizationGroupIdSourceable)
    GetMaximumCacheAgeInDays()(*int32)
    GetMaximumCacheSize()(DeliveryOptimizationMaxCacheSizeable)
    GetMinimumBatteryPercentageAllowedToUpload()(*int32)
    GetMinimumDiskSizeAllowedToPeerInGigabytes()(*int32)
    GetMinimumFileSizeToCacheInMegabytes()(*int32)
    GetMinimumRamAllowedToPeerInGigabytes()(*int32)
    GetModifyCacheLocation()(*string)
    GetRestrictPeerSelectionBy()(*DeliveryOptimizationRestrictPeerSelectionByOptions)
    GetVpnPeerCaching()(*Enablement)
    SetBackgroundDownloadFromHttpDelayInSeconds(value *int64)()
    SetBandwidthMode(value DeliveryOptimizationBandwidthable)()
    SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(value *int32)()
    SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(value *int32)()
    SetCacheServerHostNames(value []string)()
    SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)()
    SetForegroundDownloadFromHttpDelayInSeconds(value *int64)()
    SetGroupIdSource(value DeliveryOptimizationGroupIdSourceable)()
    SetMaximumCacheAgeInDays(value *int32)()
    SetMaximumCacheSize(value DeliveryOptimizationMaxCacheSizeable)()
    SetMinimumBatteryPercentageAllowedToUpload(value *int32)()
    SetMinimumDiskSizeAllowedToPeerInGigabytes(value *int32)()
    SetMinimumFileSizeToCacheInMegabytes(value *int32)()
    SetMinimumRamAllowedToPeerInGigabytes(value *int32)()
    SetModifyCacheLocation(value *string)()
    SetRestrictPeerSelectionBy(value *DeliveryOptimizationRestrictPeerSelectionByOptions)()
    SetVpnPeerCaching(value *Enablement)()
}
