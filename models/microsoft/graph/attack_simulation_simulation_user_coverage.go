package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
type AttackSimulationSimulationUserCoverage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // User in an attack simulation and training campaign.
    attackSimulationUser AttackSimulationUserable;
    // Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
    clickCount *int32;
    // Number of compromising actions by the user in attack simulation and training campaigns.
    compromisedCount *int32;
    // Date and time of latest attack simulation and training campaign that the user was included in.
    latestSimulationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of attack simulation and training campaigns that the user was included in.
    simulationCount *int32;
}
// NewAttackSimulationSimulationUserCoverage instantiates a new attackSimulationSimulationUserCoverage and sets the default values.
func NewAttackSimulationSimulationUserCoverage()(*AttackSimulationSimulationUserCoverage) {
    m := &AttackSimulationSimulationUserCoverage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttackSimulationSimulationUserCoverageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationSimulationUserCoverageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttackSimulationSimulationUserCoverage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationSimulationUserCoverage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttackSimulationUser gets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationSimulationUserCoverage) GetAttackSimulationUser()(AttackSimulationUserable) {
    if m == nil {
        return nil
    } else {
        return m.attackSimulationUser
    }
}
// GetClickCount gets the clickCount property value. Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) GetClickCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.clickCount
    }
}
// GetCompromisedCount gets the compromisedCount property value. Number of compromising actions by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) GetCompromisedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compromisedCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationSimulationUserCoverage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attackSimulationUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulationUser(val.(AttackSimulationUserable))
        }
        return nil
    }
    res["clickCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickCount(val)
        }
        return nil
    }
    res["compromisedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedCount(val)
        }
        return nil
    }
    res["latestSimulationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestSimulationDateTime(val)
        }
        return nil
    }
    res["simulationCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationCount(val)
        }
        return nil
    }
    return res
}
// GetLatestSimulationDateTime gets the latestSimulationDateTime property value. Date and time of latest attack simulation and training campaign that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) GetLatestSimulationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.latestSimulationDateTime
    }
}
// GetSimulationCount gets the simulationCount property value. Number of attack simulation and training campaigns that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) GetSimulationCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.simulationCount
    }
}
func (m *AttackSimulationSimulationUserCoverage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttackSimulationSimulationUserCoverage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attackSimulationUser", m.GetAttackSimulationUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("clickCount", m.GetClickCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("compromisedCount", m.GetCompromisedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("latestSimulationDateTime", m.GetLatestSimulationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("simulationCount", m.GetSimulationCount())
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
func (m *AttackSimulationSimulationUserCoverage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttackSimulationUser sets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationSimulationUserCoverage) SetAttackSimulationUser(value AttackSimulationUserable)() {
    if m != nil {
        m.attackSimulationUser = value
    }
}
// SetClickCount sets the clickCount property value. Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) SetClickCount(value *int32)() {
    if m != nil {
        m.clickCount = value
    }
}
// SetCompromisedCount sets the compromisedCount property value. Number of compromising actions by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) SetCompromisedCount(value *int32)() {
    if m != nil {
        m.compromisedCount = value
    }
}
// SetLatestSimulationDateTime sets the latestSimulationDateTime property value. Date and time of latest attack simulation and training campaign that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) SetLatestSimulationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.latestSimulationDateTime = value
    }
}
// SetSimulationCount sets the simulationCount property value. Number of attack simulation and training campaigns that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) SetSimulationCount(value *int32)() {
    if m != nil {
        m.simulationCount = value
    }
}
