package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppSettings 
type TeamsAppSettings struct {
    Entity
    // The isChatResourceSpecificConsentEnabled property
    isChatResourceSpecificConsentEnabled *bool
}
// NewTeamsAppSettings instantiates a new TeamsAppSettings and sets the default values.
func NewTeamsAppSettings()(*TeamsAppSettings) {
    m := &TeamsAppSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isChatResourceSpecificConsentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsChatResourceSpecificConsentEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsChatResourceSpecificConsentEnabled gets the isChatResourceSpecificConsentEnabled property value. The isChatResourceSpecificConsentEnabled property
func (m *TeamsAppSettings) GetIsChatResourceSpecificConsentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isChatResourceSpecificConsentEnabled
    }
}
// Serialize serializes information the current object
func (m *TeamsAppSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isChatResourceSpecificConsentEnabled", m.GetIsChatResourceSpecificConsentEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsChatResourceSpecificConsentEnabled sets the isChatResourceSpecificConsentEnabled property value. The isChatResourceSpecificConsentEnabled property
func (m *TeamsAppSettings) SetIsChatResourceSpecificConsentEnabled(value *bool)() {
    if m != nil {
        m.isChatResourceSpecificConsentEnabled = value
    }
}
