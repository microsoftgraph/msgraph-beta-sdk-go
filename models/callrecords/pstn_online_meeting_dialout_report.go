package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PstnOnlineMeetingDialoutReport 
type PstnOnlineMeetingDialoutReport struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPstnOnlineMeetingDialoutReport instantiates a new pstnOnlineMeetingDialoutReport and sets the default values.
func NewPstnOnlineMeetingDialoutReport()(*PstnOnlineMeetingDialoutReport) {
    m := &PstnOnlineMeetingDialoutReport{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePstnOnlineMeetingDialoutReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePstnOnlineMeetingDialoutReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPstnOnlineMeetingDialoutReport(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PstnOnlineMeetingDialoutReport) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PstnOnlineMeetingDialoutReport) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCurrency gets the currency property value. Currency used to calculate the cost of the call. For details, see ISO 4217.
func (m *PstnOnlineMeetingDialoutReport) GetCurrency()(*string) {
    val, err := m.GetBackingStore().Get("currency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationContext gets the destinationContext property value. Indicates whether the call was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
func (m *PstnOnlineMeetingDialoutReport) GetDestinationContext()(*string) {
    val, err := m.GetBackingStore().Get("destinationContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PstnOnlineMeetingDialoutReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val)
        }
        return nil
    }
    res["destinationContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationContext(val)
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
    res["totalCallCharge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCallCharge(val)
        }
        return nil
    }
    res["totalCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCalls(val)
        }
        return nil
    }
    res["totalCallSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCallSeconds(val)
        }
        return nil
    }
    res["usageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageLocation(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PstnOnlineMeetingDialoutReport) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalCallCharge gets the totalCallCharge property value. Total costs of all the calls within the selected time range, including call charges and connection fees.
func (m *PstnOnlineMeetingDialoutReport) GetTotalCallCharge()(*float64) {
    val, err := m.GetBackingStore().Get("totalCallCharge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetTotalCalls gets the totalCalls property value. Total number of dial-out calls within the selected time range.
func (m *PstnOnlineMeetingDialoutReport) GetTotalCalls()(*int32) {
    val, err := m.GetBackingStore().Get("totalCalls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalCallSeconds gets the totalCallSeconds property value. Total duration of all the calls within the selected time range, in seconds.
func (m *PstnOnlineMeetingDialoutReport) GetTotalCallSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("totalCallSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUsageLocation gets the usageLocation property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *PstnOnlineMeetingDialoutReport) GetUsageLocation()(*string) {
    val, err := m.GetBackingStore().Get("usageLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user.
func (m *PstnOnlineMeetingDialoutReport) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The unique identifier (GUID) of the user in Azure Active Directory.
func (m *PstnOnlineMeetingDialoutReport) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (sign-in name) in Azure Active Directory. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *PstnOnlineMeetingDialoutReport) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PstnOnlineMeetingDialoutReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationContext", m.GetDestinationContext())
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
        err := writer.WriteFloat64Value("totalCallCharge", m.GetTotalCallCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCalls", m.GetTotalCalls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCallSeconds", m.GetTotalCallSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usageLocation", m.GetUsageLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *PstnOnlineMeetingDialoutReport) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PstnOnlineMeetingDialoutReport) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCurrency sets the currency property value. Currency used to calculate the cost of the call. For details, see ISO 4217.
func (m *PstnOnlineMeetingDialoutReport) SetCurrency(value *string)() {
    err := m.GetBackingStore().Set("currency", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationContext sets the destinationContext property value. Indicates whether the call was Domestic (within a country or region) or International (outside a country or region) based on the user's location.
func (m *PstnOnlineMeetingDialoutReport) SetDestinationContext(value *string)() {
    err := m.GetBackingStore().Set("destinationContext", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PstnOnlineMeetingDialoutReport) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCallCharge sets the totalCallCharge property value. Total costs of all the calls within the selected time range, including call charges and connection fees.
func (m *PstnOnlineMeetingDialoutReport) SetTotalCallCharge(value *float64)() {
    err := m.GetBackingStore().Set("totalCallCharge", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCalls sets the totalCalls property value. Total number of dial-out calls within the selected time range.
func (m *PstnOnlineMeetingDialoutReport) SetTotalCalls(value *int32)() {
    err := m.GetBackingStore().Set("totalCalls", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCallSeconds sets the totalCallSeconds property value. Total duration of all the calls within the selected time range, in seconds.
func (m *PstnOnlineMeetingDialoutReport) SetTotalCallSeconds(value *int32)() {
    err := m.GetBackingStore().Set("totalCallSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetUsageLocation sets the usageLocation property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *PstnOnlineMeetingDialoutReport) SetUsageLocation(value *string)() {
    err := m.GetBackingStore().Set("usageLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user.
func (m *PstnOnlineMeetingDialoutReport) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The unique identifier (GUID) of the user in Azure Active Directory.
func (m *PstnOnlineMeetingDialoutReport) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (sign-in name) in Azure Active Directory. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *PstnOnlineMeetingDialoutReport) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// PstnOnlineMeetingDialoutReportable 
type PstnOnlineMeetingDialoutReportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCurrency()(*string)
    GetDestinationContext()(*string)
    GetOdataType()(*string)
    GetTotalCallCharge()(*float64)
    GetTotalCalls()(*int32)
    GetTotalCallSeconds()(*int32)
    GetUsageLocation()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCurrency(value *string)()
    SetDestinationContext(value *string)()
    SetOdataType(value *string)()
    SetTotalCallCharge(value *float64)()
    SetTotalCalls(value *int32)()
    SetTotalCallSeconds(value *int32)()
    SetUsageLocation(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
