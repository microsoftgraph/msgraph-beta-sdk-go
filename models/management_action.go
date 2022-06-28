package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementAction 
type ManagementAction struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description for the management action. Optional. Read-only.
    description *string
    // The display name for the management action. Optional. Read-only.
    displayName *string
    // The reference for the management template used to generate the management action. Required. Read-only.
    referenceTemplateId *string
    // The referenceTemplateVersion property
    referenceTemplateVersion *int32
}
// NewManagementAction instantiates a new ManagementAction and sets the default values.
func NewManagementAction()(*ManagementAction) {
    m := &ManagementAction{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["referenceTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateId(val)
        }
        return nil
    }
    res["referenceTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceTemplateVersion(val)
        }
        return nil
    }
    return res
}
// GetReferenceTemplateId gets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) GetReferenceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceTemplateId
    }
}
// GetReferenceTemplateVersion gets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) GetReferenceTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referenceTemplateVersion
    }
}
// Serialize serializes information the current object
func (m *ManagementAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("referenceTemplateId", m.GetReferenceTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referenceTemplateVersion", m.GetReferenceTemplateVersion())
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
func (m *ManagementAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetReferenceTemplateId sets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) SetReferenceTemplateId(value *string)() {
    if m != nil {
        m.referenceTemplateId = value
    }
}
// SetReferenceTemplateVersion sets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) SetReferenceTemplateVersion(value *int32)() {
    if m != nil {
        m.referenceTemplateVersion = value
    }
}
