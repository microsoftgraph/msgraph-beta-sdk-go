package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeamworkSoftwareUpdateHealth struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkSoftwareUpdateHealth instantiates a new TeamworkSoftwareUpdateHealth and sets the default values.
func NewTeamworkSoftwareUpdateHealth()(*TeamworkSoftwareUpdateHealth) {
    m := &TeamworkSoftwareUpdateHealth{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkSoftwareUpdateHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSoftwareUpdateHealth(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamworkSoftwareUpdateHealth) GetAdditionalData()(map[string]any) {
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
// GetAdminAgentSoftwareUpdateStatus gets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetAdminAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("adminAgentSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TeamworkSoftwareUpdateHealth) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompanyPortalSoftwareUpdateStatus gets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetCompanyPortalSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("companyPortalSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkSoftwareUpdateHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminAgentSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["companyPortalSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["firmwareSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
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
    res["operatingSystemSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["partnerAgentSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    res["teamsClientSoftwareUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareUpdateStatus(val.(TeamworkSoftwareUpdateStatusable))
        }
        return nil
    }
    return res
}
// GetFirmwareSoftwareUpdateStatus gets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetFirmwareSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("firmwareSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamworkSoftwareUpdateHealth) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystemSoftwareUpdateStatus gets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetOperatingSystemSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("operatingSystemSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// GetPartnerAgentSoftwareUpdateStatus gets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetPartnerAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("partnerAgentSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// GetTeamsClientSoftwareUpdateStatus gets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
// returns a TeamworkSoftwareUpdateStatusable when successful
func (m *TeamworkSoftwareUpdateHealth) GetTeamsClientSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable) {
    val, err := m.GetBackingStore().Get("teamsClientSoftwareUpdateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkSoftwareUpdateStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkSoftwareUpdateHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("adminAgentSoftwareUpdateStatus", m.GetAdminAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("companyPortalSoftwareUpdateStatus", m.GetCompanyPortalSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("firmwareSoftwareUpdateStatus", m.GetFirmwareSoftwareUpdateStatus())
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
        err := writer.WriteObjectValue("operatingSystemSoftwareUpdateStatus", m.GetOperatingSystemSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("partnerAgentSoftwareUpdateStatus", m.GetPartnerAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsClientSoftwareUpdateStatus", m.GetTeamsClientSoftwareUpdateStatus())
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
func (m *TeamworkSoftwareUpdateHealth) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminAgentSoftwareUpdateStatus sets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
func (m *TeamworkSoftwareUpdateHealth) SetAdminAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("adminAgentSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkSoftwareUpdateHealth) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompanyPortalSoftwareUpdateStatus sets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
func (m *TeamworkSoftwareUpdateHealth) SetCompanyPortalSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("companyPortalSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetFirmwareSoftwareUpdateStatus sets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
func (m *TeamworkSoftwareUpdateHealth) SetFirmwareSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("firmwareSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSoftwareUpdateHealth) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemSoftwareUpdateStatus sets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
func (m *TeamworkSoftwareUpdateHealth) SetOperatingSystemSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("operatingSystemSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerAgentSoftwareUpdateStatus sets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
func (m *TeamworkSoftwareUpdateHealth) SetPartnerAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("partnerAgentSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsClientSoftwareUpdateStatus sets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
func (m *TeamworkSoftwareUpdateHealth) SetTeamsClientSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)() {
    err := m.GetBackingStore().Set("teamsClientSoftwareUpdateStatus", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkSoftwareUpdateHealthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompanyPortalSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetFirmwareSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetOdataType()(*string)
    GetOperatingSystemSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetPartnerAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetTeamsClientSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    SetAdminAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompanyPortalSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetFirmwareSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetOdataType(value *string)()
    SetOperatingSystemSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetPartnerAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetTeamsClientSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
}
