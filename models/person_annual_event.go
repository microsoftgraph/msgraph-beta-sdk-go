package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAnnualEvent 
type PersonAnnualEvent struct {
    ItemFacet
}
// NewPersonAnnualEvent instantiates a new personAnnualEvent and sets the default values.
func NewPersonAnnualEvent()(*PersonAnnualEvent) {
    m := &PersonAnnualEvent{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAnnualEvent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePersonAnnualEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAnnualEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAnnualEvent(), nil
}
// GetDate gets the date property value. The date property
func (m *PersonAnnualEvent) GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("date")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PersonAnnualEvent) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAnnualEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonAnnualEvent_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PersonAnnualEvent_type))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
func (m *PersonAnnualEvent) GetTypeEscaped()(*PersonAnnualEvent_type) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PersonAnnualEvent_type)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PersonAnnualEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("date", m.GetDate())
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDate sets the date property value. The date property
func (m *PersonAnnualEvent) SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("date", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PersonAnnualEvent) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. The type property
func (m *PersonAnnualEvent) SetTypeEscaped(value *PersonAnnualEvent_type)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// PersonAnnualEventable 
type PersonAnnualEventable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetTypeEscaped()(*PersonAnnualEvent_type)
    SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetTypeEscaped(value *PersonAnnualEvent_type)()
}
