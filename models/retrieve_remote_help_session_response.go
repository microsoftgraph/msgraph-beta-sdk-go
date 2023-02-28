package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// RetrieveRemoteHelpSessionResponse remote help - response we provide back to the helper on retrieve session API call
type RetrieveRemoteHelpSessionResponse struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRetrieveRemoteHelpSessionResponse instantiates a new retrieveRemoteHelpSessionResponse and sets the default values.
func NewRetrieveRemoteHelpSessionResponse()(*RetrieveRemoteHelpSessionResponse) {
    m := &RetrieveRemoteHelpSessionResponse{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRetrieveRemoteHelpSessionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetrieveRemoteHelpSessionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetrieveRemoteHelpSessionResponse(), nil
}
// GetAcsGroupId gets the acsGroupId property value. ACS Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsGroupId()(*string) {
    val, err := m.GetBackingStore().Get("acsGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAcsHelperUserId gets the acsHelperUserId property value. Helper ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsHelperUserId()(*string) {
    val, err := m.GetBackingStore().Get("acsHelperUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAcsHelperUserToken gets the acsHelperUserToken property value. Helper ACS User Token
func (m *RetrieveRemoteHelpSessionResponse) GetAcsHelperUserToken()(*string) {
    val, err := m.GetBackingStore().Get("acsHelperUserToken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAcsSharerUserId gets the acsSharerUserId property value. Sharer ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsSharerUserId()(*string) {
    val, err := m.GetBackingStore().Get("acsSharerUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetrieveRemoteHelpSessionResponse) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *RetrieveRemoteHelpSessionResponse) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceName gets the deviceName property value. Android Device Name
func (m *RetrieveRemoteHelpSessionResponse) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetrieveRemoteHelpSessionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acsGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcsGroupId(val)
        }
        return nil
    }
    res["acsHelperUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcsHelperUserId(val)
        }
        return nil
    }
    res["acsHelperUserToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcsHelperUserToken(val)
        }
        return nil
    }
    res["acsSharerUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcsSharerUserId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
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
    res["pubSubGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPubSubGroupId(val)
        }
        return nil
    }
    res["pubSubHelperAccessUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPubSubHelperAccessUri(val)
        }
        return nil
    }
    res["sessionExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionExpirationDateTime(val)
        }
        return nil
    }
    res["sessionKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionKey(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetrieveRemoteHelpSessionResponse) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPubSubGroupId gets the pubSubGroupId property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetPubSubGroupId()(*string) {
    val, err := m.GetBackingStore().Get("pubSubGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPubSubHelperAccessUri gets the pubSubHelperAccessUri property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetPubSubHelperAccessUri()(*string) {
    val, err := m.GetBackingStore().Get("pubSubHelperAccessUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSessionExpirationDateTime gets the sessionExpirationDateTime property value. Azure Pubsub Session Expiration Date Time.
func (m *RetrieveRemoteHelpSessionResponse) GetSessionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("sessionExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSessionKey gets the sessionKey property value. The unique identifier for a session
func (m *RetrieveRemoteHelpSessionResponse) GetSessionKey()(*string) {
    val, err := m.GetBackingStore().Get("sessionKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RetrieveRemoteHelpSessionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("acsGroupId", m.GetAcsGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("acsHelperUserId", m.GetAcsHelperUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("acsHelperUserToken", m.GetAcsHelperUserToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("acsSharerUserId", m.GetAcsSharerUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
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
        err := writer.WriteStringValue("pubSubGroupId", m.GetPubSubGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pubSubHelperAccessUri", m.GetPubSubHelperAccessUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sessionExpirationDateTime", m.GetSessionExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionKey", m.GetSessionKey())
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
// SetAcsGroupId sets the acsGroupId property value. ACS Group Id
func (m *RetrieveRemoteHelpSessionResponse) SetAcsGroupId(value *string)() {
    err := m.GetBackingStore().Set("acsGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetAcsHelperUserId sets the acsHelperUserId property value. Helper ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) SetAcsHelperUserId(value *string)() {
    err := m.GetBackingStore().Set("acsHelperUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetAcsHelperUserToken sets the acsHelperUserToken property value. Helper ACS User Token
func (m *RetrieveRemoteHelpSessionResponse) SetAcsHelperUserToken(value *string)() {
    err := m.GetBackingStore().Set("acsHelperUserToken", value)
    if err != nil {
        panic(err)
    }
}
// SetAcsSharerUserId sets the acsSharerUserId property value. Sharer ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) SetAcsSharerUserId(value *string)() {
    err := m.GetBackingStore().Set("acsSharerUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetrieveRemoteHelpSessionResponse) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *RetrieveRemoteHelpSessionResponse) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceName sets the deviceName property value. Android Device Name
func (m *RetrieveRemoteHelpSessionResponse) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetrieveRemoteHelpSessionResponse) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPubSubGroupId sets the pubSubGroupId property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) SetPubSubGroupId(value *string)() {
    err := m.GetBackingStore().Set("pubSubGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetPubSubHelperAccessUri sets the pubSubHelperAccessUri property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) SetPubSubHelperAccessUri(value *string)() {
    err := m.GetBackingStore().Set("pubSubHelperAccessUri", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionExpirationDateTime sets the sessionExpirationDateTime property value. Azure Pubsub Session Expiration Date Time.
func (m *RetrieveRemoteHelpSessionResponse) SetSessionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("sessionExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionKey sets the sessionKey property value. The unique identifier for a session
func (m *RetrieveRemoteHelpSessionResponse) SetSessionKey(value *string)() {
    err := m.GetBackingStore().Set("sessionKey", value)
    if err != nil {
        panic(err)
    }
}
// RetrieveRemoteHelpSessionResponseable 
type RetrieveRemoteHelpSessionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcsGroupId()(*string)
    GetAcsHelperUserId()(*string)
    GetAcsHelperUserToken()(*string)
    GetAcsSharerUserId()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceName()(*string)
    GetOdataType()(*string)
    GetPubSubGroupId()(*string)
    GetPubSubHelperAccessUri()(*string)
    GetSessionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSessionKey()(*string)
    SetAcsGroupId(value *string)()
    SetAcsHelperUserId(value *string)()
    SetAcsHelperUserToken(value *string)()
    SetAcsSharerUserId(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceName(value *string)()
    SetOdataType(value *string)()
    SetPubSubGroupId(value *string)()
    SetPubSubHelperAccessUri(value *string)()
    SetSessionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSessionKey(value *string)()
}
