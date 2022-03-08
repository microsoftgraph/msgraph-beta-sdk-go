package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementTemplateStepVersion provides operations to manage the tenantRelationship singleton.
type ManagementTemplateStepVersion struct {
    Entity
    // 
    acceptedFor ManagementTemplateStepable;
    // 
    contentMarkdown *string;
    // 
    createdByUserId *string;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    deployments []ManagementTemplateStepDeploymentable;
    // 
    lastActionByUserId *string;
    // 
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    name *string;
    // 
    templateStep ManagementTemplateStepable;
    // 
    version *int32;
    // 
    versionInformation *string;
}
// NewManagementTemplateStepVersion instantiates a new managementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepVersionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagementTemplateStepVersion(), nil
}
// GetAcceptedFor gets the acceptedFor property value. 
func (m *ManagementTemplateStepVersion) GetAcceptedFor()(ManagementTemplateStepable) {
    if m == nil {
        return nil
    } else {
        return m.acceptedFor
    }
}
// GetContentMarkdown gets the contentMarkdown property value. 
func (m *ManagementTemplateStepVersion) GetContentMarkdown()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentMarkdown
    }
}
// GetCreatedByUserId gets the createdByUserId property value. 
func (m *ManagementTemplateStepVersion) GetCreatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByUserId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *ManagementTemplateStepVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeployments gets the deployments property value. 
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeploymentable) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedFor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedFor(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["contentMarkdown"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentMarkdown(val)
        }
        return nil
    }
    res["createdByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deployments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepDeploymentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepDeploymentable)
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["lastActionByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["templateStep"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateStep(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versionInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInformation(val)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. 
func (m *ManagementTemplateStepVersion) GetLastActionByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActionByUserId
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. 
func (m *ManagementTemplateStepVersion) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetName gets the name property value. 
func (m *ManagementTemplateStepVersion) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetTemplateStep gets the templateStep property value. 
func (m *ManagementTemplateStepVersion) GetTemplateStep()(ManagementTemplateStepable) {
    if m == nil {
        return nil
    } else {
        return m.templateStep
    }
}
// GetVersion gets the version property value. 
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetVersionInformation gets the versionInformation property value. 
func (m *ManagementTemplateStepVersion) GetVersionInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionInformation
    }
}
func (m *ManagementTemplateStepVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedFor", m.GetAcceptedFor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentMarkdown", m.GetContentMarkdown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
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
    if m.GetDeployments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStep", m.GetTemplateStep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInformation", m.GetVersionInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedFor sets the acceptedFor property value. 
func (m *ManagementTemplateStepVersion) SetAcceptedFor(value ManagementTemplateStepable)() {
    if m != nil {
        m.acceptedFor = value
    }
}
// SetContentMarkdown sets the contentMarkdown property value. 
func (m *ManagementTemplateStepVersion) SetContentMarkdown(value *string)() {
    if m != nil {
        m.contentMarkdown = value
    }
}
// SetCreatedByUserId sets the createdByUserId property value. 
func (m *ManagementTemplateStepVersion) SetCreatedByUserId(value *string)() {
    if m != nil {
        m.createdByUserId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *ManagementTemplateStepVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeployments sets the deployments property value. 
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeploymentable)() {
    if m != nil {
        m.deployments = value
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. 
func (m *ManagementTemplateStepVersion) SetLastActionByUserId(value *string)() {
    if m != nil {
        m.lastActionByUserId = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. 
func (m *ManagementTemplateStepVersion) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetName sets the name property value. 
func (m *ManagementTemplateStepVersion) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetTemplateStep sets the templateStep property value. 
func (m *ManagementTemplateStepVersion) SetTemplateStep(value ManagementTemplateStepable)() {
    if m != nil {
        m.templateStep = value
    }
}
// SetVersion sets the version property value. 
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
// SetVersionInformation sets the versionInformation property value. 
func (m *ManagementTemplateStepVersion) SetVersionInformation(value *string)() {
    if m != nil {
        m.versionInformation = value
    }
}
