package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ConditionalAccessWhatIfConditions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewConditionalAccessWhatIfConditions instantiates a new ConditionalAccessWhatIfConditions and sets the default values.
func NewConditionalAccessWhatIfConditions()(*ConditionalAccessWhatIfConditions) {
    m := &ConditionalAccessWhatIfConditions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessWhatIfConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessWhatIfConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessWhatIfConditions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessWhatIfConditions) GetAdditionalData()(map[string]any) {
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
// GetAuthenticationFlow gets the authenticationFlow property value. The authenticationFlow property
// returns a AuthenticationFlowable when successful
func (m *ConditionalAccessWhatIfConditions) GetAuthenticationFlow()(AuthenticationFlowable) {
    val, err := m.GetBackingStore().Get("authenticationFlow")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationFlowable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ConditionalAccessWhatIfConditions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClientAppType gets the clientAppType property value. The clientAppType property
// returns a *ConditionalAccessClientApp when successful
func (m *ConditionalAccessWhatIfConditions) GetClientAppType()(*ConditionalAccessClientApp) {
    val, err := m.GetBackingStore().Get("clientAppType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConditionalAccessClientApp)
    }
    return nil
}
// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *ConditionalAccessWhatIfConditions) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceInfo gets the deviceInfo property value. The deviceInfo property
// returns a DeviceInfoable when successful
func (m *ConditionalAccessWhatIfConditions) GetDeviceInfo()(DeviceInfoable) {
    val, err := m.GetBackingStore().Get("deviceInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceInfoable)
    }
    return nil
}
// GetDevicePlatform gets the devicePlatform property value. The devicePlatform property
// returns a *ConditionalAccessDevicePlatform when successful
func (m *ConditionalAccessWhatIfConditions) GetDevicePlatform()(*ConditionalAccessDevicePlatform) {
    val, err := m.GetBackingStore().Get("devicePlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConditionalAccessDevicePlatform)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessWhatIfConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationFlow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlow(val.(AuthenticationFlowable))
        }
        return nil
    }
    res["clientAppType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessClientApp)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppType(val.(*ConditionalAccessClientApp))
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["deviceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInfo(val.(DeviceInfoable))
        }
        return nil
    }
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessDevicePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val.(*ConditionalAccessDevicePlatform))
        }
        return nil
    }
    res["insiderRiskLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInsiderRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsiderRiskLevel(val.(*InsiderRiskLevel))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
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
    res["servicePrincipalRiskLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalRiskLevel(val.(*RiskLevel))
        }
        return nil
    }
    res["signInRiskLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInRiskLevel(val.(*RiskLevel))
        }
        return nil
    }
    res["userRiskLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRiskLevel(val.(*RiskLevel))
        }
        return nil
    }
    return res
}
// GetInsiderRiskLevel gets the insiderRiskLevel property value. The insiderRiskLevel property
// returns a *InsiderRiskLevel when successful
func (m *ConditionalAccessWhatIfConditions) GetInsiderRiskLevel()(*InsiderRiskLevel) {
    val, err := m.GetBackingStore().Get("insiderRiskLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InsiderRiskLevel)
    }
    return nil
}
// GetIpAddress gets the ipAddress property value. The ipAddress property
// returns a *string when successful
func (m *ConditionalAccessWhatIfConditions) GetIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("ipAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessWhatIfConditions) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePrincipalRiskLevel gets the servicePrincipalRiskLevel property value. The servicePrincipalRiskLevel property
// returns a *RiskLevel when successful
func (m *ConditionalAccessWhatIfConditions) GetServicePrincipalRiskLevel()(*RiskLevel) {
    val, err := m.GetBackingStore().Get("servicePrincipalRiskLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskLevel)
    }
    return nil
}
// GetSignInRiskLevel gets the signInRiskLevel property value. The signInRiskLevel property
// returns a *RiskLevel when successful
func (m *ConditionalAccessWhatIfConditions) GetSignInRiskLevel()(*RiskLevel) {
    val, err := m.GetBackingStore().Get("signInRiskLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskLevel)
    }
    return nil
}
// GetUserRiskLevel gets the userRiskLevel property value. The userRiskLevel property
// returns a *RiskLevel when successful
func (m *ConditionalAccessWhatIfConditions) GetUserRiskLevel()(*RiskLevel) {
    val, err := m.GetBackingStore().Get("userRiskLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RiskLevel)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessWhatIfConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationFlow", m.GetAuthenticationFlow())
        if err != nil {
            return err
        }
    }
    if m.GetClientAppType() != nil {
        cast := (*m.GetClientAppType()).String()
        err := writer.WriteStringValue("clientAppType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceInfo", m.GetDeviceInfo())
        if err != nil {
            return err
        }
    }
    if m.GetDevicePlatform() != nil {
        cast := (*m.GetDevicePlatform()).String()
        err := writer.WriteStringValue("devicePlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInsiderRiskLevel() != nil {
        cast := (*m.GetInsiderRiskLevel()).String()
        err := writer.WriteStringValue("insiderRiskLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
    if m.GetServicePrincipalRiskLevel() != nil {
        cast := (*m.GetServicePrincipalRiskLevel()).String()
        err := writer.WriteStringValue("servicePrincipalRiskLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignInRiskLevel() != nil {
        cast := (*m.GetSignInRiskLevel()).String()
        err := writer.WriteStringValue("signInRiskLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserRiskLevel() != nil {
        cast := (*m.GetUserRiskLevel()).String()
        err := writer.WriteStringValue("userRiskLevel", &cast)
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
func (m *ConditionalAccessWhatIfConditions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationFlow sets the authenticationFlow property value. The authenticationFlow property
func (m *ConditionalAccessWhatIfConditions) SetAuthenticationFlow(value AuthenticationFlowable)() {
    err := m.GetBackingStore().Set("authenticationFlow", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ConditionalAccessWhatIfConditions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClientAppType sets the clientAppType property value. The clientAppType property
func (m *ConditionalAccessWhatIfConditions) SetClientAppType(value *ConditionalAccessClientApp)() {
    err := m.GetBackingStore().Set("clientAppType", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The country property
func (m *ConditionalAccessWhatIfConditions) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceInfo sets the deviceInfo property value. The deviceInfo property
func (m *ConditionalAccessWhatIfConditions) SetDeviceInfo(value DeviceInfoable)() {
    err := m.GetBackingStore().Set("deviceInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePlatform sets the devicePlatform property value. The devicePlatform property
func (m *ConditionalAccessWhatIfConditions) SetDevicePlatform(value *ConditionalAccessDevicePlatform)() {
    err := m.GetBackingStore().Set("devicePlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetInsiderRiskLevel sets the insiderRiskLevel property value. The insiderRiskLevel property
func (m *ConditionalAccessWhatIfConditions) SetInsiderRiskLevel(value *InsiderRiskLevel)() {
    err := m.GetBackingStore().Set("insiderRiskLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddress sets the ipAddress property value. The ipAddress property
func (m *ConditionalAccessWhatIfConditions) SetIpAddress(value *string)() {
    err := m.GetBackingStore().Set("ipAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessWhatIfConditions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalRiskLevel sets the servicePrincipalRiskLevel property value. The servicePrincipalRiskLevel property
func (m *ConditionalAccessWhatIfConditions) SetServicePrincipalRiskLevel(value *RiskLevel)() {
    err := m.GetBackingStore().Set("servicePrincipalRiskLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInRiskLevel sets the signInRiskLevel property value. The signInRiskLevel property
func (m *ConditionalAccessWhatIfConditions) SetSignInRiskLevel(value *RiskLevel)() {
    err := m.GetBackingStore().Set("signInRiskLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRiskLevel sets the userRiskLevel property value. The userRiskLevel property
func (m *ConditionalAccessWhatIfConditions) SetUserRiskLevel(value *RiskLevel)() {
    err := m.GetBackingStore().Set("userRiskLevel", value)
    if err != nil {
        panic(err)
    }
}
type ConditionalAccessWhatIfConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationFlow()(AuthenticationFlowable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClientAppType()(*ConditionalAccessClientApp)
    GetCountry()(*string)
    GetDeviceInfo()(DeviceInfoable)
    GetDevicePlatform()(*ConditionalAccessDevicePlatform)
    GetInsiderRiskLevel()(*InsiderRiskLevel)
    GetIpAddress()(*string)
    GetOdataType()(*string)
    GetServicePrincipalRiskLevel()(*RiskLevel)
    GetSignInRiskLevel()(*RiskLevel)
    GetUserRiskLevel()(*RiskLevel)
    SetAuthenticationFlow(value AuthenticationFlowable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClientAppType(value *ConditionalAccessClientApp)()
    SetCountry(value *string)()
    SetDeviceInfo(value DeviceInfoable)()
    SetDevicePlatform(value *ConditionalAccessDevicePlatform)()
    SetInsiderRiskLevel(value *InsiderRiskLevel)()
    SetIpAddress(value *string)()
    SetOdataType(value *string)()
    SetServicePrincipalRiskLevel(value *RiskLevel)()
    SetSignInRiskLevel(value *RiskLevel)()
    SetUserRiskLevel(value *RiskLevel)()
}
