package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SimplePasswordSettings struct {
    PasswordSettings
}
// NewSimplePasswordSettings instantiates a new SimplePasswordSettings and sets the default values.
func NewSimplePasswordSettings()(*SimplePasswordSettings) {
    m := &SimplePasswordSettings{
        PasswordSettings: *NewPasswordSettings(),
    }
    odataTypeValue := "#microsoft.graph.industryData.simplePasswordSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSimplePasswordSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimplePasswordSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimplePasswordSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SimplePasswordSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PasswordSettings.GetFieldDeserializers()
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password for the user.
// returns a *string when successful
func (m *SimplePasswordSettings) GetPassword()(*string) {
    val, err := m.GetBackingStore().Get("password")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SimplePasswordSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PasswordSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPassword sets the password property value. The password for the user.
func (m *SimplePasswordSettings) SetPassword(value *string)() {
    err := m.GetBackingStore().Set("password", value)
    if err != nil {
        panic(err)
    }
}
type SimplePasswordSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PasswordSettingsable
    GetPassword()(*string)
    SetPassword(value *string)()
}
