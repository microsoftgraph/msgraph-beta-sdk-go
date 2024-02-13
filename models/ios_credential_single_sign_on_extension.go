package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCredentialSingleSignOnExtension represents a Credential-type Single Sign-On extension profile for iOS devices.
type IosCredentialSingleSignOnExtension struct {
    IosSingleSignOnExtension
}
// NewIosCredentialSingleSignOnExtension instantiates a new IosCredentialSingleSignOnExtension and sets the default values.
func NewIosCredentialSingleSignOnExtension()(*IosCredentialSingleSignOnExtension) {
    m := &IosCredentialSingleSignOnExtension{
        IosSingleSignOnExtension: *NewIosSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.iosCredentialSingleSignOnExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosCredentialSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosCredentialSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosCredentialSingleSignOnExtension(), nil
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
// returns a []KeyTypedValuePairable when successful
func (m *IosCredentialSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    val, err := m.GetBackingStore().Get("configurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyTypedValuePairable)
    }
    return nil
}
// GetDomains gets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
// returns a []string when successful
func (m *IosCredentialSingleSignOnExtension) GetDomains()([]string) {
    val, err := m.GetBackingStore().Get("domains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetExtensionIdentifier gets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
// returns a *string when successful
func (m *IosCredentialSingleSignOnExtension) GetExtensionIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("extensionIdentifier")
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
func (m *IosCredentialSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosSingleSignOnExtension.GetFieldDeserializers()
    res["configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyTypedValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyTypedValuePairable)
                }
            }
            m.SetConfigurations(res)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["extensionIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionIdentifier(val)
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["teamIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamIdentifier(val)
        }
        return nil
    }
    return res
}
// GetRealm gets the realm property value. Gets or sets the case-sensitive realm name for this profile.
// returns a *string when successful
func (m *IosCredentialSingleSignOnExtension) GetRealm()(*string) {
    val, err := m.GetBackingStore().Get("realm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
// returns a *string when successful
func (m *IosCredentialSingleSignOnExtension) GetTeamIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("teamIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosCredentialSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurations()))
        for i, v := range m.GetConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err = writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extensionIdentifier", m.GetExtensionIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *IosCredentialSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    err := m.GetBackingStore().Set("configurations", value)
    if err != nil {
        panic(err)
    }
}
// SetDomains sets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *IosCredentialSingleSignOnExtension) SetDomains(value []string)() {
    err := m.GetBackingStore().Set("domains", value)
    if err != nil {
        panic(err)
    }
}
// SetExtensionIdentifier sets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *IosCredentialSingleSignOnExtension) SetExtensionIdentifier(value *string)() {
    err := m.GetBackingStore().Set("extensionIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetRealm sets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *IosCredentialSingleSignOnExtension) SetRealm(value *string)() {
    err := m.GetBackingStore().Set("realm", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *IosCredentialSingleSignOnExtension) SetTeamIdentifier(value *string)() {
    err := m.GetBackingStore().Set("teamIdentifier", value)
    if err != nil {
        panic(err)
    }
}
type IosCredentialSingleSignOnExtensionable interface {
    IosSingleSignOnExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurations()([]KeyTypedValuePairable)
    GetDomains()([]string)
    GetExtensionIdentifier()(*string)
    GetRealm()(*string)
    GetTeamIdentifier()(*string)
    SetConfigurations(value []KeyTypedValuePairable)()
    SetDomains(value []string)()
    SetExtensionIdentifier(value *string)()
    SetRealm(value *string)()
    SetTeamIdentifier(value *string)()
}
