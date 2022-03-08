package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// SourceCollection provides operations to manage the compliance singleton.
type SourceCollection struct {
    Entity
    // Adds an additional source to the sourceCollection.
    additionalSources []DataSourceable;
    // Adds the results of the sourceCollection to the specified reviewSet.
    addToReviewSetOperation AddToReviewSetOperationable;
    // The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
    contentQuery *string;
    // The user who created the sourceCollection.
    createdBy IdentitySetable;
    // The date and time the sourceCollection was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Custodian sources that are included in the sourceCollection.
    custodianSources []DataSourceable;
    // When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
    dataSourceScopes *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes;
    // The description of the sourceCollection.
    description *string;
    // The display name of the sourceCollection.
    displayName *string;
    // The last estimate operation associated with the sourceCollection.
    lastEstimateStatisticsOperation EstimateStatisticsOperationable;
    // The last user who modified the sourceCollection.
    lastModifiedBy IdentitySetable;
    // The last date and time the sourceCollection was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // noncustodialDataSource sources that are included in the sourceCollection
    noncustodialSources []NoncustodialDataSourceable;
}
// NewSourceCollection instantiates a new sourceCollection and sets the default values.
func NewSourceCollection()(*SourceCollection) {
    m := &SourceCollection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSourceCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSourceCollectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSourceCollection(), nil
}
// GetAdditionalSources gets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) GetAdditionalSources()([]DataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.additionalSources
    }
}
// GetAddToReviewSetOperation gets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) GetAddToReviewSetOperation()(AddToReviewSetOperationable) {
    if m == nil {
        return nil
    } else {
        return m.addToReviewSetOperation
    }
}
// GetContentQuery gets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) GetContentQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentQuery
    }
}
// GetCreatedBy gets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustodianSources gets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) GetCustodianSources()([]DataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.custodianSources
    }
}
// GetDataSourceScopes gets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) GetDataSourceScopes()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes) {
    if m == nil {
        return nil
    } else {
        return m.dataSourceScopes
    }
}
// GetDescription gets the description property value. The description of the sourceCollection.
func (m *SourceCollection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SourceCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetAdditionalSources(res)
        }
        return nil
    }
    res["addToReviewSetOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAddToReviewSetOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToReviewSetOperation(val.(AddToReviewSetOperationable))
        }
        return nil
    }
    res["contentQuery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["custodianSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetCustodianSources(res)
        }
        return nil
    }
    res["dataSourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseDataSourceScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSourceScopes(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes))
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
    res["lastEstimateStatisticsOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEstimateStatisticsOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEstimateStatisticsOperation(val.(EstimateStatisticsOperationable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["noncustodialSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NoncustodialDataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(NoncustodialDataSourceable)
            }
            m.SetNoncustodialSources(res)
        }
        return nil
    }
    return res
}
// GetLastEstimateStatisticsOperation gets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable) {
    if m == nil {
        return nil
    } else {
        return m.lastEstimateStatisticsOperation
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) GetLastModifiedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNoncustodialSources gets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) GetNoncustodialSources()([]NoncustodialDataSourceable) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialSources
    }
}
func (m *SourceCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SourceCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalSources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalSources()))
        for i, v := range m.GetAdditionalSources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("additionalSources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("addToReviewSetOperation", m.GetAddToReviewSetOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
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
    if m.GetCustodianSources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustodianSources()))
        for i, v := range m.GetCustodianSources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("custodianSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataSourceScopes() != nil {
        cast := (*m.GetDataSourceScopes()).String()
        err = writer.WriteStringValue("dataSourceScopes", &cast)
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
        err = writer.WriteObjectValue("lastEstimateStatisticsOperation", m.GetLastEstimateStatisticsOperation())
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
    if m.GetNoncustodialSources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNoncustodialSources()))
        for i, v := range m.GetNoncustodialSources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalSources sets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) SetAdditionalSources(value []DataSourceable)() {
    if m != nil {
        m.additionalSources = value
    }
}
// SetAddToReviewSetOperation sets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) SetAddToReviewSetOperation(value AddToReviewSetOperationable)() {
    if m != nil {
        m.addToReviewSetOperation = value
    }
}
// SetContentQuery sets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) SetContentQuery(value *string)() {
    if m != nil {
        m.contentQuery = value
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustodianSources sets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) SetCustodianSources(value []DataSourceable)() {
    if m != nil {
        m.custodianSources = value
    }
}
// SetDataSourceScopes sets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) SetDataSourceScopes(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.DataSourceScopes)() {
    if m != nil {
        m.dataSourceScopes = value
    }
}
// SetDescription sets the description property value. The description of the sourceCollection.
func (m *SourceCollection) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastEstimateStatisticsOperation sets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)() {
    if m != nil {
        m.lastEstimateStatisticsOperation = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) SetLastModifiedBy(value IdentitySetable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNoncustodialSources sets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) SetNoncustodialSources(value []NoncustodialDataSourceable)() {
    if m != nil {
        m.noncustodialSources = value
    }
}
