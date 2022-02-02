package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Incident 
type Incident struct {
    Entity
    // 
    alerts []Alert_v2;
    // 
    assignedTo *string;
    // 
    classification *AlertClassification_v2;
    // 
    comments []AlertComment_v2;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    determination *AlertDetermination_v2;
    // 
    displayName *string;
    // 
    incidentWebUrl *string;
    // 
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    redirectIncidentId *string;
    // 
    severity *AlertSeverity_v2;
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
// GetAlerts gets the alerts property value. 
func (m *Incident) GetAlerts()([]Alert_v2) {
    if m == nil {
        return nil
    } else {
        return m.alerts
    }
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
func (m *Incident) GetClassification()(*AlertClassification_v2) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetComments gets the comments property value. 
func (m *Incident) GetComments()([]AlertComment_v2) {
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
func (m *Incident) GetDetermination()(*AlertDetermination_v2) {
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
func (m *Incident) GetSeverity()(*AlertSeverity_v2) {
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
    res["alerts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlert_v2() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alert_v2, len(val))
            for i, v := range val {
                res[i] = *(v.(*Alert_v2))
            }
            m.SetAlerts(res)
        }
        return nil
    }
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
        val, err := n.GetEnumValue(ParseAlertClassification_v2)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AlertClassification_v2)
            m.SetClassification(&cast)
        }
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlertComment_v2() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertComment_v2, len(val))
            for i, v := range val {
                res[i] = *(v.(*AlertComment_v2))
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
        val, err := n.GetEnumValue(ParseAlertDetermination_v2)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AlertDetermination_v2)
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
        val, err := n.GetEnumValue(ParseAlertSeverity_v2)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AlertSeverity_v2)
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
    if m.GetAlerts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
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
    if m.GetComments() != nil {
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
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. 
func (m *Incident) SetAlerts(value []Alert_v2)() {
    if m != nil {
        m.alerts = value
    }
}
// SetAssignedTo sets the assignedTo property value. 
func (m *Incident) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetClassification sets the classification property value. 
func (m *Incident) SetClassification(value *AlertClassification_v2)() {
    if m != nil {
        m.classification = value
    }
}
// SetComments sets the comments property value. 
func (m *Incident) SetComments(value []AlertComment_v2)() {
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
func (m *Incident) SetDetermination(value *AlertDetermination_v2)() {
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
func (m *Incident) SetSeverity(value *AlertSeverity_v2)() {
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
