package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SensitivityLabel 
type SensitivityLabel struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The color property
    color *string;
    // The contentFormats property
    contentFormats []string;
    // The description property
    description *string;
    // The hasProtection property
    hasProtection *bool;
    // The isActive property
    isActive *bool;
    // The isAppliable property
    isAppliable *bool;
    // The name property
    name *string;
    // The parent property
    parent SensitivityLabelable;
    // The sensitivity property
    sensitivity *int32;
    // The tooltip property
    tooltip *string;
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetColor gets the color property value. The color property
func (m *SensitivityLabel) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetContentFormats gets the contentFormats property value. The contentFormats property
func (m *SensitivityLabel) GetContentFormats()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentFormats
    }
}
// GetDescription gets the description property value. The description property
func (m *SensitivityLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["contentFormats"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentFormats(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["hasProtection"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasProtection(val)
        }
        return nil
    }
    res["isActive"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["isAppliable"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAppliable(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parent"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(SensitivityLabelable))
        }
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val)
        }
        return nil
    }
    res["tooltip"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHasProtection gets the hasProtection property value. The hasProtection property
func (m *SensitivityLabel) GetHasProtection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasProtection
    }
}
// GetIsActive gets the isActive property value. The isActive property
func (m *SensitivityLabel) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetIsAppliable gets the isAppliable property value. The isAppliable property
func (m *SensitivityLabel) GetIsAppliable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAppliable
    }
}
// GetName gets the name property value. The name property
func (m *SensitivityLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParent gets the parent property value. The parent property
func (m *SensitivityLabel) GetParent()(SensitivityLabelable) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetSensitivity gets the sensitivity property value. The sensitivity property
func (m *SensitivityLabel) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetTooltip gets the tooltip property value. The tooltip property
func (m *SensitivityLabel) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
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
// SetColor sets the color property value. The color property
func (m *SensitivityLabel) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
// SetContentFormats sets the contentFormats property value. The contentFormats property
func (m *SensitivityLabel) SetContentFormats(value []string)() {
    if m != nil {
        m.contentFormats = value
    }
}
// SetDescription sets the description property value. The description property
func (m *SensitivityLabel) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetHasProtection sets the hasProtection property value. The hasProtection property
func (m *SensitivityLabel) SetHasProtection(value *bool)() {
    if m != nil {
        m.hasProtection = value
    }
}
// SetIsActive sets the isActive property value. The isActive property
func (m *SensitivityLabel) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetIsAppliable sets the isAppliable property value. The isAppliable property
func (m *SensitivityLabel) SetIsAppliable(value *bool)() {
    if m != nil {
        m.isAppliable = value
    }
}
// SetName sets the name property value. The name property
func (m *SensitivityLabel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParent sets the parent property value. The parent property
func (m *SensitivityLabel) SetParent(value SensitivityLabelable)() {
    if m != nil {
        m.parent = value
    }
}
// SetSensitivity sets the sensitivity property value. The sensitivity property
func (m *SensitivityLabel) SetSensitivity(value *int32)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetTooltip sets the tooltip property value. The tooltip property
func (m *SensitivityLabel) SetTooltip(value *string)() {
    if m != nil {
        m.tooltip = value
    }
}
