package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Roadmap struct {
    ChangeItemBase
}
// NewRoadmap instantiates a new Roadmap and sets the default values.
func NewRoadmap()(*Roadmap) {
    m := &Roadmap{
        ChangeItemBase: *NewChangeItemBase(),
    }
    odataTypeValue := "#microsoft.graph.roadmap"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRoadmapFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoadmapFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoadmap(), nil
}
// GetCategory gets the category property value. Indicates the category with which this item is associated. Supports $filter (eq, ne, in) and $orderby.
// returns a *string when successful
func (m *Roadmap) GetCategory()(*string) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetChangeItemState gets the changeItemState property value. The changeItemState property
// returns a *ChangeItemState when successful
func (m *Roadmap) GetChangeItemState()(*ChangeItemState) {
    val, err := m.GetBackingStore().Get("changeItemState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChangeItemState)
    }
    return nil
}
// GetDeliveryStage gets the deliveryStage property value. The deliveryStage property
// returns a *RoadmapItemDeliveryStage when successful
func (m *Roadmap) GetDeliveryStage()(*RoadmapItemDeliveryStage) {
    val, err := m.GetBackingStore().Get("deliveryStage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RoadmapItemDeliveryStage)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Roadmap) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeItemBase.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["changeItemState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeItemState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeItemState(val.(*ChangeItemState))
        }
        return nil
    }
    res["deliveryStage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoadmapItemDeliveryStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryStage(val.(*RoadmapItemDeliveryStage))
        }
        return nil
    }
    res["gotoLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGotoLink(val)
        }
        return nil
    }
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    return res
}
// GetGotoLink gets the gotoLink property value. Link to the feature page in the Microsoft Entra admin center. Supports $filter (eq, ne, in) and $orderby.
// returns a *string when successful
func (m *Roadmap) GetGotoLink()(*string) {
    val, err := m.GetBackingStore().Get("gotoLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublishedDateTime gets the publishedDateTime property value. Feature planned release date. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
// returns a *Time when successful
func (m *Roadmap) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("publishedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Roadmap) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeItemBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    if m.GetChangeItemState() != nil {
        cast := (*m.GetChangeItemState()).String()
        err = writer.WriteStringValue("changeItemState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryStage() != nil {
        cast := (*m.GetDeliveryStage()).String()
        err = writer.WriteStringValue("deliveryStage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gotoLink", m.GetGotoLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. Indicates the category with which this item is associated. Supports $filter (eq, ne, in) and $orderby.
func (m *Roadmap) SetCategory(value *string)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetChangeItemState sets the changeItemState property value. The changeItemState property
func (m *Roadmap) SetChangeItemState(value *ChangeItemState)() {
    err := m.GetBackingStore().Set("changeItemState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeliveryStage sets the deliveryStage property value. The deliveryStage property
func (m *Roadmap) SetDeliveryStage(value *RoadmapItemDeliveryStage)() {
    err := m.GetBackingStore().Set("deliveryStage", value)
    if err != nil {
        panic(err)
    }
}
// SetGotoLink sets the gotoLink property value. Link to the feature page in the Microsoft Entra admin center. Supports $filter (eq, ne, in) and $orderby.
func (m *Roadmap) SetGotoLink(value *string)() {
    err := m.GetBackingStore().Set("gotoLink", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishedDateTime sets the publishedDateTime property value. Feature planned release date. Supports $filter (eq, ne, gt, lt, le and ge on year(), month(), day(), hour(), minute(), and second() built in functions) and $orderby.
func (m *Roadmap) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("publishedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type Roadmapable interface {
    ChangeItemBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*string)
    GetChangeItemState()(*ChangeItemState)
    GetDeliveryStage()(*RoadmapItemDeliveryStage)
    GetGotoLink()(*string)
    GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCategory(value *string)()
    SetChangeItemState(value *ChangeItemState)()
    SetDeliveryStage(value *RoadmapItemDeliveryStage)()
    SetGotoLink(value *string)()
    SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
