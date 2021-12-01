package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Incident 
type Incident struct {
    Entity
    // 
    assignedTo *string;
    // 
    classification *M365AlertClassification;
    // 
    comments []M365AlertComment;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    determination *M365AlertDetermination;
    // 
    displayName *string;
    // 
    incidentWebUrl *string;
    // 
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    redirectIncidentId *string;
    // 
    severity *M365AlertSeverity;
    // 
    status *IncidentStatus;
    // 
    tags []string;
}
// NewIncident instantiates a new incident and sets the default values.
func NewIncident()(*Incident) {
    m := &Incident{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignedTo gets the assignedTo property value. 
func (m *Incident) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetClassification gets the classification property value. 
func (m *Incident) GetClassification()(*M365AlertClassification) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetComments gets the comments property value. 
func (m *Incident) GetComments()([]M365AlertComment) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *Incident) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDetermination gets the determination property value. 
func (m *Incident) GetDetermination()(*M365AlertDetermination) {
    if m == nil {
        return nil
    } else {
        return m.determination
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Incident) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIncidentWebUrl gets the incidentWebUrl property value. 
func (m *Incident) GetIncidentWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.incidentWebUrl
    }
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. 
func (m *Incident) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetRedirectIncidentId gets the redirectIncidentId property value. 
func (m *Incident) GetRedirectIncidentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectIncidentId
    }
}
// GetSeverity gets the severity property value. 
func (m *Incident) GetSeverity()(*M365AlertSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetStatus gets the status property value. 
func (m *Incident) GetStatus()(*IncidentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTags gets the tags property value. 
func (m *Incident) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Incident) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseM365AlertClassification)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(M365AlertClassification)
            m.SetClassification(&cast)
        }
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewM365AlertComment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]M365AlertComment, len(val))
            for i, v := range val {
                res[i] = *(v.(*M365AlertComment))
            }
            m.SetComments(res)
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
    res["determination"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseM365AlertDetermination)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(M365AlertDetermination)
            m.SetDetermination(&cast)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["incidentWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentWebUrl(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["redirectIncidentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectIncidentId(val)
        }
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseM365AlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(M365AlertSeverity)
            m.SetSeverity(&cast)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncidentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(IncidentStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
func (m *Incident) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Incident) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetClassification() != nil {
        cast := m.GetClassification().String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("comments", cast)
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
    if m.GetDetermination() != nil {
        cast := m.GetDetermination().String()
        err = writer.WriteStringValue("determination", &cast)
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
        err = writer.WriteStringValue("incidentWebUrl", m.GetIncidentWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectIncidentId", m.GetRedirectIncidentId())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := m.GetSeverity().String()
        err = writer.WriteStringValue("severity", &cast)
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
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. 
func (m *Incident) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetClassification sets the classification property value. 
func (m *Incident) SetClassification(value *M365AlertClassification)() {
    if m != nil {
        m.classification = value
    }
}
// SetComments sets the comments property value. 
func (m *Incident) SetComments(value []M365AlertComment)() {
    if m != nil {
        m.comments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *Incident) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDetermination sets the determination property value. 
func (m *Incident) SetDetermination(value *M365AlertDetermination)() {
    if m != nil {
        m.determination = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Incident) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIncidentWebUrl sets the incidentWebUrl property value. 
func (m *Incident) SetIncidentWebUrl(value *string)() {
    if m != nil {
        m.incidentWebUrl = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. 
func (m *Incident) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetRedirectIncidentId sets the redirectIncidentId property value. 
func (m *Incident) SetRedirectIncidentId(value *string)() {
    if m != nil {
        m.redirectIncidentId = value
    }
}
// SetSeverity sets the severity property value. 
func (m *Incident) SetSeverity(value *M365AlertSeverity)() {
    if m != nil {
        m.severity = value
    }
}
// SetStatus sets the status property value. 
func (m *Incident) SetStatus(value *IncidentStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTags sets the tags property value. 
func (m *Incident) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
