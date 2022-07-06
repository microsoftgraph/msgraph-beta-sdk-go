package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileEasEmailProfileBase 
type AndroidWorkProfileEasEmailProfileBase struct {
    DeviceConfiguration
    // Authentication method for Exchange ActiveSync. Possible values are: usernameAndPassword, certificate, derivedCredential.
    authenticationMethod *EasAuthenticationMethod
    // Duration of time email should be synced to. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
    durationOfEmailToSync *EmailSyncDuration
    // Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
    emailAddressSource *UserEmailSource
    // Exchange location (URL) that the mail app connects to.
    hostName *string
    // Identity certificate.
    identityCertificate AndroidWorkProfileCertificateProfileBaseable
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // The type property
    type_escaped *string
    // Username attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: username, userPrincipalName, samAccountName, primarySmtpAddress.
    usernameSource *AndroidUsernameSource
}
// NewAndroidWorkProfileEasEmailProfileBase instantiates a new AndroidWorkProfileEasEmailProfileBase and sets the default values.
func NewAndroidWorkProfileEasEmailProfileBase()(*AndroidWorkProfileEasEmailProfileBase) {
    m := &AndroidWorkProfileEasEmailProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
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
                mappingStr := *mappingValue
                switch mappingStr {
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
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication method for Exchange ActiveSync. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *AndroidWorkProfileEasEmailProfileBase) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Duration of time email should be synced to. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
func (m *AndroidWorkProfileEasEmailProfileBase) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    if m == nil {
        return nil
    } else {
        return m.durationOfEmailToSync
    }
}
// GetEmailAddressSource gets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *AndroidWorkProfileEasEmailProfileBase) GetEmailAddressSource()(*UserEmailSource) {
    if m == nil {
        return nil
    } else {
        return m.emailAddressSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileEasEmailProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateAndroidWorkProfileCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidWorkProfileCertificateProfileBaseable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
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
func (m *AndroidWorkProfileEasEmailProfileBase) GetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostName
    }
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *AndroidWorkProfileEasEmailProfileBase) GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable) {
    if m == nil {
        return nil
    } else {
        return m.identityCertificate
    }
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidWorkProfileEasEmailProfileBase) GetRequireSsl()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireSsl
    }
}
// GetType gets the type property value. The type property
func (m *AndroidWorkProfileEasEmailProfileBase) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUsernameSource gets the usernameSource property value. Username attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: username, userPrincipalName, samAccountName, primarySmtpAddress.
func (m *AndroidWorkProfileEasEmailProfileBase) GetUsernameSource()(*AndroidUsernameSource) {
    if m == nil {
        return nil
    } else {
        return m.usernameSource
    }
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
    {
        err = writer.WriteStringValue("type", m.GetType())
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
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication method for Exchange ActiveSync. Possible values are: usernameAndPassword, certificate, derivedCredential.
func (m *AndroidWorkProfileEasEmailProfileBase) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    if m != nil {
        m.authenticationMethod = value
    }
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Duration of time email should be synced to. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
func (m *AndroidWorkProfileEasEmailProfileBase) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    if m != nil {
        m.durationOfEmailToSync = value
    }
}
// SetEmailAddressSource sets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *AndroidWorkProfileEasEmailProfileBase) SetEmailAddressSource(value *UserEmailSource)() {
    if m != nil {
        m.emailAddressSource = value
    }
}
// SetHostName sets the hostName property value. Exchange location (URL) that the mail app connects to.
func (m *AndroidWorkProfileEasEmailProfileBase) SetHostName(value *string)() {
    if m != nil {
        m.hostName = value
    }
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidWorkProfileEasEmailProfileBase) SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)() {
    if m != nil {
        m.identityCertificate = value
    }
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidWorkProfileEasEmailProfileBase) SetRequireSsl(value *bool)() {
    if m != nil {
        m.requireSsl = value
    }
}
// SetType sets the type property value. The type property
func (m *AndroidWorkProfileEasEmailProfileBase) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUsernameSource sets the usernameSource property value. Username attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: username, userPrincipalName, samAccountName, primarySmtpAddress.
func (m *AndroidWorkProfileEasEmailProfileBase) SetUsernameSource(value *AndroidUsernameSource)() {
    if m != nil {
        m.usernameSource = value
    }
}
