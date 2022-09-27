package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceLogCollectionResponse windows Log Collection request entity.
type DeviceLogCollectionResponse struct {
    Entity
    // The error code, if any. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
    errorCode *int64
    // The DateTime of the expiration of the logs
    expirationDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The UPN for who initiated the request
    initiatedByUserPrincipalName *string
    // The device Id
    managedDeviceId *string
    // The DateTime the request was received
    receivedDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateTime of the request
    requestedDateTimeUTC *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The size of the logs. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    size *float64
    // The status of the log collection request
    status *string
}
// NewDeviceLogCollectionResponse instantiates a new deviceLogCollectionResponse and sets the default values.
func NewDeviceLogCollectionResponse()(*DeviceLogCollectionResponse) {
    m := &DeviceLogCollectionResponse{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceLogCollectionResponse";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceLogCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceLogCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceLogCollectionResponse(), nil
}
// GetErrorCode gets the errorCode property value. The error code, if any. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *DeviceLogCollectionResponse) GetErrorCode()(*int64) {
    return m.errorCode
}
// GetExpirationDateTimeUTC gets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs
func (m *DeviceLogCollectionResponse) GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTimeUTC
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceLogCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetErrorCode)
    res["expirationDateTimeUTC"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTimeUTC)
    res["initiatedByUserPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitiatedByUserPrincipalName)
    res["managedDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceId)
    res["receivedDateTimeUTC"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReceivedDateTimeUTC)
    res["requestedDateTimeUTC"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRequestedDateTimeUTC)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetSize)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    return res
}
// GetInitiatedByUserPrincipalName gets the initiatedByUserPrincipalName property value. The UPN for who initiated the request
func (m *DeviceLogCollectionResponse) GetInitiatedByUserPrincipalName()(*string) {
    return m.initiatedByUserPrincipalName
}
// GetManagedDeviceId gets the managedDeviceId property value. The device Id
func (m *DeviceLogCollectionResponse) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetReceivedDateTimeUTC gets the receivedDateTimeUTC property value. The DateTime the request was received
func (m *DeviceLogCollectionResponse) GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTimeUTC
}
// GetRequestedDateTimeUTC gets the requestedDateTimeUTC property value. The DateTime of the request
func (m *DeviceLogCollectionResponse) GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestedDateTimeUTC
}
// GetSize gets the size property value. The size of the logs. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *DeviceLogCollectionResponse) GetSize()(*float64) {
    return m.size
}
// GetStatus gets the status property value. The status of the log collection request
func (m *DeviceLogCollectionResponse) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceLogCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTimeUTC", m.GetExpirationDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUserPrincipalName", m.GetInitiatedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTimeUTC", m.GetReceivedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedDateTimeUTC", m.GetRequestedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorCode sets the errorCode property value. The error code, if any. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *DeviceLogCollectionResponse) SetErrorCode(value *int64)() {
    m.errorCode = value
}
// SetExpirationDateTimeUTC sets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs
func (m *DeviceLogCollectionResponse) SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTimeUTC = value
}
// SetInitiatedByUserPrincipalName sets the initiatedByUserPrincipalName property value. The UPN for who initiated the request
func (m *DeviceLogCollectionResponse) SetInitiatedByUserPrincipalName(value *string)() {
    m.initiatedByUserPrincipalName = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The device Id
func (m *DeviceLogCollectionResponse) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetReceivedDateTimeUTC sets the receivedDateTimeUTC property value. The DateTime the request was received
func (m *DeviceLogCollectionResponse) SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTimeUTC = value
}
// SetRequestedDateTimeUTC sets the requestedDateTimeUTC property value. The DateTime of the request
func (m *DeviceLogCollectionResponse) SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedDateTimeUTC = value
}
// SetSize sets the size property value. The size of the logs. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *DeviceLogCollectionResponse) SetSize(value *float64)() {
    m.size = value
}
// SetStatus sets the status property value. The status of the log collection request
func (m *DeviceLogCollectionResponse) SetStatus(value *string)() {
    m.status = value
}
