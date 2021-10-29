package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Simulation struct {
    Entity
    // The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
    attackTechnique *SimulationAttackTechnique;
    // Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
    attackType *SimulationAttackType;
    // Flag representing if artifacts were cleaned up in the attack simulation and training campaign.
    cleanupArtifacts *bool;
    // Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity of the user who created the attack simulation and training campaign.
    createdBy *EmailIdentity;
    // Date and time of creation of the attack simulation and training campaign.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the attack simulation and training campaign.
    description *string;
    // Display name of the attack simulation and training campaign. Supports $filter and $orderby.
    displayName *string;
    // Flag representing whether to enable or disable timezone-aware delivery of phishing payload in the attack simulation and training campaign.
    enableRegionTimezoneDelivery *bool;
    // Flag representing inclusion of all the users of a tenant in the attack simulation and training campaign.
    includeAllAccountTargets *bool;
    // Flag representing if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
    isAutomated *bool;
    // Identity of the user who most recently modified the attack simulation and training campaign.
    lastModifiedBy *EmailIdentity;
    // Date and time of the most recent modification of the attack simulation and training campaign.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
    launchDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Mode of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: real, preview, unknownFutureValue.
    mode *SimulationMode;
    // Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
    payloadDeliveryPlatform *PayloadDeliveryPlatform;
    // Source of phishing payload in the attack simulation and training campaign. Possible values are: unknown, global, tenant, unknownFutureValue.
    payloadSource *PayloadSource;
    // Report of the attack simulation and training campaign.
    report *SimulationReport;
    // Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, inProgress, scheduled, completed, partiallyCompleted, failed, cancelled, excluded, deleted, included, unknownFutureValue.
    status *SimulationStatus;
    // Preference of the tenant admin to assign training to users in the attack simulation and training campaign. Possible values are: unknown, auto, manual, unknownFutureValue.
    trainingAssignmentPreference *TrainingAssignmentPreference;
    // Preference of the tenant admin for the source of training content to assign to users in the attack simulation and training campaign. Possible values are: unknown, microsoft, custom, noTraining, unknownFutureValue.
    trainingContentPreference *TrainingContentPreference;
    // Date and time before which the trainings need to be completed by users in the attack simulation and training campaign.
    trainingDueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new simulation and sets the default values.
func NewSimulation()(*Simulation) {
    m := &Simulation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the attackTechnique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
func (m *Simulation) GetAttackTechnique()(*SimulationAttackTechnique) {
    if m == nil {
        return nil
    } else {
        return m.attackTechnique
    }
}
// Gets the attackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Simulation) GetAttackType()(*SimulationAttackType) {
    if m == nil {
        return nil
    } else {
        return m.attackType
    }
}
// Gets the cleanupArtifacts property value. Flag representing if artifacts were cleaned up in the attack simulation and training campaign.
func (m *Simulation) GetCleanupArtifacts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cleanupArtifacts
    }
}
// Gets the completionDateTime property value. Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completionDateTime
    }
}
// Gets the createdBy property value. Identity of the user who created the attack simulation and training campaign.
func (m *Simulation) GetCreatedBy()(*EmailIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. Date and time of creation of the attack simulation and training campaign.
func (m *Simulation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description of the attack simulation and training campaign.
func (m *Simulation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enableRegionTimezoneDelivery property value. Flag representing whether to enable or disable timezone-aware delivery of phishing payload in the attack simulation and training campaign.
func (m *Simulation) GetEnableRegionTimezoneDelivery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableRegionTimezoneDelivery
    }
}
// Gets the includeAllAccountTargets property value. Flag representing inclusion of all the users of a tenant in the attack simulation and training campaign.
func (m *Simulation) GetIncludeAllAccountTargets()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeAllAccountTargets
    }
}
// Gets the isAutomated property value. Flag representing if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
func (m *Simulation) GetIsAutomated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomated
    }
}
// Gets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign.
func (m *Simulation) GetLastModifiedBy()(*EmailIdentity) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. Date and time of the most recent modification of the attack simulation and training campaign.
func (m *Simulation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the launchDateTime property value. Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) GetLaunchDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.launchDateTime
    }
}
// Gets the mode property value. Mode of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: real, preview, unknownFutureValue.
func (m *Simulation) GetMode()(*SimulationMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// Gets the payloadDeliveryPlatform property value. Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Simulation) GetPayloadDeliveryPlatform()(*PayloadDeliveryPlatform) {
    if m == nil {
        return nil
    } else {
        return m.payloadDeliveryPlatform
    }
}
// Gets the payloadSource property value. Source of phishing payload in the attack simulation and training campaign. Possible values are: unknown, global, tenant, unknownFutureValue.
func (m *Simulation) GetPayloadSource()(*PayloadSource) {
    if m == nil {
        return nil
    } else {
        return m.payloadSource
    }
}
// Gets the report property value. Report of the attack simulation and training campaign.
func (m *Simulation) GetReport()(*SimulationReport) {
    if m == nil {
        return nil
    } else {
        return m.report
    }
}
// Gets the status property value. Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, inProgress, scheduled, completed, partiallyCompleted, failed, cancelled, excluded, deleted, included, unknownFutureValue.
func (m *Simulation) GetStatus()(*SimulationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the trainingAssignmentPreference property value. Preference of the tenant admin to assign training to users in the attack simulation and training campaign. Possible values are: unknown, auto, manual, unknownFutureValue.
func (m *Simulation) GetTrainingAssignmentPreference()(*TrainingAssignmentPreference) {
    if m == nil {
        return nil
    } else {
        return m.trainingAssignmentPreference
    }
}
// Gets the trainingContentPreference property value. Preference of the tenant admin for the source of training content to assign to users in the attack simulation and training campaign. Possible values are: unknown, microsoft, custom, noTraining, unknownFutureValue.
func (m *Simulation) GetTrainingContentPreference()(*TrainingContentPreference) {
    if m == nil {
        return nil
    } else {
        return m.trainingContentPreference
    }
}
// Gets the trainingDueDateTime property value. Date and time before which the trainings need to be completed by users in the attack simulation and training campaign.
func (m *Simulation) GetTrainingDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.trainingDueDateTime
    }
}
// The deserialization information for the current model
func (m *Simulation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attackTechnique"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackTechnique)
        if err != nil {
            return err
        }
        cast := val.(SimulationAttackTechnique)
        m.SetAttackTechnique(&cast)
        return nil
    }
    res["attackType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackType)
        if err != nil {
            return err
        }
        cast := val.(SimulationAttackType)
        m.SetAttackType(&cast)
        return nil
    }
    res["cleanupArtifacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCleanupArtifacts(val)
        return nil
    }
    res["completionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCompletionDateTime(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*EmailIdentity))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enableRegionTimezoneDelivery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableRegionTimezoneDelivery(val)
        return nil
    }
    res["includeAllAccountTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIncludeAllAccountTargets(val)
        return nil
    }
    res["isAutomated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAutomated(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailIdentity() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*EmailIdentity))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["launchDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLaunchDateTime(val)
        return nil
    }
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationMode)
        if err != nil {
            return err
        }
        cast := val.(SimulationMode)
        m.SetMode(&cast)
        return nil
    }
    res["payloadDeliveryPlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadDeliveryPlatform)
        if err != nil {
            return err
        }
        cast := val.(PayloadDeliveryPlatform)
        m.SetPayloadDeliveryPlatform(&cast)
        return nil
    }
    res["payloadSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadSource)
        if err != nil {
            return err
        }
        cast := val.(PayloadSource)
        m.SetPayloadSource(&cast)
        return nil
    }
    res["report"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationReport() })
        if err != nil {
            return err
        }
        m.SetReport(val.(*SimulationReport))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationStatus)
        if err != nil {
            return err
        }
        cast := val.(SimulationStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["trainingAssignmentPreference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingAssignmentPreference)
        if err != nil {
            return err
        }
        cast := val.(TrainingAssignmentPreference)
        m.SetTrainingAssignmentPreference(&cast)
        return nil
    }
    res["trainingContentPreference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingContentPreference)
        if err != nil {
            return err
        }
        cast := val.(TrainingContentPreference)
        m.SetTrainingContentPreference(&cast)
        return nil
    }
    res["trainingDueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTrainingDueDateTime(val)
        return nil
    }
    return res
}
func (m *Simulation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Simulation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttackTechnique() != nil {
        cast := m.GetAttackTechnique().String()
        err = writer.WriteStringValue("attackTechnique", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttackType() != nil {
        cast := m.GetAttackType().String()
        err = writer.WriteStringValue("attackType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cleanupArtifacts", m.GetCleanupArtifacts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableRegionTimezoneDelivery", m.GetEnableRegionTimezoneDelivery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAllAccountTargets", m.GetIncludeAllAccountTargets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutomated", m.GetIsAutomated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("launchDateTime", m.GetLaunchDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMode() != nil {
        cast := m.GetMode().String()
        err = writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPayloadDeliveryPlatform() != nil {
        cast := m.GetPayloadDeliveryPlatform().String()
        err = writer.WriteStringValue("payloadDeliveryPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPayloadSource() != nil {
        cast := m.GetPayloadSource().String()
        err = writer.WriteStringValue("payloadSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("report", m.GetReport())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrainingAssignmentPreference() != nil {
        cast := m.GetTrainingAssignmentPreference().String()
        err = writer.WriteStringValue("trainingAssignmentPreference", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrainingContentPreference() != nil {
        cast := m.GetTrainingContentPreference().String()
        err = writer.WriteStringValue("trainingContentPreference", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("trainingDueDateTime", m.GetTrainingDueDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the attackTechnique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue. For more information on the types of social engineering attack techniques, see simulations.
// Parameters:
//  - value : Value to set for the attackTechnique property.
func (m *Simulation) SetAttackTechnique(value *SimulationAttackTechnique)() {
    m.attackTechnique = value
}
// Sets the attackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
// Parameters:
//  - value : Value to set for the attackType property.
func (m *Simulation) SetAttackType(value *SimulationAttackType)() {
    m.attackType = value
}
// Sets the cleanupArtifacts property value. Flag representing if artifacts were cleaned up in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the cleanupArtifacts property.
func (m *Simulation) SetCleanupArtifacts(value *bool)() {
    m.cleanupArtifacts = value
}
// Sets the completionDateTime property value. Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the completionDateTime property.
func (m *Simulation) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// Sets the createdBy property value. Identity of the user who created the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *Simulation) SetCreatedBy(value *EmailIdentity)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. Date and time of creation of the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Simulation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description of the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the description property.
func (m *Simulation) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name of the attack simulation and training campaign. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Simulation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enableRegionTimezoneDelivery property value. Flag representing whether to enable or disable timezone-aware delivery of phishing payload in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the enableRegionTimezoneDelivery property.
func (m *Simulation) SetEnableRegionTimezoneDelivery(value *bool)() {
    m.enableRegionTimezoneDelivery = value
}
// Sets the includeAllAccountTargets property value. Flag representing inclusion of all the users of a tenant in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the includeAllAccountTargets property.
func (m *Simulation) SetIncludeAllAccountTargets(value *bool)() {
    m.includeAllAccountTargets = value
}
// Sets the isAutomated property value. Flag representing if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the isAutomated property.
func (m *Simulation) SetIsAutomated(value *bool)() {
    m.isAutomated = value
}
// Sets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *Simulation) SetLastModifiedBy(value *EmailIdentity)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. Date and time of the most recent modification of the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Simulation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the launchDateTime property value. Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the launchDateTime property.
func (m *Simulation) SetLaunchDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.launchDateTime = value
}
// Sets the mode property value. Mode of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: real, preview, unknownFutureValue.
// Parameters:
//  - value : Value to set for the mode property.
func (m *Simulation) SetMode(value *SimulationMode)() {
    m.mode = value
}
// Sets the payloadDeliveryPlatform property value. Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
// Parameters:
//  - value : Value to set for the payloadDeliveryPlatform property.
func (m *Simulation) SetPayloadDeliveryPlatform(value *PayloadDeliveryPlatform)() {
    m.payloadDeliveryPlatform = value
}
// Sets the payloadSource property value. Source of phishing payload in the attack simulation and training campaign. Possible values are: unknown, global, tenant, unknownFutureValue.
// Parameters:
//  - value : Value to set for the payloadSource property.
func (m *Simulation) SetPayloadSource(value *PayloadSource)() {
    m.payloadSource = value
}
// Sets the report property value. Report of the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the report property.
func (m *Simulation) SetReport(value *SimulationReport)() {
    m.report = value
}
// Sets the status property value. Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, inProgress, scheduled, completed, partiallyCompleted, failed, cancelled, excluded, deleted, included, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *Simulation) SetStatus(value *SimulationStatus)() {
    m.status = value
}
// Sets the trainingAssignmentPreference property value. Preference of the tenant admin to assign training to users in the attack simulation and training campaign. Possible values are: unknown, auto, manual, unknownFutureValue.
// Parameters:
//  - value : Value to set for the trainingAssignmentPreference property.
func (m *Simulation) SetTrainingAssignmentPreference(value *TrainingAssignmentPreference)() {
    m.trainingAssignmentPreference = value
}
// Sets the trainingContentPreference property value. Preference of the tenant admin for the source of training content to assign to users in the attack simulation and training campaign. Possible values are: unknown, microsoft, custom, noTraining, unknownFutureValue.
// Parameters:
//  - value : Value to set for the trainingContentPreference property.
func (m *Simulation) SetTrainingContentPreference(value *TrainingContentPreference)() {
    m.trainingContentPreference = value
}
// Sets the trainingDueDateTime property value. Date and time before which the trainings need to be completed by users in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the trainingDueDateTime property.
func (m *Simulation) SetTrainingDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.trainingDueDateTime = value
}
