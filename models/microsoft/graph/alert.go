package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// Alert 
type Alert struct {
    Entity
    // 
    actorDisplayName *string;
    // 
    alertWebUrl *string;
    // 
    assignedTo *string;
    // 
    category *string;
    // 
    classification *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertClassification;
    // 
    comments []AlertComment;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    description *string;
    // 
    detectionSource *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DetectionSource;
    // 
    detectorId *string;
    // 
    determination *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertDetermination;
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
    recommendedActions *string;
    // 
    resolvedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    serviceSource *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ServiceSource;
    // 
    severity *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertSeverity;
    // 
    status *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertStatus;
    // 
    tenantId *string;
    // 
    threatDisplayName *string;
    // 
    threatFamilyName *string;
    // 
    title *string;
}
// NewAlert instantiates a new alert and sets the default values.
func NewAlert()(*Alert) {
    m := &Alert{
        Entity: *NewEntity(),
    }
    return m
}
// GetActorDisplayName gets the actorDisplayName property value. 
func (m *Alert) GetActorDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actorDisplayName
    }
}
// GetAlertWebUrl gets the alertWebUrl property value. 
func (m *Alert) GetAlertWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alertWebUrl
    }
}
// GetAssignedTo gets the assignedTo property value. 
func (m *Alert) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetCategory gets the category property value. 
func (m *Alert) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetClassification gets the classification property value. 
func (m *Alert) GetClassification()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertClassification) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetComments gets the comments property value. 
func (m *Alert) GetComments()([]AlertComment) {
    if m == nil {
        return nil
    } else {
        return m.comments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *Alert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. 
func (m *Alert) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetectionSource gets the detectionSource property value. 
func (m *Alert) GetDetectionSource()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DetectionSource) {
    if m == nil {
        return nil
    } else {
        return m.detectionSource
    }
}
// GetDetectorId gets the detectorId property value. 
func (m *Alert) GetDetectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectorId
    }
}
// GetDetermination gets the determination property value. 
func (m *Alert) GetDetermination()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertDetermination) {
    if m == nil {
        return nil
    } else {
        return m.determination
    }
}
// GetFirstActivityDateTime gets the firstActivityDateTime property value. 
func (m *Alert) GetFirstActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstActivityDateTime
    }
}
// GetIncidentId gets the incidentId property value. 
func (m *Alert) GetIncidentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.incidentId
    }
}
// GetIncidentWebUrl gets the incidentWebUrl property value. 
func (m *Alert) GetIncidentWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.incidentWebUrl
    }
}
// GetLastActivityDateTime gets the lastActivityDateTime property value. 
func (m *Alert) GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDateTime
    }
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. 
func (m *Alert) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetMitreTechniques gets the mitreTechniques property value. 
func (m *Alert) GetMitreTechniques()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mitreTechniques
    }
}
// GetProviderAlertId gets the providerAlertId property value. 
func (m *Alert) GetProviderAlertId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerAlertId
    }
}
// GetRecommendedActions gets the recommendedActions property value. 
func (m *Alert) GetRecommendedActions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedActions
    }
}
// GetResolvedDateTime gets the resolvedDateTime property value. 
func (m *Alert) GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.resolvedDateTime
    }
}
// GetServiceSource gets the serviceSource property value. 
func (m *Alert) GetServiceSource()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ServiceSource) {
    if m == nil {
        return nil
    } else {
        return m.serviceSource
    }
}
// GetSeverity gets the severity property value. 
func (m *Alert) GetSeverity()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetStatus gets the status property value. 
func (m *Alert) GetStatus()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTenantId gets the tenantId property value. 
func (m *Alert) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetThreatDisplayName gets the threatDisplayName property value. 
func (m *Alert) GetThreatDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatDisplayName
    }
}
// GetThreatFamilyName gets the threatFamilyName property value. 
func (m *Alert) GetThreatFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threatFamilyName
    }
}
// GetTitle gets the title property value. 
func (m *Alert) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Alert) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseAlertClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertClassification))
        }
        return nil
    }
    res["comments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlertComment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertComment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AlertComment))
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
    res["detectionSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseDetectionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionSource(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DetectionSource))
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
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseAlertDetermination)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetermination(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertDetermination))
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
    res["recommendedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedActions(val)
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
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseServiceSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceSource(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ServiceSource))
        }
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertSeverity))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertStatus))
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
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
func (m *Alert) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Alert) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        cast := (*m.GetClassification()).String()
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
    if m.GetDetectionSource() != nil {
        cast := (*m.GetDetectionSource()).String()
        err = writer.WriteStringValue("detectionSource", &cast)
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
        cast := (*m.GetDetermination()).String()
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
        err = writer.WriteStringValue("recommendedActions", m.GetRecommendedActions())
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
        cast := (*m.GetServiceSource()).String()
        err = writer.WriteStringValue("serviceSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
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
// SetActorDisplayName sets the actorDisplayName property value. 
func (m *Alert) SetActorDisplayName(value *string)() {
    if m != nil {
        m.actorDisplayName = value
    }
}
// SetAlertWebUrl sets the alertWebUrl property value. 
func (m *Alert) SetAlertWebUrl(value *string)() {
    if m != nil {
        m.alertWebUrl = value
    }
}
// SetAssignedTo sets the assignedTo property value. 
func (m *Alert) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetCategory sets the category property value. 
func (m *Alert) SetCategory(value *string)() {
    if m != nil {
        m.category = value
    }
}
// SetClassification sets the classification property value. 
func (m *Alert) SetClassification(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertClassification)() {
    if m != nil {
        m.classification = value
    }
}
// SetComments sets the comments property value. 
func (m *Alert) SetComments(value []AlertComment)() {
    if m != nil {
        m.comments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *Alert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. 
func (m *Alert) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetectionSource sets the detectionSource property value. 
func (m *Alert) SetDetectionSource(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DetectionSource)() {
    if m != nil {
        m.detectionSource = value
    }
}
// SetDetectorId sets the detectorId property value. 
func (m *Alert) SetDetectorId(value *string)() {
    if m != nil {
        m.detectorId = value
    }
}
// SetDetermination sets the determination property value. 
func (m *Alert) SetDetermination(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertDetermination)() {
    if m != nil {
        m.determination = value
    }
}
// SetFirstActivityDateTime sets the firstActivityDateTime property value. 
func (m *Alert) SetFirstActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstActivityDateTime = value
    }
}
// SetIncidentId sets the incidentId property value. 
func (m *Alert) SetIncidentId(value *string)() {
    if m != nil {
        m.incidentId = value
    }
}
// SetIncidentWebUrl sets the incidentWebUrl property value. 
func (m *Alert) SetIncidentWebUrl(value *string)() {
    if m != nil {
        m.incidentWebUrl = value
    }
}
// SetLastActivityDateTime sets the lastActivityDateTime property value. 
func (m *Alert) SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActivityDateTime = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. 
func (m *Alert) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetMitreTechniques sets the mitreTechniques property value. 
func (m *Alert) SetMitreTechniques(value []string)() {
    if m != nil {
        m.mitreTechniques = value
    }
}
// SetProviderAlertId sets the providerAlertId property value. 
func (m *Alert) SetProviderAlertId(value *string)() {
    if m != nil {
        m.providerAlertId = value
    }
}
// SetRecommendedActions sets the recommendedActions property value. 
func (m *Alert) SetRecommendedActions(value *string)() {
    if m != nil {
        m.recommendedActions = value
    }
}
// SetResolvedDateTime sets the resolvedDateTime property value. 
func (m *Alert) SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.resolvedDateTime = value
    }
}
// SetServiceSource sets the serviceSource property value. 
func (m *Alert) SetServiceSource(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ServiceSource)() {
    if m != nil {
        m.serviceSource = value
    }
}
// SetSeverity sets the severity property value. 
func (m *Alert) SetSeverity(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertSeverity)() {
    if m != nil {
        m.severity = value
    }
}
// SetStatus sets the status property value. 
func (m *Alert) SetStatus(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AlertStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *Alert) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetThreatDisplayName sets the threatDisplayName property value. 
func (m *Alert) SetThreatDisplayName(value *string)() {
    if m != nil {
        m.threatDisplayName = value
    }
}
// SetThreatFamilyName sets the threatFamilyName property value. 
func (m *Alert) SetThreatFamilyName(value *string)() {
    if m != nil {
        m.threatFamilyName = value
    }
}
// SetTitle sets the title property value. 
func (m *Alert) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
