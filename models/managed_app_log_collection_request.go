package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppLogCollectionRequest the Managed App log collection response
type ManagedAppLogCollectionRequest struct {
    Entity
}
// NewManagedAppLogCollectionRequest instantiates a new ManagedAppLogCollectionRequest and sets the default values.
func NewManagedAppLogCollectionRequest()(*ManagedAppLogCollectionRequest) {
    m := &ManagedAppLogCollectionRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAppLogCollectionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedAppLogCollectionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppLogCollectionRequest(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. DateTime of when the log upload request was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
// returns a *Time when successful
func (m *ManagedAppLogCollectionRequest) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDateTime")
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
func (m *ManagedAppLogCollectionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["managedAppRegistrationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedAppRegistrationId(val)
        }
        return nil
    }
    res["requestedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedBy(val)
        }
        return nil
    }
    res["requestedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedByUserPrincipalName(val)
        }
        return nil
    }
    res["requestedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["uploadedLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppLogUploadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppLogUploadable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAppLogUploadable)
                }
            }
            m.SetUploadedLogs(res)
        }
        return nil
    }
    res["userLogUploadConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppLogUploadConsent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserLogUploadConsent(val.(*ManagedAppLogUploadConsent))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetManagedAppRegistrationId gets the managedAppRegistrationId property value. The unique identifier of the app instance for which diagnostic logs were collected. Read-only.
// returns a *string when successful
func (m *ManagedAppLogCollectionRequest) GetManagedAppRegistrationId()(*string) {
    val, err := m.GetBackingStore().Get("managedAppRegistrationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedBy gets the requestedBy property value. The user principal name associated with the request for the managed application log collection. Read-only.
// returns a *string when successful
func (m *ManagedAppLogCollectionRequest) GetRequestedBy()(*string) {
    val, err := m.GetBackingStore().Get("requestedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedByUserPrincipalName gets the requestedByUserPrincipalName property value. The user principal name associated with the request for the managed application log collection. Read-only.
// returns a *string when successful
func (m *ManagedAppLogCollectionRequest) GetRequestedByUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("requestedByUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedDateTime gets the requestedDateTime property value. DateTime of when the log upload request was received. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
// returns a *Time when successful
func (m *ManagedAppLogCollectionRequest) GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. Indicates the status for the app log collection request - pending, completed or failed. Default is pending.
// returns a *string when successful
func (m *ManagedAppLogCollectionRequest) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUploadedLogs gets the uploadedLogs property value. The collection of log upload results as reported by each component on the device. Such components can be the application itself, the Mobile Application Management (MAM) SDK, and other on-device components that are requested to upload diagnostic logs. Read-only.
// returns a []ManagedAppLogUploadable when successful
func (m *ManagedAppLogCollectionRequest) GetUploadedLogs()([]ManagedAppLogUploadable) {
    val, err := m.GetBackingStore().Get("uploadedLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppLogUploadable)
    }
    return nil
}
// GetUserLogUploadConsent gets the userLogUploadConsent property value. Represents the current consent status of the associated `managedAppLogCollectionRequest`.
// returns a *ManagedAppLogUploadConsent when successful
func (m *ManagedAppLogCollectionRequest) GetUserLogUploadConsent()(*ManagedAppLogUploadConsent) {
    val, err := m.GetBackingStore().Get("userLogUploadConsent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppLogUploadConsent)
    }
    return nil
}
// GetVersion gets the version property value. Version of the entity.
// returns a *string when successful
func (m *ManagedAppLogCollectionRequest) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedAppLogCollectionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedAppRegistrationId", m.GetManagedAppRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedBy", m.GetRequestedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedByUserPrincipalName", m.GetRequestedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedDateTime", m.GetRequestedDateTime())
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
    if m.GetUploadedLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUploadedLogs()))
        for i, v := range m.GetUploadedLogs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("uploadedLogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserLogUploadConsent() != nil {
        cast := (*m.GetUserLogUploadConsent()).String()
        err = writer.WriteStringValue("userLogUploadConsent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. DateTime of when the log upload request was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *ManagedAppLogCollectionRequest) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedAppRegistrationId sets the managedAppRegistrationId property value. The unique identifier of the app instance for which diagnostic logs were collected. Read-only.
func (m *ManagedAppLogCollectionRequest) SetManagedAppRegistrationId(value *string)() {
    err := m.GetBackingStore().Set("managedAppRegistrationId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedBy sets the requestedBy property value. The user principal name associated with the request for the managed application log collection. Read-only.
func (m *ManagedAppLogCollectionRequest) SetRequestedBy(value *string)() {
    err := m.GetBackingStore().Set("requestedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedByUserPrincipalName sets the requestedByUserPrincipalName property value. The user principal name associated with the request for the managed application log collection. Read-only.
func (m *ManagedAppLogCollectionRequest) SetRequestedByUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("requestedByUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedDateTime sets the requestedDateTime property value. DateTime of when the log upload request was received. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *ManagedAppLogCollectionRequest) SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Indicates the status for the app log collection request - pending, completed or failed. Default is pending.
func (m *ManagedAppLogCollectionRequest) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadedLogs sets the uploadedLogs property value. The collection of log upload results as reported by each component on the device. Such components can be the application itself, the Mobile Application Management (MAM) SDK, and other on-device components that are requested to upload diagnostic logs. Read-only.
func (m *ManagedAppLogCollectionRequest) SetUploadedLogs(value []ManagedAppLogUploadable)() {
    err := m.GetBackingStore().Set("uploadedLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetUserLogUploadConsent sets the userLogUploadConsent property value. Represents the current consent status of the associated `managedAppLogCollectionRequest`.
func (m *ManagedAppLogCollectionRequest) SetUserLogUploadConsent(value *ManagedAppLogUploadConsent)() {
    err := m.GetBackingStore().Set("userLogUploadConsent", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppLogCollectionRequest) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type ManagedAppLogCollectionRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedAppRegistrationId()(*string)
    GetRequestedBy()(*string)
    GetRequestedByUserPrincipalName()(*string)
    GetRequestedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    GetUploadedLogs()([]ManagedAppLogUploadable)
    GetUserLogUploadConsent()(*ManagedAppLogUploadConsent)
    GetVersion()(*string)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedAppRegistrationId(value *string)()
    SetRequestedBy(value *string)()
    SetRequestedByUserPrincipalName(value *string)()
    SetRequestedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
    SetUploadedLogs(value []ManagedAppLogUploadable)()
    SetUserLogUploadConsent(value *ManagedAppLogUploadConsent)()
    SetVersion(value *string)()
}
