package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationInfo 
type AttackSimulationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attackSimDateTime property
    attackSimDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The attackSimDurationTime property
    attackSimDurationTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The attackSimId property
    attackSimId *string
    // The attackSimUserId property
    attackSimUserId *string
}
// NewAttackSimulationInfo instantiates a new attackSimulationInfo and sets the default values.
func NewAttackSimulationInfo()(*AttackSimulationInfo) {
    m := &AttackSimulationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttackSimulationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttackSimDateTime gets the attackSimDateTime property value. The attackSimDateTime property
func (m *AttackSimulationInfo) GetAttackSimDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.attackSimDateTime
    }
}
// GetAttackSimDurationTime gets the attackSimDurationTime property value. The attackSimDurationTime property
func (m *AttackSimulationInfo) GetAttackSimDurationTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.attackSimDurationTime
    }
}
// GetAttackSimId gets the attackSimId property value. The attackSimId property
func (m *AttackSimulationInfo) GetAttackSimId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attackSimId
    }
}
// GetAttackSimUserId gets the attackSimUserId property value. The attackSimUserId property
func (m *AttackSimulationInfo) GetAttackSimUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attackSimUserId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attackSimDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimDateTime(val)
        }
        return nil
    }
    res["attackSimDurationTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimDurationTime(val)
        }
        return nil
    }
    res["attackSimId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimId(val)
        }
        return nil
    }
    res["attackSimUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimUserId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AttackSimulationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("attackSimDateTime", m.GetAttackSimDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("attackSimDurationTime", m.GetAttackSimDurationTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attackSimId", m.GetAttackSimId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attackSimUserId", m.GetAttackSimUserId())
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
func (m *AttackSimulationInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttackSimDateTime sets the attackSimDateTime property value. The attackSimDateTime property
func (m *AttackSimulationInfo) SetAttackSimDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.attackSimDateTime = value
    }
}
// SetAttackSimDurationTime sets the attackSimDurationTime property value. The attackSimDurationTime property
func (m *AttackSimulationInfo) SetAttackSimDurationTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.attackSimDurationTime = value
    }
}
// SetAttackSimId sets the attackSimId property value. The attackSimId property
func (m *AttackSimulationInfo) SetAttackSimId(value *string)() {
    if m != nil {
        m.attackSimId = value
    }
}
// SetAttackSimUserId sets the attackSimUserId property value. The attackSimUserId property
func (m *AttackSimulationInfo) SetAttackSimUserId(value *string)() {
    if m != nil {
        m.attackSimUserId = value
    }
}
