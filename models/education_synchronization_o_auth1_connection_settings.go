package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationOAuth1ConnectionSettings 
type EducationSynchronizationOAuth1ConnectionSettings struct {
    EducationSynchronizationConnectionSettings
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewEducationSynchronizationOAuth1ConnectionSettings instantiates a new EducationSynchronizationOAuth1ConnectionSettings and sets the default values.
func NewEducationSynchronizationOAuth1ConnectionSettings()(*EducationSynchronizationOAuth1ConnectionSettings) {
    m := &EducationSynchronizationOAuth1ConnectionSettings{
        EducationSynchronizationConnectionSettings: *NewEducationSynchronizationConnectionSettings(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationSynchronizationOAuth1ConnectionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationOAuth1ConnectionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationOAuth1ConnectionSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationOAuth1ConnectionSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationOAuth1ConnectionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationConnectionSettings.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EducationSynchronizationOAuth1ConnectionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationConnectionSettings.Serialize(writer)
    if err != nil {
        return err
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
func (m *EducationSynchronizationOAuth1ConnectionSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
