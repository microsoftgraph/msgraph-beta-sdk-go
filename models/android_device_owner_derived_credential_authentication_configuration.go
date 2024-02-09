package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration android COBO Derived Credential profile.
type AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration struct {
    DeviceConfiguration
}
// NewAndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration instantiates a new AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration and sets the default values.
func NewAndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration()(*AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) {
    m := &AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerDerivedCredentialAuthenticationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration(), nil
}
// GetCertificateAccessType gets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
// returns a *AndroidDeviceOwnerCertificateAccessType when successful
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType) {
    val, err := m.GetBackingStore().Get("certificateAccessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerCertificateAccessType)
    }
    return nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
// returns a DeviceManagementDerivedCredentialSettingsable when successful
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentialSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["certificateAccessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerCertificateAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateAccessType(val.(*AndroidDeviceOwnerCertificateAccessType))
        }
        return nil
    }
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    res["silentCertificateAccessDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerSilentCertificateAccessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerSilentCertificateAccessable)
                }
            }
            m.SetSilentCertificateAccessDetails(res)
        }
        return nil
    }
    return res
}
// GetSilentCertificateAccessDetails gets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
// returns a []AndroidDeviceOwnerSilentCertificateAccessable when successful
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable) {
    val, err := m.GetBackingStore().Get("silentCertificateAccessDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerSilentCertificateAccessable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateAccessType() != nil {
        cast := (*m.GetCertificateAccessType()).String()
        err = writer.WriteStringValue("certificateAccessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSilentCertificateAccessDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSilentCertificateAccessDetails()))
        for i, v := range m.GetSilentCertificateAccessDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("silentCertificateAccessDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAccessType sets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)() {
    err := m.GetBackingStore().Set("certificateAccessType", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentialSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSilentCertificateAccessDetails sets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration) SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)() {
    err := m.GetBackingStore().Set("silentCertificateAccessDetails", value)
    if err != nil {
        panic(err)
    }
}
type AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
}
