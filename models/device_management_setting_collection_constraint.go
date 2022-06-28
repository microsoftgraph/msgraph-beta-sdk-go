package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingCollectionConstraint 
type DeviceManagementSettingCollectionConstraint struct {
    DeviceManagementConstraint
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The maximum number of elements in the collection
    maximumLength *int32
    // The minimum number of elements in the collection
    minimumLength *int32
}
// NewDeviceManagementSettingCollectionConstraint instantiates a new DeviceManagementSettingCollectionConstraint and sets the default values.
func NewDeviceManagementSettingCollectionConstraint()(*DeviceManagementSettingCollectionConstraint) {
    m := &DeviceManagementSettingCollectionConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingCollectionConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingCollectionConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingCollectionConstraint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingCollectionConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingCollectionConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["maximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumLength(val)
        }
        return nil
    }
    res["minimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLength(val)
        }
        return nil
    }
    return res
}
// GetMaximumLength gets the maximumLength property value. The maximum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) GetMaximumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumLength
    }
}
// GetMinimumLength gets the minimumLength property value. The minimum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) GetMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumLength
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingCollectionConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumLength", m.GetMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumLength", m.GetMinimumLength())
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
func (m *DeviceManagementSettingCollectionConstraint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumLength sets the maximumLength property value. The maximum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) SetMaximumLength(value *int32)() {
    if m != nil {
        m.maximumLength = value
    }
}
// SetMinimumLength sets the minimumLength property value. The minimum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) SetMinimumLength(value *int32)() {
    if m != nil {
        m.minimumLength = value
    }
}
