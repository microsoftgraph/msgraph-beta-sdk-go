package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CloudPcPartnerAgentInstallResult 
type CloudPcPartnerAgentInstallResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcPartnerAgentInstallResult instantiates a new cloudPcPartnerAgentInstallResult and sets the default values.
func NewCloudPcPartnerAgentInstallResult()(*CloudPcPartnerAgentInstallResult) {
    m := &CloudPcPartnerAgentInstallResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcPartnerAgentInstallResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcPartnerAgentInstallResult) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *CloudPcPartnerAgentInstallResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcPartnerAgentInstallResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["installStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcPartnerAgentInstallStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallStatus(val.(*CloudPcPartnerAgentInstallStatus))
        }
        return nil
    }
    res["isThirdPartyPartner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsThirdPartyPartner(val)
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
    res["partnerAgentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcPartnerAgentName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentName(val.(*CloudPcPartnerAgentName))
        }
        return nil
    }
    res["retriable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetriable(val)
        }
        return nil
    }
    return res
}
// GetInstallStatus gets the installStatus property value. The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling, uninstallFailed and licensed. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) GetInstallStatus()(*CloudPcPartnerAgentInstallStatus) {
    val, err := m.GetBackingStore().Get("installStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcPartnerAgentInstallStatus)
    }
    return nil
}
// GetIsThirdPartyPartner gets the isThirdPartyPartner property value. Indicates if the partner agent is a third party. When 'TRUE' the agent is a third-party (non-Microsoft) agent and when 'FALSE' the agent is a Microsoft agent or isn't known.  The default value is 'FALSE'
func (m *CloudPcPartnerAgentInstallResult) GetIsThirdPartyPartner()(*bool) {
    val, err := m.GetBackingStore().Get("isThirdPartyPartner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcPartnerAgentInstallResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPartnerAgentName gets the partnerAgentName property value. The name of the partner agent, whether first party or third party. Possible values for third-party partners are Citrix and VMware. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) GetPartnerAgentName()(*CloudPcPartnerAgentName) {
    val, err := m.GetBackingStore().Get("partnerAgentName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcPartnerAgentName)
    }
    return nil
}
// GetRetriable gets the retriable property value. Indicates if the partner agent is a third party. When 'TRUE' the agent is a third-party (non-Microsoft) agent and when 'FALSE' the agent is a Microsoft agent or isn't known. The default value is 'FALSE'
func (m *CloudPcPartnerAgentInstallResult) GetRetriable()(*bool) {
    val, err := m.GetBackingStore().Get("retriable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcPartnerAgentInstallResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInstallStatus() != nil {
        cast := (*m.GetInstallStatus()).String()
        err := writer.WriteStringValue("installStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isThirdPartyPartner", m.GetIsThirdPartyPartner())
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
    if m.GetPartnerAgentName() != nil {
        cast := (*m.GetPartnerAgentName()).String()
        err := writer.WriteStringValue("partnerAgentName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retriable", m.GetRetriable())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcPartnerAgentInstallResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CloudPcPartnerAgentInstallResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetInstallStatus sets the installStatus property value. The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling, uninstallFailed and licensed. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) SetInstallStatus(value *CloudPcPartnerAgentInstallStatus)() {
    err := m.GetBackingStore().Set("installStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetIsThirdPartyPartner sets the isThirdPartyPartner property value. Indicates if the partner agent is a third party. When 'TRUE' the agent is a third-party (non-Microsoft) agent and when 'FALSE' the agent is a Microsoft agent or isn't known.  The default value is 'FALSE'
func (m *CloudPcPartnerAgentInstallResult) SetIsThirdPartyPartner(value *bool)() {
    err := m.GetBackingStore().Set("isThirdPartyPartner", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcPartnerAgentInstallResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerAgentName sets the partnerAgentName property value. The name of the partner agent, whether first party or third party. Possible values for third-party partners are Citrix and VMware. Read-Only.
func (m *CloudPcPartnerAgentInstallResult) SetPartnerAgentName(value *CloudPcPartnerAgentName)() {
    err := m.GetBackingStore().Set("partnerAgentName", value)
    if err != nil {
        panic(err)
    }
}
// SetRetriable sets the retriable property value. Indicates if the partner agent is a third party. When 'TRUE' the agent is a third-party (non-Microsoft) agent and when 'FALSE' the agent is a Microsoft agent or isn't known. The default value is 'FALSE'
func (m *CloudPcPartnerAgentInstallResult) SetRetriable(value *bool)() {
    err := m.GetBackingStore().Set("retriable", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcPartnerAgentInstallResultable 
type CloudPcPartnerAgentInstallResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetInstallStatus()(*CloudPcPartnerAgentInstallStatus)
    GetIsThirdPartyPartner()(*bool)
    GetOdataType()(*string)
    GetPartnerAgentName()(*CloudPcPartnerAgentName)
    GetRetriable()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetInstallStatus(value *CloudPcPartnerAgentInstallStatus)()
    SetIsThirdPartyPartner(value *bool)()
    SetOdataType(value *string)()
    SetPartnerAgentName(value *CloudPcPartnerAgentName)()
    SetRetriable(value *bool)()
}
