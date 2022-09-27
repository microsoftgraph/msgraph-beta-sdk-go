package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAnnualEvent 
type PersonAnnualEvent struct {
    ItemFacet
    // The date property
    date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The displayName property
    displayName *string
    // The type property
    type_escaped *PersonAnnualEventType
}
// NewPersonAnnualEvent instantiates a new PersonAnnualEvent and sets the default values.
func NewPersonAnnualEvent()(*PersonAnnualEvent) {
    m := &PersonAnnualEvent{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAnnualEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonAnnualEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAnnualEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAnnualEvent(), nil
}
// GetDate gets the date property value. The date property
func (m *PersonAnnualEvent) GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.date
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PersonAnnualEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAnnualEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["date"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetDate)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePersonAnnualEventType , m.SetType)
    return res
}
// GetType gets the type property value. The type property
func (m *PersonAnnualEvent) GetType()(*PersonAnnualEventType) {
    return m.type_escaped
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDate sets the date property value. The date property
func (m *PersonAnnualEvent) SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.date = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PersonAnnualEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetType sets the type property value. The type property
func (m *PersonAnnualEvent) SetType(value *PersonAnnualEventType)() {
    m.type_escaped = value
}
