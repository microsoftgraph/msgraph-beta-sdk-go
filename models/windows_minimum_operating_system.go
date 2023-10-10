package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsMinimumOperatingSystem the minimum operating system required for a Windows mobile app.
type WindowsMinimumOperatingSystem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsMinimumOperatingSystem instantiates a new windowsMinimumOperatingSystem and sets the default values.
func NewWindowsMinimumOperatingSystem()(*WindowsMinimumOperatingSystem) {
    m := &WindowsMinimumOperatingSystem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *WindowsMinimumOperatingSystem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
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
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetV100 gets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV100()(*bool) {
    val, err := m.GetBackingStore().Get("v10_0")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101607 gets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) GetV101607()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1607")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101703 gets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) GetV101703()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1703")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101709 gets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) GetV101709()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1709")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101803 gets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) GetV101803()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1803")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101809 gets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) GetV101809()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1809")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101903 gets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) GetV101903()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1903")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV101909 gets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) GetV101909()(*bool) {
    val, err := m.GetBackingStore().Get("v10_1909")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV102004 gets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) GetV102004()(*bool) {
    val, err := m.GetBackingStore().Get("v10_2004")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV1021H1 gets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) GetV1021H1()(*bool) {
    val, err := m.GetBackingStore().Get("v10_21H1")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV102H20 gets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) GetV102H20()(*bool) {
    val, err := m.GetBackingStore().Get("v10_2H20")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV80 gets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV80()(*bool) {
    val, err := m.GetBackingStore().Get("v8_0")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetV81 gets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) GetV81()(*bool) {
    val, err := m.GetBackingStore().Get("v8_1")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsMinimumOperatingSystem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetV100 sets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV100(value *bool)() {
    err := m.GetBackingStore().Set("v10_0", value)
    if err != nil {
        panic(err)
    }
}
// SetV101607 sets the v10_1607 property value. Windows 10 1607 or later.
func (m *WindowsMinimumOperatingSystem) SetV101607(value *bool)() {
    err := m.GetBackingStore().Set("v10_1607", value)
    if err != nil {
        panic(err)
    }
}
// SetV101703 sets the v10_1703 property value. Windows 10 1703 or later.
func (m *WindowsMinimumOperatingSystem) SetV101703(value *bool)() {
    err := m.GetBackingStore().Set("v10_1703", value)
    if err != nil {
        panic(err)
    }
}
// SetV101709 sets the v10_1709 property value. Windows 10 1709 or later.
func (m *WindowsMinimumOperatingSystem) SetV101709(value *bool)() {
    err := m.GetBackingStore().Set("v10_1709", value)
    if err != nil {
        panic(err)
    }
}
// SetV101803 sets the v10_1803 property value. Windows 10 1803 or later.
func (m *WindowsMinimumOperatingSystem) SetV101803(value *bool)() {
    err := m.GetBackingStore().Set("v10_1803", value)
    if err != nil {
        panic(err)
    }
}
// SetV101809 sets the v10_1809 property value. Windows 10 1809 or later.
func (m *WindowsMinimumOperatingSystem) SetV101809(value *bool)() {
    err := m.GetBackingStore().Set("v10_1809", value)
    if err != nil {
        panic(err)
    }
}
// SetV101903 sets the v10_1903 property value. Windows 10 1903 or later.
func (m *WindowsMinimumOperatingSystem) SetV101903(value *bool)() {
    err := m.GetBackingStore().Set("v10_1903", value)
    if err != nil {
        panic(err)
    }
}
// SetV101909 sets the v10_1909 property value. Windows 10 1909 or later.
func (m *WindowsMinimumOperatingSystem) SetV101909(value *bool)() {
    err := m.GetBackingStore().Set("v10_1909", value)
    if err != nil {
        panic(err)
    }
}
// SetV102004 sets the v10_2004 property value. Windows 10 2004 or later.
func (m *WindowsMinimumOperatingSystem) SetV102004(value *bool)() {
    err := m.GetBackingStore().Set("v10_2004", value)
    if err != nil {
        panic(err)
    }
}
// SetV1021H1 sets the v10_21H1 property value. Windows 10 21H1 or later.
func (m *WindowsMinimumOperatingSystem) SetV1021H1(value *bool)() {
    err := m.GetBackingStore().Set("v10_21H1", value)
    if err != nil {
        panic(err)
    }
}
// SetV102H20 sets the v10_2H20 property value. Windows 10 2H20 or later.
func (m *WindowsMinimumOperatingSystem) SetV102H20(value *bool)() {
    err := m.GetBackingStore().Set("v10_2H20", value)
    if err != nil {
        panic(err)
    }
}
// SetV80 sets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV80(value *bool)() {
    err := m.GetBackingStore().Set("v8_0", value)
    if err != nil {
        panic(err)
    }
}
// SetV81 sets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) SetV81(value *bool)() {
    err := m.GetBackingStore().Set("v8_1", value)
    if err != nil {
        panic(err)
    }
}
// WindowsMinimumOperatingSystemable 
type WindowsMinimumOperatingSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetV100()(*bool)
    GetV101607()(*bool)
    GetV101703()(*bool)
    GetV101709()(*bool)
    GetV101803()(*bool)
    GetV101809()(*bool)
    GetV101903()(*bool)
    GetV101909()(*bool)
    GetV102004()(*bool)
    GetV1021H1()(*bool)
    GetV102H20()(*bool)
    GetV80()(*bool)
    GetV81()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetV100(value *bool)()
    SetV101607(value *bool)()
    SetV101703(value *bool)()
    SetV101709(value *bool)()
    SetV101803(value *bool)()
    SetV101809(value *bool)()
    SetV101903(value *bool)()
    SetV101909(value *bool)()
    SetV102004(value *bool)()
    SetV1021H1(value *bool)()
    SetV102H20(value *bool)()
    SetV80(value *bool)()
    SetV81(value *bool)()
}
