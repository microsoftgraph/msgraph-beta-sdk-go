package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationInfo 
type AttackSimulationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the date time of the attack simulation.
    attackSimDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the duration (in time) for the attack simulation
    attackSimDurationTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Specifies the activity id for the attack simulation.
    attackSimId *string
    // Specifies the user id of the user who got the attack simulation email
    attackSimUserId *string
    // The OdataType property
    odataType *string
}
// NewAttackSimulationInfo instantiates a new attackSimulationInfo and sets the default values.
func NewAttackSimulationInfo()(*AttackSimulationInfo) {
    m := &AttackSimulationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.security.attackSimulationInfo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttackSimulationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttackSimDateTime gets the attackSimDateTime property value. Specifies the date time of the attack simulation.
func (m *AttackSimulationInfo) GetAttackSimDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.attackSimDateTime
}
// GetAttackSimDurationTime gets the attackSimDurationTime property value. Specifies the duration (in time) for the attack simulation
func (m *AttackSimulationInfo) GetAttackSimDurationTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.attackSimDurationTime
}
// GetAttackSimId gets the attackSimId property value. Specifies the activity id for the attack simulation.
func (m *AttackSimulationInfo) GetAttackSimId()(*string) {
    return m.attackSimId
}
// GetAttackSimUserId gets the attackSimUserId property value. Specifies the user id of the user who got the attack simulation email
func (m *AttackSimulationInfo) GetAttackSimUserId()(*string) {
    return m.attackSimUserId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attackSimDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAttackSimDateTime)
    res["attackSimDurationTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetAttackSimDurationTime)
    res["attackSimId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAttackSimId)
    res["attackSimUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAttackSimUserId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationInfo) GetOdataType()(*string) {
    return m.odataType
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetAttackSimDateTime sets the attackSimDateTime property value. Specifies the date time of the attack simulation.
func (m *AttackSimulationInfo) SetAttackSimDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.attackSimDateTime = value
}
// SetAttackSimDurationTime sets the attackSimDurationTime property value. Specifies the duration (in time) for the attack simulation
func (m *AttackSimulationInfo) SetAttackSimDurationTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.attackSimDurationTime = value
}
// SetAttackSimId sets the attackSimId property value. Specifies the activity id for the attack simulation.
func (m *AttackSimulationInfo) SetAttackSimId(value *string)() {
    m.attackSimId = value
}
// SetAttackSimUserId sets the attackSimUserId property value. Specifies the user id of the user who got the attack simulation email
func (m *AttackSimulationInfo) SetAttackSimUserId(value *string)() {
    m.attackSimUserId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationInfo) SetOdataType(value *string)() {
    m.odataType = value
}
