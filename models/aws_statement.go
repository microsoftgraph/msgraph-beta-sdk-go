package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AwsStatement 
type AwsStatement struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAwsStatement instantiates a new awsStatement and sets the default values.
func NewAwsStatement()(*AwsStatement) {
    m := &AwsStatement{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAwsStatementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsStatementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsStatement(), nil
}
// GetActions gets the actions property value. The actions property
func (m *AwsStatement) GetActions()([]string) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AwsStatement) GetAdditionalData()(map[string]any) {
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
func (m *AwsStatement) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCondition gets the condition property value. The condition property
func (m *AwsStatement) GetCondition()(AwsConditionable) {
    val, err := m.GetBackingStore().Get("condition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsConditionable)
    }
    return nil
}
// GetEffect gets the effect property value. The effect property
func (m *AwsStatement) GetEffect()(*AwsStatementEffect) {
    val, err := m.GetBackingStore().Get("effect")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsStatementEffect)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsStatement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetActions(res)
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsConditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val.(AwsConditionable))
        }
        return nil
    }
    res["effect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsStatementEffect)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEffect(val.(*AwsStatementEffect))
        }
        return nil
    }
    res["notActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNotActions(res)
        }
        return nil
    }
    res["notResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetNotResources(res)
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
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetResources(res)
        }
        return nil
    }
    res["statementId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatementId(val)
        }
        return nil
    }
    return res
}
// GetNotActions gets the notActions property value. The notActions property
func (m *AwsStatement) GetNotActions()([]string) {
    val, err := m.GetBackingStore().Get("notActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNotResources gets the notResources property value. The notResources property
func (m *AwsStatement) GetNotResources()([]string) {
    val, err := m.GetBackingStore().Get("notResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AwsStatement) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResources gets the resources property value. The resources property
func (m *AwsStatement) GetResources()([]string) {
    val, err := m.GetBackingStore().Get("resources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetStatementId gets the statementId property value. The statementId property
func (m *AwsStatement) GetStatementId()(*string) {
    val, err := m.GetBackingStore().Get("statementId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsStatement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        err := writer.WriteCollectionOfStringValues("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("condition", m.GetCondition())
        if err != nil {
            return err
        }
    }
    if m.GetEffect() != nil {
        cast := (*m.GetEffect()).String()
        err := writer.WriteStringValue("effect", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotActions() != nil {
        err := writer.WriteCollectionOfStringValues("notActions", m.GetNotActions())
        if err != nil {
            return err
        }
    }
    if m.GetNotResources() != nil {
        err := writer.WriteCollectionOfStringValues("notResources", m.GetNotResources())
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
    if m.GetResources() != nil {
        err := writer.WriteCollectionOfStringValues("resources", m.GetResources())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statementId", m.GetStatementId())
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
// SetActions sets the actions property value. The actions property
func (m *AwsStatement) SetActions(value []string)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AwsStatement) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AwsStatement) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCondition sets the condition property value. The condition property
func (m *AwsStatement) SetCondition(value AwsConditionable)() {
    err := m.GetBackingStore().Set("condition", value)
    if err != nil {
        panic(err)
    }
}
// SetEffect sets the effect property value. The effect property
func (m *AwsStatement) SetEffect(value *AwsStatementEffect)() {
    err := m.GetBackingStore().Set("effect", value)
    if err != nil {
        panic(err)
    }
}
// SetNotActions sets the notActions property value. The notActions property
func (m *AwsStatement) SetNotActions(value []string)() {
    err := m.GetBackingStore().Set("notActions", value)
    if err != nil {
        panic(err)
    }
}
// SetNotResources sets the notResources property value. The notResources property
func (m *AwsStatement) SetNotResources(value []string)() {
    err := m.GetBackingStore().Set("notResources", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AwsStatement) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResources sets the resources property value. The resources property
func (m *AwsStatement) SetResources(value []string)() {
    err := m.GetBackingStore().Set("resources", value)
    if err != nil {
        panic(err)
    }
}
// SetStatementId sets the statementId property value. The statementId property
func (m *AwsStatement) SetStatementId(value *string)() {
    err := m.GetBackingStore().Set("statementId", value)
    if err != nil {
        panic(err)
    }
}
// AwsStatementable 
type AwsStatementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCondition()(AwsConditionable)
    GetEffect()(*AwsStatementEffect)
    GetNotActions()([]string)
    GetNotResources()([]string)
    GetOdataType()(*string)
    GetResources()([]string)
    GetStatementId()(*string)
    SetActions(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCondition(value AwsConditionable)()
    SetEffect(value *AwsStatementEffect)()
    SetNotActions(value []string)()
    SetNotResources(value []string)()
    SetOdataType(value *string)()
    SetResources(value []string)()
    SetStatementId(value *string)()
}
