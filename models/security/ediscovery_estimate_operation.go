package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryEstimateOperation 
type EdiscoveryEstimateOperation struct {
    CaseOperation
    // The indexedItemCount property
    indexedItemCount *int64
    // The indexedItemsSize property
    indexedItemsSize *int64
    // The mailboxCount property
    mailboxCount *int32
    // The search property
    search EdiscoverySearchable
    // The siteCount property
    siteCount *int32
    // The unindexedItemCount property
    unindexedItemCount *int64
    // The unindexedItemsSize property
    unindexedItemsSize *int64
}
// NewEdiscoveryEstimateOperation instantiates a new ediscoveryEstimateOperation and sets the default values.
func NewEdiscoveryEstimateOperation()(*EdiscoveryEstimateOperation) {
    m := &EdiscoveryEstimateOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryEstimateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryEstimateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryEstimateOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryEstimateOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
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
// GetIndexedItemCount gets the indexedItemCount property value. The indexedItemCount property
func (m *EdiscoveryEstimateOperation) GetIndexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemCount
    }
}
// GetIndexedItemsSize gets the indexedItemsSize property value. The indexedItemsSize property
func (m *EdiscoveryEstimateOperation) GetIndexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemsSize
    }
}
// GetMailboxCount gets the mailboxCount property value. The mailboxCount property
func (m *EdiscoveryEstimateOperation) GetMailboxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mailboxCount
    }
}
// GetSearch gets the search property value. The search property
func (m *EdiscoveryEstimateOperation) GetSearch()(EdiscoverySearchable) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// GetSiteCount gets the siteCount property value. The siteCount property
func (m *EdiscoveryEstimateOperation) GetSiteCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.siteCount
    }
}
// GetUnindexedItemCount gets the unindexedItemCount property value. The unindexedItemCount property
func (m *EdiscoveryEstimateOperation) GetUnindexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemCount
    }
}
// GetUnindexedItemsSize gets the unindexedItemsSize property value. The unindexedItemsSize property
func (m *EdiscoveryEstimateOperation) GetUnindexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemsSize
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryEstimateOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("search", m.GetSearch())
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
// SetIndexedItemCount sets the indexedItemCount property value. The indexedItemCount property
func (m *EdiscoveryEstimateOperation) SetIndexedItemCount(value *int64)() {
    if m != nil {
        m.indexedItemCount = value
    }
}
// SetIndexedItemsSize sets the indexedItemsSize property value. The indexedItemsSize property
func (m *EdiscoveryEstimateOperation) SetIndexedItemsSize(value *int64)() {
    if m != nil {
        m.indexedItemsSize = value
    }
}
// SetMailboxCount sets the mailboxCount property value. The mailboxCount property
func (m *EdiscoveryEstimateOperation) SetMailboxCount(value *int32)() {
    if m != nil {
        m.mailboxCount = value
    }
}
// SetSearch sets the search property value. The search property
func (m *EdiscoveryEstimateOperation) SetSearch(value EdiscoverySearchable)() {
    if m != nil {
        m.search = value
    }
}
// SetSiteCount sets the siteCount property value. The siteCount property
func (m *EdiscoveryEstimateOperation) SetSiteCount(value *int32)() {
    if m != nil {
        m.siteCount = value
    }
}
// SetUnindexedItemCount sets the unindexedItemCount property value. The unindexedItemCount property
func (m *EdiscoveryEstimateOperation) SetUnindexedItemCount(value *int64)() {
    if m != nil {
        m.unindexedItemCount = value
    }
}
// SetUnindexedItemsSize sets the unindexedItemsSize property value. The unindexedItemsSize property
func (m *EdiscoveryEstimateOperation) SetUnindexedItemsSize(value *int64)() {
    if m != nil {
        m.unindexedItemsSize = value
    }
}
