package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Alert_v2 
type Alert_v2 struct {
    Entity
    // 
    aadTenantId *string;
    // 
    actorDisplayName *string;
    // 
    alertWebUrl *string;
    // 
    assignedTo *string;
    // 
    category *string;
    // 
    classification *AlertClassification_v2;
    // 
    comments []AlertComment_v2;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    description *string;
    // 
    detectorId *string;
    // 
    determination *AlertDetermination_v2;
    // 
    firstActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    incidentId *string;
    // 
    incidentWebUrl *string;
    // 
    lastActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    mitreTechniques []string;
    // 
    providerAlertId *string;
    // 
    resolvedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    serviceSource *ServiceSource;
    // 
    severity *AlertSeverity_v2;
    // 
    status *AlertStatus_v2;
    // 
    threatDisplayName *string;
    // 
    threatFamilyName *string;
    // 
    title *string;
}
// NewAlert_v2 instantiates a new alert_v2 and sets the default values.
func NewAlert_v2()(*Alert_v2) {
    m := &Alert_v2{
        Entity: *NewEntity(),
    }
    return m
}
// GetAadTenantId gets the aadTenantId property value. 
func (m *Alert_v2) GetAadTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aadTenantId
    }
}
// GetActorDisplayName gets the actorDisplayName property value. 
func (m *Alert_v2) GetActorDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actorDisplayName
    }
}
// GetAlertWebUrl gets the alertWebUrl property value. 
func (m *Alert_v2) GetAlertWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alertWebUrl
    }
}
// GetAssignedTo gets the assignedTo property value. 
func (m *Alert_v2) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetCategory gets the category property value. 
func (m *Alert_v2) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetClassification gets the classification property value. 
func (m *Alert_v2) GetClassification()(*AlertClassification_v2) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetComments gets the comments property value. 
func (m *Alert_v2) GetComments()([]AlertComment_v2) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *Alert_v2) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. 
func (m *Alert_v2) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetectorId gets the detectorId property value. 
func (m *Alert_v2) GetDetectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectorId
    }
}
// GetDetermination gets the determination property value. 
func (m *Alert_v2) GetDetermination()(*AlertDetermination_v2) {
    if m == nil {
        return nil
    } else {
        return m.determination
    }
}
// GetFirstActivityDateTime gets the firstActivityDateTime property value. 
func (m *Alert_v2) GetFirstActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstActivityDateTime
    }
}
// GetIncidentId gets the incidentId property value. 
func (m *Alert_v2) GetIncidentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.incidentId
    }
}
// GetIncidentWebUrl gets the incidentWebUrl property value. 
func (m *Alert_v2) GetIncidentWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.incidentWebUrl
    }
}
// GetLastActivityDateTime gets the lastActivityDateTime property value. 
func (m *Alert_v2) GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDateTime
    }
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. 
func (m *Alert_v2) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetMitreTechniques gets the mitreTechniques property value. 
func (m *Alert_v2) GetMitreTechniques()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mitreTechniques
    }
}
// GetProviderAlertId gets the providerAlertId property value. 
func (m *Alert_v2) GetProviderAlertId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerAlertId
    }
}
// GetResolvedDateTime gets the resolvedDateTime property value. 
func (m *Alert_v2) GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.resolvedDateTime
    }
}
// GetServiceSource gets the serviceSource property value. 
func (m *Alert_v2) GetServiceSource()(*ServiceSource) {
    if m == nil {
        return nil
    } else {
        return m.serviceSource
    }
}
// GetSeverity gets the severity property value. 
func (m *Alert_v2) GetSeverity()(*AlertSeverity_v2) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetStatus gets the status property value. 
func (m *Alert_v2) GetStatus()(*AlertStatus_v2) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetThreatDisplayName gets the threatDisplayName property value. 
func (m *Alert_v2) GetThreatDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatDisplayName
    }
}
// GetThreatFamilyName gets the threatFamilyName property value. 
func (m *Alert_v2) GetThreatFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatFamilyName
    }
}
// GetTitle gets the title property value. 
func (m *Alert_v2) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Alert_v2) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadTenantId(val)
        }
        return nil
    }
    res["actorDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorDisplayName(val)
        }
        return nil
    }
    res["alertWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertWebUrl(val)
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
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["detectorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectorId(val)
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
    res["firstActivityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstActivityDateTime(val)
        }
        return nil
    }
    res["incidentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentId(val)
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
    res["lastActivityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDateTime(val)
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
    res["mitreTechniques"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMitreTechniques(res)
        }
        return nil
    }
    res["providerAlertId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderAlertId(val)
        }
        return nil
    }
    res["resolvedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedDateTime(val)
        }
        return nil
    }
    res["serviceSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ServiceSource)
            m.SetServiceSource(&cast)
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
        val, err := n.GetEnumValue(ParseAlertStatus_v2)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AlertStatus_v2)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["threatDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatDisplayName(val)
        }
        return nil
    }
    res["threatFamilyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatFamilyName(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
func (m *Alert_v2) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Alert_v2) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aadTenantId", m.GetAadTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("actorDisplayName", m.GetActorDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertWebUrl", m.GetAlertWebUrl())
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
    {
        err = writer.WriteStringValue("category", m.GetCategory())
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetDetermination() != nil {
        cast := m.GetDetermination().String()
        err = writer.WriteStringValue("determination", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstActivityDateTime", m.GetFirstActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("incidentId", m.GetIncidentId())
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
        err = writer.WriteTimeValue("lastActivityDateTime", m.GetLastActivityDateTime())
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
    if m.GetMitreTechniques() != nil {
        err = writer.WriteCollectionOfStringValues("mitreTechniques", m.GetMitreTechniques())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerAlertId", m.GetProviderAlertId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("resolvedDateTime", m.GetResolvedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetServiceSource() != nil {
        cast := m.GetServiceSource().String()
        err = writer.WriteStringValue("serviceSource", &cast)
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
        err = writer.WriteStringValue("threatDisplayName", m.GetThreatDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("threatFamilyName", m.GetThreatFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadTenantId sets the aadTenantId property value. 
func (m *Alert_v2) SetAadTenantId(value *string)() {
    if m != nil {
        m.aadTenantId = value
    }
}
// SetActorDisplayName sets the actorDisplayName property value. 
func (m *Alert_v2) SetActorDisplayName(value *string)() {
    if m != nil {
        m.actorDisplayName = value
    }
}
// SetAlertWebUrl sets the alertWebUrl property value. 
func (m *Alert_v2) SetAlertWebUrl(value *string)() {
    if m != nil {
        m.alertWebUrl = value
    }
}
// SetAssignedTo sets the assignedTo property value. 
func (m *Alert_v2) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetCategory sets the category property value. 
func (m *Alert_v2) SetCategory(value *string)() {
    if m != nil {
        m.category = value
    }
}
// SetClassification sets the classification property value. 
func (m *Alert_v2) SetClassification(value *AlertClassification_v2)() {
    if m != nil {
        m.classification = value
    }
}
// SetComments sets the comments property value. 
func (m *Alert_v2) SetComments(value []AlertComment_v2)() {
    if m != nil {
        m.comments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *Alert_v2) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. 
func (m *Alert_v2) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetectorId sets the detectorId property value. 
func (m *Alert_v2) SetDetectorId(value *string)() {
    if m != nil {
        m.detectorId = value
    }
}
// SetDetermination sets the determination property value. 
func (m *Alert_v2) SetDetermination(value *AlertDetermination_v2)() {
    if m != nil {
        m.determination = value
    }
}
// SetFirstActivityDateTime sets the firstActivityDateTime property value. 
func (m *Alert_v2) SetFirstActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstActivityDateTime = value
    }
}
// SetIncidentId sets the incidentId property value. 
func (m *Alert_v2) SetIncidentId(value *string)() {
    if m != nil {
        m.incidentId = value
    }
}
// SetIncidentWebUrl sets the incidentWebUrl property value. 
func (m *Alert_v2) SetIncidentWebUrl(value *string)() {
    if m != nil {
        m.incidentWebUrl = value
    }
}
// SetLastActivityDateTime sets the lastActivityDateTime property value. 
func (m *Alert_v2) SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActivityDateTime = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. 
func (m *Alert_v2) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetMitreTechniques sets the mitreTechniques property value. 
func (m *Alert_v2) SetMitreTechniques(value []string)() {
    if m != nil {
        m.mitreTechniques = value
    }
}
// SetProviderAlertId sets the providerAlertId property value. 
func (m *Alert_v2) SetProviderAlertId(value *string)() {
    if m != nil {
        m.providerAlertId = value
    }
}
// SetResolvedDateTime sets the resolvedDateTime property value. 
func (m *Alert_v2) SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.resolvedDateTime = value
    }
}
// SetServiceSource sets the serviceSource property value. 
func (m *Alert_v2) SetServiceSource(value *ServiceSource)() {
    if m != nil {
        m.serviceSource = value
    }
}
// SetSeverity sets the severity property value. 
func (m *Alert_v2) SetSeverity(value *AlertSeverity_v2)() {
    if m != nil {
        m.severity = value
    }
}
// SetStatus sets the status property value. 
func (m *Alert_v2) SetStatus(value *AlertStatus_v2)() {
    if m != nil {
        m.status = value
    }
}
// SetThreatDisplayName sets the threatDisplayName property value. 
func (m *Alert_v2) SetThreatDisplayName(value *string)() {
    if m != nil {
        m.threatDisplayName = value
    }
}
// SetThreatFamilyName sets the threatFamilyName property value. 
func (m *Alert_v2) SetThreatFamilyName(value *string)() {
    if m != nil {
        m.threatFamilyName = value
    }
}
// SetTitle sets the title property value. 
func (m *Alert_v2) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
