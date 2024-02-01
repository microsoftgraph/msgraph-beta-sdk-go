package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PrivateAccessDetails 
type PrivateAccessDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPrivateAccessDetails instantiates a new privateAccessDetails and sets the default values.
func NewPrivateAccessDetails()(*PrivateAccessDetails) {
    m := &PrivateAccessDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrivateAccessDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivateAccessDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivateAccessDetails(), nil
}
// GetAccessType gets the accessType property value. The accessType property
func (m *PrivateAccessDetails) GetAccessType()(*AccessType) {
    val, err := m.GetBackingStore().Get("accessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessType)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateAccessDetails) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PrivateAccessDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConnectionStatus gets the connectionStatus property value. The connectionStatus property
func (m *PrivateAccessDetails) GetConnectionStatus()(*ConnectionStatus) {
    val, err := m.GetBackingStore().Get("connectionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConnectionStatus)
    }
    return nil
}
// GetConnectorId gets the connectorId property value. The connectorId property
func (m *PrivateAccessDetails) GetConnectorId()(*string) {
    val, err := m.GetBackingStore().Get("connectorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectorIp gets the connectorIp property value. The connectorIp property
func (m *PrivateAccessDetails) GetConnectorIp()(*string) {
    val, err := m.GetBackingStore().Get("connectorIp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectorName gets the connectorName property value. The connectorName property
func (m *PrivateAccessDetails) GetConnectorName()(*string) {
    val, err := m.GetBackingStore().Get("connectorName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivateAccessDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*AccessType))
        }
        return nil
    }
    res["connectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionStatus(val.(*ConnectionStatus))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
        }
        return nil
    }
    res["connectorIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorIp(val)
        }
        return nil
    }
    res["connectorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorName(val)
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
    res["processingRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingRegion(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrivateAccessDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProcessingRegion gets the processingRegion property value. The processingRegion property
func (m *PrivateAccessDetails) GetProcessingRegion()(*string) {
    val, err := m.GetBackingStore().Get("processingRegion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivateAccessDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectionStatus() != nil {
        cast := (*m.GetConnectionStatus()).String()
        err := writer.WriteStringValue("connectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorId", m.GetConnectorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorIp", m.GetConnectorIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorName", m.GetConnectorName())
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
        err := writer.WriteStringValue("processingRegion", m.GetProcessingRegion())
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
// SetAccessType sets the accessType property value. The accessType property
func (m *PrivateAccessDetails) SetAccessType(value *AccessType)() {
    err := m.GetBackingStore().Set("accessType", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateAccessDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PrivateAccessDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConnectionStatus sets the connectionStatus property value. The connectionStatus property
func (m *PrivateAccessDetails) SetConnectionStatus(value *ConnectionStatus)() {
    err := m.GetBackingStore().Set("connectionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorId sets the connectorId property value. The connectorId property
func (m *PrivateAccessDetails) SetConnectorId(value *string)() {
    err := m.GetBackingStore().Set("connectorId", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorIp sets the connectorIp property value. The connectorIp property
func (m *PrivateAccessDetails) SetConnectorIp(value *string)() {
    err := m.GetBackingStore().Set("connectorIp", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorName sets the connectorName property value. The connectorName property
func (m *PrivateAccessDetails) SetConnectorName(value *string)() {
    err := m.GetBackingStore().Set("connectorName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrivateAccessDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessingRegion sets the processingRegion property value. The processingRegion property
func (m *PrivateAccessDetails) SetProcessingRegion(value *string)() {
    err := m.GetBackingStore().Set("processingRegion", value)
    if err != nil {
        panic(err)
    }
}
// PrivateAccessDetailsable 
type PrivateAccessDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessType()(*AccessType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConnectionStatus()(*ConnectionStatus)
    GetConnectorId()(*string)
    GetConnectorIp()(*string)
    GetConnectorName()(*string)
    GetOdataType()(*string)
    GetProcessingRegion()(*string)
    SetAccessType(value *AccessType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConnectionStatus(value *ConnectionStatus)()
    SetConnectorId(value *string)()
    SetConnectorIp(value *string)()
    SetConnectorName(value *string)()
    SetOdataType(value *string)()
    SetProcessingRegion(value *string)()
}
