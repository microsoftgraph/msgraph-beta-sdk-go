package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PstnBlockedUsersLogRow struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPstnBlockedUsersLogRow instantiates a new PstnBlockedUsersLogRow and sets the default values.
func NewPstnBlockedUsersLogRow()(*PstnBlockedUsersLogRow) {
    m := &PstnBlockedUsersLogRow{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePstnBlockedUsersLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePstnBlockedUsersLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPstnBlockedUsersLogRow(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PstnBlockedUsersLogRow) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *PstnBlockedUsersLogRow) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockDateTime gets the blockDateTime property value. The date and time when the user was blocked/unblocked from making PSTN calls. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *PstnBlockedUsersLogRow) GetBlockDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("blockDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetBlockReason gets the blockReason property value. The reason why the user is blocked/unblocked from making calls.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetBlockReason()(*string) {
    val, err := m.GetBackingStore().Get("blockReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PstnBlockedUsersLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDateTime(val)
        }
        return nil
    }
    res["blockReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockReason(val)
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
    res["remediationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationId(val)
        }
        return nil
    }
    res["userBlockMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePstnUserBlockMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserBlockMode(val.(*PstnUserBlockMode))
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
    res["userTelephoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTelephoneNumber(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationId gets the remediationId property value. Unique identifier (GUID) for the blocking/unblocking action.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetRemediationId()(*string) {
    val, err := m.GetBackingStore().Get("remediationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserBlockMode gets the userBlockMode property value. Indicates whether the user is blocked or unblocked from making PSTN calls in Microsoft Teams. The possible values are: blocked, unblocked, unknownFutureValue.
// returns a *PstnUserBlockMode when successful
func (m *PstnBlockedUsersLogRow) GetUserBlockMode()(*PstnUserBlockMode) {
    val, err := m.GetBackingStore().Get("userBlockMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PstnUserBlockMode)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The unique identifier (GUID) of the user in Microsoft Entra ID.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserTelephoneNumber gets the userTelephoneNumber property value. User's blocked number. For details, see E.164.
// returns a *string when successful
func (m *PstnBlockedUsersLogRow) GetUserTelephoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("userTelephoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PstnBlockedUsersLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("blockDateTime", m.GetBlockDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blockReason", m.GetBlockReason())
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
        err := writer.WriteStringValue("remediationId", m.GetRemediationId())
        if err != nil {
            return err
        }
    }
    if m.GetUserBlockMode() != nil {
        cast := (*m.GetUserBlockMode()).String()
        err := writer.WriteStringValue("userBlockMode", &cast)
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
        err := writer.WriteStringValue("userTelephoneNumber", m.GetUserTelephoneNumber())
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
func (m *PstnBlockedUsersLogRow) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PstnBlockedUsersLogRow) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockDateTime sets the blockDateTime property value. The date and time when the user was blocked/unblocked from making PSTN calls. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PstnBlockedUsersLogRow) SetBlockDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("blockDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockReason sets the blockReason property value. The reason why the user is blocked/unblocked from making calls.
func (m *PstnBlockedUsersLogRow) SetBlockReason(value *string)() {
    err := m.GetBackingStore().Set("blockReason", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PstnBlockedUsersLogRow) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationId sets the remediationId property value. Unique identifier (GUID) for the blocking/unblocking action.
func (m *PstnBlockedUsersLogRow) SetRemediationId(value *string)() {
    err := m.GetBackingStore().Set("remediationId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserBlockMode sets the userBlockMode property value. Indicates whether the user is blocked or unblocked from making PSTN calls in Microsoft Teams. The possible values are: blocked, unblocked, unknownFutureValue.
func (m *PstnBlockedUsersLogRow) SetUserBlockMode(value *PstnUserBlockMode)() {
    err := m.GetBackingStore().Set("userBlockMode", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user.
func (m *PstnBlockedUsersLogRow) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The unique identifier (GUID) of the user in Microsoft Entra ID.
func (m *PstnBlockedUsersLogRow) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (sign-in name) in Microsoft Entra ID. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *PstnBlockedUsersLogRow) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserTelephoneNumber sets the userTelephoneNumber property value. User's blocked number. For details, see E.164.
func (m *PstnBlockedUsersLogRow) SetUserTelephoneNumber(value *string)() {
    err := m.GetBackingStore().Set("userTelephoneNumber", value)
    if err != nil {
        panic(err)
    }
}
type PstnBlockedUsersLogRowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBlockReason()(*string)
    GetOdataType()(*string)
    GetRemediationId()(*string)
    GetUserBlockMode()(*PstnUserBlockMode)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUserTelephoneNumber()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBlockReason(value *string)()
    SetOdataType(value *string)()
    SetRemediationId(value *string)()
    SetUserBlockMode(value *PstnUserBlockMode)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserTelephoneNumber(value *string)()
}
