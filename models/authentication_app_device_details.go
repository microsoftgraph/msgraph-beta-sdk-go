package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationAppDeviceDetails 
type AuthenticationAppDeviceDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appVersion property
    appVersion *string
    // The clientApp property
    clientApp *string
    // The deviceId property
    deviceId *string
    // The OdataType property
    odataType *string
    // The operatingSystem property
    operatingSystem *string
}
// NewAuthenticationAppDeviceDetails instantiates a new authenticationAppDeviceDetails and sets the default values.
func NewAuthenticationAppDeviceDetails()(*AuthenticationAppDeviceDetails) {
    m := &AuthenticationAppDeviceDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.authenticationAppDeviceDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAppDeviceDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationAppDeviceDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppVersion gets the appVersion property value. The appVersion property
func (m *AuthenticationAppDeviceDetails) GetAppVersion()(*string) {
    return m.appVersion
}
// GetClientApp gets the clientApp property value. The clientApp property
func (m *AuthenticationAppDeviceDetails) GetClientApp()(*string) {
    return m.clientApp
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *AuthenticationAppDeviceDetails) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationAppDeviceDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppVersion)
    res["clientApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientApp)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["operatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOperatingSystem)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationAppDeviceDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystem gets the operatingSystem property value. The operatingSystem property
func (m *AuthenticationAppDeviceDetails) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// Serialize serializes information the current object
func (m *AuthenticationAppDeviceDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientApp", m.GetClientApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
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
    {
        err := writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
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
func (m *AuthenticationAppDeviceDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppVersion sets the appVersion property value. The appVersion property
func (m *AuthenticationAppDeviceDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetClientApp sets the clientApp property value. The clientApp property
func (m *AuthenticationAppDeviceDetails) SetClientApp(value *string)() {
    m.clientApp = value
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *AuthenticationAppDeviceDetails) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAppDeviceDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystem sets the operatingSystem property value. The operatingSystem property
func (m *AuthenticationAppDeviceDetails) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
