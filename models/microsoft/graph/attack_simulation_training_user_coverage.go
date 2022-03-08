package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
type AttackSimulationTrainingUserCoverage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // User in an attack simulation and training campaign.
    attackSimulationUser AttackSimulationUserable;
    // List of assigned trainings' and their statuses for the user.
    userTrainings []UserTrainingStatusInfoable;
}
// NewAttackSimulationTrainingUserCoverage instantiates a new attackSimulationTrainingUserCoverage and sets the default values.
func NewAttackSimulationTrainingUserCoverage()(*AttackSimulationTrainingUserCoverage) {
    m := &AttackSimulationTrainingUserCoverage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttackSimulationTrainingUserCoverage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationTrainingUserCoverage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttackSimulationUser gets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationTrainingUserCoverage) GetAttackSimulationUser()(AttackSimulationUserable) {
    if m == nil {
        return nil
    } else {
        return m.attackSimulationUser
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationTrainingUserCoverage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["userTrainings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserTrainingStatusInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserTrainingStatusInfoable, len(val))
            for i, v := range val {
                res[i] = v.(UserTrainingStatusInfoable)
            }
            m.SetUserTrainings(res)
        }
        return nil
    }
    return res
}
// GetUserTrainings gets the userTrainings property value. List of assigned trainings' and their statuses for the user.
func (m *AttackSimulationTrainingUserCoverage) GetUserTrainings()([]UserTrainingStatusInfoable) {
    if m == nil {
        return nil
    } else {
        return m.userTrainings
    }
}
func (m *AttackSimulationTrainingUserCoverage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttackSimulationTrainingUserCoverage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attackSimulationUser", m.GetAttackSimulationUser())
        if err != nil {
            return err
        }
    }
    if m.GetUserTrainings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserTrainings()))
        for i, v := range m.GetUserTrainings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userTrainings", cast)
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
func (m *AttackSimulationTrainingUserCoverage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttackSimulationUser sets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationTrainingUserCoverage) SetAttackSimulationUser(value AttackSimulationUserable)() {
    if m != nil {
        m.attackSimulationUser = value
    }
}
// SetUserTrainings sets the userTrainings property value. List of assigned trainings' and their statuses for the user.
func (m *AttackSimulationTrainingUserCoverage) SetUserTrainings(value []UserTrainingStatusInfoable)() {
    if m != nil {
        m.userTrainings = value
    }
}
