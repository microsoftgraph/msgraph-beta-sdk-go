package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleEligibilitySchedule provides operations to manage the collection of accessReview entities.
type UnifiedRoleEligibilitySchedule struct {
    UnifiedRoleScheduleBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
    memberType *string
    // The schedule object of the eligible role assignment request.
    scheduleInfo RequestScheduleable
}
// NewUnifiedRoleEligibilitySchedule instantiates a new unifiedRoleEligibilitySchedule and sets the default values.
func NewUnifiedRoleEligibilitySchedule()(*UnifiedRoleEligibilitySchedule) {
    m := &UnifiedRoleEligibilitySchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleEligibilitySchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleEligibilitySchedule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleEligibilitySchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleBase.GetFieldDeserializers()
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val)
        }
        return nil
    }
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetMemberType gets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleEligibilitySchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// GetScheduleInfo gets the scheduleInfo property value. The schedule object of the eligible role assignment request.
func (m *UnifiedRoleEligibilitySchedule) GetScheduleInfo()(RequestScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleEligibilitySchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
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
func (m *UnifiedRoleEligibilitySchedule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMemberType sets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleEligibilitySchedule) SetMemberType(value *string)() {
    if m != nil {
        m.memberType = value
    }
}
// SetScheduleInfo sets the scheduleInfo property value. The schedule object of the eligible role assignment request.
func (m *UnifiedRoleEligibilitySchedule) SetScheduleInfo(value RequestScheduleable)() {
    if m != nil {
        m.scheduleInfo = value
    }
}
