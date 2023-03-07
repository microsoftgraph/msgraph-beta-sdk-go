package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MobileAppIntentAndStateDetail mobile App Intent and Install State for a given device.
type MobileAppIntentAndStateDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMobileAppIntentAndStateDetail instantiates a new mobileAppIntentAndStateDetail and sets the default values.
func NewMobileAppIntentAndStateDetail()(*MobileAppIntentAndStateDetail) {
    m := &MobileAppIntentAndStateDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppIntentAndStateDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppIntentAndStateDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppIntentAndStateDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppIntentAndStateDetail) GetAdditionalData()(map[string]any) {
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
// GetApplicationId gets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) GetApplicationId()(*string) {
    val, err := m.GetBackingStore().Get("applicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MobileAppIntentAndStateDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) GetDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("displayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndStateDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["installState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["mobileAppIntent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIntent(val.(*MobileAppIntent))
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
    res["supportedDeviceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppSupportedDeviceTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppSupportedDeviceTypeable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppSupportedDeviceTypeable)
            }
            m.SetSupportedDeviceTypes(res)
        }
        return nil
    }
    return res
}
// GetInstallState gets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppIntentAndStateDetail) GetInstallState()(*ResultantAppState) {
    val, err := m.GetBackingStore().Get("installState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ResultantAppState)
    }
    return nil
}
// GetMobileAppIntent gets the mobileAppIntent property value. Indicates the status of the mobile app on the device.
func (m *MobileAppIntentAndStateDetail) GetMobileAppIntent()(*MobileAppIntent) {
    val, err := m.GetBackingStore().Get("mobileAppIntent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppIntent)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MobileAppIntentAndStateDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportedDeviceTypes gets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) GetSupportedDeviceTypes()([]MobileAppSupportedDeviceTypeable) {
    val, err := m.GetBackingStore().Get("supportedDeviceTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppSupportedDeviceTypeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppIntentAndStateDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := (*m.GetInstallState()).String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntent() != nil {
        cast := (*m.GetMobileAppIntent()).String()
        err := writer.WriteStringValue("mobileAppIntent", &cast)
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
    if m.GetSupportedDeviceTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSupportedDeviceTypes()))
        for i, v := range m.GetSupportedDeviceTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("supportedDeviceTypes", cast)
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
func (m *MobileAppIntentAndStateDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationId sets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) SetApplicationId(value *string)() {
    err := m.GetBackingStore().Set("applicationId", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MobileAppIntentAndStateDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) SetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("displayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallState sets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppIntentAndStateDetail) SetInstallState(value *ResultantAppState)() {
    err := m.GetBackingStore().Set("installState", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppIntent sets the mobileAppIntent property value. Indicates the status of the mobile app on the device.
func (m *MobileAppIntentAndStateDetail) SetMobileAppIntent(value *MobileAppIntent)() {
    err := m.GetBackingStore().Set("mobileAppIntent", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppIntentAndStateDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedDeviceTypes sets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) SetSupportedDeviceTypes(value []MobileAppSupportedDeviceTypeable)() {
    err := m.GetBackingStore().Set("supportedDeviceTypes", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppIntentAndStateDetailable 
type MobileAppIntentAndStateDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetInstallState()(*ResultantAppState)
    GetMobileAppIntent()(*MobileAppIntent)
    GetOdataType()(*string)
    GetSupportedDeviceTypes()([]MobileAppSupportedDeviceTypeable)
    SetApplicationId(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetInstallState(value *ResultantAppState)()
    SetMobileAppIntent(value *MobileAppIntent)()
    SetOdataType(value *string)()
    SetSupportedDeviceTypes(value []MobileAppSupportedDeviceTypeable)()
}
