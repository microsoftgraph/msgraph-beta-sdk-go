package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomersSpendingBudgetSecurityRequirement struct {
    SecurityRequirement
}
// NewCustomersSpendingBudgetSecurityRequirement instantiates a new CustomersSpendingBudgetSecurityRequirement and sets the default values.
func NewCustomersSpendingBudgetSecurityRequirement()(*CustomersSpendingBudgetSecurityRequirement) {
    m := &CustomersSpendingBudgetSecurityRequirement{
        SecurityRequirement: *NewSecurityRequirement(),
    }
    return m
}
// CreateCustomersSpendingBudgetSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomersSpendingBudgetSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomersSpendingBudgetSecurityRequirement(), nil
}
// GetCustomersWithSpendBudgetCount gets the customersWithSpendBudgetCount property value. The number of customers with a spending budget set.
// returns a *int64 when successful
func (m *CustomersSpendingBudgetSecurityRequirement) GetCustomersWithSpendBudgetCount()(*int64) {
    val, err := m.GetBackingStore().Get("customersWithSpendBudgetCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomersSpendingBudgetSecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityRequirement.GetFieldDeserializers()
    res["customersWithSpendBudgetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomersWithSpendBudgetCount(val)
        }
        return nil
    }
    res["totalCustomersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCustomersCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCustomersCount gets the totalCustomersCount property value. The total number of customers associated with the partner.
// returns a *int64 when successful
func (m *CustomersSpendingBudgetSecurityRequirement) GetTotalCustomersCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalCustomersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomersSpendingBudgetSecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("customersWithSpendBudgetCount", m.GetCustomersWithSpendBudgetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalCustomersCount", m.GetTotalCustomersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomersWithSpendBudgetCount sets the customersWithSpendBudgetCount property value. The number of customers with a spending budget set.
func (m *CustomersSpendingBudgetSecurityRequirement) SetCustomersWithSpendBudgetCount(value *int64)() {
    err := m.GetBackingStore().Set("customersWithSpendBudgetCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCustomersCount sets the totalCustomersCount property value. The total number of customers associated with the partner.
func (m *CustomersSpendingBudgetSecurityRequirement) SetTotalCustomersCount(value *int64)() {
    err := m.GetBackingStore().Set("totalCustomersCount", value)
    if err != nil {
        panic(err)
    }
}
type CustomersSpendingBudgetSecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecurityRequirementable
    GetCustomersWithSpendBudgetCount()(*int64)
    GetTotalCustomersCount()(*int64)
    SetCustomersWithSpendBudgetCount(value *int64)()
    SetTotalCustomersCount(value *int64)()
}
