package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ZebraFotaDeploymentStatus describes the status for a single FOTA deployment.
type ZebraFotaDeploymentStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewZebraFotaDeploymentStatus instantiates a new zebraFotaDeploymentStatus and sets the default values.
func NewZebraFotaDeploymentStatus()(*ZebraFotaDeploymentStatus) {
    m := &ZebraFotaDeploymentStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateZebraFotaDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeploymentStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentStatus) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ZebraFotaDeploymentStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCancelRequested gets the cancelRequested property value. A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not guarantee that the deployment was canceled.
func (m *ZebraFotaDeploymentStatus) GetCancelRequested()(*bool) {
    val, err := m.GetBackingStore().Get("cancelRequested")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCompleteOrCanceledDateTime gets the completeOrCanceledDateTime property value. The date and time when this deployment was completed or canceled. The actual date time is determined by the value of state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this property holds the completion date/time. If the deployment is not completed before the deployment end date, then completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation that is in progress can continue past the deployment end date.
func (m *ZebraFotaDeploymentStatus) GetCompleteOrCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completeOrCanceledDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. An error code indicating the failure reason, when the deployment state is createFailed. Possible values: See zebraFotaErrorCode enum.
func (m *ZebraFotaDeploymentStatus) GetErrorCode()(*ZebraFotaErrorCode) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ZebraFotaErrorCode)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cancelRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelRequested(val)
        }
        return nil
    }
    res["completeOrCanceledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteOrCanceledDateTime(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaErrorCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val.(*ZebraFotaErrorCode))
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaDeploymentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ZebraFotaDeploymentState))
        }
        return nil
    }
    res["totalAwaitingInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAwaitingInstall(val)
        }
        return nil
    }
    res["totalCanceled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCanceled(val)
        }
        return nil
    }
    res["totalCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCreated(val)
        }
        return nil
    }
    res["totalDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDevices(val)
        }
        return nil
    }
    res["totalDownloading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDownloading(val)
        }
        return nil
    }
    res["totalFailedDownload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFailedDownload(val)
        }
        return nil
    }
    res["totalFailedInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFailedInstall(val)
        }
        return nil
    }
    res["totalScheduled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalScheduled(val)
        }
        return nil
    }
    res["totalSucceededInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSucceededInstall(val)
        }
        return nil
    }
    res["totalUnknown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnknown(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time when the deployment status was updated from Zebra
func (m *ZebraFotaDeploymentStatus) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetState gets the state property value. Represents the state of Zebra FOTA deployment.
func (m *ZebraFotaDeploymentStatus) GetState()(*ZebraFotaDeploymentState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ZebraFotaDeploymentState)
    }
    return nil
}
// GetTotalAwaitingInstall gets the totalAwaitingInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalAwaitingInstall()(*int32) {
    val, err := m.GetBackingStore().Get("totalAwaitingInstall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalCanceled gets the totalCanceled property value. An integer that indicates the total number of devices where installation was canceled.
func (m *ZebraFotaDeploymentStatus) GetTotalCanceled()(*int32) {
    val, err := m.GetBackingStore().Get("totalCanceled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalCreated gets the totalCreated property value. An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs that did not reach the devices.
func (m *ZebraFotaDeploymentStatus) GetTotalCreated()(*int32) {
    val, err := m.GetBackingStore().Get("totalCreated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalDevices gets the totalDevices property value. An integer that indicates the total number of devices in the deployment.
func (m *ZebraFotaDeploymentStatus) GetTotalDevices()(*int32) {
    val, err := m.GetBackingStore().Get("totalDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalDownloading gets the totalDownloading property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalDownloading()(*int32) {
    val, err := m.GetBackingStore().Get("totalDownloading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalFailedDownload gets the totalFailedDownload property value. An integer that indicates the total number of devices that have failed to download the new OS file.
func (m *ZebraFotaDeploymentStatus) GetTotalFailedDownload()(*int32) {
    val, err := m.GetBackingStore().Get("totalFailedDownload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalFailedInstall gets the totalFailedInstall property value. An integer that indicates the total number of devices that have failed to install the new OS file.
func (m *ZebraFotaDeploymentStatus) GetTotalFailedInstall()(*int32) {
    val, err := m.GetBackingStore().Get("totalFailedInstall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalScheduled gets the totalScheduled property value. An integer that indicates the total number of devices that received the json and are scheduled.
func (m *ZebraFotaDeploymentStatus) GetTotalScheduled()(*int32) {
    val, err := m.GetBackingStore().Get("totalScheduled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalSucceededInstall gets the totalSucceededInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalSucceededInstall()(*int32) {
    val, err := m.GetBackingStore().Get("totalSucceededInstall")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalUnknown gets the totalUnknown property value. An integer that indicates the total number of devices where no deployment status or end state has not received, even after the scheduled end date was reached.
func (m *ZebraFotaDeploymentStatus) GetTotalUnknown()(*int32) {
    val, err := m.GetBackingStore().Get("totalUnknown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ZebraFotaDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("cancelRequested", m.GetCancelRequested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completeOrCanceledDateTime", m.GetCompleteOrCanceledDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCode() != nil {
        cast := (*m.GetErrorCode()).String()
        err := writer.WriteStringValue("errorCode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalAwaitingInstall", m.GetTotalAwaitingInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCanceled", m.GetTotalCanceled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCreated", m.GetTotalCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDevices", m.GetTotalDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDownloading", m.GetTotalDownloading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalFailedDownload", m.GetTotalFailedDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalFailedInstall", m.GetTotalFailedInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalScheduled", m.GetTotalScheduled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalSucceededInstall", m.GetTotalSucceededInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalUnknown", m.GetTotalUnknown())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ZebraFotaDeploymentStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCancelRequested sets the cancelRequested property value. A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not guarantee that the deployment was canceled.
func (m *ZebraFotaDeploymentStatus) SetCancelRequested(value *bool)() {
    err := m.GetBackingStore().Set("cancelRequested", value)
    if err != nil {
        panic(err)
    }
}
// SetCompleteOrCanceledDateTime sets the completeOrCanceledDateTime property value. The date and time when this deployment was completed or canceled. The actual date time is determined by the value of state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this property holds the completion date/time. If the deployment is not completed before the deployment end date, then completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation that is in progress can continue past the deployment end date.
func (m *ZebraFotaDeploymentStatus) SetCompleteOrCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completeOrCanceledDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. An error code indicating the failure reason, when the deployment state is createFailed. Possible values: See zebraFotaErrorCode enum.
func (m *ZebraFotaDeploymentStatus) SetErrorCode(value *ZebraFotaErrorCode)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time when the deployment status was updated from Zebra
func (m *ZebraFotaDeploymentStatus) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Represents the state of Zebra FOTA deployment.
func (m *ZebraFotaDeploymentStatus) SetState(value *ZebraFotaDeploymentState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalAwaitingInstall sets the totalAwaitingInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalAwaitingInstall(value *int32)() {
    err := m.GetBackingStore().Set("totalAwaitingInstall", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCanceled sets the totalCanceled property value. An integer that indicates the total number of devices where installation was canceled.
func (m *ZebraFotaDeploymentStatus) SetTotalCanceled(value *int32)() {
    err := m.GetBackingStore().Set("totalCanceled", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCreated sets the totalCreated property value. An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs that did not reach the devices.
func (m *ZebraFotaDeploymentStatus) SetTotalCreated(value *int32)() {
    err := m.GetBackingStore().Set("totalCreated", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDevices sets the totalDevices property value. An integer that indicates the total number of devices in the deployment.
func (m *ZebraFotaDeploymentStatus) SetTotalDevices(value *int32)() {
    err := m.GetBackingStore().Set("totalDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDownloading sets the totalDownloading property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalDownloading(value *int32)() {
    err := m.GetBackingStore().Set("totalDownloading", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalFailedDownload sets the totalFailedDownload property value. An integer that indicates the total number of devices that have failed to download the new OS file.
func (m *ZebraFotaDeploymentStatus) SetTotalFailedDownload(value *int32)() {
    err := m.GetBackingStore().Set("totalFailedDownload", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalFailedInstall sets the totalFailedInstall property value. An integer that indicates the total number of devices that have failed to install the new OS file.
func (m *ZebraFotaDeploymentStatus) SetTotalFailedInstall(value *int32)() {
    err := m.GetBackingStore().Set("totalFailedInstall", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalScheduled sets the totalScheduled property value. An integer that indicates the total number of devices that received the json and are scheduled.
func (m *ZebraFotaDeploymentStatus) SetTotalScheduled(value *int32)() {
    err := m.GetBackingStore().Set("totalScheduled", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalSucceededInstall sets the totalSucceededInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalSucceededInstall(value *int32)() {
    err := m.GetBackingStore().Set("totalSucceededInstall", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUnknown sets the totalUnknown property value. An integer that indicates the total number of devices where no deployment status or end state has not received, even after the scheduled end date was reached.
func (m *ZebraFotaDeploymentStatus) SetTotalUnknown(value *int32)() {
    err := m.GetBackingStore().Set("totalUnknown", value)
    if err != nil {
        panic(err)
    }
}
// ZebraFotaDeploymentStatusable 
type ZebraFotaDeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCancelRequested()(*bool)
    GetCompleteOrCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrorCode()(*ZebraFotaErrorCode)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetState()(*ZebraFotaDeploymentState)
    GetTotalAwaitingInstall()(*int32)
    GetTotalCanceled()(*int32)
    GetTotalCreated()(*int32)
    GetTotalDevices()(*int32)
    GetTotalDownloading()(*int32)
    GetTotalFailedDownload()(*int32)
    GetTotalFailedInstall()(*int32)
    GetTotalScheduled()(*int32)
    GetTotalSucceededInstall()(*int32)
    GetTotalUnknown()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCancelRequested(value *bool)()
    SetCompleteOrCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrorCode(value *ZebraFotaErrorCode)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetState(value *ZebraFotaDeploymentState)()
    SetTotalAwaitingInstall(value *int32)()
    SetTotalCanceled(value *int32)()
    SetTotalCreated(value *int32)()
    SetTotalDevices(value *int32)()
    SetTotalDownloading(value *int32)()
    SetTotalFailedDownload(value *int32)()
    SetTotalFailedInstall(value *int32)()
    SetTotalScheduled(value *int32)()
    SetTotalSucceededInstall(value *int32)()
    SetTotalUnknown(value *int32)()
}
