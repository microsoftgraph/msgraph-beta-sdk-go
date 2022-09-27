package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileEasEmailProfileBase 
type AndroidWorkProfileEasEmailProfileBase struct {
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
    identityCertificate AndroidWorkProfileCertificateProfileBaseable
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // Android username source.
    usernameSource *AndroidUsernameSource
}
// NewAndroidWorkProfileEasEmailProfileBase instantiates a new AndroidWorkProfileEasEmailProfileBase and sets the default values.
func NewAndroidWorkProfileEasEmailProfileBase()(*AndroidWorkProfileEasEmailProfileBase) {
    m := &AndroidWorkProfileEasEmailProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileEasEmailProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfileEasEmailProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileEasEmailProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidWorkProfileGmailEasConfiguration":
                        return NewAndroidWorkProfileGmailEasConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileNineWorkEasConfiguration":
                        return NewAndroidWorkProfileNineWorkEasConfiguration(), nil
                }
            }
        }
    }
    return NewAndroidWorkProfileEasEmailProfileBase(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidWorkProfileEasEmailProfileBase) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    return m.authenticationMethod
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidWorkProfileEasEmailProfileBase) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    return m.durationOfEmailToSync
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidWorkProfileEasEmailProfileBase) GetEmailAddressSource()(*UserEmailSource) {
    return m.emailAddressSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileEasEmailProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEasAuthenticationMethod , m.SetAuthenticationMethod)
    res["durationOfEmailToSync"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEmailSyncDuration , m.SetDurationOfEmailToSync)
    res["emailAddressSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserEmailSource , m.SetEmailAddressSource)
    res["hostName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHostName)
    res["identityCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue , m.SetIdentityCertificate)
    res["requireSsl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireSsl)
    res["usernameSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidUsernameSource , m.SetUsernameSource)
    return res
}
// GetHostName gets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidWorkProfileEasEmailProfileBase) GetHostName()(*string) {
    return m.hostName
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *AndroidWorkProfileEasEmailProfileBase) GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidWorkProfileEasEmailProfileBase) GetRequireSsl()(*bool) {
    return m.requireSsl
}
// GetUsernameSource gets the usernameSource property value. Android username source.
func (m *AndroidWorkProfileEasEmailProfileBase) GetUsernameSource()(*AndroidUsernameSource) {
    return m.usernameSource
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileEasEmailProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidWorkProfileEasEmailProfileBase) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidWorkProfileEasEmailProfileBase) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    m.durationOfEmailToSync = value
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidWorkProfileEasEmailProfileBase) SetEmailAddressSource(value *UserEmailSource)() {
    m.emailAddressSource = value
}
// SetHostName sets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidWorkProfileEasEmailProfileBase) SetHostName(value *string)() {
    m.hostName = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidWorkProfileEasEmailProfileBase) SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidWorkProfileEasEmailProfileBase) SetRequireSsl(value *bool)() {
    m.requireSsl = value
}
// SetUsernameSource sets the usernameSource property value. Android username source.
func (m *AndroidWorkProfileEasEmailProfileBase) SetUsernameSource(value *AndroidUsernameSource)() {
    m.usernameSource = value
}
