package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetrieveRemoteHelpSessionResponse remote help - response we provide back to the helper on retrieve session API call
type RetrieveRemoteHelpSessionResponse struct {
    // ACS Group Id
    acsGroupId *string
    // Helper ACS User Id
    acsHelperUserId *string
    // Helper ACS User Token
    acsHelperUserToken *string
    // Sharer ACS User Id
    acsSharerUserId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Android Device Name
    deviceName *string
    // The OdataType property
    odataType *string
    // Azure Pubsub Group Id
    pubSubGroupId *string
    // Azure Pubsub Group Id
    pubSubHelperAccessUri *string
    // Azure Pubsub Session Expiration Date Time.
    sessionExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier for a session
    sessionKey *string
}
// NewRetrieveRemoteHelpSessionResponse instantiates a new retrieveRemoteHelpSessionResponse and sets the default values.
func NewRetrieveRemoteHelpSessionResponse()(*RetrieveRemoteHelpSessionResponse) {
    m := &RetrieveRemoteHelpSessionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.retrieveRemoteHelpSessionResponse";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRetrieveRemoteHelpSessionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetrieveRemoteHelpSessionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetrieveRemoteHelpSessionResponse(), nil
}
// GetAcsGroupId gets the acsGroupId property value. ACS Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsGroupId()(*string) {
    return m.acsGroupId
}
// GetAcsHelperUserId gets the acsHelperUserId property value. Helper ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsHelperUserId()(*string) {
    return m.acsHelperUserId
}
// GetAcsHelperUserToken gets the acsHelperUserToken property value. Helper ACS User Token
func (m *RetrieveRemoteHelpSessionResponse) GetAcsHelperUserToken()(*string) {
    return m.acsHelperUserToken
}
// GetAcsSharerUserId gets the acsSharerUserId property value. Sharer ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) GetAcsSharerUserId()(*string) {
    return m.acsSharerUserId
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetrieveRemoteHelpSessionResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceName gets the deviceName property value. Android Device Name
func (m *RetrieveRemoteHelpSessionResponse) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetrieveRemoteHelpSessionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acsGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAcsGroupId)
    res["acsHelperUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAcsHelperUserId)
    res["acsHelperUserToken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAcsHelperUserToken)
    res["acsSharerUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAcsSharerUserId)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["pubSubGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPubSubGroupId)
    res["pubSubHelperAccessUri"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPubSubHelperAccessUri)
    res["sessionExpirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetSessionExpirationDateTime)
    res["sessionKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSessionKey)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetrieveRemoteHelpSessionResponse) GetOdataType()(*string) {
    return m.odataType
}
// GetPubSubGroupId gets the pubSubGroupId property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetPubSubGroupId()(*string) {
    return m.pubSubGroupId
}
// GetPubSubHelperAccessUri gets the pubSubHelperAccessUri property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) GetPubSubHelperAccessUri()(*string) {
    return m.pubSubHelperAccessUri
}
// GetSessionExpirationDateTime gets the sessionExpirationDateTime property value. Azure Pubsub Session Expiration Date Time.
func (m *RetrieveRemoteHelpSessionResponse) GetSessionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sessionExpirationDateTime
}
// GetSessionKey gets the sessionKey property value. The unique identifier for a session
func (m *RetrieveRemoteHelpSessionResponse) GetSessionKey()(*string) {
    return m.sessionKey
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
    m.acsGroupId = value
}
// SetAcsHelperUserId sets the acsHelperUserId property value. Helper ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) SetAcsHelperUserId(value *string)() {
    m.acsHelperUserId = value
}
// SetAcsHelperUserToken sets the acsHelperUserToken property value. Helper ACS User Token
func (m *RetrieveRemoteHelpSessionResponse) SetAcsHelperUserToken(value *string)() {
    m.acsHelperUserToken = value
}
// SetAcsSharerUserId sets the acsSharerUserId property value. Sharer ACS User Id
func (m *RetrieveRemoteHelpSessionResponse) SetAcsSharerUserId(value *string)() {
    m.acsSharerUserId = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetrieveRemoteHelpSessionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceName sets the deviceName property value. Android Device Name
func (m *RetrieveRemoteHelpSessionResponse) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetrieveRemoteHelpSessionResponse) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPubSubGroupId sets the pubSubGroupId property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) SetPubSubGroupId(value *string)() {
    m.pubSubGroupId = value
}
// SetPubSubHelperAccessUri sets the pubSubHelperAccessUri property value. Azure Pubsub Group Id
func (m *RetrieveRemoteHelpSessionResponse) SetPubSubHelperAccessUri(value *string)() {
    m.pubSubHelperAccessUri = value
}
// SetSessionExpirationDateTime sets the sessionExpirationDateTime property value. Azure Pubsub Session Expiration Date Time.
func (m *RetrieveRemoteHelpSessionResponse) SetSessionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sessionExpirationDateTime = value
}
// SetSessionKey sets the sessionKey property value. The unique identifier for a session
func (m *RetrieveRemoteHelpSessionResponse) SetSessionKey(value *string)() {
    m.sessionKey = value
}
