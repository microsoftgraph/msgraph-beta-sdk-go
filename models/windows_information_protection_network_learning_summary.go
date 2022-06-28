package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionNetworkLearningSummary windows Information Protection Network learning Summary entity.
type WindowsInformationProtectionNetworkLearningSummary struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device Count
    deviceCount *int32
    // Website url
    url *string
}
// NewWindowsInformationProtectionNetworkLearningSummary instantiates a new windowsInformationProtectionNetworkLearningSummary and sets the default values.
func NewWindowsInformationProtectionNetworkLearningSummary()(*WindowsInformationProtectionNetworkLearningSummary) {
    m := &WindowsInformationProtectionNetworkLearningSummary{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionNetworkLearningSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionNetworkLearningSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceCount gets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionNetworkLearningSummary) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionNetworkLearningSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. Website url
func (m *WindowsInformationProtectionNetworkLearningSummary) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionNetworkLearningSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
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
func (m *WindowsInformationProtectionNetworkLearningSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceCount sets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionNetworkLearningSummary) SetDeviceCount(value *int32)() {
    if m != nil {
        m.deviceCount = value
    }
}
// SetUrl sets the url property value. Website url
func (m *WindowsInformationProtectionNetworkLearningSummary) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
