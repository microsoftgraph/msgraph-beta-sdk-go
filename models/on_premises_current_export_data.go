package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnPremisesCurrentExportData 
type OnPremisesCurrentExportData struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnPremisesCurrentExportData instantiates a new onPremisesCurrentExportData and sets the default values.
func NewOnPremisesCurrentExportData()(*OnPremisesCurrentExportData) {
    m := &OnPremisesCurrentExportData{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesCurrentExportDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesCurrentExportDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesCurrentExportData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesCurrentExportData) GetAdditionalData()(map[string]any) {
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
func (m *OnPremisesCurrentExportData) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClientMachineName gets the clientMachineName property value. The name of the onPremises client machine that ran the last export.
func (m *OnPremisesCurrentExportData) GetClientMachineName()(*string) {
    val, err := m.GetBackingStore().Get("clientMachineName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesCurrentExportData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientMachineName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientMachineName(val)
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
    res["pendingObjectsAddition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingObjectsAddition(val)
        }
        return nil
    }
    res["pendingObjectsDeletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingObjectsDeletion(val)
        }
        return nil
    }
    res["pendingObjectsUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingObjectsUpdate(val)
        }
        return nil
    }
    res["serviceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccount(val)
        }
        return nil
    }
    res["successfulLinksProvisioningCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulLinksProvisioningCount(val)
        }
        return nil
    }
    res["successfulObjectsProvisioningCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulObjectsProvisioningCount(val)
        }
        return nil
    }
    res["totalConnectorSpaceObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalConnectorSpaceObjects(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesCurrentExportData) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPendingObjectsAddition gets the pendingObjectsAddition property value. The count of pending adds from on-premises directory.
func (m *OnPremisesCurrentExportData) GetPendingObjectsAddition()(*int32) {
    val, err := m.GetBackingStore().Get("pendingObjectsAddition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPendingObjectsDeletion gets the pendingObjectsDeletion property value. The count of pending deletes from on-premises directory.
func (m *OnPremisesCurrentExportData) GetPendingObjectsDeletion()(*int32) {
    val, err := m.GetBackingStore().Get("pendingObjectsDeletion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPendingObjectsUpdate gets the pendingObjectsUpdate property value. The count of pending updates from on-premises directory.
func (m *OnPremisesCurrentExportData) GetPendingObjectsUpdate()(*int32) {
    val, err := m.GetBackingStore().Get("pendingObjectsUpdate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetServiceAccount gets the serviceAccount property value. The name of the dirsync service account that is configured to connect to the directory.
func (m *OnPremisesCurrentExportData) GetServiceAccount()(*string) {
    val, err := m.GetBackingStore().Get("serviceAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessfulLinksProvisioningCount gets the successfulLinksProvisioningCount property value. The count of updated links during the current directory sync export run.
func (m *OnPremisesCurrentExportData) GetSuccessfulLinksProvisioningCount()(*int64) {
    val, err := m.GetBackingStore().Get("successfulLinksProvisioningCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSuccessfulObjectsProvisioningCount gets the successfulObjectsProvisioningCount property value. The count of objects that were successfully provisioned during the current directory sync export run.
func (m *OnPremisesCurrentExportData) GetSuccessfulObjectsProvisioningCount()(*int32) {
    val, err := m.GetBackingStore().Get("successfulObjectsProvisioningCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalConnectorSpaceObjects gets the totalConnectorSpaceObjects property value. The total number of objects in the AAD Connector Space.
func (m *OnPremisesCurrentExportData) GetTotalConnectorSpaceObjects()(*int32) {
    val, err := m.GetBackingStore().Get("totalConnectorSpaceObjects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPremisesCurrentExportData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientMachineName", m.GetClientMachineName())
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
        err := writer.WriteInt32Value("pendingObjectsAddition", m.GetPendingObjectsAddition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingObjectsDeletion", m.GetPendingObjectsDeletion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingObjectsUpdate", m.GetPendingObjectsUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceAccount", m.GetServiceAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("successfulLinksProvisioningCount", m.GetSuccessfulLinksProvisioningCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("successfulObjectsProvisioningCount", m.GetSuccessfulObjectsProvisioningCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalConnectorSpaceObjects", m.GetTotalConnectorSpaceObjects())
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
func (m *OnPremisesCurrentExportData) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OnPremisesCurrentExportData) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClientMachineName sets the clientMachineName property value. The name of the onPremises client machine that ran the last export.
func (m *OnPremisesCurrentExportData) SetClientMachineName(value *string)() {
    err := m.GetBackingStore().Set("clientMachineName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesCurrentExportData) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingObjectsAddition sets the pendingObjectsAddition property value. The count of pending adds from on-premises directory.
func (m *OnPremisesCurrentExportData) SetPendingObjectsAddition(value *int32)() {
    err := m.GetBackingStore().Set("pendingObjectsAddition", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingObjectsDeletion sets the pendingObjectsDeletion property value. The count of pending deletes from on-premises directory.
func (m *OnPremisesCurrentExportData) SetPendingObjectsDeletion(value *int32)() {
    err := m.GetBackingStore().Set("pendingObjectsDeletion", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingObjectsUpdate sets the pendingObjectsUpdate property value. The count of pending updates from on-premises directory.
func (m *OnPremisesCurrentExportData) SetPendingObjectsUpdate(value *int32)() {
    err := m.GetBackingStore().Set("pendingObjectsUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceAccount sets the serviceAccount property value. The name of the dirsync service account that is configured to connect to the directory.
func (m *OnPremisesCurrentExportData) SetServiceAccount(value *string)() {
    err := m.GetBackingStore().Set("serviceAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulLinksProvisioningCount sets the successfulLinksProvisioningCount property value. The count of updated links during the current directory sync export run.
func (m *OnPremisesCurrentExportData) SetSuccessfulLinksProvisioningCount(value *int64)() {
    err := m.GetBackingStore().Set("successfulLinksProvisioningCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulObjectsProvisioningCount sets the successfulObjectsProvisioningCount property value. The count of objects that were successfully provisioned during the current directory sync export run.
func (m *OnPremisesCurrentExportData) SetSuccessfulObjectsProvisioningCount(value *int32)() {
    err := m.GetBackingStore().Set("successfulObjectsProvisioningCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalConnectorSpaceObjects sets the totalConnectorSpaceObjects property value. The total number of objects in the AAD Connector Space.
func (m *OnPremisesCurrentExportData) SetTotalConnectorSpaceObjects(value *int32)() {
    err := m.GetBackingStore().Set("totalConnectorSpaceObjects", value)
    if err != nil {
        panic(err)
    }
}
// OnPremisesCurrentExportDataable 
type OnPremisesCurrentExportDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClientMachineName()(*string)
    GetOdataType()(*string)
    GetPendingObjectsAddition()(*int32)
    GetPendingObjectsDeletion()(*int32)
    GetPendingObjectsUpdate()(*int32)
    GetServiceAccount()(*string)
    GetSuccessfulLinksProvisioningCount()(*int64)
    GetSuccessfulObjectsProvisioningCount()(*int32)
    GetTotalConnectorSpaceObjects()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClientMachineName(value *string)()
    SetOdataType(value *string)()
    SetPendingObjectsAddition(value *int32)()
    SetPendingObjectsDeletion(value *int32)()
    SetPendingObjectsUpdate(value *int32)()
    SetServiceAccount(value *string)()
    SetSuccessfulLinksProvisioningCount(value *int64)()
    SetSuccessfulObjectsProvisioningCount(value *int32)()
    SetTotalConnectorSpaceObjects(value *int32)()
}
