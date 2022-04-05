package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParentLabelDetails 
type ParentLabelDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The color that the user interface should display for the label, if configured.
    color *string;
    // The admin-defined description for the label.
    description *string;
    // The label ID is a globally unique identifier (GUID).
    id *string;
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
    isActive *bool;
    // The plaintext name of the label.
    name *string;
    // The parent property
    parent ParentLabelDetailsable;
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32;
    // The tooltip that should be displayed for the label in a user interface.
    tooltip *string;
}
// NewParentLabelDetails instantiates a new parentLabelDetails and sets the default values.
func NewParentLabelDetails()(*ParentLabelDetails) {
    m := &ParentLabelDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParentLabelDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParentLabelDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParentLabelDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParentLabelDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetColor gets the color property value. The color that the user interface should display for the label, if configured.
func (m *ParentLabelDetails) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *ParentLabelDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParentLabelDetails) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["id"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
        val, err := n.GetObjectValue(CreateParentLabelDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(ParentLabelDetailsable))
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
// GetId gets the id property value. The label ID is a globally unique identifier (GUID).
func (m *ParentLabelDetails) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
func (m *ParentLabelDetails) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetName gets the name property value. The plaintext name of the label.
func (m *ParentLabelDetails) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParent gets the parent property value. The parent property
func (m *ParentLabelDetails) GetParent()(ParentLabelDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *ParentLabelDetails) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a user interface.
func (m *ParentLabelDetails) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
// Serialize serializes information the current object
func (m *ParentLabelDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sensitivity", m.GetSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tooltip", m.GetTooltip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParentLabelDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetColor sets the color property value. The color that the user interface should display for the label, if configured.
func (m *ParentLabelDetails) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *ParentLabelDetails) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetId sets the id property value. The label ID is a globally unique identifier (GUID).
func (m *ParentLabelDetails) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
func (m *ParentLabelDetails) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetName sets the name property value. The plaintext name of the label.
func (m *ParentLabelDetails) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParent sets the parent property value. The parent property
func (m *ParentLabelDetails) SetParent(value ParentLabelDetailsable)() {
    if m != nil {
        m.parent = value
    }
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *ParentLabelDetails) SetSensitivity(value *int32)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a user interface.
func (m *ParentLabelDetails) SetTooltip(value *string)() {
    if m != nil {
        m.tooltip = value
    }
}
