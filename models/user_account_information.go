// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserAccountInformation struct {
    ItemFacet
}
// NewUserAccountInformation instantiates a new UserAccountInformation and sets the default values.
func NewUserAccountInformation()(*UserAccountInformation) {
    m := &UserAccountInformation{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.userAccountInformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserAccountInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserAccountInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAccountInformation(), nil
}
// GetAgeGroup gets the ageGroup property value. Shows the age group of user. Allowed values null, minor, notAdult and adult are generated by the directory and can't be changed.
// returns a *string when successful
func (m *UserAccountInformation) GetAgeGroup()(*string) {
    val, err := m.GetBackingStore().Get("ageGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountryCode gets the countryCode property value. Contains the two-character country code associated with the users' account.
// returns a *string when successful
func (m *UserAccountInformation) GetCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("countryCode")
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
func (m *UserAccountInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["ageGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgeGroup(val)
        }
        return nil
    }
    res["countryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["preferredLanguageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguageTag(val.(LocaleInfoable))
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
// GetPreferredLanguageTag gets the preferredLanguageTag property value. Contains the language the user has associated as preferred for the account.
// returns a LocaleInfoable when successful
func (m *UserAccountInformation) GetPreferredLanguageTag()(LocaleInfoable) {
    val, err := m.GetBackingStore().Get("preferredLanguageTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LocaleInfoable)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user associated with the account.
// returns a *string when successful
func (m *UserAccountInformation) GetUserPrincipalName()(*string) {
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
func (m *UserAccountInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ageGroup", m.GetAgeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("preferredLanguageTag", m.GetPreferredLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgeGroup sets the ageGroup property value. Shows the age group of user. Allowed values null, minor, notAdult and adult are generated by the directory and can't be changed.
func (m *UserAccountInformation) SetAgeGroup(value *string)() {
    err := m.GetBackingStore().Set("ageGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetCountryCode sets the countryCode property value. Contains the two-character country code associated with the users' account.
func (m *UserAccountInformation) SetCountryCode(value *string)() {
    err := m.GetBackingStore().Set("countryCode", value)
    if err != nil {
        panic(err)
    }
}
// SetPreferredLanguageTag sets the preferredLanguageTag property value. Contains the language the user has associated as preferred for the account.
func (m *UserAccountInformation) SetPreferredLanguageTag(value LocaleInfoable)() {
    err := m.GetBackingStore().Set("preferredLanguageTag", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user associated with the account.
func (m *UserAccountInformation) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type UserAccountInformationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgeGroup()(*string)
    GetCountryCode()(*string)
    GetPreferredLanguageTag()(LocaleInfoable)
    GetUserPrincipalName()(*string)
    SetAgeGroup(value *string)()
    SetCountryCode(value *string)()
    SetPreferredLanguageTag(value LocaleInfoable)()
    SetUserPrincipalName(value *string)()
}
