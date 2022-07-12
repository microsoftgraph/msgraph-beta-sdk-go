package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskAppBase the base class for a type of apps
type WindowsKioskAppBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of Windows kiosk app.
    appType *WindowsKioskAppType
    // Allow the app to be auto-launched in multi-app kiosk mode
    autoLaunch *bool
    // Represents the friendly name of an app
    name *string
    // The tile size of Windows app in the start layout.
    startLayoutTileSize *WindowsAppStartLayoutTileSize
    // The type property
    type_escaped *string
}
// NewWindowsKioskAppBase instantiates a new windowsKioskAppBase and sets the default values.
func NewWindowsKioskAppBase()(*WindowsKioskAppBase) {
    m := &WindowsKioskAppBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    typeValue := "#microsoft.graph.windowsKioskAppBase";
    m.SetType(&typeValue);
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
                mappingStr := *mappingValue
                switch mappingStr {
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
func (m *WindowsKioskAppBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppType gets the appType property value. The type of Windows kiosk app.
func (m *WindowsKioskAppBase) GetAppType()(*WindowsKioskAppType) {
    if m == nil {
        return nil
    } else {
        return m.appType
    }
}
// GetAutoLaunch gets the autoLaunch property value. Allow the app to be auto-launched in multi-app kiosk mode
func (m *WindowsKioskAppBase) GetAutoLaunch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoLaunch
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Represents the friendly name of an app
func (m *WindowsKioskAppBase) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetStartLayoutTileSize gets the startLayoutTileSize property value. The tile size of Windows app in the start layout.
func (m *WindowsKioskAppBase) GetStartLayoutTileSize()(*WindowsAppStartLayoutTileSize) {
    if m == nil {
        return nil
    } else {
        return m.startLayoutTileSize
    }
}
// GetType gets the type property value. The type property
func (m *WindowsKioskAppBase) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
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
    if m.GetStartLayoutTileSize() != nil {
        cast := (*m.GetStartLayoutTileSize()).String()
        err := writer.WriteStringValue("startLayoutTileSize", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *WindowsKioskAppBase) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppType sets the appType property value. The type of Windows kiosk app.
func (m *WindowsKioskAppBase) SetAppType(value *WindowsKioskAppType)() {
    if m != nil {
        m.appType = value
    }
}
// SetAutoLaunch sets the autoLaunch property value. Allow the app to be auto-launched in multi-app kiosk mode
func (m *WindowsKioskAppBase) SetAutoLaunch(value *bool)() {
    if m != nil {
        m.autoLaunch = value
    }
}
// SetName sets the name property value. Represents the friendly name of an app
func (m *WindowsKioskAppBase) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetStartLayoutTileSize sets the startLayoutTileSize property value. The tile size of Windows app in the start layout.
func (m *WindowsKioskAppBase) SetStartLayoutTileSize(value *WindowsAppStartLayoutTileSize)() {
    if m != nil {
        m.startLayoutTileSize = value
    }
}
// SetType sets the type property value. The type property
func (m *WindowsKioskAppBase) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
