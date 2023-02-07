package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMinimumOperatingSystem the minimum operating system required for a Windows mobile app.
type WindowsMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Windows version 10.0 or later.
    v10_0 *bool
    // Windows 10 1607 or later.
    v10_1607 *bool
    // Windows 10 1703 or later.
    v10_1703 *bool
    // Windows 10 1709 or later.
    v10_1709 *bool
    // Windows 10 1803 or later.
    v10_1803 *bool
    // Windows 10 1809 or later.
    v10_1809 *bool
    // Windows 10 1903 or later.
    v10_1903 *bool
    // Windows 10 1909 or later.
    v10_1909 *bool
    // Windows 10 2004 or later.
    v10_2004 *bool
    // Windows 10 21H1 or later.
    v10_21H1 *bool
    // Windows 10 2H20 or later.
    v10_2H20 *bool
    // Windows version 8.0 or later.
    v8_0 *bool
    // Windows version 8.1 or later.
    v8_1 *bool
}
// NewWindowsMinimumOperatingSystem instantiates a new windowsMinimumOperatingSystem and sets the default values.
func NewWindowsMinimumOperatingSystem()(*WindowsMinimumOperatingSystem) {
    m := &WindowsMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["v10_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV100(val)
        }
        return nil
    }
    res["v10_1607"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101607(val)
        }
        return nil
    }
    res["v10_1703"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101703(val)
        }
        return nil
    }
    res["v10_1709"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101709(val)
        }
        return nil
    }
    res["v10_1803"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101803(val)
        }
        return nil
    }
    res["v10_1809"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101809(val)
        }
        return nil
    }
    res["v10_1903"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101903(val)
        }
        return nil
    }
    res["v10_1909"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV101909(val)
        }
        return nil
    }
    res["v10_2004"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV102004(val)
        }
        return nil
    }
    res["v10_21H1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1021H1(val)
        }
        return nil
    }
    res["v10_2H20"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV102H20(val)
        }
        return nil
    }
    res["v8_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV80(val)
        }
        return nil
    }
    res["v8_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV81(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV100 gets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV100()(*bool) {
    return m.v10_0
}
// GetV101607 gets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) GetV101607()(*bool) {
    return m.v10_1607
}
// GetV101703 gets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) GetV101703()(*bool) {
    return m.v10_1703
}
// GetV101709 gets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) GetV101709()(*bool) {
    return m.v10_1709
}
// GetV101803 gets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) GetV101803()(*bool) {
    return m.v10_1803
}
// GetV101809 gets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) GetV101809()(*bool) {
    return m.v10_1809
}
// GetV101903 gets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) GetV101903()(*bool) {
    return m.v10_1903
}
// GetV101909 gets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) GetV101909()(*bool) {
    return m.v10_1909
}
// GetV102004 gets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) GetV102004()(*bool) {
    return m.v10_2004
}
// GetV1021H1 gets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) GetV1021H1()(*bool) {
    return m.v10_21H1
}
// GetV102H20 gets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) GetV102H20()(*bool) {
    return m.v10_2H20
}
// GetV80 gets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV80()(*bool) {
    return m.v8_0
}
// GetV81 gets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) GetV81()(*bool) {
    return m.v8_1
}
// Serialize serializes information the current object
func (m *WindowsMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_0", m.GetV100())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1607", m.GetV101607())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1703", m.GetV101703())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1709", m.GetV101709())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1803", m.GetV101803())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1809", m.GetV101809())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1903", m.GetV101903())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_1909", m.GetV101909())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_2004", m.GetV102004())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_21H1", m.GetV1021H1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_2H20", m.GetV102H20())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_0", m.GetV80())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_1", m.GetV81())
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
func (m *WindowsMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV100 sets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV100(value *bool)() {
    m.v10_0 = value
}
// SetV101607 sets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) SetV101607(value *bool)() {
    m.v10_1607 = value
}
// SetV101703 sets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) SetV101703(value *bool)() {
    m.v10_1703 = value
}
// SetV101709 sets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) SetV101709(value *bool)() {
    m.v10_1709 = value
}
// SetV101803 sets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) SetV101803(value *bool)() {
    m.v10_1803 = value
}
// SetV101809 sets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) SetV101809(value *bool)() {
    m.v10_1809 = value
}
// SetV101903 sets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) SetV101903(value *bool)() {
    m.v10_1903 = value
}
// SetV101909 sets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) SetV101909(value *bool)() {
    m.v10_1909 = value
}
// SetV102004 sets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) SetV102004(value *bool)() {
    m.v10_2004 = value
}
// SetV1021H1 sets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) SetV1021H1(value *bool)() {
    m.v10_21H1 = value
}
// SetV102H20 sets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) SetV102H20(value *bool)() {
    m.v10_2H20 = value
}
// SetV80 sets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV80(value *bool)() {
    m.v8_0 = value
}
// SetV81 sets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) SetV81(value *bool)() {
    m.v8_1 = value
}
