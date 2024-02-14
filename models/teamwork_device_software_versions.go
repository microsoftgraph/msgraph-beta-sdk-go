package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeamworkDeviceSoftwareVersions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkDeviceSoftwareVersions instantiates a new TeamworkDeviceSoftwareVersions and sets the default values.
func NewTeamworkDeviceSoftwareVersions()(*TeamworkDeviceSoftwareVersions) {
    m := &TeamworkDeviceSoftwareVersions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceSoftwareVersions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamworkDeviceSoftwareVersions) GetAdditionalData()(map[string]any) {
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
// GetAdminAgentSoftwareVersion gets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
// returns a *string when successful
func (m *TeamworkDeviceSoftwareVersions) GetAdminAgentSoftwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("adminAgentSoftwareVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TeamworkDeviceSoftwareVersions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkDeviceSoftwareVersions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminAgentSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareVersion(val)
        }
        return nil
    }
    res["firmwareSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareVersion(val)
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
    res["operatingSystemSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareVersion(val)
        }
        return nil
    }
    res["partnerAgentSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareVersion(val)
        }
        return nil
    }
    res["teamsClientSoftwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareVersion(val)
        }
        return nil
    }
    return res
}
// GetFirmwareSoftwareVersion gets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
// returns a *string when successful
func (m *TeamworkDeviceSoftwareVersions) GetFirmwareSoftwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("firmwareSoftwareVersion")
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
func (m *TeamworkDeviceSoftwareVersions) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystemSoftwareVersion gets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
// returns a *string when successful
func (m *TeamworkDeviceSoftwareVersions) GetOperatingSystemSoftwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystemSoftwareVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPartnerAgentSoftwareVersion gets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
// returns a *string when successful
func (m *TeamworkDeviceSoftwareVersions) GetPartnerAgentSoftwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("partnerAgentSoftwareVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamsClientSoftwareVersion gets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
// returns a *string when successful
func (m *TeamworkDeviceSoftwareVersions) GetTeamsClientSoftwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("teamsClientSoftwareVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceSoftwareVersions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adminAgentSoftwareVersion", m.GetAdminAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareSoftwareVersion", m.GetFirmwareSoftwareVersion())
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
        err := writer.WriteStringValue("operatingSystemSoftwareVersion", m.GetOperatingSystemSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("partnerAgentSoftwareVersion", m.GetPartnerAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsClientSoftwareVersion", m.GetTeamsClientSoftwareVersion())
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
func (m *TeamworkDeviceSoftwareVersions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminAgentSoftwareVersion sets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetAdminAgentSoftwareVersion(value *string)() {
    err := m.GetBackingStore().Set("adminAgentSoftwareVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkDeviceSoftwareVersions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFirmwareSoftwareVersion sets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetFirmwareSoftwareVersion(value *string)() {
    err := m.GetBackingStore().Set("firmwareSoftwareVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDeviceSoftwareVersions) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemSoftwareVersion sets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
func (m *TeamworkDeviceSoftwareVersions) SetOperatingSystemSoftwareVersion(value *string)() {
    err := m.GetBackingStore().Set("operatingSystemSoftwareVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerAgentSoftwareVersion sets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetPartnerAgentSoftwareVersion(value *string)() {
    err := m.GetBackingStore().Set("partnerAgentSoftwareVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsClientSoftwareVersion sets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetTeamsClientSoftwareVersion(value *string)() {
    err := m.GetBackingStore().Set("teamsClientSoftwareVersion", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkDeviceSoftwareVersionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminAgentSoftwareVersion()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFirmwareSoftwareVersion()(*string)
    GetOdataType()(*string)
    GetOperatingSystemSoftwareVersion()(*string)
    GetPartnerAgentSoftwareVersion()(*string)
    GetTeamsClientSoftwareVersion()(*string)
    SetAdminAgentSoftwareVersion(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFirmwareSoftwareVersion(value *string)()
    SetOdataType(value *string)()
    SetOperatingSystemSoftwareVersion(value *string)()
    SetPartnerAgentSoftwareVersion(value *string)()
    SetTeamsClientSoftwareVersion(value *string)()
}
