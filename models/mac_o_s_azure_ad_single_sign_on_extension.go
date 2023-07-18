package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSAzureAdSingleSignOnExtension represents an Azure AD-type Single Sign-On extension profile for macOS devices.
type MacOSAzureAdSingleSignOnExtension struct {
    MacOSSingleSignOnExtension
}
// NewMacOSAzureAdSingleSignOnExtension instantiates a new macOSAzureAdSingleSignOnExtension and sets the default values.
func NewMacOSAzureAdSingleSignOnExtension()(*MacOSAzureAdSingleSignOnExtension) {
    m := &MacOSAzureAdSingleSignOnExtension{
        MacOSSingleSignOnExtension: *NewMacOSSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.macOSAzureAdSingleSignOnExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSAzureAdSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSAzureAdSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSAzureAdSingleSignOnExtension(), nil
}
// GetBundleIdAccessControlList gets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *MacOSAzureAdSingleSignOnExtension) GetBundleIdAccessControlList()([]string) {
    val, err := m.GetBackingStore().Get("bundleIdAccessControlList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSAzureAdSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    val, err := m.GetBackingStore().Get("configurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyTypedValuePairable)
    }
    return nil
}
// GetEnableSharedDeviceMode gets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *MacOSAzureAdSingleSignOnExtension) GetEnableSharedDeviceMode()(*bool) {
    val, err := m.GetBackingStore().Get("enableSharedDeviceMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSAzureAdSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSSingleSignOnExtension.GetFieldDeserializers()
    res["bundleIdAccessControlList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBundleIdAccessControlList(res)
        }
        return nil
    }
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
    res["enableSharedDeviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSharedDeviceMode(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSAzureAdSingleSignOnExtension) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSAzureAdSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBundleIdAccessControlList() != nil {
        err = writer.WriteCollectionOfStringValues("bundleIdAccessControlList", m.GetBundleIdAccessControlList())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteBoolValue("enableSharedDeviceMode", m.GetEnableSharedDeviceMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleIdAccessControlList sets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *MacOSAzureAdSingleSignOnExtension) SetBundleIdAccessControlList(value []string)() {
    err := m.GetBackingStore().Set("bundleIdAccessControlList", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSAzureAdSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    err := m.GetBackingStore().Set("configurations", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSharedDeviceMode sets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *MacOSAzureAdSingleSignOnExtension) SetEnableSharedDeviceMode(value *bool)() {
    err := m.GetBackingStore().Set("enableSharedDeviceMode", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSAzureAdSingleSignOnExtension) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// MacOSAzureAdSingleSignOnExtensionable 
type MacOSAzureAdSingleSignOnExtensionable interface {
    MacOSSingleSignOnExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBundleIdAccessControlList()([]string)
    GetConfigurations()([]KeyTypedValuePairable)
    GetEnableSharedDeviceMode()(*bool)
    GetOdataType()(*string)
    SetBundleIdAccessControlList(value []string)()
    SetConfigurations(value []KeyTypedValuePairable)()
    SetEnableSharedDeviceMode(value *bool)()
    SetOdataType(value *string)()
}
