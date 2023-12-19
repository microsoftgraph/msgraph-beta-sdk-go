package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CustomExtensionHandlerInstance 
type CustomExtensionHandlerInstance struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomExtensionHandlerInstance instantiates a new customExtensionHandlerInstance and sets the default values.
func NewCustomExtensionHandlerInstance()(*CustomExtensionHandlerInstance) {
    m := &CustomExtensionHandlerInstance{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomExtensionHandlerInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionHandlerInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionHandlerInstance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionHandlerInstance) GetAdditionalData()(map[string]any) {
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
func (m *CustomExtensionHandlerInstance) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCustomExtensionId gets the customExtensionId property value. Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
func (m *CustomExtensionHandlerInstance) GetCustomExtensionId()(*string) {
    val, err := m.GetBackingStore().Get("customExtensionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalCorrelationId gets the externalCorrelationId property value. The unique run ID for the logic app.
func (m *CustomExtensionHandlerInstance) GetExternalCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("externalCorrelationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandlerInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customExtensionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionId(val)
        }
        return nil
    }
    res["externalCorrelationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalCorrelationId(val)
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
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomExtensionHandlerInstance_stage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*CustomExtensionHandlerInstance_stage))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomExtensionHandlerInstance_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CustomExtensionHandlerInstance_status))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomExtensionHandlerInstance) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStage gets the stage property value. Indicates the stage of the request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) GetStage()(*CustomExtensionHandlerInstance_stage) {
    val, err := m.GetBackingStore().Get("stage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CustomExtensionHandlerInstance_stage)
    }
    return nil
}
// GetStatus gets the status property value. Status of the request to run the access package custom extension workflow that is associated with the logic app. The possible values are: requestSent, requestReceived, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) GetStatus()(*CustomExtensionHandlerInstance_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CustomExtensionHandlerInstance_status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomExtensionHandlerInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customExtensionId", m.GetCustomExtensionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalCorrelationId", m.GetExternalCorrelationId())
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
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CustomExtensionHandlerInstance) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomExtensionHandlerInstance) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCustomExtensionId sets the customExtensionId property value. Identifier of the customAccessPackageWorkflowExtension triggered at this instance.
func (m *CustomExtensionHandlerInstance) SetCustomExtensionId(value *string)() {
    err := m.GetBackingStore().Set("customExtensionId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalCorrelationId sets the externalCorrelationId property value. The unique run ID for the logic app.
func (m *CustomExtensionHandlerInstance) SetExternalCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("externalCorrelationId", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionHandlerInstance) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStage sets the stage property value. Indicates the stage of the request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) SetStage(value *CustomExtensionHandlerInstance_stage)() {
    err := m.GetBackingStore().Set("stage", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Status of the request to run the access package custom extension workflow that is associated with the logic app. The possible values are: requestSent, requestReceived, unknownFutureValue.
func (m *CustomExtensionHandlerInstance) SetStatus(value *CustomExtensionHandlerInstance_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// CustomExtensionHandlerInstanceable 
type CustomExtensionHandlerInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCustomExtensionId()(*string)
    GetExternalCorrelationId()(*string)
    GetOdataType()(*string)
    GetStage()(*CustomExtensionHandlerInstance_stage)
    GetStatus()(*CustomExtensionHandlerInstance_status)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCustomExtensionId(value *string)()
    SetExternalCorrelationId(value *string)()
    SetOdataType(value *string)()
    SetStage(value *CustomExtensionHandlerInstance_stage)()
    SetStatus(value *CustomExtensionHandlerInstance_status)()
}
