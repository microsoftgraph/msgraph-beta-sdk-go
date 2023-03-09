package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppProductCodeDetection 
type Win32LobAppProductCodeDetection struct {
    Win32LobAppDetection
}
// NewWin32LobAppProductCodeDetection instantiates a new Win32LobAppProductCodeDetection and sets the default values.
func NewWin32LobAppProductCodeDetection()(*Win32LobAppProductCodeDetection) {
    m := &Win32LobAppProductCodeDetection{
        Win32LobAppDetection: *NewWin32LobAppDetection(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppProductCodeDetection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppProductCodeDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppProductCodeDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppProductCodeDetection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppProductCodeDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppDetection.GetFieldDeserializers()
    res["productCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductCode(val)
        }
        return nil
    }
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    res["productVersionOperator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDetectionOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersionOperator(val.(*Win32LobAppDetectionOperator))
        }
        return nil
    }
    return res
}
// GetProductCode gets the productCode property value. The product code of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) GetProductCode()(*string) {
    val, err := m.GetBackingStore().Get("productCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductVersion gets the productVersion property value. The product version of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) GetProductVersion()(*string) {
    val, err := m.GetBackingStore().Get("productVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductVersionOperator gets the productVersionOperator property value. Contains properties for detection operator.
func (m *Win32LobAppProductCodeDetection) GetProductVersionOperator()(*Win32LobAppDetectionOperator) {
    val, err := m.GetBackingStore().Get("productVersionOperator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppDetectionOperator)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobAppProductCodeDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppDetection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("productCode", m.GetProductCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    if m.GetProductVersionOperator() != nil {
        cast := (*m.GetProductVersionOperator()).String()
        err = writer.WriteStringValue("productVersionOperator", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProductCode sets the productCode property value. The product code of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) SetProductCode(value *string)() {
    err := m.GetBackingStore().Set("productCode", value)
    if err != nil {
        panic(err)
    }
}
// SetProductVersion sets the productVersion property value. The product version of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) SetProductVersion(value *string)() {
    err := m.GetBackingStore().Set("productVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetProductVersionOperator sets the productVersionOperator property value. Contains properties for detection operator.
func (m *Win32LobAppProductCodeDetection) SetProductVersionOperator(value *Win32LobAppDetectionOperator)() {
    err := m.GetBackingStore().Set("productVersionOperator", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppProductCodeDetectionable 
type Win32LobAppProductCodeDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppDetectionable
    GetProductCode()(*string)
    GetProductVersion()(*string)
    GetProductVersionOperator()(*Win32LobAppDetectionOperator)
    SetProductCode(value *string)()
    SetProductVersion(value *string)()
    SetProductVersionOperator(value *Win32LobAppDetectionOperator)()
}
