package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StrongAuthenticationPhoneAppDetail 
type StrongAuthenticationPhoneAppDetail struct {
    Entity
}
// NewStrongAuthenticationPhoneAppDetail instantiates a new strongAuthenticationPhoneAppDetail and sets the default values.
func NewStrongAuthenticationPhoneAppDetail()(*StrongAuthenticationPhoneAppDetail) {
    m := &StrongAuthenticationPhoneAppDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateStrongAuthenticationPhoneAppDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStrongAuthenticationPhoneAppDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStrongAuthenticationPhoneAppDetail(), nil
}
// GetAuthenticationType gets the authenticationType property value. The authenticationType property
func (m *StrongAuthenticationPhoneAppDetail) GetAuthenticationType()(*string) {
    val, err := m.GetBackingStore().Get("authenticationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthenticatorFlavor gets the authenticatorFlavor property value. The authenticatorFlavor property
func (m *StrongAuthenticationPhoneAppDetail) GetAuthenticatorFlavor()(*string) {
    val, err := m.GetBackingStore().Get("authenticatorFlavor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The deviceName property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceTag gets the deviceTag property value. The deviceTag property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceTag()(*string) {
    val, err := m.GetBackingStore().Get("deviceTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceToken gets the deviceToken property value. The deviceToken property
func (m *StrongAuthenticationPhoneAppDetail) GetDeviceToken()(*string) {
    val, err := m.GetBackingStore().Get("deviceToken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StrongAuthenticationPhoneAppDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val)
        }
        return nil
    }
    res["authenticatorFlavor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticatorFlavor(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    res["deviceToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceToken(val)
        }
        return nil
    }
    res["hashFunction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashFunction(val)
        }
        return nil
    }
    res["lastAuthenticatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAuthenticatedDateTime(val)
        }
        return nil
    }
    res["notificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationType(val)
        }
        return nil
    }
    res["oathSecretKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathSecretKey(val)
        }
        return nil
    }
    res["oathTokenMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOathTokenMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathTokenMetadata(val.(OathTokenMetadataable))
        }
        return nil
    }
    res["oathTokenTimeDriftInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOathTokenTimeDriftInSeconds(val)
        }
        return nil
    }
    res["phoneAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneAppVersion(val)
        }
        return nil
    }
    res["tenantDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDeviceId(val)
        }
        return nil
    }
    res["tokenGenerationIntervalInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenGenerationIntervalInSeconds(val)
        }
        return nil
    }
    return res
}
// GetHashFunction gets the hashFunction property value. The hashFunction property
func (m *StrongAuthenticationPhoneAppDetail) GetHashFunction()(*string) {
    val, err := m.GetBackingStore().Get("hashFunction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastAuthenticatedDateTime gets the lastAuthenticatedDateTime property value. The lastAuthenticatedDateTime property
func (m *StrongAuthenticationPhoneAppDetail) GetLastAuthenticatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastAuthenticatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNotificationType gets the notificationType property value. The notificationType property
func (m *StrongAuthenticationPhoneAppDetail) GetNotificationType()(*string) {
    val, err := m.GetBackingStore().Get("notificationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOathSecretKey gets the oathSecretKey property value. The oathSecretKey property
func (m *StrongAuthenticationPhoneAppDetail) GetOathSecretKey()(*string) {
    val, err := m.GetBackingStore().Get("oathSecretKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOathTokenMetadata gets the oathTokenMetadata property value. The oathTokenMetadata property
func (m *StrongAuthenticationPhoneAppDetail) GetOathTokenMetadata()(OathTokenMetadataable) {
    val, err := m.GetBackingStore().Get("oathTokenMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OathTokenMetadataable)
    }
    return nil
}
// GetOathTokenTimeDriftInSeconds gets the oathTokenTimeDriftInSeconds property value. The oathTokenTimeDriftInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) GetOathTokenTimeDriftInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("oathTokenTimeDriftInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPhoneAppVersion gets the phoneAppVersion property value. The phoneAppVersion property
func (m *StrongAuthenticationPhoneAppDetail) GetPhoneAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("phoneAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantDeviceId gets the tenantDeviceId property value. The tenantDeviceId property
func (m *StrongAuthenticationPhoneAppDetail) GetTenantDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("tenantDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTokenGenerationIntervalInSeconds gets the tokenGenerationIntervalInSeconds property value. The tokenGenerationIntervalInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) GetTokenGenerationIntervalInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("tokenGenerationIntervalInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *StrongAuthenticationPhoneAppDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationType", m.GetAuthenticationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticatorFlavor", m.GetAuthenticatorFlavor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceToken", m.GetDeviceToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hashFunction", m.GetHashFunction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastAuthenticatedDateTime", m.GetLastAuthenticatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationType", m.GetNotificationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("oathSecretKey", m.GetOathSecretKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("oathTokenMetadata", m.GetOathTokenMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("oathTokenTimeDriftInSeconds", m.GetOathTokenTimeDriftInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneAppVersion", m.GetPhoneAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDeviceId", m.GetTenantDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("tokenGenerationIntervalInSeconds", m.GetTokenGenerationIntervalInSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationType sets the authenticationType property value. The authenticationType property
func (m *StrongAuthenticationPhoneAppDetail) SetAuthenticationType(value *string)() {
    err := m.GetBackingStore().Set("authenticationType", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticatorFlavor sets the authenticatorFlavor property value. The authenticatorFlavor property
func (m *StrongAuthenticationPhoneAppDetail) SetAuthenticatorFlavor(value *string)() {
    err := m.GetBackingStore().Set("authenticatorFlavor", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceTag sets the deviceTag property value. The deviceTag property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceTag(value *string)() {
    err := m.GetBackingStore().Set("deviceTag", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceToken sets the deviceToken property value. The deviceToken property
func (m *StrongAuthenticationPhoneAppDetail) SetDeviceToken(value *string)() {
    err := m.GetBackingStore().Set("deviceToken", value)
    if err != nil {
        panic(err)
    }
}
// SetHashFunction sets the hashFunction property value. The hashFunction property
func (m *StrongAuthenticationPhoneAppDetail) SetHashFunction(value *string)() {
    err := m.GetBackingStore().Set("hashFunction", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAuthenticatedDateTime sets the lastAuthenticatedDateTime property value. The lastAuthenticatedDateTime property
func (m *StrongAuthenticationPhoneAppDetail) SetLastAuthenticatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAuthenticatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationType sets the notificationType property value. The notificationType property
func (m *StrongAuthenticationPhoneAppDetail) SetNotificationType(value *string)() {
    err := m.GetBackingStore().Set("notificationType", value)
    if err != nil {
        panic(err)
    }
}
// SetOathSecretKey sets the oathSecretKey property value. The oathSecretKey property
func (m *StrongAuthenticationPhoneAppDetail) SetOathSecretKey(value *string)() {
    err := m.GetBackingStore().Set("oathSecretKey", value)
    if err != nil {
        panic(err)
    }
}
// SetOathTokenMetadata sets the oathTokenMetadata property value. The oathTokenMetadata property
func (m *StrongAuthenticationPhoneAppDetail) SetOathTokenMetadata(value OathTokenMetadataable)() {
    err := m.GetBackingStore().Set("oathTokenMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetOathTokenTimeDriftInSeconds sets the oathTokenTimeDriftInSeconds property value. The oathTokenTimeDriftInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) SetOathTokenTimeDriftInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("oathTokenTimeDriftInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneAppVersion sets the phoneAppVersion property value. The phoneAppVersion property
func (m *StrongAuthenticationPhoneAppDetail) SetPhoneAppVersion(value *string)() {
    err := m.GetBackingStore().Set("phoneAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDeviceId sets the tenantDeviceId property value. The tenantDeviceId property
func (m *StrongAuthenticationPhoneAppDetail) SetTenantDeviceId(value *string)() {
    err := m.GetBackingStore().Set("tenantDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenGenerationIntervalInSeconds sets the tokenGenerationIntervalInSeconds property value. The tokenGenerationIntervalInSeconds property
func (m *StrongAuthenticationPhoneAppDetail) SetTokenGenerationIntervalInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("tokenGenerationIntervalInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// StrongAuthenticationPhoneAppDetailable 
type StrongAuthenticationPhoneAppDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationType()(*string)
    GetAuthenticatorFlavor()(*string)
    GetDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetDeviceName()(*string)
    GetDeviceTag()(*string)
    GetDeviceToken()(*string)
    GetHashFunction()(*string)
    GetLastAuthenticatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotificationType()(*string)
    GetOathSecretKey()(*string)
    GetOathTokenMetadata()(OathTokenMetadataable)
    GetOathTokenTimeDriftInSeconds()(*int32)
    GetPhoneAppVersion()(*string)
    GetTenantDeviceId()(*string)
    GetTokenGenerationIntervalInSeconds()(*int32)
    SetAuthenticationType(value *string)()
    SetAuthenticatorFlavor(value *string)()
    SetDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetDeviceName(value *string)()
    SetDeviceTag(value *string)()
    SetDeviceToken(value *string)()
    SetHashFunction(value *string)()
    SetLastAuthenticatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotificationType(value *string)()
    SetOathSecretKey(value *string)()
    SetOathTokenMetadata(value OathTokenMetadataable)()
    SetOathTokenTimeDriftInSeconds(value *int32)()
    SetPhoneAppVersion(value *string)()
    SetTenantDeviceId(value *string)()
    SetTokenGenerationIntervalInSeconds(value *int32)()
}
