package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectionRule 
type DetectionRule struct {
    ProtectionRule
}
// NewDetectionRule instantiates a new detectionRule and sets the default values.
func NewDetectionRule()(*DetectionRule) {
    m := &DetectionRule{
        ProtectionRule: *NewProtectionRule(),
    }
    odataTypeValue := "#microsoft.graph.security.detectionRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDetectionRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectionRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetectionRule(), nil
}
// GetDetectionAction gets the detectionAction property value. The detectionAction property
func (m *DetectionRule) GetDetectionAction()(DetectionActionable) {
    val, err := m.GetBackingStore().Get("detectionAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DetectionActionable)
    }
    return nil
}
// GetDetectorId gets the detectorId property value. The detectorId property
func (m *DetectionRule) GetDetectorId()(*string) {
    val, err := m.GetBackingStore().Get("detectorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectionRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionRule.GetFieldDeserializers()
    res["detectionAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionAction(val.(DetectionActionable))
        }
        return nil
    }
    res["detectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectorId(val)
        }
        return nil
    }
    res["lastRunDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRunDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDetails(val.(RunDetailsable))
        }
        return nil
    }
    res["queryCondition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateQueryConditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryCondition(val.(QueryConditionable))
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRuleScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RuleScheduleable))
        }
        return nil
    }
    return res
}
// GetLastRunDetails gets the lastRunDetails property value. The lastRunDetails property
func (m *DetectionRule) GetLastRunDetails()(RunDetailsable) {
    val, err := m.GetBackingStore().Get("lastRunDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RunDetailsable)
    }
    return nil
}
// GetQueryCondition gets the queryCondition property value. The queryCondition property
func (m *DetectionRule) GetQueryCondition()(QueryConditionable) {
    val, err := m.GetBackingStore().Get("queryCondition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(QueryConditionable)
    }
    return nil
}
// GetSchedule gets the schedule property value. The schedule property
func (m *DetectionRule) GetSchedule()(RuleScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RuleScheduleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DetectionRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("detectionAction", m.GetDetectionAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("detectorId", m.GetDetectorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastRunDetails", m.GetLastRunDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("queryCondition", m.GetQueryCondition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionAction sets the detectionAction property value. The detectionAction property
func (m *DetectionRule) SetDetectionAction(value DetectionActionable)() {
    err := m.GetBackingStore().Set("detectionAction", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectorId sets the detectorId property value. The detectorId property
func (m *DetectionRule) SetDetectorId(value *string)() {
    err := m.GetBackingStore().Set("detectorId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRunDetails sets the lastRunDetails property value. The lastRunDetails property
func (m *DetectionRule) SetLastRunDetails(value RunDetailsable)() {
    err := m.GetBackingStore().Set("lastRunDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetQueryCondition sets the queryCondition property value. The queryCondition property
func (m *DetectionRule) SetQueryCondition(value QueryConditionable)() {
    err := m.GetBackingStore().Set("queryCondition", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. The schedule property
func (m *DetectionRule) SetSchedule(value RuleScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// DetectionRuleable 
type DetectionRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionRuleable
    GetDetectionAction()(DetectionActionable)
    GetDetectorId()(*string)
    GetLastRunDetails()(RunDetailsable)
    GetQueryCondition()(QueryConditionable)
    GetSchedule()(RuleScheduleable)
    SetDetectionAction(value DetectionActionable)()
    SetDetectorId(value *string)()
    SetLastRunDetails(value RunDetailsable)()
    SetQueryCondition(value QueryConditionable)()
    SetSchedule(value RuleScheduleable)()
}
