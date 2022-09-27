package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RetentionEvent provides operations to manage the collection of activityStatistics entities.
type RetentionEvent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The user who created the retentionEvent.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date time when the retentionEvent was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional information about the event.
    description *string
    // Name of the event.
    displayName *string
    // The eventPropagationResults property
    eventPropagationResults []EventPropagationResultable
    // Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
    eventQueries []EventQueryable
    // The eventStatus property
    eventStatus RetentionEventStatusable
    // Optional time when the event should be triggered.
    eventTriggerDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user who last modified the retentionEvent.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The latest date time when the retentionEvent was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last time the status of the event was updated.
    lastStatusUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the event that will start the retention period for labels that use this event type when an event is created.
    retentionEventType RetentionEventTypeable
}
// NewRetentionEvent instantiates a new retentionEvent and sets the default values.
func NewRetentionEvent()(*RetentionEvent) {
    m := &RetentionEvent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.retentionEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRetentionEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEvent(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created the retentionEvent.
func (m *RetentionEvent) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the retentionEvent was created.
func (m *RetentionEvent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Optional information about the event.
func (m *RetentionEvent) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the event.
func (m *RetentionEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetEventPropagationResults gets the eventPropagationResults property value. The eventPropagationResults property
func (m *RetentionEvent) GetEventPropagationResults()([]EventPropagationResultable) {
    return m.eventPropagationResults
}
// GetEventQueries gets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
func (m *RetentionEvent) GetEventQueries()([]EventQueryable) {
    return m.eventQueries
}
// GetEventStatus gets the eventStatus property value. The eventStatus property
func (m *RetentionEvent) GetEventStatus()(RetentionEventStatusable) {
    return m.eventStatus
}
// GetEventTriggerDateTime gets the eventTriggerDateTime property value. Optional time when the event should be triggered.
func (m *RetentionEvent) GetEventTriggerDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventTriggerDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["eventPropagationResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEventPropagationResultFromDiscriminatorValue , m.SetEventPropagationResults)
    res["eventQueries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEventQueryFromDiscriminatorValue , m.SetEventQueries)
    res["eventStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRetentionEventStatusFromDiscriminatorValue , m.SetEventStatus)
    res["eventTriggerDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEventTriggerDateTime)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["lastStatusUpdateDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastStatusUpdateDateTime)
    res["retentionEventType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRetentionEventTypeFromDiscriminatorValue , m.SetRetentionEventType)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the retentionEvent.
func (m *RetentionEvent) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
func (m *RetentionEvent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastStatusUpdateDateTime gets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
func (m *RetentionEvent) GetLastStatusUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastStatusUpdateDateTime
}
// GetRetentionEventType gets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *RetentionEvent) GetRetentionEventType()(RetentionEventTypeable) {
    return m.retentionEventType
}
// Serialize serializes information the current object
func (m *RetentionEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEventPropagationResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEventPropagationResults())
        err = writer.WriteCollectionOfObjectValues("eventPropagationResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventQueries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEventQueries())
        err = writer.WriteCollectionOfObjectValues("eventQueries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventStatus", m.GetEventStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventTriggerDateTime", m.GetEventTriggerDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    {
        err = writer.WriteTimeValue("lastStatusUpdateDateTime", m.GetLastStatusUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionEventType", m.GetRetentionEventType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created the retentionEvent.
func (m *RetentionEvent) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the retentionEvent was created.
func (m *RetentionEvent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Optional information about the event.
func (m *RetentionEvent) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the event.
func (m *RetentionEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEventPropagationResults sets the eventPropagationResults property value. The eventPropagationResults property
func (m *RetentionEvent) SetEventPropagationResults(value []EventPropagationResultable)() {
    m.eventPropagationResults = value
}
// SetEventQueries sets the eventQueries property value. Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information associated with a retention event.
func (m *RetentionEvent) SetEventQueries(value []EventQueryable)() {
    m.eventQueries = value
}
// SetEventStatus sets the eventStatus property value. The eventStatus property
func (m *RetentionEvent) SetEventStatus(value RetentionEventStatusable)() {
    m.eventStatus = value
}
// SetEventTriggerDateTime sets the eventTriggerDateTime property value. Optional time when the event should be triggered.
func (m *RetentionEvent) SetEventTriggerDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventTriggerDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the retentionEvent.
func (m *RetentionEvent) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The latest date time when the retentionEvent was modified.
func (m *RetentionEvent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastStatusUpdateDateTime sets the lastStatusUpdateDateTime property value. Last time the status of the event was updated.
func (m *RetentionEvent) SetLastStatusUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastStatusUpdateDateTime = value
}
// SetRetentionEventType sets the retentionEventType property value. Specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *RetentionEvent) SetRetentionEventType(value RetentionEventTypeable)() {
    m.retentionEventType = value
}
