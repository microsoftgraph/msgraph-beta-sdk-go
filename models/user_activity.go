package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserActivity 
type UserActivity struct {
    Entity
}
// NewUserActivity instantiates a new userActivity and sets the default values.
func NewUserActivity()(*UserActivity) {
    m := &UserActivity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserActivity(), nil
}
// GetActivationUrl gets the activationUrl property value. The activationUrl property
func (m *UserActivity) GetActivationUrl()(*string) {
    val, err := m.GetBackingStore().Get("activationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetActivitySourceHost gets the activitySourceHost property value. The activitySourceHost property
func (m *UserActivity) GetActivitySourceHost()(*string) {
    val, err := m.GetBackingStore().Get("activitySourceHost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppActivityId gets the appActivityId property value. The appActivityId property
func (m *UserActivity) GetAppActivityId()(*string) {
    val, err := m.GetBackingStore().Get("appActivityId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppDisplayName gets the appDisplayName property value. The appDisplayName property
func (m *UserActivity) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *UserActivity) GetContentInfo()(Jsonable) {
    val, err := m.GetBackingStore().Get("contentInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Jsonable)
    }
    return nil
}
// GetContentUrl gets the contentUrl property value. The contentUrl property
func (m *UserActivity) GetContentUrl()(*string) {
    val, err := m.GetBackingStore().Get("contentUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *UserActivity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *UserActivity) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFallbackUrl gets the fallbackUrl property value. The fallbackUrl property
func (m *UserActivity) GetFallbackUrl()(*string) {
    val, err := m.GetBackingStore().Get("fallbackUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationUrl(val)
        }
        return nil
    }
    res["activitySourceHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivitySourceHost(val)
        }
        return nil
    }
    res["appActivityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActivityId(val)
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(Jsonable))
        }
        return nil
    }
    res["contentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["fallbackUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFallbackUrl(val)
        }
        return nil
    }
    res["historyItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityHistoryItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityHistoryItemable, len(val))
            for i, v := range val {
                res[i] = v.(ActivityHistoryItemable)
            }
            m.SetHistoryItems(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Status))
        }
        return nil
    }
    res["userTimezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTimezone(val)
        }
        return nil
    }
    res["visualElements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVisualInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisualElements(val.(VisualInfoable))
        }
        return nil
    }
    return res
}
// GetHistoryItems gets the historyItems property value. The historyItems property
func (m *UserActivity) GetHistoryItems()([]ActivityHistoryItemable) {
    val, err := m.GetBackingStore().Get("historyItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActivityHistoryItemable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserActivity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *UserActivity) GetStatus()(*Status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Status)
    }
    return nil
}
// GetUserTimezone gets the userTimezone property value. The userTimezone property
func (m *UserActivity) GetUserTimezone()(*string) {
    val, err := m.GetBackingStore().Get("userTimezone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVisualElements gets the visualElements property value. The visualElements property
func (m *UserActivity) GetVisualElements()(VisualInfoable) {
    val, err := m.GetBackingStore().Get("visualElements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VisualInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activationUrl", m.GetActivationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activitySourceHost", m.GetActivitySourceHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appActivityId", m.GetAppActivityId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fallbackUrl", m.GetFallbackUrl())
        if err != nil {
            return err
        }
    }
    if m.GetHistoryItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistoryItems()))
        for i, v := range m.GetHistoryItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("historyItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userTimezone", m.GetUserTimezone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("visualElements", m.GetVisualElements())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationUrl sets the activationUrl property value. The activationUrl property
func (m *UserActivity) SetActivationUrl(value *string)() {
    err := m.GetBackingStore().Set("activationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetActivitySourceHost sets the activitySourceHost property value. The activitySourceHost property
func (m *UserActivity) SetActivitySourceHost(value *string)() {
    err := m.GetBackingStore().Set("activitySourceHost", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActivityId sets the appActivityId property value. The appActivityId property
func (m *UserActivity) SetAppActivityId(value *string)() {
    err := m.GetBackingStore().Set("appActivityId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppDisplayName sets the appDisplayName property value. The appDisplayName property
func (m *UserActivity) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *UserActivity) SetContentInfo(value Jsonable)() {
    err := m.GetBackingStore().Set("contentInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetContentUrl sets the contentUrl property value. The contentUrl property
func (m *UserActivity) SetContentUrl(value *string)() {
    err := m.GetBackingStore().Set("contentUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *UserActivity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *UserActivity) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFallbackUrl sets the fallbackUrl property value. The fallbackUrl property
func (m *UserActivity) SetFallbackUrl(value *string)() {
    err := m.GetBackingStore().Set("fallbackUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetHistoryItems sets the historyItems property value. The historyItems property
func (m *UserActivity) SetHistoryItems(value []ActivityHistoryItemable)() {
    err := m.GetBackingStore().Set("historyItems", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserActivity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *UserActivity) SetStatus(value *Status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUserTimezone sets the userTimezone property value. The userTimezone property
func (m *UserActivity) SetUserTimezone(value *string)() {
    err := m.GetBackingStore().Set("userTimezone", value)
    if err != nil {
        panic(err)
    }
}
// SetVisualElements sets the visualElements property value. The visualElements property
func (m *UserActivity) SetVisualElements(value VisualInfoable)() {
    err := m.GetBackingStore().Set("visualElements", value)
    if err != nil {
        panic(err)
    }
}
// UserActivityable 
type UserActivityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationUrl()(*string)
    GetActivitySourceHost()(*string)
    GetAppActivityId()(*string)
    GetAppDisplayName()(*string)
    GetContentInfo()(Jsonable)
    GetContentUrl()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFallbackUrl()(*string)
    GetHistoryItems()([]ActivityHistoryItemable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*Status)
    GetUserTimezone()(*string)
    GetVisualElements()(VisualInfoable)
    SetActivationUrl(value *string)()
    SetActivitySourceHost(value *string)()
    SetAppActivityId(value *string)()
    SetAppDisplayName(value *string)()
    SetContentInfo(value Jsonable)()
    SetContentUrl(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFallbackUrl(value *string)()
    SetHistoryItems(value []ActivityHistoryItemable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *Status)()
    SetUserTimezone(value *string)()
    SetVisualElements(value VisualInfoable)()
}
