package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskActiveDirectoryGroup 
type WindowsKioskActiveDirectoryGroup struct {
    WindowsKioskUser
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the AD group that will be locked to this kiosk configuration
    groupName *string
}
// NewWindowsKioskActiveDirectoryGroup instantiates a new WindowsKioskActiveDirectoryGroup and sets the default values.
func NewWindowsKioskActiveDirectoryGroup()(*WindowsKioskActiveDirectoryGroup) {
    m := &WindowsKioskActiveDirectoryGroup{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsKioskActiveDirectoryGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskActiveDirectoryGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskActiveDirectoryGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskActiveDirectoryGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskActiveDirectoryGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The name of the AD group that will be locked to this kiosk configuration
func (m *WindowsKioskActiveDirectoryGroup) GetGroupName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupName
    }
}
// Serialize serializes information the current object
func (m *WindowsKioskActiveDirectoryGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskActiveDirectoryGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupName sets the groupName property value. The name of the AD group that will be locked to this kiosk configuration
func (m *WindowsKioskActiveDirectoryGroup) SetGroupName(value *string)() {
    if m != nil {
        m.groupName = value
    }
}
