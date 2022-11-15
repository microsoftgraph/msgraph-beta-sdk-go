package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserActivity provides operations to manage the collection of accessReviewDecision entities.
type UserActivity struct {
    Entity
    // The activationUrl property
    activationUrl *string
    // The activitySourceHost property
    activitySourceHost *string
    // The appActivityId property
    appActivityId *string
    // The appDisplayName property
    appDisplayName *string
    // The contentInfo property
    contentInfo Jsonable
    // The contentUrl property
    contentUrl *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fallbackUrl property
    fallbackUrl *string
    // The historyItems property
    historyItems []ActivityHistoryItemable
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *Status
    // The userTimezone property
    userTimezone *string
    // The visualElements property
    visualElements VisualInfoable
}
// NewUserActivity instantiates a new userActivity and sets the default values.
func NewUserActivity()(*UserActivity) {
    m := &UserActivity{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userActivity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserActivity(), nil
}
// GetActivationUrl gets the activationUrl property value. The activationUrl property
func (m *UserActivity) GetActivationUrl()(*string) {
    return m.activationUrl
}
// GetActivitySourceHost gets the activitySourceHost property value. The activitySourceHost property
func (m *UserActivity) GetActivitySourceHost()(*string) {
    return m.activitySourceHost
}
// GetAppActivityId gets the appActivityId property value. The appActivityId property
func (m *UserActivity) GetAppActivityId()(*string) {
    return m.appActivityId
}
// GetAppDisplayName gets the appDisplayName property value. The appDisplayName property
func (m *UserActivity) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *UserActivity) GetContentInfo()(Jsonable) {
    return m.contentInfo
}
// GetContentUrl gets the contentUrl property value. The contentUrl property
func (m *UserActivity) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *UserActivity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *UserActivity) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFallbackUrl gets the fallbackUrl property value. The fallbackUrl property
func (m *UserActivity) GetFallbackUrl()(*string) {
    return m.fallbackUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivationUrl)
    res["activitySourceHost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivitySourceHost)
    res["appActivityId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppActivityId)
    res["appDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppDisplayName)
    res["contentInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateJsonFromDiscriminatorValue , m.SetContentInfo)
    res["contentUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentUrl)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["fallbackUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFallbackUrl)
    res["historyItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateActivityHistoryItemFromDiscriminatorValue , m.SetHistoryItems)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseStatus , m.SetStatus)
    res["userTimezone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserTimezone)
    res["visualElements"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVisualInfoFromDiscriminatorValue , m.SetVisualElements)
    return res
}
// GetHistoryItems gets the historyItems property value. The historyItems property
func (m *UserActivity) GetHistoryItems()([]ActivityHistoryItemable) {
    return m.historyItems
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserActivity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetStatus gets the status property value. The status property
func (m *UserActivity) GetStatus()(*Status) {
    return m.status
}
// GetUserTimezone gets the userTimezone property value. The userTimezone property
func (m *UserActivity) GetUserTimezone()(*string) {
    return m.userTimezone
}
// GetVisualElements gets the visualElements property value. The visualElements property
func (m *UserActivity) GetVisualElements()(VisualInfoable) {
    return m.visualElements
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHistoryItems())
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
    m.activationUrl = value
}
// SetActivitySourceHost sets the activitySourceHost property value. The activitySourceHost property
func (m *UserActivity) SetActivitySourceHost(value *string)() {
    m.activitySourceHost = value
}
// SetAppActivityId sets the appActivityId property value. The appActivityId property
func (m *UserActivity) SetAppActivityId(value *string)() {
    m.appActivityId = value
}
// SetAppDisplayName sets the appDisplayName property value. The appDisplayName property
func (m *UserActivity) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *UserActivity) SetContentInfo(value Jsonable)() {
    m.contentInfo = value
}
// SetContentUrl sets the contentUrl property value. The contentUrl property
func (m *UserActivity) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *UserActivity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *UserActivity) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetFallbackUrl sets the fallbackUrl property value. The fallbackUrl property
func (m *UserActivity) SetFallbackUrl(value *string)() {
    m.fallbackUrl = value
}
// SetHistoryItems sets the historyItems property value. The historyItems property
func (m *UserActivity) SetHistoryItems(value []ActivityHistoryItemable)() {
    m.historyItems = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UserActivity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *UserActivity) SetStatus(value *Status)() {
    m.status = value
}
// SetUserTimezone sets the userTimezone property value. The userTimezone property
func (m *UserActivity) SetUserTimezone(value *string)() {
    m.userTimezone = value
}
// SetVisualElements sets the visualElements property value. The visualElements property
func (m *UserActivity) SetVisualElements(value VisualInfoable)() {
    m.visualElements = value
}
