package getoffice365activationsusercounts

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365ActivationsUserCounts 
type GetOffice365ActivationsUserCounts struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of users who have activated the product.
    activated *int64;
    // The number of users have been assigned for the product license.
    assigned *int64;
    // The product type such as 'Microsoft 365 ProPlus' or 'Project Client'.
    productType *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of users who have used the product on a shared computer.
    sharedComputerActivation *int64;
}
// NewGetOffice365ActivationsUserCounts instantiates a new getOffice365ActivationsUserCounts and sets the default values.
func NewGetOffice365ActivationsUserCounts()(*GetOffice365ActivationsUserCounts) {
    m := &GetOffice365ActivationsUserCounts{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActivated gets the activated property value. The number of users who have activated the product.
func (m *GetOffice365ActivationsUserCounts) GetActivated()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activated
    }
}
// GetAssigned gets the assigned property value. The number of users have been assigned for the product license.
func (m *GetOffice365ActivationsUserCounts) GetAssigned()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.assigned
    }
}
// GetProductType gets the productType property value. The product type such as 'Microsoft 365 ProPlus' or 'Project Client'.
func (m *GetOffice365ActivationsUserCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActivationsUserCounts) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharedComputerActivation gets the sharedComputerActivation property value. The number of users who have used the product on a shared computer.
func (m *GetOffice365ActivationsUserCounts) GetSharedComputerActivation()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedComputerActivation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365ActivationsUserCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivated(val)
        }
        return nil
    }
    res["assigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigned(val)
        }
        return nil
    }
    res["productType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductType(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharedComputerActivation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedComputerActivation(val)
        }
        return nil
    }
    return res
}
func (m *GetOffice365ActivationsUserCounts) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActivated sets the activated property value. The number of users who have activated the product.
func (m *GetOffice365ActivationsUserCounts) SetActivated(value *int64)() {
    m.activated = value
}
// SetAssigned sets the assigned property value. The number of users have been assigned for the product license.
func (m *GetOffice365ActivationsUserCounts) SetAssigned(value *int64)() {
    m.assigned = value
}
// SetProductType sets the productType property value. The product type such as 'Microsoft 365 ProPlus' or 'Project Client'.
func (m *GetOffice365ActivationsUserCounts) SetProductType(value *string)() {
    m.productType = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActivationsUserCounts) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetSharedComputerActivation sets the sharedComputerActivation property value. The number of users who have used the product on a shared computer.
func (m *GetOffice365ActivationsUserCounts) SetSharedComputerActivation(value *int64)() {
    m.sharedComputerActivation = value
}
