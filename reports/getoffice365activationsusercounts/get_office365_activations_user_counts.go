package getoffice365activationsusercounts

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365ActivationsUserCounts struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    activated *int64;
    assigned *int64;
    productType *string;
    reportRefreshDate *string;
    sharedComputerActivation *int64;
}
func NewGetOffice365ActivationsUserCounts()(*GetOffice365ActivationsUserCounts) {
    m := &GetOffice365ActivationsUserCounts{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365ActivationsUserCounts) GetActivated()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activated
    }
}
func (m *GetOffice365ActivationsUserCounts) GetAssigned()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.assigned
    }
}
func (m *GetOffice365ActivationsUserCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
func (m *GetOffice365ActivationsUserCounts) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365ActivationsUserCounts) GetSharedComputerActivation()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedComputerActivation
    }
}
func (m *GetOffice365ActivationsUserCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetActivated(val)
        return nil
    }
    res["assigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAssigned(val)
        return nil
    }
    res["productType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductType(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["sharedComputerActivation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedComputerActivation(val)
        return nil
    }
    return res
}
func (m *GetOffice365ActivationsUserCounts) IsNil()(bool) {
    return m == nil
}
func (m *GetOffice365ActivationsUserCounts) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("activated", m.GetActivated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("assigned", m.GetAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productType", m.GetProductType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedComputerActivation", m.GetSharedComputerActivation())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetOffice365ActivationsUserCounts) SetActivated(value *int64)() {
    m.activated = value
}
func (m *GetOffice365ActivationsUserCounts) SetAssigned(value *int64)() {
    m.assigned = value
}
func (m *GetOffice365ActivationsUserCounts) SetProductType(value *string)() {
    m.productType = value
}
func (m *GetOffice365ActivationsUserCounts) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365ActivationsUserCounts) SetSharedComputerActivation(value *int64)() {
    m.sharedComputerActivation = value
}
