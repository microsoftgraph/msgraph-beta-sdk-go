package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSImportedPFXCertificateProfile macOS PFX Import certificate profile
type MacOSImportedPFXCertificateProfile struct {
    MacOSCertificateProfileBase
}
// NewMacOSImportedPFXCertificateProfile instantiates a new MacOSImportedPFXCertificateProfile and sets the default values.
func NewMacOSImportedPFXCertificateProfile()(*MacOSImportedPFXCertificateProfile) {
    m := &MacOSImportedPFXCertificateProfile{
        MacOSCertificateProfileBase: *NewMacOSCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.macOSImportedPFXCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSImportedPFXCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSImportedPFXCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSImportedPFXCertificateProfile(), nil
}
// GetDeploymentChannel gets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
// returns a *AppleDeploymentChannel when successful
func (m *MacOSImportedPFXCertificateProfile) GetDeploymentChannel()(*AppleDeploymentChannel) {
    val, err := m.GetBackingStore().Get("deploymentChannel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleDeploymentChannel)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSImportedPFXCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSCertificateProfileBase.GetFieldDeserializers()
    res["deploymentChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleDeploymentChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentChannel(val.(*AppleDeploymentChannel))
        }
        return nil
    }
    res["intendedPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIntendedPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedPurpose(val.(*IntendedPurpose))
        }
        return nil
    }
    res["managedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCertificateStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceCertificateStateable)
                }
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    return res
}
// GetIntendedPurpose gets the intendedPurpose property value. PFX Import Options.
// returns a *IntendedPurpose when successful
func (m *MacOSImportedPFXCertificateProfile) GetIntendedPurpose()(*IntendedPurpose) {
    val, err := m.GetBackingStore().Get("intendedPurpose")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IntendedPurpose)
    }
    return nil
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
// returns a []ManagedDeviceCertificateStateable when successful
func (m *MacOSImportedPFXCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSImportedPFXCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentChannel() != nil {
        cast := (*m.GetDeploymentChannel()).String()
        err = writer.WriteStringValue("deploymentChannel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntendedPurpose() != nil {
        cast := (*m.GetIntendedPurpose()).String()
        err = writer.WriteStringValue("intendedPurpose", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCertificateStates()))
        for i, v := range m.GetManagedDeviceCertificateStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentChannel sets the deploymentChannel property value. Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel, userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
func (m *MacOSImportedPFXCertificateProfile) SetDeploymentChannel(value *AppleDeploymentChannel)() {
    err := m.GetBackingStore().Set("deploymentChannel", value)
    if err != nil {
        panic(err)
    }
}
// SetIntendedPurpose sets the intendedPurpose property value. PFX Import Options.
func (m *MacOSImportedPFXCertificateProfile) SetIntendedPurpose(value *IntendedPurpose)() {
    err := m.GetBackingStore().Set("intendedPurpose", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *MacOSImportedPFXCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
type MacOSImportedPFXCertificateProfileable interface {
    MacOSCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentChannel()(*AppleDeploymentChannel)
    GetIntendedPurpose()(*IntendedPurpose)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    SetDeploymentChannel(value *AppleDeploymentChannel)()
    SetIntendedPurpose(value *IntendedPurpose)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
}
