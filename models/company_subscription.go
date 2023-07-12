package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanySubscription 
type CompanySubscription struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewCompanySubscription instantiates a new companySubscription and sets the default values.
func NewCompanySubscription()(*CompanySubscription) {
    m := &CompanySubscription{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCompanySubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanySubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanySubscription(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *CompanySubscription) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanySubscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isTrial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTrial(val)
        }
        return nil
    }
    res["nextLifecycleDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextLifecycleDateTime(val)
        }
        return nil
    }
    res["ocpSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOcpSubscriptionId(val)
        }
        return nil
    }
    res["serviceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePlanInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePlanInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServicePlanInfoable)
                }
            }
            m.SetServiceStatus(res)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["skuPartNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuPartNumber(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenses(val)
        }
        return nil
    }
    return res
}
// GetIsTrial gets the isTrial property value. The isTrial property
func (m *CompanySubscription) GetIsTrial()(*bool) {
    val, err := m.GetBackingStore().Get("isTrial")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetNextLifecycleDateTime gets the nextLifecycleDateTime property value. The nextLifecycleDateTime property
func (m *CompanySubscription) GetNextLifecycleDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("nextLifecycleDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOcpSubscriptionId gets the ocpSubscriptionId property value. The ocpSubscriptionId property
func (m *CompanySubscription) GetOcpSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("ocpSubscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServiceStatus gets the serviceStatus property value. The serviceStatus property
func (m *CompanySubscription) GetServiceStatus()([]ServicePlanInfoable) {
    val, err := m.GetBackingStore().Get("serviceStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicePlanInfoable)
    }
    return nil
}
// GetSkuId gets the skuId property value. The skuId property
func (m *CompanySubscription) GetSkuId()(*string) {
    val, err := m.GetBackingStore().Get("skuId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkuPartNumber gets the skuPartNumber property value. The skuPartNumber property
func (m *CompanySubscription) GetSkuPartNumber()(*string) {
    val, err := m.GetBackingStore().Get("skuPartNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *CompanySubscription) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalLicenses gets the totalLicenses property value. The totalLicenses property
func (m *CompanySubscription) GetTotalLicenses()(*int32) {
    val, err := m.GetBackingStore().Get("totalLicenses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CompanySubscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTrial", m.GetIsTrial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextLifecycleDateTime", m.GetNextLifecycleDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ocpSubscriptionId", m.GetOcpSubscriptionId())
        if err != nil {
            return err
        }
    }
    if m.GetServiceStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceStatus()))
        for i, v := range m.GetServiceStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("serviceStatus", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuPartNumber", m.GetSkuPartNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenses", m.GetTotalLicenses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *CompanySubscription) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTrial sets the isTrial property value. The isTrial property
func (m *CompanySubscription) SetIsTrial(value *bool)() {
    err := m.GetBackingStore().Set("isTrial", value)
    if err != nil {
        panic(err)
    }
}
// SetNextLifecycleDateTime sets the nextLifecycleDateTime property value. The nextLifecycleDateTime property
func (m *CompanySubscription) SetNextLifecycleDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("nextLifecycleDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOcpSubscriptionId sets the ocpSubscriptionId property value. The ocpSubscriptionId property
func (m *CompanySubscription) SetOcpSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("ocpSubscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceStatus sets the serviceStatus property value. The serviceStatus property
func (m *CompanySubscription) SetServiceStatus(value []ServicePlanInfoable)() {
    err := m.GetBackingStore().Set("serviceStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuId sets the skuId property value. The skuId property
func (m *CompanySubscription) SetSkuId(value *string)() {
    err := m.GetBackingStore().Set("skuId", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuPartNumber sets the skuPartNumber property value. The skuPartNumber property
func (m *CompanySubscription) SetSkuPartNumber(value *string)() {
    err := m.GetBackingStore().Set("skuPartNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *CompanySubscription) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicenses sets the totalLicenses property value. The totalLicenses property
func (m *CompanySubscription) SetTotalLicenses(value *int32)() {
    err := m.GetBackingStore().Set("totalLicenses", value)
    if err != nil {
        panic(err)
    }
}
// CompanySubscriptionable 
type CompanySubscriptionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsTrial()(*bool)
    GetNextLifecycleDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOcpSubscriptionId()(*string)
    GetServiceStatus()([]ServicePlanInfoable)
    GetSkuId()(*string)
    GetSkuPartNumber()(*string)
    GetStatus()(*string)
    GetTotalLicenses()(*int32)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsTrial(value *bool)()
    SetNextLifecycleDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOcpSubscriptionId(value *string)()
    SetServiceStatus(value []ServicePlanInfoable)()
    SetSkuId(value *string)()
    SetSkuPartNumber(value *string)()
    SetStatus(value *string)()
    SetTotalLicenses(value *int32)()
}
