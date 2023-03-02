package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PayloadByFilter this entity represents a single payload with requested assignment filter Id
type PayloadByFilter struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPayloadByFilter instantiates a new payloadByFilter and sets the default values.
func NewPayloadByFilter()(*PayloadByFilter) {
    m := &PayloadByFilter{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePayloadByFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadByFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadByFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadByFilter) GetAdditionalData()(map[string]any) {
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
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *PayloadByFilter) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    val, err := m.GetBackingStore().Get("assignmentFilterType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAndAppManagementAssignmentFilterType)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PayloadByFilter) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadByFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
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
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["payloadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssociatedAssignmentPayloadType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadType(val.(*AssociatedAssignmentPayloadType))
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The Azure AD security group ID
func (m *PayloadByFilter) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PayloadByFilter) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadId gets the payloadId property value. The policy identifier
func (m *PayloadByFilter) GetPayloadId()(*string) {
    val, err := m.GetBackingStore().Get("payloadId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadType gets the payloadType property value. This enum represents associated assignment payload type
func (m *PayloadByFilter) GetPayloadType()(*AssociatedAssignmentPayloadType) {
    val, err := m.GetBackingStore().Get("payloadType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AssociatedAssignmentPayloadType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PayloadByFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadType() != nil {
        cast := (*m.GetPayloadType()).String()
        err := writer.WriteStringValue("payloadType", &cast)
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
func (m *PayloadByFilter) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *PayloadByFilter) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    err := m.GetBackingStore().Set("assignmentFilterType", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PayloadByFilter) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetGroupId sets the groupId property value. The Azure AD security group ID
func (m *PayloadByFilter) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PayloadByFilter) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadId sets the payloadId property value. The policy identifier
func (m *PayloadByFilter) SetPayloadId(value *string)() {
    err := m.GetBackingStore().Set("payloadId", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadType sets the payloadType property value. This enum represents associated assignment payload type
func (m *PayloadByFilter) SetPayloadType(value *AssociatedAssignmentPayloadType)() {
    err := m.GetBackingStore().Set("payloadType", value)
    if err != nil {
        panic(err)
    }
}
// PayloadByFilterable 
type PayloadByFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetGroupId()(*string)
    GetOdataType()(*string)
    GetPayloadId()(*string)
    GetPayloadType()(*AssociatedAssignmentPayloadType)
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetGroupId(value *string)()
    SetOdataType(value *string)()
    SetPayloadId(value *string)()
    SetPayloadType(value *AssociatedAssignmentPayloadType)()
}
