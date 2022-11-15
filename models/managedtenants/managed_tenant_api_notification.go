package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenantApiNotification provides operations to manage the collection of accessReviewDecision entities.
type ManagedTenantApiNotification struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The alert property
    alert ManagedTenantAlertable
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The isAcknowledged property
    isAcknowledged *bool
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The message property
    message *string
    // The title property
    title *string
    // The userId property
    userId *string
}
// NewManagedTenantApiNotification instantiates a new managedTenantApiNotification and sets the default values.
func NewManagedTenantApiNotification()(*ManagedTenantApiNotification) {
    m := &ManagedTenantApiNotification{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managedTenantApiNotification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedTenantApiNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantApiNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantApiNotification(), nil
}
// GetAlert gets the alert property value. The alert property
func (m *ManagedTenantApiNotification) GetAlert()(ManagedTenantAlertable) {
    return m.alert
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantApiNotification) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantApiNotification) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantApiNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alert"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedTenantAlertFromDiscriminatorValue , m.SetAlert)
    res["createdByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedByUserId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["isAcknowledged"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAcknowledged)
    res["lastActionByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastActionByUserId)
    res["lastActionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastActionDateTime)
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetIsAcknowledged gets the isAcknowledged property value. The isAcknowledged property
func (m *ManagedTenantApiNotification) GetIsAcknowledged()(*bool) {
    return m.isAcknowledged
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantApiNotification) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantApiNotification) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetMessage gets the message property value. The message property
func (m *ManagedTenantApiNotification) GetMessage()(*string) {
    return m.message
}
// GetTitle gets the title property value. The title property
func (m *ManagedTenantApiNotification) GetTitle()(*string) {
    return m.title
}
// GetUserId gets the userId property value. The userId property
func (m *ManagedTenantApiNotification) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *ManagedTenantApiNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alert", m.GetAlert())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
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
        err = writer.WriteBoolValue("isAcknowledged", m.GetIsAcknowledged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlert sets the alert property value. The alert property
func (m *ManagedTenantApiNotification) SetAlert(value ManagedTenantAlertable)() {
    m.alert = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantApiNotification) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantApiNotification) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIsAcknowledged sets the isAcknowledged property value. The isAcknowledged property
func (m *ManagedTenantApiNotification) SetIsAcknowledged(value *bool)() {
    m.isAcknowledged = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantApiNotification) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantApiNotification) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *ManagedTenantApiNotification) SetMessage(value *string)() {
    m.message = value
}
// SetTitle sets the title property value. The title property
func (m *ManagedTenantApiNotification) SetTitle(value *string)() {
    m.title = value
}
// SetUserId sets the userId property value. The userId property
func (m *ManagedTenantApiNotification) SetUserId(value *string)() {
    m.userId = value
}
