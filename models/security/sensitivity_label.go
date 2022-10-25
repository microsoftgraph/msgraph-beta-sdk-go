package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SensitivityLabel provides operations to manage the collection of accessReview entities.
type SensitivityLabel struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The color that the UI should display for the label, if configured.
    color *string
    // Returns the supported content formats for the label.
    contentFormats []string
    // The admin-defined description for the label.
    description *string
    // Indicates whether the label has protection actions configured.
    hasProtection *bool
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
    isActive *bool
    // Indicates whether the label can be applied to content. False if the label is a parent with child labels.
    isAppliable *bool
    // The plaintext name of the label.
    name *string
    // The parent label associated with a child label. Null if the label has no parent.
    parent SensitivityLabelable
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32
    // The tooltip that should be displayed for the label in a UI.
    tooltip *string
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.sensitivityLabel";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetColor gets the color property value. The color that the UI should display for the label, if configured.
func (m *SensitivityLabel) GetColor()(*string) {
    return m.color
}
// GetContentFormats gets the contentFormats property value. Returns the supported content formats for the label.
func (m *SensitivityLabel) GetContentFormats()([]string) {
    return m.contentFormats
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *SensitivityLabel) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetColor)
    res["contentFormats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetContentFormats)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["hasProtection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasProtection)
    res["isActive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsActive)
    res["isAppliable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAppliable)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["parent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSensitivityLabelFromDiscriminatorValue , m.SetParent)
    res["sensitivity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSensitivity)
    res["tooltip"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTooltip)
    return res
}
// GetHasProtection gets the hasProtection property value. Indicates whether the label has protection actions configured.
func (m *SensitivityLabel) GetHasProtection()(*bool) {
    return m.hasProtection
}
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
func (m *SensitivityLabel) GetIsActive()(*bool) {
    return m.isActive
}
// GetIsAppliable gets the isAppliable property value. Indicates whether the label can be applied to content. False if the label is a parent with child labels.
func (m *SensitivityLabel) GetIsAppliable()(*bool) {
    return m.isAppliable
}
// GetName gets the name property value. The plaintext name of the label.
func (m *SensitivityLabel) GetName()(*string) {
    return m.name
}
// GetParent gets the parent property value. The parent label associated with a child label. Null if the label has no parent.
func (m *SensitivityLabel) GetParent()(SensitivityLabelable) {
    return m.parent
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *SensitivityLabel) GetSensitivity()(*int32) {
    return m.sensitivity
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *SensitivityLabel) GetTooltip()(*string) {
    return m.tooltip
}
// Serialize serializes information the current object
func (m *SensitivityLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetContentFormats() != nil {
        err = writer.WriteCollectionOfStringValues("contentFormats", m.GetContentFormats())
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
        err = writer.WriteBoolValue("hasProtection", m.GetHasProtection())
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
        err = writer.WriteBoolValue("isAppliable", m.GetIsAppliable())
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
func (m *SensitivityLabel) SetColor(value *string)() {
    m.color = value
}
// SetContentFormats sets the contentFormats property value. Returns the supported content formats for the label.
func (m *SensitivityLabel) SetContentFormats(value []string)() {
    m.contentFormats = value
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *SensitivityLabel) SetDescription(value *string)() {
    m.description = value
}
// SetHasProtection sets the hasProtection property value. Indicates whether the label has protection actions configured.
func (m *SensitivityLabel) SetHasProtection(value *bool)() {
    m.hasProtection = value
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
func (m *SensitivityLabel) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetIsAppliable sets the isAppliable property value. Indicates whether the label can be applied to content. False if the label is a parent with child labels.
func (m *SensitivityLabel) SetIsAppliable(value *bool)() {
    m.isAppliable = value
}
// SetName sets the name property value. The plaintext name of the label.
func (m *SensitivityLabel) SetName(value *string)() {
    m.name = value
}
// SetParent sets the parent property value. The parent label associated with a child label. Null if the label has no parent.
func (m *SensitivityLabel) SetParent(value SensitivityLabelable)() {
    m.parent = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *SensitivityLabel) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *SensitivityLabel) SetTooltip(value *string)() {
    m.tooltip = value
}
