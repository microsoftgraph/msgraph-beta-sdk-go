package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementIntent 
type ManagementIntent struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name for the management intent. Optional. Read-only.
    displayName *string
    // A flag indicating whether the management intent is global. Required. Read-only.
    isGlobal *bool
}
// NewManagementIntent instantiates a new ManagementIntent and sets the default values.
func NewManagementIntent()(*ManagementIntent) {
    m := &ManagementIntent{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementIntent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementIntent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isGlobal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobal(val)
        }
        return nil
    }
    return res
}
// GetIsGlobal gets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) GetIsGlobal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobal
    }
}
// Serialize serializes information the current object
func (m *ManagementIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isGlobal", m.GetIsGlobal())
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
func (m *ManagementIntent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsGlobal sets the isGlobal property value. A flag indicating whether the management intent is global. Required. Read-only.
func (m *ManagementIntent) SetIsGlobal(value *bool)() {
    if m != nil {
        m.isGlobal = value
    }
}
