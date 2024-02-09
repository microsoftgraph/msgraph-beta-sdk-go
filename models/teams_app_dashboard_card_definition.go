package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamsAppDashboardCardDefinition struct {
    Entity
}
// NewTeamsAppDashboardCardDefinition instantiates a new TeamsAppDashboardCardDefinition and sets the default values.
func NewTeamsAppDashboardCardDefinition()(*TeamsAppDashboardCardDefinition) {
    m := &TeamsAppDashboardCardDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppDashboardCardDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsAppDashboardCardDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppDashboardCardDefinition(), nil
}
// GetContentSource gets the contentSource property value. The contentSource property
// returns a TeamsAppDashboardCardContentSourceable when successful
func (m *TeamsAppDashboardCardDefinition) GetContentSource()(TeamsAppDashboardCardContentSourceable) {
    val, err := m.GetBackingStore().Get("contentSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsAppDashboardCardContentSourceable)
    }
    return nil
}
// GetDefaultSize gets the defaultSize property value. The defaultSize property
// returns a *TeamsAppDashboardCardSize when successful
func (m *TeamsAppDashboardCardDefinition) GetDefaultSize()(*TeamsAppDashboardCardSize) {
    val, err := m.GetBackingStore().Get("defaultSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamsAppDashboardCardSize)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TeamsAppDashboardCardDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *TeamsAppDashboardCardDefinition) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsAppDashboardCardDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppDashboardCardContentSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSource(val.(TeamsAppDashboardCardContentSourceable))
        }
        return nil
    }
    res["defaultSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppDashboardCardSize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSize(val.(*TeamsAppDashboardCardSize))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["icon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppDashboardCardIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIcon(val.(TeamsAppDashboardCardIconable))
        }
        return nil
    }
    res["pickerGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPickerGroupId(val)
        }
        return nil
    }
    return res
}
// GetIcon gets the icon property value. The icon property
// returns a TeamsAppDashboardCardIconable when successful
func (m *TeamsAppDashboardCardDefinition) GetIcon()(TeamsAppDashboardCardIconable) {
    val, err := m.GetBackingStore().Get("icon")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsAppDashboardCardIconable)
    }
    return nil
}
// GetPickerGroupId gets the pickerGroupId property value. The pickerGroupId property
// returns a *string when successful
func (m *TeamsAppDashboardCardDefinition) GetPickerGroupId()(*string) {
    val, err := m.GetBackingStore().Get("pickerGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamsAppDashboardCardDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentSource", m.GetContentSource())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultSize() != nil {
        cast := (*m.GetDefaultSize()).String()
        err = writer.WriteStringValue("defaultSize", &cast)
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
    {
        err = writer.WriteObjectValue("icon", m.GetIcon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("pickerGroupId", m.GetPickerGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentSource sets the contentSource property value. The contentSource property
func (m *TeamsAppDashboardCardDefinition) SetContentSource(value TeamsAppDashboardCardContentSourceable)() {
    err := m.GetBackingStore().Set("contentSource", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultSize sets the defaultSize property value. The defaultSize property
func (m *TeamsAppDashboardCardDefinition) SetDefaultSize(value *TeamsAppDashboardCardSize)() {
    err := m.GetBackingStore().Set("defaultSize", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *TeamsAppDashboardCardDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TeamsAppDashboardCardDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIcon sets the icon property value. The icon property
func (m *TeamsAppDashboardCardDefinition) SetIcon(value TeamsAppDashboardCardIconable)() {
    err := m.GetBackingStore().Set("icon", value)
    if err != nil {
        panic(err)
    }
}
// SetPickerGroupId sets the pickerGroupId property value. The pickerGroupId property
func (m *TeamsAppDashboardCardDefinition) SetPickerGroupId(value *string)() {
    err := m.GetBackingStore().Set("pickerGroupId", value)
    if err != nil {
        panic(err)
    }
}
type TeamsAppDashboardCardDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentSource()(TeamsAppDashboardCardContentSourceable)
    GetDefaultSize()(*TeamsAppDashboardCardSize)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIcon()(TeamsAppDashboardCardIconable)
    GetPickerGroupId()(*string)
    SetContentSource(value TeamsAppDashboardCardContentSourceable)()
    SetDefaultSize(value *TeamsAppDashboardCardSize)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIcon(value TeamsAppDashboardCardIconable)()
    SetPickerGroupId(value *string)()
}
