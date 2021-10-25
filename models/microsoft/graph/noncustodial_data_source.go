package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type NoncustodialDataSource struct {
    DataSourceContainer
    applyHoldToSource *bool;
    dataSource *DataSource;
}
func NewNoncustodialDataSource()(*NoncustodialDataSource) {
    m := &NoncustodialDataSource{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    return m
}
func (m *NoncustodialDataSource) GetApplyHoldToSource()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyHoldToSource
    }
}
func (m *NoncustodialDataSource) GetDataSource()(*DataSource) {
    if m == nil {
        return nil
    } else {
        return m.dataSource
    }
}
func (m *NoncustodialDataSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["applyHoldToSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplyHoldToSource(val)
        return nil
    }
    res["dataSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSource() })
        if err != nil {
            return err
        }
        m.SetDataSource(val.(*DataSource))
        return nil
    }
    return res
}
func (m *NoncustodialDataSource) IsNil()(bool) {
    return m == nil
}
func (m *NoncustodialDataSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applyHoldToSource", m.GetApplyHoldToSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataSource", m.GetDataSource())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *NoncustodialDataSource) SetApplyHoldToSource(value *bool)() {
    m.applyHoldToSource = value
}
func (m *NoncustodialDataSource) SetDataSource(value *DataSource)() {
    m.dataSource = value
}
