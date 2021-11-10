package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserSimulationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of trainings assigned to a user in an attack simulation and training campaign.
    assignedTrainingsCount *int32;
    // Number of trainings completed by a user in an attack simulation and training campaign.
    completedTrainingsCount *int32;
    // Date and time of the compromising online action by a user in an attack simulation and training campaign.
    compromisedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of trainings in progress by a user in an attack simulation and training campaign.
    inProgressTrainingsCount *int32;
    // Flag representing if user was compromised in an attack simulation and training campaign.
    isCompromised *bool;
    // Date and time when user reported delivered payload as phish in the attack simulation and training campaign.
    reportedPhishDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of simulation events of a user in the attack simulation and training campaign.
    simulationEvents []UserSimulationEventInfo;
    // User in an attack simulation and training campaign.
    simulationUser *AttackSimulationUser;
    // List of training events of a user in the attack simulation and training campaign.
    trainingEvents []UserTrainingEventInfo;
}
// Instantiates a new userSimulationDetails and sets the default values.
func NewUserSimulationDetails()(*UserSimulationDetails) {
    m := &UserSimulationDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignedTrainingsCount property value. Number of trainings assigned to a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetAssignedTrainingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.assignedTrainingsCount
    }
}
// Gets the completedTrainingsCount property value. Number of trainings completed by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetCompletedTrainingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.completedTrainingsCount
    }
}
// Gets the compromisedDateTime property value. Date and time of the compromising online action by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetCompromisedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.compromisedDateTime
    }
}
// Gets the inProgressTrainingsCount property value. Number of trainings in progress by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetInProgressTrainingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inProgressTrainingsCount
    }
}
// Gets the isCompromised property value. Flag representing if user was compromised in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetIsCompromised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompromised
    }
}
// Gets the reportedPhishDateTime property value. Date and time when user reported delivered payload as phish in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetReportedPhishDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reportedPhishDateTime
    }
}
// Gets the simulationEvents property value. List of simulation events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetSimulationEvents()([]UserSimulationEventInfo) {
    if m == nil {
        return nil
    } else {
        return m.simulationEvents
    }
}
// Gets the simulationUser property value. User in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetSimulationUser()(*AttackSimulationUser) {
    if m == nil {
        return nil
    } else {
        return m.simulationUser
    }
}
// Gets the trainingEvents property value. List of training events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetTrainingEvents()([]UserTrainingEventInfo) {
    if m == nil {
        return nil
    } else {
        return m.trainingEvents
    }
}
// The deserialization information for the current model
func (m *UserSimulationDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedTrainingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTrainingsCount(val)
        }
        return nil
    }
    res["completedTrainingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedTrainingsCount(val)
        }
        return nil
    }
    res["compromisedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedDateTime(val)
        }
        return nil
    }
    res["inProgressTrainingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInProgressTrainingsCount(val)
        }
        return nil
    }
    res["isCompromised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompromised(val)
        }
        return nil
    }
    res["reportedPhishDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportedPhishDateTime(val)
        }
        return nil
    }
    res["simulationEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSimulationEventInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSimulationEventInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSimulationEventInfo))
            }
            m.SetSimulationEvents(res)
        }
        return nil
    }
    res["simulationUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttackSimulationUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationUser(val.(*AttackSimulationUser))
        }
        return nil
    }
    res["trainingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserTrainingEventInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserTrainingEventInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserTrainingEventInfo))
            }
            m.SetTrainingEvents(res)
        }
        return nil
    }
    return res
}
func (m *UserSimulationDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserSimulationDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignedTrainingsCount", m.GetAssignedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("completedTrainingsCount", m.GetCompletedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("compromisedDateTime", m.GetCompromisedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgressTrainingsCount", m.GetInProgressTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCompromised", m.GetIsCompromised())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("reportedPhishDateTime", m.GetReportedPhishDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSimulationEvents()))
        for i, v := range m.GetSimulationEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("simulationEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("simulationUser", m.GetSimulationUser())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTrainingEvents()))
        for i, v := range m.GetTrainingEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("trainingEvents", cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserSimulationDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignedTrainingsCount property value. Number of trainings assigned to a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the assignedTrainingsCount property.
func (m *UserSimulationDetails) SetAssignedTrainingsCount(value *int32)() {
    m.assignedTrainingsCount = value
}
// Sets the completedTrainingsCount property value. Number of trainings completed by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the completedTrainingsCount property.
func (m *UserSimulationDetails) SetCompletedTrainingsCount(value *int32)() {
    m.completedTrainingsCount = value
}
// Sets the compromisedDateTime property value. Date and time of the compromising online action by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the compromisedDateTime property.
func (m *UserSimulationDetails) SetCompromisedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.compromisedDateTime = value
}
// Sets the inProgressTrainingsCount property value. Number of trainings in progress by a user in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the inProgressTrainingsCount property.
func (m *UserSimulationDetails) SetInProgressTrainingsCount(value *int32)() {
    m.inProgressTrainingsCount = value
}
// Sets the isCompromised property value. Flag representing if user was compromised in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the isCompromised property.
func (m *UserSimulationDetails) SetIsCompromised(value *bool)() {
    m.isCompromised = value
}
// Sets the reportedPhishDateTime property value. Date and time when user reported delivered payload as phish in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the reportedPhishDateTime property.
func (m *UserSimulationDetails) SetReportedPhishDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reportedPhishDateTime = value
}
// Sets the simulationEvents property value. List of simulation events of a user in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the simulationEvents property.
func (m *UserSimulationDetails) SetSimulationEvents(value []UserSimulationEventInfo)() {
    m.simulationEvents = value
}
// Sets the simulationUser property value. User in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the simulationUser property.
func (m *UserSimulationDetails) SetSimulationUser(value *AttackSimulationUser)() {
    m.simulationUser = value
}
// Sets the trainingEvents property value. List of training events of a user in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the trainingEvents property.
func (m *UserSimulationDetails) SetTrainingEvents(value []UserTrainingEventInfo)() {
    m.trainingEvents = value
}
