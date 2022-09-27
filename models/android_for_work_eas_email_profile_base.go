package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEasEmailProfileBase 
type AndroidForWorkEasEmailProfileBase struct {
    DeviceConfiguration
    // Exchange Active Sync authentication method.
    authenticationMethod *EasAuthenticationMethod
    // Possible values for email sync duration.
    durationOfEmailToSync *EmailSyncDuration
    // Possible values for username source or email source.
    emailAddressSource *UserEmailSource
    // Exchange location (URL) that the mail app connects to.
    hostName *string
    // Identity certificate.
    identityCertificate AndroidForWorkCertificateProfileBaseable
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // Android username source.
    usernameSource *AndroidUsernameSource
}
// NewAndroidForWorkEasEmailProfileBase instantiates a new AndroidForWorkEasEmailProfileBase and sets the default values.
func NewAndroidForWorkEasEmailProfileBase()(*AndroidForWorkEasEmailProfileBase) {
    m := &AndroidForWorkEasEmailProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkEasEmailProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkEasEmailProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
func (m *AndroidForWorkEasEmailProfileBase) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    return m.authenticationMethod
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidForWorkEasEmailProfileBase) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    return m.durationOfEmailToSync
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidForWorkEasEmailProfileBase) GetEmailAddressSource()(*UserEmailSource) {
    return m.emailAddressSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkEasEmailProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEasAuthenticationMethod , m.SetAuthenticationMethod)
    res["durationOfEmailToSync"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEmailSyncDuration , m.SetDurationOfEmailToSync)
    res["emailAddressSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserEmailSource , m.SetEmailAddressSource)
    res["hostName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHostName)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidForWorkCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["requireSsl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireSsl)
    res["usernameSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidUsernameSource , m.SetUsernameSource)
    return res
}
// GetHostName gets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidForWorkEasEmailProfileBase) GetHostName()(*string) {
    return m.hostName
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *AndroidForWorkEasEmailProfileBase) GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidForWorkEasEmailProfileBase) GetRequireSsl()(*bool) {
    return m.requireSsl
}
// GetUsernameSource gets the usernameSource property value. Android username source.
func (m *AndroidForWorkEasEmailProfileBase) GetUsernameSource()(*AndroidUsernameSource) {
    return m.usernameSource
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
    m.authenticationMethod = value
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidForWorkEasEmailProfileBase) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    m.durationOfEmailToSync = value
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidForWorkEasEmailProfileBase) SetEmailAddressSource(value *UserEmailSource)() {
    m.emailAddressSource = value
}
// SetHostName sets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidForWorkEasEmailProfileBase) SetHostName(value *string)() {
    m.hostName = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidForWorkEasEmailProfileBase) SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidForWorkEasEmailProfileBase) SetRequireSsl(value *bool)() {
    m.requireSsl = value
}
// SetUsernameSource sets the usernameSource property value. Android username source.
func (m *AndroidForWorkEasEmailProfileBase) SetUsernameSource(value *AndroidUsernameSource)() {
    m.usernameSource = value
}
