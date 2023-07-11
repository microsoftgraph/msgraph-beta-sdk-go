package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EstimateStatisticsOperation 
type EstimateStatisticsOperation struct {
    CaseOperation
}
// NewEstimateStatisticsOperation instantiates a new estimateStatisticsOperation and sets the default values.
func NewEstimateStatisticsOperation()(*EstimateStatisticsOperation) {
    m := &EstimateStatisticsOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEstimateStatisticsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEstimateStatisticsOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEstimateStatisticsOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EstimateStatisticsOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["indexedItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemCount(val)
        }
        return nil
    }
    res["indexedItemsSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemsSize(val)
        }
        return nil
    }
    res["mailboxCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["siteCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCount(val)
        }
        return nil
    }
    res["sourceCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(SourceCollectionable))
        }
        return nil
    }
    res["unindexedItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemCount(val)
        }
        return nil
    }
    res["unindexedItemsSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemsSize(val)
        }
        return nil
    }
    return res
}
// GetIndexedItemCount gets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemCount()(*int64) {
    val, err := m.GetBackingStore().Get("indexedItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetIndexedItemsSize gets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemsSize()(*int64) {
    val, err := m.GetBackingStore().Get("indexedItemsSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetMailboxCount gets the mailboxCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetMailboxCount()(*int32) {
    val, err := m.GetBackingStore().Get("mailboxCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EstimateStatisticsOperation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteCount gets the siteCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetSiteCount()(*int32) {
    val, err := m.GetBackingStore().Get("siteCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSourceCollection gets the sourceCollection property value. eDiscovery collection, commonly known as a search.
func (m *EstimateStatisticsOperation) GetSourceCollection()(SourceCollectionable) {
    val, err := m.GetBackingStore().Get("sourceCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SourceCollectionable)
    }
    return nil
}
// GetUnindexedItemCount gets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemCount()(*int64) {
    val, err := m.GetBackingStore().Get("unindexedItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUnindexedItemsSize gets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemsSize()(*int64) {
    val, err := m.GetBackingStore().Get("unindexedItemsSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EstimateStatisticsOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("indexedItemCount", m.GetIndexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("indexedItemsSize", m.GetIndexedItemsSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mailboxCount", m.GetMailboxCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("siteCount", m.GetSiteCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemCount", m.GetUnindexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemsSize", m.GetUnindexedItemsSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIndexedItemCount sets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) SetIndexedItemCount(value *int64)() {
    err := m.GetBackingStore().Set("indexedItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIndexedItemsSize sets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) SetIndexedItemsSize(value *int64)() {
    err := m.GetBackingStore().Set("indexedItemsSize", value)
    if err != nil {
        panic(err)
    }
}
// SetMailboxCount sets the mailboxCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) SetMailboxCount(value *int32)() {
    err := m.GetBackingStore().Set("mailboxCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EstimateStatisticsOperation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteCount sets the siteCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) SetSiteCount(value *int32)() {
    err := m.GetBackingStore().Set("siteCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceCollection sets the sourceCollection property value. eDiscovery collection, commonly known as a search.
func (m *EstimateStatisticsOperation) SetSourceCollection(value SourceCollectionable)() {
    err := m.GetBackingStore().Set("sourceCollection", value)
    if err != nil {
        panic(err)
    }
}
// SetUnindexedItemCount sets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
func (m *EstimateStatisticsOperation) SetUnindexedItemCount(value *int64)() {
    err := m.GetBackingStore().Set("unindexedItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnindexedItemsSize sets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
func (m *EstimateStatisticsOperation) SetUnindexedItemsSize(value *int64)() {
    err := m.GetBackingStore().Set("unindexedItemsSize", value)
    if err != nil {
        panic(err)
    }
}
// EstimateStatisticsOperationable 
type EstimateStatisticsOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndexedItemCount()(*int64)
    GetIndexedItemsSize()(*int64)
    GetMailboxCount()(*int32)
    GetOdataType()(*string)
    GetSiteCount()(*int32)
    GetSourceCollection()(SourceCollectionable)
    GetUnindexedItemCount()(*int64)
    GetUnindexedItemsSize()(*int64)
    SetIndexedItemCount(value *int64)()
    SetIndexedItemsSize(value *int64)()
    SetMailboxCount(value *int32)()
    SetOdataType(value *string)()
    SetSiteCount(value *int32)()
    SetSourceCollection(value SourceCollectionable)()
    SetUnindexedItemCount(value *int64)()
    SetUnindexedItemsSize(value *int64)()
}
