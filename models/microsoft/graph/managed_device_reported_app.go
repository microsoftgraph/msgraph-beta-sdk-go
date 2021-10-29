package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceReportedApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The application or bundle identifier of the application
    appId *string;
}
// Instantiates a new managedDeviceReportedApp and sets the default values.
func NewManagedDeviceReportedApp()(*ManagedDeviceReportedApp) {
    m := &ManagedDeviceReportedApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceReportedApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appId property value. The application or bundle identifier of the application
func (m *ManagedDeviceReportedApp) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceReportedApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceReportedApp) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceReportedApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagedDeviceReportedApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appId property value. The application or bundle identifier of the application
// Parameters:
//  - value : Value to set for the appId property.
func (m *ManagedDeviceReportedApp) SetAppId(value *string)() {
    m.appId = value
}
