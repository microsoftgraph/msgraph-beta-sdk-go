package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationOAuth1ConnectionSettings 
type EducationSynchronizationOAuth1ConnectionSettings struct {
    EducationSynchronizationConnectionSettings
}
// NewEducationSynchronizationOAuth1ConnectionSettings instantiates a new educationSynchronizationOAuth1ConnectionSettings and sets the default values.
func NewEducationSynchronizationOAuth1ConnectionSettings()(*EducationSynchronizationOAuth1ConnectionSettings) {
    m := &EducationSynchronizationOAuth1ConnectionSettings{
        EducationSynchronizationConnectionSettings: *NewEducationSynchronizationConnectionSettings(),
    }
    odataTypeValue := "#microsoft.graph.educationSynchronizationOAuth1ConnectionSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationSynchronizationOAuth1ConnectionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationOAuth1ConnectionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationOAuth1ConnectionSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationOAuth1ConnectionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationConnectionSettings.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationOAuth1ConnectionSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationOAuth1ConnectionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationConnectionSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationOAuth1ConnectionSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EducationSynchronizationOAuth1ConnectionSettingsable 
type EducationSynchronizationOAuth1ConnectionSettingsable interface {
    EducationSynchronizationConnectionSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
