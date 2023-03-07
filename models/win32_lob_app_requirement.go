package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// Win32LobAppRequirement base class to detect a Win32 App
type Win32LobAppRequirement struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWin32LobAppRequirement instantiates a new win32LobAppRequirement and sets the default values.
func NewWin32LobAppRequirement()(*Win32LobAppRequirement) {
    m := &Win32LobAppRequirement{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWin32LobAppRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                switch *mappingValue {
                    case "#microsoft.graph.win32LobAppFileSystemRequirement":
                        return NewWin32LobAppFileSystemRequirement(), nil
                    case "#microsoft.graph.win32LobAppPowerShellScriptRequirement":
                        return NewWin32LobAppPowerShellScriptRequirement(), nil
                    case "#microsoft.graph.win32LobAppRegistryRequirement":
                        return NewWin32LobAppRegistryRequirement(), nil
                }
            }
        }
    }
    return NewWin32LobAppRequirement(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppRequirement) GetAdditionalData()(map[string]any) {
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
func (m *Win32LobAppRequirement) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDetectionValue gets the detectionValue property value. The detection value
func (m *Win32LobAppRequirement) GetDetectionValue()(*string) {
    val, err := m.GetBackingStore().Get("detectionValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionValue(val)
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDetectionOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Win32LobAppDetectionOperator))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Win32LobAppRequirement) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperator gets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRequirement) GetOperator()(*Win32LobAppDetectionOperator) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppDetectionOperator)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobAppRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detectionValue", m.GetDetectionValue())
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
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
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
func (m *Win32LobAppRequirement) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *Win32LobAppRequirement) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDetectionValue sets the detectionValue property value. The detection value
func (m *Win32LobAppRequirement) SetDetectionValue(value *string)() {
    err := m.GetBackingStore().Set("detectionValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Win32LobAppRequirement) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRequirement) SetOperator(value *Win32LobAppDetectionOperator)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppRequirementable 
type Win32LobAppRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDetectionValue()(*string)
    GetOdataType()(*string)
    GetOperator()(*Win32LobAppDetectionOperator)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDetectionValue(value *string)()
    SetOdataType(value *string)()
    SetOperator(value *Win32LobAppDetectionOperator)()
}
