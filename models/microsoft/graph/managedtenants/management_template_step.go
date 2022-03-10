package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ManagementTemplateStep provides operations to manage the tenantRelationship singleton.
type ManagementTemplateStep struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    acceptedVersion ManagementTemplateStepVersionable;
    // 
    category *ManagementCategory;
    // 
    createdByUserId *string;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    description *string;
    // 
    displayName *string;
    // 
    lastActionByUserId *string;
    // 
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    managementTemplate ManagementTemplateable;
    // 
    portalLink i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable;
    // 
    priority *int32;
    // 
    versions []ManagementTemplateStepVersionable;
}
// NewManagementTemplateStep instantiates a new managementTemplateStep and sets the default values.
func NewManagementTemplateStep()(*ManagementTemplateStep) {
    m := &ManagementTemplateStep{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagementTemplateStep(), nil
}
// GetAcceptedVersion gets the acceptedVersion property value. 
func (m *ManagementTemplateStep) GetAcceptedVersion()(ManagementTemplateStepVersionable) {
    if m == nil {
        return nil
    } else {
        return m.acceptedVersion
    }
}
// GetCategory gets the category property value. 
func (m *ManagementTemplateStep) GetCategory()(*ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCreatedByUserId gets the createdByUserId property value. 
func (m *ManagementTemplateStep) GetCreatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByUserId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *ManagementTemplateStep) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. 
func (m *ManagementTemplateStep) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *ManagementTemplateStep) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedVersion(val.(ManagementTemplateStepVersionable))
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ManagementCategory))
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
    res["managementTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplate(val.(ManagementTemplateable))
        }
        return nil
    }
    res["portalLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateActionUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalLink(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable))
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
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. 
func (m *ManagementTemplateStep) GetLastActionByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActionByUserId
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. 
func (m *ManagementTemplateStep) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetManagementTemplate gets the managementTemplate property value. 
func (m *ManagementTemplateStep) GetManagementTemplate()(ManagementTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplate
    }
}
// GetPortalLink gets the portalLink property value. 
func (m *ManagementTemplateStep) GetPortalLink()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable) {
    if m == nil {
        return nil
    } else {
        return m.portalLink
    }
}
// GetPriority gets the priority property value. 
func (m *ManagementTemplateStep) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetVersions gets the versions property value. 
func (m *ManagementTemplateStep) GetVersions()([]ManagementTemplateStepVersionable) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
func (m *ManagementTemplateStep) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedVersion", m.GetAcceptedVersion())
        if err != nil {
            return err
        }
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
        err = writer.WriteObjectValue("managementTemplate", m.GetManagementTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("portalLink", m.GetPortalLink())
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
    if m.GetVersions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedVersion sets the acceptedVersion property value. 
func (m *ManagementTemplateStep) SetAcceptedVersion(value ManagementTemplateStepVersionable)() {
    if m != nil {
        m.acceptedVersion = value
    }
}
// SetCategory sets the category property value. 
func (m *ManagementTemplateStep) SetCategory(value *ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCreatedByUserId sets the createdByUserId property value. 
func (m *ManagementTemplateStep) SetCreatedByUserId(value *string)() {
    if m != nil {
        m.createdByUserId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *ManagementTemplateStep) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. 
func (m *ManagementTemplateStep) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *ManagementTemplateStep) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. 
func (m *ManagementTemplateStep) SetLastActionByUserId(value *string)() {
    if m != nil {
        m.lastActionByUserId = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. 
func (m *ManagementTemplateStep) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetManagementTemplate sets the managementTemplate property value. 
func (m *ManagementTemplateStep) SetManagementTemplate(value ManagementTemplateable)() {
    if m != nil {
        m.managementTemplate = value
    }
}
// SetPortalLink sets the portalLink property value. 
func (m *ManagementTemplateStep) SetPortalLink(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable)() {
    if m != nil {
        m.portalLink = value
    }
}
// SetPriority sets the priority property value. 
func (m *ManagementTemplateStep) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetVersions sets the versions property value. 
func (m *ManagementTemplateStep) SetVersions(value []ManagementTemplateStepVersionable)() {
    if m != nil {
        m.versions = value
    }
}
