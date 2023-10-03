package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionHandler 
type CustomExtensionHandler struct {
    Entity
}
// NewCustomExtensionHandler instantiates a new customExtensionHandler and sets the default values.
func NewCustomExtensionHandler()(*CustomExtensionHandler) {
    m := &CustomExtensionHandler{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionHandler(), nil
}
// GetCustomExtension gets the customExtension property value. Indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
func (m *CustomExtensionHandler) GetCustomExtension()(CustomAccessPackageWorkflowExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomAccessPackageWorkflowExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(CustomAccessPackageWorkflowExtensionable))
        }
        return nil
    }
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    return res
}
// GetStage gets the stage property value. Indicates the stage of the access package assignment request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandler) GetStage()(*AccessPackageCustomExtensionStage) {
    val, err := m.GetBackingStore().Get("stage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessPackageCustomExtensionStage)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. Indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
func (m *CustomExtensionHandler) SetCustomExtension(value CustomAccessPackageWorkflowExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetStage sets the stage property value. Indicates the stage of the access package assignment request workflow when the access package custom extension runs. The possible values are: assignmentRequestCreated, assignmentRequestApproved, assignmentRequestGranted, assignmentRequestRemoved, assignmentFourteenDaysBeforeExpiration, assignmentOneDayBeforeExpiration, unknownFutureValue.
func (m *CustomExtensionHandler) SetStage(value *AccessPackageCustomExtensionStage)() {
    err := m.GetBackingStore().Set("stage", value)
    if err != nil {
        panic(err)
    }
}
// CustomExtensionHandlerable 
type CustomExtensionHandlerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtension()(CustomAccessPackageWorkflowExtensionable)
    GetStage()(*AccessPackageCustomExtensionStage)
    SetCustomExtension(value CustomAccessPackageWorkflowExtensionable)()
    SetStage(value *AccessPackageCustomExtensionStage)()
}
