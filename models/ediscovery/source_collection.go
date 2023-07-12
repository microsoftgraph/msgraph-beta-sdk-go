package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SourceCollection 
type SourceCollection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewSourceCollection instantiates a new sourceCollection and sets the default values.
func NewSourceCollection()(*SourceCollection) {
    m := &SourceCollection{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSourceCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSourceCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSourceCollection(), nil
}
// GetAdditionalSources gets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) GetAdditionalSources()([]DataSourceable) {
    val, err := m.GetBackingStore().Get("additionalSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DataSourceable)
    }
    return nil
}
// GetAddToReviewSetOperation gets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) GetAddToReviewSetOperation()(AddToReviewSetOperationable) {
    val, err := m.GetBackingStore().Get("addToReviewSetOperation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AddToReviewSetOperationable)
    }
    return nil
}
// GetContentQuery gets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) GetContentQuery()(*string) {
    val, err := m.GetBackingStore().Get("contentQuery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustodianSources gets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) GetCustodianSources()([]DataSourceable) {
    val, err := m.GetBackingStore().Get("custodianSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DataSourceable)
    }
    return nil
}
// GetDataSourceScopes gets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) GetDataSourceScopes()(*DataSourceScopes) {
    val, err := m.GetBackingStore().Get("dataSourceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DataSourceScopes)
    }
    return nil
}
// GetDescription gets the description property value. The description of the sourceCollection.
func (m *SourceCollection) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SourceCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DataSourceable)
                }
            }
            m.SetAdditionalSources(res)
        }
        return nil
    }
    res["addToReviewSetOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAddToReviewSetOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToReviewSetOperation(val.(AddToReviewSetOperationable))
        }
        return nil
    }
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["custodianSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DataSourceable)
                }
            }
            m.SetCustodianSources(res)
        }
        return nil
    }
    res["dataSourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSourceScopes(val.(*DataSourceScopes))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastEstimateStatisticsOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEstimateStatisticsOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEstimateStatisticsOperation(val.(EstimateStatisticsOperationable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["noncustodialSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NoncustodialDataSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NoncustodialDataSourceable)
                }
            }
            m.SetNoncustodialSources(res)
        }
        return nil
    }
    return res
}
// GetLastEstimateStatisticsOperation gets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable) {
    val, err := m.GetBackingStore().Get("lastEstimateStatisticsOperation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EstimateStatisticsOperationable)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNoncustodialSources gets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) GetNoncustodialSources()([]NoncustodialDataSourceable) {
    val, err := m.GetBackingStore().Get("noncustodialSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NoncustodialDataSourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SourceCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalSources()))
        for i, v := range m.GetAdditionalSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodianSources()))
        for i, v := range m.GetCustodianSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialSources()))
        for i, v := range m.GetNoncustodialSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("additionalSources", value)
    if err != nil {
        panic(err)
    }
}
// SetAddToReviewSetOperation sets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) SetAddToReviewSetOperation(value AddToReviewSetOperationable)() {
    err := m.GetBackingStore().Set("addToReviewSetOperation", value)
    if err != nil {
        panic(err)
    }
}
// SetContentQuery sets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) SetContentQuery(value *string)() {
    err := m.GetBackingStore().Set("contentQuery", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustodianSources sets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) SetCustodianSources(value []DataSourceable)() {
    err := m.GetBackingStore().Set("custodianSources", value)
    if err != nil {
        panic(err)
    }
}
// SetDataSourceScopes sets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) SetDataSourceScopes(value *DataSourceScopes)() {
    err := m.GetBackingStore().Set("dataSourceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the sourceCollection.
func (m *SourceCollection) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastEstimateStatisticsOperation sets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)() {
    err := m.GetBackingStore().Set("lastEstimateStatisticsOperation", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNoncustodialSources sets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) SetNoncustodialSources(value []NoncustodialDataSourceable)() {
    err := m.GetBackingStore().Set("noncustodialSources", value)
    if err != nil {
        panic(err)
    }
}
// SourceCollectionable 
type SourceCollectionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalSources()([]DataSourceable)
    GetAddToReviewSetOperation()(AddToReviewSetOperationable)
    GetContentQuery()(*string)
    GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodianSources()([]DataSourceable)
    GetDataSourceScopes()(*DataSourceScopes)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable)
    GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoncustodialSources()([]NoncustodialDataSourceable)
    SetAdditionalSources(value []DataSourceable)()
    SetAddToReviewSetOperation(value AddToReviewSetOperationable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodianSources(value []DataSourceable)()
    SetDataSourceScopes(value *DataSourceScopes)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)()
    SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoncustodialSources(value []NoncustodialDataSourceable)()
}
