package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BusinessScenarioProperties 
type BusinessScenarioProperties struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBusinessScenarioProperties instantiates a new businessScenarioProperties and sets the default values.
func NewBusinessScenarioProperties()(*BusinessScenarioProperties) {
    m := &BusinessScenarioProperties{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBusinessScenarioPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioProperties(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessScenarioProperties) GetAdditionalData()(map[string]any) {
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
func (m *BusinessScenarioProperties) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExternalBucketId gets the externalBucketId property value. The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be placed in the corresponding plannerBucket in the target plan. Required.
func (m *BusinessScenarioProperties) GetExternalBucketId()(*string) {
    val, err := m.GetBackingStore().Get("externalBucketId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalContextId gets the externalContextId property value. The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by their externalContextId. Optional.
func (m *BusinessScenarioProperties) GetExternalContextId()(*string) {
    val, err := m.GetBackingStore().Get("externalContextId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalObjectId gets the externalObjectId property value. Application-specific identifier for the task. Every task for the same scenario must have a unique identifier specified for this property. Required.
func (m *BusinessScenarioProperties) GetExternalObjectId()(*string) {
    val, err := m.GetBackingStore().Get("externalObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalObjectVersion gets the externalObjectVersion property value. Application-specific version of the task. Optional.
func (m *BusinessScenarioProperties) GetExternalObjectVersion()(*string) {
    val, err := m.GetBackingStore().Get("externalObjectVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalBucketId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalBucketId(val)
        }
        return nil
    }
    res["externalContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalContextId(val)
        }
        return nil
    }
    res["externalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectId(val)
        }
        return nil
    }
    res["externalObjectVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectVersion(val)
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
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BusinessScenarioProperties) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebUrl gets the webUrl property value. The URL to the application-specific experience for this task. Optional.
func (m *BusinessScenarioProperties) GetWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("webUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BusinessScenarioProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalBucketId", m.GetExternalBucketId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalContextId", m.GetExternalContextId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalObjectId", m.GetExternalObjectId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalObjectVersion", m.GetExternalObjectVersion())
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
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *BusinessScenarioProperties) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BusinessScenarioProperties) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExternalBucketId sets the externalBucketId property value. The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be placed in the corresponding plannerBucket in the target plan. Required.
func (m *BusinessScenarioProperties) SetExternalBucketId(value *string)() {
    err := m.GetBackingStore().Set("externalBucketId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalContextId sets the externalContextId property value. The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by their externalContextId. Optional.
func (m *BusinessScenarioProperties) SetExternalContextId(value *string)() {
    err := m.GetBackingStore().Set("externalContextId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalObjectId sets the externalObjectId property value. Application-specific identifier for the task. Every task for the same scenario must have a unique identifier specified for this property. Required.
func (m *BusinessScenarioProperties) SetExternalObjectId(value *string)() {
    err := m.GetBackingStore().Set("externalObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalObjectVersion sets the externalObjectVersion property value. Application-specific version of the task. Optional.
func (m *BusinessScenarioProperties) SetExternalObjectVersion(value *string)() {
    err := m.GetBackingStore().Set("externalObjectVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BusinessScenarioProperties) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetWebUrl sets the webUrl property value. The URL to the application-specific experience for this task. Optional.
func (m *BusinessScenarioProperties) SetWebUrl(value *string)() {
    err := m.GetBackingStore().Set("webUrl", value)
    if err != nil {
        panic(err)
    }
}
// BusinessScenarioPropertiesable 
type BusinessScenarioPropertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExternalBucketId()(*string)
    GetExternalContextId()(*string)
    GetExternalObjectId()(*string)
    GetExternalObjectVersion()(*string)
    GetOdataType()(*string)
    GetWebUrl()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExternalBucketId(value *string)()
    SetExternalContextId(value *string)()
    SetExternalObjectId(value *string)()
    SetExternalObjectVersion(value *string)()
    SetOdataType(value *string)()
    SetWebUrl(value *string)()
}
