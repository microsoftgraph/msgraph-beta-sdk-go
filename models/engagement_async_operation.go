package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EngagementAsyncOperation 
type EngagementAsyncOperation struct {
    LongRunningOperation
}
// NewEngagementAsyncOperation instantiates a new engagementAsyncOperation and sets the default values.
func NewEngagementAsyncOperation()(*EngagementAsyncOperation) {
    m := &EngagementAsyncOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// CreateEngagementAsyncOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEngagementAsyncOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEngagementAsyncOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EngagementAsyncOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEngagementAsyncOperation_operationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*EngagementAsyncOperation_operationType))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetOperationType gets the operationType property value. The type of the long-running operation. The possible values are: createCommunity, unknownFutureValue.
func (m *EngagementAsyncOperation) GetOperationType()(*EngagementAsyncOperation_operationType) {
    val, err := m.GetBackingStore().Get("operationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EngagementAsyncOperation_operationType)
    }
    return nil
}
// GetResourceId gets the resourceId property value. The ID of the object created or modified as a result of this async operation.
func (m *EngagementAsyncOperation) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EngagementAsyncOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOperationType sets the operationType property value. The type of the long-running operation. The possible values are: createCommunity, unknownFutureValue.
func (m *EngagementAsyncOperation) SetOperationType(value *EngagementAsyncOperation_operationType)() {
    err := m.GetBackingStore().Set("operationType", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. The ID of the object created or modified as a result of this async operation.
func (m *EngagementAsyncOperation) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// EngagementAsyncOperationable 
type EngagementAsyncOperationable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOperationType()(*EngagementAsyncOperation_operationType)
    GetResourceId()(*string)
    SetOperationType(value *EngagementAsyncOperation_operationType)()
    SetResourceId(value *string)()
}
