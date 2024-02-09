package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEasEmailProfileBase base for Android For Work EAS Email profiles
type AndroidForWorkEasEmailProfileBase struct {
    DeviceConfiguration
}
// NewAndroidForWorkEasEmailProfileBase instantiates a new AndroidForWorkEasEmailProfileBase and sets the default values.
func NewAndroidForWorkEasEmailProfileBase()(*AndroidForWorkEasEmailProfileBase) {
    m := &AndroidForWorkEasEmailProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkEasEmailProfileBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidForWorkEasEmailProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidForWorkEasEmailProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidForWorkGmailEasConfiguration":
                        return NewAndroidForWorkGmailEasConfiguration(), nil
                    case "#microsoft.graph.androidForWorkNineWorkEasConfiguration":
                        return NewAndroidForWorkNineWorkEasConfiguration(), nil
                }
            }
        }
    }
    return NewAndroidForWorkEasEmailProfileBase(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Exchange Active Sync authentication method.
// returns a *EasAuthenticationMethod when successful
func (m *AndroidForWorkEasEmailProfileBase) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EasAuthenticationMethod)
    }
    return nil
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
// returns a *EmailSyncDuration when successful
func (m *AndroidForWorkEasEmailProfileBase) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    val, err := m.GetBackingStore().Get("durationOfEmailToSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailSyncDuration)
    }
    return nil
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
// returns a *UserEmailSource when successful
func (m *AndroidForWorkEasEmailProfileBase) GetEmailAddressSource()(*UserEmailSource) {
    val, err := m.GetBackingStore().Get("emailAddressSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserEmailSource)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidForWorkEasEmailProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEasAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*EasAuthenticationMethod))
        }
        return nil
    }
    res["durationOfEmailToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailSyncDuration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationOfEmailToSync(val.(*EmailSyncDuration))
        }
        return nil
    }
    res["emailAddressSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserEmailSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddressSource(val.(*UserEmailSource))
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidForWorkCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidForWorkCertificateProfileBaseable))
        }
        return nil
    }
    res["requireSsl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSsl(val)
        }
        return nil
    }
    res["usernameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidUsernameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameSource(val.(*AndroidUsernameSource))
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Exchange location (URL) that the mail app connects to.
// returns a *string when successful
func (m *AndroidForWorkEasEmailProfileBase) GetHostName()(*string) {
    val, err := m.GetBackingStore().Get("hostName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
// returns a AndroidForWorkCertificateProfileBaseable when successful
func (m *AndroidForWorkEasEmailProfileBase) GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidForWorkCertificateProfileBaseable)
    }
    return nil
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
// returns a *bool when successful
func (m *AndroidForWorkEasEmailProfileBase) GetRequireSsl()(*bool) {
    val, err := m.GetBackingStore().Get("requireSsl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUsernameSource gets the usernameSource property value. Android username source.
// returns a *AndroidUsernameSource when successful
func (m *AndroidForWorkEasEmailProfileBase) GetUsernameSource()(*AndroidUsernameSource) {
    val, err := m.GetBackingStore().Get("usernameSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidUsernameSource)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidForWorkEasEmailProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDurationOfEmailToSync() != nil {
        cast := (*m.GetDurationOfEmailToSync()).String()
        err = writer.WriteStringValue("durationOfEmailToSync", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailAddressSource() != nil {
        cast := (*m.GetEmailAddressSource()).String()
        err = writer.WriteStringValue("emailAddressSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostName", m.GetHostName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireSsl", m.GetRequireSsl())
        if err != nil {
            return err
        }
    }
    if m.GetUsernameSource() != nil {
        cast := (*m.GetUsernameSource()).String()
        err = writer.WriteStringValue("usernameSource", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidForWorkEasEmailProfileBase) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidForWorkEasEmailProfileBase) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    err := m.GetBackingStore().Set("durationOfEmailToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidForWorkEasEmailProfileBase) SetEmailAddressSource(value *UserEmailSource)() {
    err := m.GetBackingStore().Set("emailAddressSource", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidForWorkEasEmailProfileBase) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidForWorkEasEmailProfileBase) SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidForWorkEasEmailProfileBase) SetRequireSsl(value *bool)() {
    err := m.GetBackingStore().Set("requireSsl", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameSource sets the usernameSource property value. Android username source.
func (m *AndroidForWorkEasEmailProfileBase) SetUsernameSource(value *AndroidUsernameSource)() {
    err := m.GetBackingStore().Set("usernameSource", value)
    if err != nil {
        panic(err)
    }
}
type AndroidForWorkEasEmailProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*EasAuthenticationMethod)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetHostName()(*string)
    GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable)
    GetRequireSsl()(*bool)
    GetUsernameSource()(*AndroidUsernameSource)
    SetAuthenticationMethod(value *EasAuthenticationMethod)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetHostName(value *string)()
    SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)()
    SetRequireSsl(value *bool)()
    SetUsernameSource(value *AndroidUsernameSource)()
}
