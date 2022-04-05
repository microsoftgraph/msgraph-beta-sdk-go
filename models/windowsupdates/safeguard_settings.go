package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SafeguardSettings 
type SafeguardSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of safeguards to ignore per device.
    disabledSafeguardProfiles []SafeguardProfileable;
}
// NewSafeguardSettings instantiates a new safeguardSettings and sets the default values.
func NewSafeguardSettings()(*SafeguardSettings) {
    m := &SafeguardSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSafeguardSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSafeguardSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSafeguardSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SafeguardSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisabledSafeguardProfiles gets the disabledSafeguardProfiles property value. List of safeguards to ignore per device.
func (m *SafeguardSettings) GetDisabledSafeguardProfiles()([]SafeguardProfileable) {
    if m == nil {
        return nil
    } else {
        return m.disabledSafeguardProfiles
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SafeguardSettings) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["disabledSafeguardProfiles"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSafeguardProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SafeguardProfileable, len(val))
            for i, v := range val {
                res[i] = v.(SafeguardProfileable)
            }
            m.SetDisabledSafeguardProfiles(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SafeguardSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDisabledSafeguardProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDisabledSafeguardProfiles()))
        for i, v := range m.GetDisabledSafeguardProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("disabledSafeguardProfiles", cast)
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
func (m *SafeguardSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisabledSafeguardProfiles sets the disabledSafeguardProfiles property value. List of safeguards to ignore per device.
func (m *SafeguardSettings) SetDisabledSafeguardProfiles(value []SafeguardProfileable)() {
    if m != nil {
        m.disabledSafeguardProfiles = value
    }
}
