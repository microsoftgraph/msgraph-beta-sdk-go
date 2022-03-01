package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementTemplate 
type ManagementTemplate struct {
    Entity
    // The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory;
    // 
    createdByUserId *string;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description for the management template. Optional. Read-only.
    description *string;
    // The display name for the management template. Required. Read-only.
    displayName *string;
    // 
    informationLinks []ActionUrl;
    // 
    lastActionByUserId *string;
    // 
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    managementTemplateCollections []ManagementTemplateCollection;
    // 
    managementTemplateSteps []ManagementTemplateStep;
    // The collection of parameters used by the management template. Optional. Read-only.
    parameters []TemplateParameter;
    // 
    priority *int32;
    // 
    provider *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider;
    // 
    userImpact *string;
    // 
    version *int32;
    // The collection of workload actions associated with the management template. Optional. Read-only.
    workloadActions []WorkloadAction;
}
// NewManagementTemplate instantiates a new managementTemplate and sets the default values.
func NewManagementTemplate()(*ManagementTemplate) {
    m := &ManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCreatedByUserId gets the createdByUserId property value. 
func (m *ManagementTemplate) GetCreatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByUserId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *ManagementTemplate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetInformationLinks gets the informationLinks property value. 
func (m *ManagementTemplate) GetInformationLinks()([]ActionUrl) {
    if m == nil {
        return nil
    } else {
        return m.informationLinks
    }
}
// GetLastActionByUserId gets the lastActionByUserId property value. 
func (m *ManagementTemplate) GetLastActionByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActionByUserId
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. 
func (m *ManagementTemplate) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. 
func (m *ManagementTemplate) GetManagementTemplateCollections()([]ManagementTemplateCollection) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateCollections
    }
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. 
func (m *ManagementTemplate) GetManagementTemplateSteps()([]ManagementTemplateStep) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateSteps
    }
}
// GetParameters gets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) GetParameters()([]TemplateParameter) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// GetPriority gets the priority property value. 
func (m *ManagementTemplate) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetProvider gets the provider property value. 
func (m *ManagementTemplate) GetProvider()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// GetUserImpact gets the userImpact property value. 
func (m *ManagementTemplate) GetUserImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userImpact
    }
}
// GetVersion gets the version property value. 
func (m *ManagementTemplate) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) GetWorkloadActions()([]WorkloadAction) {
    if m == nil {
        return nil
    } else {
        return m.workloadActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory))
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
    res["informationLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActionUrl() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionUrl, len(val))
            for i, v := range val {
                res[i] = *(v.(*ActionUrl))
            }
            m.SetInformationLinks(res)
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
    res["managementTemplateCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollection, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateCollection))
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStep))
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["parameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemplateParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemplateParameter, len(val))
            for i, v := range val {
                res[i] = *(v.(*TemplateParameter))
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["provider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider))
        }
        return nil
    }
    res["userImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserImpact(val)
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
    res["workloadActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkloadAction))
            }
            m.SetWorkloadActions(res)
        }
        return nil
    }
    return res
}
func (m *ManagementTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
    if m.GetInformationLinks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInformationLinks()))
        for i, v := range m.GetInformationLinks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("informationLinks", cast)
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
    if m.GetManagementTemplateCollections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetParameters() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetProvider() != nil {
        cast := (*m.GetProvider()).String()
        err = writer.WriteStringValue("provider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userImpact", m.GetUserImpact())
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
    if m.GetWorkloadActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkloadActions()))
        for i, v := range m.GetWorkloadActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("workloadActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCreatedByUserId sets the createdByUserId property value. 
func (m *ManagementTemplate) SetCreatedByUserId(value *string)() {
    if m != nil {
        m.createdByUserId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *ManagementTemplate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInformationLinks sets the informationLinks property value. 
func (m *ManagementTemplate) SetInformationLinks(value []ActionUrl)() {
    if m != nil {
        m.informationLinks = value
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. 
func (m *ManagementTemplate) SetLastActionByUserId(value *string)() {
    if m != nil {
        m.lastActionByUserId = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. 
func (m *ManagementTemplate) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. 
func (m *ManagementTemplate) SetManagementTemplateCollections(value []ManagementTemplateCollection)() {
    if m != nil {
        m.managementTemplateCollections = value
    }
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. 
func (m *ManagementTemplate) SetManagementTemplateSteps(value []ManagementTemplateStep)() {
    if m != nil {
        m.managementTemplateSteps = value
    }
}
// SetParameters sets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) SetParameters(value []TemplateParameter)() {
    if m != nil {
        m.parameters = value
    }
}
// SetPriority sets the priority property value. 
func (m *ManagementTemplate) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetProvider sets the provider property value. 
func (m *ManagementTemplate) SetProvider(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)() {
    if m != nil {
        m.provider = value
    }
}
// SetUserImpact sets the userImpact property value. 
func (m *ManagementTemplate) SetUserImpact(value *string)() {
    if m != nil {
        m.userImpact = value
    }
}
// SetVersion sets the version property value. 
func (m *ManagementTemplate) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) SetWorkloadActions(value []WorkloadAction)() {
    if m != nil {
        m.workloadActions = value
    }
}
