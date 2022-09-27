package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppProductCodeDetection 
type Win32LobAppProductCodeDetection struct {
    Win32LobAppDetection
    // The product code of Win32 Line of Business (LoB) app.
    productCode *string
    // The product version of Win32 Line of Business (LoB) app.
    productVersion *string
    // Contains properties for detection operator.
    productVersionOperator *Win32LobAppDetectionOperator
}
// NewWin32LobAppProductCodeDetection instantiates a new Win32LobAppProductCodeDetection and sets the default values.
func NewWin32LobAppProductCodeDetection()(*Win32LobAppProductCodeDetection) {
    m := &Win32LobAppProductCodeDetection{
        Win32LobAppDetection: *NewWin32LobAppDetection(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppProductCodeDetection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWin32LobAppProductCodeDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppProductCodeDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppProductCodeDetection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppProductCodeDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppDetection.GetFieldDeserializers()
    res["productCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProductCode)
    res["productVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProductVersion)
    res["productVersionOperator"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWin32LobAppDetectionOperator , m.SetProductVersionOperator)
    return res
}
// GetProductCode gets the productCode property value. The product code of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) GetProductCode()(*string) {
    return m.productCode
}
// GetProductVersion gets the productVersion property value. The product version of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) GetProductVersion()(*string) {
    return m.productVersion
}
// GetProductVersionOperator gets the productVersionOperator property value. Contains properties for detection operator.
func (m *Win32LobAppProductCodeDetection) GetProductVersionOperator()(*Win32LobAppDetectionOperator) {
    return m.productVersionOperator
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
    m.productCode = value
}
// SetProductVersion sets the productVersion property value. The product version of Win32 Line of Business (LoB) app.
func (m *Win32LobAppProductCodeDetection) SetProductVersion(value *string)() {
    m.productVersion = value
}
// SetProductVersionOperator sets the productVersionOperator property value. Contains properties for detection operator.
func (m *Win32LobAppProductCodeDetection) SetProductVersionOperator(value *Win32LobAppDetectionOperator)() {
    m.productVersionOperator = value
}
