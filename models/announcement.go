package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Announcement struct {
    ChangeItemBase
}
// NewAnnouncement instantiates a new Announcement and sets the default values.
func NewAnnouncement()(*Announcement) {
    m := &Announcement{
        ChangeItemBase: *NewChangeItemBase(),
    }
    odataTypeValue := "#microsoft.graph.announcement"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAnnouncementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnnouncementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnnouncement(), nil
}
// GetAnnouncementDateTime gets the announcementDateTime property value. Change announcement date. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
// returns a *Time when successful
func (m *Announcement) GetAnnouncementDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("announcementDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetChangeType gets the changeType property value. The changeType property
// returns a *ChangeAnnouncementChangeType when successful
func (m *Announcement) GetChangeType()(*ChangeAnnouncementChangeType) {
    val, err := m.GetBackingStore().Get("changeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChangeAnnouncementChangeType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Announcement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeItemBase.GetFieldDeserializers()
    res["announcementDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnnouncementDateTime(val)
        }
        return nil
    }
    res["changeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeAnnouncementChangeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeType(val.(*ChangeAnnouncementChangeType))
        }
        return nil
    }
    res["impactLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactLink(val)
        }
        return nil
    }
    res["isCustomerActionRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomerActionRequired(val)
        }
        return nil
    }
    res["targetDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDateTime(val)
        }
        return nil
    }
    return res
}
// GetImpactLink gets the impactLink property value. Change impact URL. Supports $filter (eq, ne, in) and $orderby.
// returns a *string when successful
func (m *Announcement) GetImpactLink()(*string) {
    val, err := m.GetBackingStore().Get("impactLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsCustomerActionRequired gets the isCustomerActionRequired property value. Indicates whether the customer needs to take any action for this change. Supports $filter (eq, ne).
// returns a *bool when successful
func (m *Announcement) GetIsCustomerActionRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isCustomerActionRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTargetDateTime gets the targetDateTime property value. Date on which the change rolls out. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
// returns a *Time when successful
func (m *Announcement) GetTargetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("targetDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Announcement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeItemBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("announcementDateTime", m.GetAnnouncementDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetChangeType() != nil {
        cast := (*m.GetChangeType()).String()
        err = writer.WriteStringValue("changeType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactLink", m.GetImpactLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomerActionRequired", m.GetIsCustomerActionRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("targetDateTime", m.GetTargetDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnnouncementDateTime sets the announcementDateTime property value. Change announcement date. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
func (m *Announcement) SetAnnouncementDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("announcementDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetChangeType sets the changeType property value. The changeType property
func (m *Announcement) SetChangeType(value *ChangeAnnouncementChangeType)() {
    err := m.GetBackingStore().Set("changeType", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactLink sets the impactLink property value. Change impact URL. Supports $filter (eq, ne, in) and $orderby.
func (m *Announcement) SetImpactLink(value *string)() {
    err := m.GetBackingStore().Set("impactLink", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCustomerActionRequired sets the isCustomerActionRequired property value. Indicates whether the customer needs to take any action for this change. Supports $filter (eq, ne).
func (m *Announcement) SetIsCustomerActionRequired(value *bool)() {
    err := m.GetBackingStore().Set("isCustomerActionRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetDateTime sets the targetDateTime property value. Date on which the change rolls out. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
func (m *Announcement) SetTargetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("targetDateTime", value)
    if err != nil {
        panic(err)
    }
}
type Announcementable interface {
    ChangeItemBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnnouncementDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetChangeType()(*ChangeAnnouncementChangeType)
    GetImpactLink()(*string)
    GetIsCustomerActionRequired()(*bool)
    GetTargetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAnnouncementDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetChangeType(value *ChangeAnnouncementChangeType)()
    SetImpactLink(value *string)()
    SetIsCustomerActionRequired(value *bool)()
    SetTargetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
