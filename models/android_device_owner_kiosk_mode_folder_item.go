package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeFolderItem 
type AndroidDeviceOwnerKioskModeFolderItem struct {
    AndroidDeviceOwnerKioskModeHomeScreenItem
    // The type property
    type_escaped *string
}
// NewAndroidDeviceOwnerKioskModeFolderItem instantiates a new AndroidDeviceOwnerKioskModeFolderItem and sets the default values.
func NewAndroidDeviceOwnerKioskModeFolderItem()(*AndroidDeviceOwnerKioskModeFolderItem) {
    m := &AndroidDeviceOwnerKioskModeFolderItem{
        AndroidDeviceOwnerKioskModeHomeScreenItem: *NewAndroidDeviceOwnerKioskModeHomeScreenItem(),
    }
    return m
}
// CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidDeviceOwnerKioskModeApp":
                        return NewAndroidDeviceOwnerKioskModeApp(), nil
                    case "#microsoft.graph.androidDeviceOwnerKioskModeWeblink":
                        return NewAndroidDeviceOwnerKioskModeWeblink(), nil
                }
            }
        }
    }
    return NewAndroidDeviceOwnerKioskModeFolderItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeFolderItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeHomeScreenItem.GetFieldDeserializers()
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
// GetType gets the type property value. The type property
func (m *AndroidDeviceOwnerKioskModeFolderItem) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeFolderItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. The type property
func (m *AndroidDeviceOwnerKioskModeFolderItem) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
