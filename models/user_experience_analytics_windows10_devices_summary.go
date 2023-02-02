package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWindows10DevicesSummary the user experience analytics work from anywhere Windows 10 devices summary.
type UserExperienceAnalyticsWindows10DevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The count of Windows 10 devices that have unsupported OS versions.
    unsupportedOSversionDeviceCount *int32
}
// NewUserExperienceAnalyticsWindows10DevicesSummary instantiates a new userExperienceAnalyticsWindows10DevicesSummary and sets the default values.
func NewUserExperienceAnalyticsWindows10DevicesSummary()(*UserExperienceAnalyticsWindows10DevicesSummary) {
    m := &UserExperienceAnalyticsWindows10DevicesSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsWindows10DevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWindows10DevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWindows10DevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["unsupportedOSversionDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsupportedOSversionDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetUnsupportedOSversionDeviceCount gets the unsupportedOSversionDeviceCount property value. The count of Windows 10 devices that have unsupported OS versions.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetUnsupportedOSversionDeviceCount()(*int32) {
    return m.unsupportedOSversionDeviceCount
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWindows10DevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unsupportedOSversionDeviceCount", m.GetUnsupportedOSversionDeviceCount())
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
func (m *UserExperienceAnalyticsWindows10DevicesSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsWindows10DevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUnsupportedOSversionDeviceCount sets the unsupportedOSversionDeviceCount property value. The count of Windows 10 devices that have unsupported OS versions.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) SetUnsupportedOSversionDeviceCount(value *int32)() {
    m.unsupportedOSversionDeviceCount = value
}
