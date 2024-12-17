package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OidcInboundClaimMappingOverride struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOidcInboundClaimMappingOverride instantiates a new OidcInboundClaimMappingOverride and sets the default values.
func NewOidcInboundClaimMappingOverride()(*OidcInboundClaimMappingOverride) {
    m := &OidcInboundClaimMappingOverride{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOidcInboundClaimMappingOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOidcInboundClaimMappingOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOidcInboundClaimMappingOverride(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OidcInboundClaimMappingOverride) GetAdditionalData()(map[string]any) {
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
// GetAddress gets the address property value. End-user's preferred postal address. The value of the address member is a JSON RFC8259 structure containing some or all of the members defined in the resource type
// returns a OidcAddressInboundClaimsable when successful
func (m *OidcInboundClaimMappingOverride) GetAddress()(OidcAddressInboundClaimsable) {
    val, err := m.GetBackingStore().Get("address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OidcAddressInboundClaimsable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OidcInboundClaimMappingOverride) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEmail gets the email property value. End-user's preferred e-mail address. Its value MUST conform to the RFC 5322 addr-spec syntax.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailVerified gets the email_verified property value. True if the end-user's e-mail address has been verified by the identity provider; otherwise, false. When this claim value is true, this means that your identity provider took affirmative steps to ensure that this e-mail address was controlled by the end-user at the time the verification was performed. If this claim value is false, or not mapped with any claim of the identity provider, the user is asked to verify email during sign-up if email is required in the user flow.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetEmailVerified()(*string) {
    val, err := m.GetBackingStore().Get("email_verified")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFamilyName gets the family_name property value. Surname(s) or family name of the end-user.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetFamilyName()(*string) {
    val, err := m.GetBackingStore().Get("family_name")
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
func (m *OidcInboundClaimMappingOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOidcAddressInboundClaimsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(OidcAddressInboundClaimsable))
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["email_verified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailVerified(val)
        }
        return nil
    }
    res["family_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFamilyName(val)
        }
        return nil
    }
    res["given_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["phone_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["phone_number_verified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumberVerified(val)
        }
        return nil
    }
    res["sub"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSub(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the given_name property value. Given name(s) or first name(s) of the end-user.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetGivenName()(*string) {
    val, err := m.GetBackingStore().Get("given_name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetName gets the name property value. End-user's full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the end-user's locale and preferences.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
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
func (m *OidcInboundClaimMappingOverride) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhoneNumber gets the phone_number property value. The claim provides the phone number for the user.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phone_number")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhoneNumberVerified gets the phone_number_verified property value. True if the end-user's phone number has been verified; otherwise, false. When this claim value is true, this means that your identity provider took affirmative steps to verify the phone number.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetPhoneNumberVerified()(*string) {
    val, err := m.GetBackingStore().Get("phone_number_verified")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSub gets the sub property value. Subject - Identifier for the end-user at the Issuer.
// returns a *string when successful
func (m *OidcInboundClaimMappingOverride) GetSub()(*string) {
    val, err := m.GetBackingStore().Get("sub")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OidcInboundClaimMappingOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email_verified", m.GetEmailVerified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("family_name", m.GetFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("given_name", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("phone_number", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phone_number_verified", m.GetPhoneNumberVerified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sub", m.GetSub())
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
func (m *OidcInboundClaimMappingOverride) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddress sets the address property value. End-user's preferred postal address. The value of the address member is a JSON RFC8259 structure containing some or all of the members defined in the resource type
func (m *OidcInboundClaimMappingOverride) SetAddress(value OidcAddressInboundClaimsable)() {
    err := m.GetBackingStore().Set("address", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OidcInboundClaimMappingOverride) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEmail sets the email property value. End-user's preferred e-mail address. Its value MUST conform to the RFC 5322 addr-spec syntax.
func (m *OidcInboundClaimMappingOverride) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailVerified sets the email_verified property value. True if the end-user's e-mail address has been verified by the identity provider; otherwise, false. When this claim value is true, this means that your identity provider took affirmative steps to ensure that this e-mail address was controlled by the end-user at the time the verification was performed. If this claim value is false, or not mapped with any claim of the identity provider, the user is asked to verify email during sign-up if email is required in the user flow.
func (m *OidcInboundClaimMappingOverride) SetEmailVerified(value *string)() {
    err := m.GetBackingStore().Set("email_verified", value)
    if err != nil {
        panic(err)
    }
}
// SetFamilyName sets the family_name property value. Surname(s) or family name of the end-user.
func (m *OidcInboundClaimMappingOverride) SetFamilyName(value *string)() {
    err := m.GetBackingStore().Set("family_name", value)
    if err != nil {
        panic(err)
    }
}
// SetGivenName sets the given_name property value. Given name(s) or first name(s) of the end-user.
func (m *OidcInboundClaimMappingOverride) SetGivenName(value *string)() {
    err := m.GetBackingStore().Set("given_name", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. End-user's full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the end-user's locale and preferences.
func (m *OidcInboundClaimMappingOverride) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OidcInboundClaimMappingOverride) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phone_number property value. The claim provides the phone number for the user.
func (m *OidcInboundClaimMappingOverride) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phone_number", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumberVerified sets the phone_number_verified property value. True if the end-user's phone number has been verified; otherwise, false. When this claim value is true, this means that your identity provider took affirmative steps to verify the phone number.
func (m *OidcInboundClaimMappingOverride) SetPhoneNumberVerified(value *string)() {
    err := m.GetBackingStore().Set("phone_number_verified", value)
    if err != nil {
        panic(err)
    }
}
// SetSub sets the sub property value. Subject - Identifier for the end-user at the Issuer.
func (m *OidcInboundClaimMappingOverride) SetSub(value *string)() {
    err := m.GetBackingStore().Set("sub", value)
    if err != nil {
        panic(err)
    }
}
type OidcInboundClaimMappingOverrideable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(OidcAddressInboundClaimsable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEmail()(*string)
    GetEmailVerified()(*string)
    GetFamilyName()(*string)
    GetGivenName()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetPhoneNumber()(*string)
    GetPhoneNumberVerified()(*string)
    GetSub()(*string)
    SetAddress(value OidcAddressInboundClaimsable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEmail(value *string)()
    SetEmailVerified(value *string)()
    SetFamilyName(value *string)()
    SetGivenName(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPhoneNumber(value *string)()
    SetPhoneNumberVerified(value *string)()
    SetSub(value *string)()
}
