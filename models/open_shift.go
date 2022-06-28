package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenShift 
type OpenShift struct {
    ChangeTrackedEntity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // An unpublished open shift.
    draftOpenShift OpenShiftItemable
    // The isStagedForDeletion property
    isStagedForDeletion *bool
    // ID for the scheduling group that the open shift belongs to.
    schedulingGroupId *string
    // A published open shift.
    sharedOpenShift OpenShiftItemable
}
// NewOpenShift instantiates a new OpenShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOpenShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenShift(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OpenShift) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDraftOpenShift gets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) GetDraftOpenShift()(OpenShiftItemable) {
    if m == nil {
        return nil
    } else {
        return m.draftOpenShift
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftOpenShift(val.(OpenShiftItemable))
        }
        return nil
    }
    res["isStagedForDeletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStagedForDeletion(val)
        }
        return nil
    }
    res["schedulingGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["sharedOpenShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOpenShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedOpenShift(val.(OpenShiftItemable))
        }
        return nil
    }
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) GetIsStagedForDeletion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isStagedForDeletion
    }
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
// GetSharedOpenShift gets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) GetSharedOpenShift()(OpenShiftItemable) {
    if m == nil {
        return nil
    } else {
        return m.sharedOpenShift
    }
}
// Serialize serializes information the current object
func (m *OpenShift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftOpenShift", m.GetDraftOpenShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isStagedForDeletion", m.GetIsStagedForDeletion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedOpenShift", m.GetSharedOpenShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OpenShift) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDraftOpenShift sets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) SetDraftOpenShift(value OpenShiftItemable)() {
    if m != nil {
        m.draftOpenShift = value
    }
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) SetIsStagedForDeletion(value *bool)() {
    if m != nil {
        m.isStagedForDeletion = value
    }
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    if m != nil {
        m.schedulingGroupId = value
    }
}
// SetSharedOpenShift sets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) SetSharedOpenShift(value OpenShiftItemable)() {
    if m != nil {
        m.sharedOpenShift = value
    }
}
