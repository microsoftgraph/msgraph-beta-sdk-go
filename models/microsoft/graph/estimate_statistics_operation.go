package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EstimateStatisticsOperation struct {
    CaseOperation
    indexedItemCount *int64;
    indexedItemsSize *int64;
    mailboxCount *int32;
    siteCount *int32;
    sourceCollection *SourceCollection;
    unindexedItemCount *int64;
    unindexedItemsSize *int64;
}
func NewEstimateStatisticsOperation()(*EstimateStatisticsOperation) {
    m := &EstimateStatisticsOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
func (m *EstimateStatisticsOperation) GetIndexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemCount
    }
}
func (m *EstimateStatisticsOperation) GetIndexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemsSize
    }
}
func (m *EstimateStatisticsOperation) GetMailboxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mailboxCount
    }
}
func (m *EstimateStatisticsOperation) GetSiteCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.siteCount
    }
}
func (m *EstimateStatisticsOperation) GetSourceCollection()(*SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollection
    }
}
func (m *EstimateStatisticsOperation) GetUnindexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemCount
    }
}
func (m *EstimateStatisticsOperation) GetUnindexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemsSize
    }
}
func (m *EstimateStatisticsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["indexedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIndexedItemCount(val)
        return nil
    }
    res["indexedItemsSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIndexedItemsSize(val)
        return nil
    }
    res["mailboxCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMailboxCount(val)
        return nil
    }
    res["siteCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSiteCount(val)
        return nil
    }
    res["sourceCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        m.SetSourceCollection(val.(*SourceCollection))
        return nil
    }
    res["unindexedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUnindexedItemCount(val)
        return nil
    }
    res["unindexedItemsSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUnindexedItemsSize(val)
        return nil
    }
    return res
}
func (m *EstimateStatisticsOperation) IsNil()(bool) {
    return m == nil
}
func (m *EstimateStatisticsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *EstimateStatisticsOperation) SetIndexedItemCount(value *int64)() {
    m.indexedItemCount = value
}
func (m *EstimateStatisticsOperation) SetIndexedItemsSize(value *int64)() {
    m.indexedItemsSize = value
}
func (m *EstimateStatisticsOperation) SetMailboxCount(value *int32)() {
    m.mailboxCount = value
}
func (m *EstimateStatisticsOperation) SetSiteCount(value *int32)() {
    m.siteCount = value
}
func (m *EstimateStatisticsOperation) SetSourceCollection(value *SourceCollection)() {
    m.sourceCollection = value
}
func (m *EstimateStatisticsOperation) SetUnindexedItemCount(value *int64)() {
    m.unindexedItemCount = value
}
func (m *EstimateStatisticsOperation) SetUnindexedItemsSize(value *int64)() {
    m.unindexedItemsSize = value
}
