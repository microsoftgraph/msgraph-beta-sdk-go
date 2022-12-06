package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainSecurityProfile provides operations to manage the collection of accessReviewDecision entities.
type DomainSecurityProfile struct {
    Entity
    // The activityGroupNames property
    activityGroupNames []string
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The countHits property
    countHits *int32
    // The countInOrg property
    countInOrg *int32
    // The domainCategories property
    domainCategories []ReputationCategoryable
    // The domainRegisteredDateTime property
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The registrant property
    registrant DomainRegistrantable
    // The riskScore property
    riskScore *string
    // The tags property
    tags []string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewDomainSecurityProfile instantiates a new domainSecurityProfile and sets the default values.
func NewDomainSecurityProfile()(*DomainSecurityProfile) {
    m := &DomainSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDomainSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainSecurityProfile(), nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The activityGroupNames property
func (m *DomainSecurityProfile) GetActivityGroupNames()([]string) {
    return m.activityGroupNames
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *DomainSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *DomainSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetCountHits gets the countHits property value. The countHits property
func (m *DomainSecurityProfile) GetCountHits()(*int32) {
    return m.countHits
}
// GetCountInOrg gets the countInOrg property value. The countInOrg property
func (m *DomainSecurityProfile) GetCountInOrg()(*int32) {
    return m.countInOrg
}
// GetDomainCategories gets the domainCategories property value. The domainCategories property
func (m *DomainSecurityProfile) GetDomainCategories()([]ReputationCategoryable) {
    return m.domainCategories
}
// GetDomainRegisteredDateTime gets the domainRegisteredDateTime property value. The domainRegisteredDateTime property
func (m *DomainSecurityProfile) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.domainRegisteredDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityGroupNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetActivityGroupNames)
    res["azureSubscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureSubscriptionId)
    res["azureTenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureTenantId)
    res["countHits"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCountHits)
    res["countInOrg"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCountInOrg)
    res["domainCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateReputationCategoryFromDiscriminatorValue , m.SetDomainCategories)
    res["domainRegisteredDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDomainRegisteredDateTime)
    res["firstSeenDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetFirstSeenDateTime)
    res["lastSeenDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSeenDateTime)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["registrant"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDomainRegistrantFromDiscriminatorValue , m.SetRegistrant)
    res["riskScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRiskScore)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTags)
    res["vendorInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue , m.SetVendorInformation)
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DomainSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *DomainSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetName gets the name property value. The name property
func (m *DomainSecurityProfile) GetName()(*string) {
    return m.name
}
// GetRegistrant gets the registrant property value. The registrant property
func (m *DomainSecurityProfile) GetRegistrant()(DomainRegistrantable) {
    return m.registrant
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *DomainSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetTags gets the tags property value. The tags property
func (m *DomainSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *DomainSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *DomainSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityGroupNames() != nil {
        err = writer.WriteCollectionOfStringValues("activityGroupNames", m.GetActivityGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countHits", m.GetCountHits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countInOrg", m.GetCountInOrg())
        if err != nil {
            return err
        }
    }
    if m.GetDomainCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDomainCategories())
        err = writer.WriteCollectionOfObjectValues("domainCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("domainRegisteredDateTime", m.GetDomainRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
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
        err = writer.WriteObjectValue("registrant", m.GetRegistrant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
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
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityGroupNames sets the activityGroupNames property value. The activityGroupNames property
func (m *DomainSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *DomainSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *DomainSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetCountHits sets the countHits property value. The countHits property
func (m *DomainSecurityProfile) SetCountHits(value *int32)() {
    m.countHits = value
}
// SetCountInOrg sets the countInOrg property value. The countInOrg property
func (m *DomainSecurityProfile) SetCountInOrg(value *int32)() {
    m.countInOrg = value
}
// SetDomainCategories sets the domainCategories property value. The domainCategories property
func (m *DomainSecurityProfile) SetDomainCategories(value []ReputationCategoryable)() {
    m.domainCategories = value
}
// SetDomainRegisteredDateTime sets the domainRegisteredDateTime property value. The domainRegisteredDateTime property
func (m *DomainSecurityProfile) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DomainSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *DomainSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetName sets the name property value. The name property
func (m *DomainSecurityProfile) SetName(value *string)() {
    m.name = value
}
// SetRegistrant sets the registrant property value. The registrant property
func (m *DomainSecurityProfile) SetRegistrant(value DomainRegistrantable)() {
    m.registrant = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *DomainSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetTags sets the tags property value. The tags property
func (m *DomainSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *DomainSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
