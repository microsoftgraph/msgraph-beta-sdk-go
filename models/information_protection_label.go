package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionLabel 
type InformationProtectionLabel struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewInformationProtectionLabel instantiates a new informationProtectionLabel and sets the default values.
func NewInformationProtectionLabel()(*InformationProtectionLabel) {
    m := &InformationProtectionLabel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionLabel(), nil
}
// GetColor gets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) GetColor()(*string) {
    val, err := m.GetBackingStore().Get("color")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
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
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParentLabelDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(ParentLabelDetailsable))
        }
        return nil
    }
    res["sensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val)
        }
        return nil
    }
    res["tooltip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTooltip(val)
        }
        return nil
    }
    return res
}
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) GetIsActive()(*bool) {
    val, err := m.GetBackingStore().Get("isActive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParent gets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) GetParent()(ParentLabelDetailsable) {
    val, err := m.GetBackingStore().Get("parent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ParentLabelDetailsable)
    }
    return nil
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) GetSensitivity()(*int32) {
    val, err := m.GetBackingStore().Get("sensitivity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) GetTooltip()(*string) {
    val, err := m.GetBackingStore().Get("tooltip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InformationProtectionLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
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
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sensitivity", m.GetSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tooltip", m.GetTooltip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) SetColor(value *string)() {
    err := m.GetBackingStore().Set("color", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) SetIsActive(value *bool)() {
    err := m.GetBackingStore().Set("isActive", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetParent sets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) SetParent(value ParentLabelDetailsable)() {
    err := m.GetBackingStore().Set("parent", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) SetSensitivity(value *int32)() {
    err := m.GetBackingStore().Set("sensitivity", value)
    if err != nil {
        panic(err)
    }
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) SetTooltip(value *string)() {
    err := m.GetBackingStore().Set("tooltip", value)
    if err != nil {
        panic(err)
    }
}
// InformationProtectionLabelable 
type InformationProtectionLabelable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetDescription()(*string)
    GetIsActive()(*bool)
    GetName()(*string)
    GetParent()(ParentLabelDetailsable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetDescription(value *string)()
    SetIsActive(value *bool)()
    SetName(value *string)()
    SetParent(value ParentLabelDetailsable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
