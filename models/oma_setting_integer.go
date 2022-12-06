package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingInteger 
type OmaSettingInteger struct {
    OmaSetting
    // By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
    isReadOnly *bool
    // Value.
    value *int32
}
// NewOmaSettingInteger instantiates a new OmaSettingInteger and sets the default values.
func NewOmaSettingInteger()(*OmaSettingInteger) {
    m := &OmaSettingInteger{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingInteger";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOmaSettingIntegerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingIntegerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingInteger(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingInteger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["isReadOnly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReadOnly)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetValue)
    return res
}
// GetIsReadOnly gets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) GetIsReadOnly()(*bool) {
    return m.isReadOnly
}
// GetValue gets the value property value. Value.
func (m *OmaSettingInteger) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingInteger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsReadOnly sets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) SetIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
// SetValue sets the value property value. Value.
func (m *OmaSettingInteger) SetValue(value *int32)() {
    m.value = value
}
