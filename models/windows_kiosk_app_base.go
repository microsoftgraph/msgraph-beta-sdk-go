package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsKioskAppBase the base class for a type of apps
type WindowsKioskAppBase struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsKioskAppBase instantiates a new WindowsKioskAppBase and sets the default values.
func NewWindowsKioskAppBase()(*WindowsKioskAppBase) {
    m := &WindowsKioskAppBase{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsKioskAppBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskAppBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsKioskDesktopApp":
                        return NewWindowsKioskDesktopApp(), nil
                    case "#microsoft.graph.windowsKioskUWPApp":
                        return NewWindowsKioskUWPApp(), nil
                    case "#microsoft.graph.windowsKioskWin32App":
                        return NewWindowsKioskWin32App(), nil
                }
            }
        }
    }
    return NewWindowsKioskAppBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskAppBase) GetAdditionalData()(map[string]any) {
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
// GetAppType gets the appType property value. The type of Windows kiosk app.
func (m *WindowsKioskAppBase) GetAppType()(*WindowsKioskAppType) {
    val, err := m.GetBackingStore().Get("appType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsKioskAppType)
    }
    return nil
}
// GetAutoLaunch gets the autoLaunch property value. Allow the app to be auto-launched in multi-app kiosk mode
func (m *WindowsKioskAppBase) GetAutoLaunch()(*bool) {
    val, err := m.GetBackingStore().Get("autoLaunch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WindowsKioskAppBase) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskAppBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsKioskAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppType(val.(*WindowsKioskAppType))
        }
        return nil
    }
    res["autoLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoLaunch(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
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
    res["startLayoutTileSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAppStartLayoutTileSize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartLayoutTileSize(val.(*WindowsAppStartLayoutTileSize))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Represents the friendly name of an app
func (m *WindowsKioskAppBase) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsKioskAppBase) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartLayoutTileSize gets the startLayoutTileSize property value. The tile size of Windows app in the start layout.
func (m *WindowsKioskAppBase) GetStartLayoutTileSize()(*WindowsAppStartLayoutTileSize) {
    val, err := m.GetBackingStore().Get("startLayoutTileSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAppStartLayoutTileSize)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskAppBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppType() != nil {
        cast := (*m.GetAppType()).String()
        err := writer.WriteStringValue("appType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoLaunch", m.GetAutoLaunch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStartLayoutTileSize() != nil {
        cast := (*m.GetStartLayoutTileSize()).String()
        err := writer.WriteStringValue("startLayoutTileSize", &cast)
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
func (m *WindowsKioskAppBase) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppType sets the appType property value. The type of Windows kiosk app.
func (m *WindowsKioskAppBase) SetAppType(value *WindowsKioskAppType)() {
    err := m.GetBackingStore().Set("appType", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoLaunch sets the autoLaunch property value. Allow the app to be auto-launched in multi-app kiosk mode
func (m *WindowsKioskAppBase) SetAutoLaunch(value *bool)() {
    err := m.GetBackingStore().Set("autoLaunch", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WindowsKioskAppBase) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetName sets the name property value. Represents the friendly name of an app
func (m *WindowsKioskAppBase) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsKioskAppBase) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStartLayoutTileSize sets the startLayoutTileSize property value. The tile size of Windows app in the start layout.
func (m *WindowsKioskAppBase) SetStartLayoutTileSize(value *WindowsAppStartLayoutTileSize)() {
    err := m.GetBackingStore().Set("startLayoutTileSize", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskAppBaseable 
type WindowsKioskAppBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppType()(*WindowsKioskAppType)
    GetAutoLaunch()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetName()(*string)
    GetOdataType()(*string)
    GetStartLayoutTileSize()(*WindowsAppStartLayoutTileSize)
    SetAppType(value *WindowsKioskAppType)()
    SetAutoLaunch(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetStartLayoutTileSize(value *WindowsAppStartLayoutTileSize)()
}
