package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementTroubleshootingErrorDetails object containing detailed information about the error and its remediation.
type DeviceManagementTroubleshootingErrorDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementTroubleshootingErrorDetails instantiates a new deviceManagementTroubleshootingErrorDetails and sets the default values.
func NewDeviceManagementTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    m := &DeviceManagementTroubleshootingErrorDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTroubleshootingErrorDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorDetails) GetAdditionalData()(map[string]any) {
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
func (m *DeviceManagementTroubleshootingErrorDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetContext gets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetContext()(*string) {
    val, err := m.GetBackingStore().Get("context")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailure gets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailure()(*string) {
    val, err := m.GetBackingStore().Get("failure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailureDetails gets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) GetFailureDetails()(*string) {
    val, err := m.GetBackingStore().Get("failureDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingErrorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    res["failure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailure(val)
        }
        return nil
    }
    res["failureDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureDetails(val)
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
    res["remediation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediation(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingErrorResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingErrorResourceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTroubleshootingErrorResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediation gets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) GetRemediation()(*string) {
    val, err := m.GetBackingStore().Get("remediation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResources gets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) GetResources()([]DeviceManagementTroubleshootingErrorResourceable) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTroubleshootingErrorResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingErrorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failure", m.GetFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failureDetails", m.GetFailureDetails())
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
        err := writer.WriteStringValue("remediation", m.GetRemediation())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("resources", cast)
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
func (m *DeviceManagementTroubleshootingErrorDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceManagementTroubleshootingErrorDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetContext sets the context property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetContext(value *string)() {
    err := m.GetBackingStore().Set("context", value)
    if err != nil {
        panic(err)
    }
}
// SetFailure sets the failure property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailure(value *string)() {
    err := m.GetBackingStore().Set("failure", value)
    if err != nil {
        panic(err)
    }
}
// SetFailureDetails sets the failureDetails property value. The detailed description of what went wrong.
func (m *DeviceManagementTroubleshootingErrorDetails) SetFailureDetails(value *string)() {
    err := m.GetBackingStore().Set("failureDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediation sets the remediation property value. The detailed description of how to remediate this issue.
func (m *DeviceManagementTroubleshootingErrorDetails) SetRemediation(value *string)() {
    err := m.GetBackingStore().Set("remediation", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. Links to helpful documentation about this failure.
func (m *DeviceManagementTroubleshootingErrorDetails) SetResources(value []DeviceManagementTroubleshootingErrorResourceable)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementTroubleshootingErrorDetailsable 
type DeviceManagementTroubleshootingErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetContext()(*string)
    GetFailure()(*string)
    GetFailureDetails()(*string)
    GetOdataType()(*string)
    GetRemediation()(*string)
    GetResources()([]DeviceManagementTroubleshootingErrorResourceable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetContext(value *string)()
    SetFailure(value *string)()
    SetFailureDetails(value *string)()
    SetOdataType(value *string)()
    SetRemediation(value *string)()
    SetResources(value []DeviceManagementTroubleshootingErrorResourceable)()
}
