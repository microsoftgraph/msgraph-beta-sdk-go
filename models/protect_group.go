package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectGroup 
type ProtectGroup struct {
    LabelActionBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowEmailFromGuestUsers property
    allowEmailFromGuestUsers *bool
    // The allowGuestUsers property
    allowGuestUsers *bool
    // The privacy property
    privacy *GroupPrivacy
}
// NewProtectGroup instantiates a new ProtectGroup and sets the default values.
func NewProtectGroup()(*ProtectGroup) {
    m := &ProtectGroup{
        LabelActionBase: *NewLabelActionBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProtectGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowEmailFromGuestUsers gets the allowEmailFromGuestUsers property value. The allowEmailFromGuestUsers property
func (m *ProtectGroup) GetAllowEmailFromGuestUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEmailFromGuestUsers
    }
}
// GetAllowGuestUsers gets the allowGuestUsers property value. The allowGuestUsers property
func (m *ProtectGroup) GetAllowGuestUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowGuestUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["allowEmailFromGuestUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmailFromGuestUsers(val)
        }
        return nil
    }
    res["allowGuestUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowGuestUsers(val)
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPrivacy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val.(*GroupPrivacy))
        }
        return nil
    }
    return res
}
// GetPrivacy gets the privacy property value. The privacy property
func (m *ProtectGroup) GetPrivacy()(*GroupPrivacy) {
    if m == nil {
        return nil
    } else {
        return m.privacy
    }
}
// Serialize serializes information the current object
func (m *ProtectGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowEmailFromGuestUsers", m.GetAllowEmailFromGuestUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowGuestUsers", m.GetAllowGuestUsers())
        if err != nil {
            return err
        }
    }
    if m.GetPrivacy() != nil {
        cast := (*m.GetPrivacy()).String()
        err = writer.WriteStringValue("privacy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowEmailFromGuestUsers sets the allowEmailFromGuestUsers property value. The allowEmailFromGuestUsers property
func (m *ProtectGroup) SetAllowEmailFromGuestUsers(value *bool)() {
    if m != nil {
        m.allowEmailFromGuestUsers = value
    }
}
// SetAllowGuestUsers sets the allowGuestUsers property value. The allowGuestUsers property
func (m *ProtectGroup) SetAllowGuestUsers(value *bool)() {
    if m != nil {
        m.allowGuestUsers = value
    }
}
// SetPrivacy sets the privacy property value. The privacy property
func (m *ProtectGroup) SetPrivacy(value *GroupPrivacy)() {
    if m != nil {
        m.privacy = value
    }
}
