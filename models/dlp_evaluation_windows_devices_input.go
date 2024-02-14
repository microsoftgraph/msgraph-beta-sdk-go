package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DlpEvaluationWindowsDevicesInput struct {
    DlpEvaluationInput
}
// NewDlpEvaluationWindowsDevicesInput instantiates a new DlpEvaluationWindowsDevicesInput and sets the default values.
func NewDlpEvaluationWindowsDevicesInput()(*DlpEvaluationWindowsDevicesInput) {
    m := &DlpEvaluationWindowsDevicesInput{
        DlpEvaluationInput: *NewDlpEvaluationInput(),
    }
    odataTypeValue := "#microsoft.graph.dlpEvaluationWindowsDevicesInput"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDlpEvaluationWindowsDevicesInputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDlpEvaluationWindowsDevicesInputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpEvaluationWindowsDevicesInput(), nil
}
// GetContentProperties gets the contentProperties property value. The contentProperties property
// returns a ContentPropertiesable when successful
func (m *DlpEvaluationWindowsDevicesInput) GetContentProperties()(ContentPropertiesable) {
    val, err := m.GetBackingStore().Get("contentProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContentPropertiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DlpEvaluationWindowsDevicesInput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpEvaluationInput.GetFieldDeserializers()
    res["contentProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentPropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentProperties(val.(ContentPropertiesable))
        }
        return nil
    }
    res["sharedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val)
        }
        return nil
    }
    return res
}
// GetSharedBy gets the sharedBy property value. The sharedBy property
// returns a *string when successful
func (m *DlpEvaluationWindowsDevicesInput) GetSharedBy()(*string) {
    val, err := m.GetBackingStore().Get("sharedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DlpEvaluationWindowsDevicesInput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpEvaluationInput.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentProperties", m.GetContentProperties())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentProperties sets the contentProperties property value. The contentProperties property
func (m *DlpEvaluationWindowsDevicesInput) SetContentProperties(value ContentPropertiesable)() {
    err := m.GetBackingStore().Set("contentProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedBy sets the sharedBy property value. The sharedBy property
func (m *DlpEvaluationWindowsDevicesInput) SetSharedBy(value *string)() {
    err := m.GetBackingStore().Set("sharedBy", value)
    if err != nil {
        panic(err)
    }
}
type DlpEvaluationWindowsDevicesInputable interface {
    DlpEvaluationInputable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentProperties()(ContentPropertiesable)
    GetSharedBy()(*string)
    SetContentProperties(value ContentPropertiesable)()
    SetSharedBy(value *string)()
}
