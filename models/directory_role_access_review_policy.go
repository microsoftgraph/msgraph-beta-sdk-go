package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryRoleAccessReviewPolicy 
type DirectoryRoleAccessReviewPolicy struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The settings property
    settings AccessReviewScheduleSettingsable
}
// NewDirectoryRoleAccessReviewPolicy instantiates a new DirectoryRoleAccessReviewPolicy and sets the default values.
func NewDirectoryRoleAccessReviewPolicy()(*DirectoryRoleAccessReviewPolicy) {
    m := &DirectoryRoleAccessReviewPolicy{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDirectoryRoleAccessReviewPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryRoleAccessReviewPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAccessReviewPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectoryRoleAccessReviewPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryRoleAccessReviewPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewScheduleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(AccessReviewScheduleSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings property
func (m *DirectoryRoleAccessReviewPolicy) GetSettings()(AccessReviewScheduleSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Serialize serializes information the current object
func (m *DirectoryRoleAccessReviewPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
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
func (m *DirectoryRoleAccessReviewPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSettings sets the settings property value. The settings property
func (m *DirectoryRoleAccessReviewPolicy) SetSettings(value AccessReviewScheduleSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
