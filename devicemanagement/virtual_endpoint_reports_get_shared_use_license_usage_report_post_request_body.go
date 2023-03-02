package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody 
type VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
    // The select property
    SelectEscaped []string
}
// NewVirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody instantiates a new VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody and sets the default values.
func NewVirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody()(*VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) {
    m := &VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["groupBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupBy(res)
        }
        return nil
    }
    res["orderBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["reportName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcReportName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["skip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. The filter property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetFilter()(*string) {
    val, err := m.GetBackingStore().Get("filter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetGroupBy gets the groupBy property value. The groupBy property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetGroupBy()([]string) {
    val, err := m.GetBackingStore().Get("groupBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOrderBy gets the orderBy property value. The orderBy property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetOrderBy()([]string) {
    val, err := m.GetBackingStore().Get("orderBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetReportName gets the reportName property value. The reportName property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetReportName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName) {
    val, err := m.GetBackingStore().Get("reportName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)
    }
    return nil
}
// GetSearch gets the search property value. The search property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetSearch()(*string) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSelect gets the select property value. The select property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetSelect()([]string) {
    val, err := m.GetBackingStore().Get("selectEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSkip gets the skip property value. The skip property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetSkip()(*int32) {
    val, err := m.GetBackingStore().Get("skip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTop gets the top property value. The top property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) GetTop()(*int32) {
    val, err := m.GetBackingStore().Get("top")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetGroupBy() != nil {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
        if err != nil {
            return err
        }
    }
    if m.GetOrderBy() != nil {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    if m.GetReportName() != nil {
        cast := (*m.GetReportName()).String()
        err := writer.WriteStringValue("reportName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFilter sets the filter property value. The filter property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetFilter(value *string)() {
    err := m.GetBackingStore().Set("filter", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetGroupBy(value []string)() {
    err := m.GetBackingStore().Set("groupBy", value)
    if err != nil {
        panic(err)
    }
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetOrderBy(value []string)() {
    err := m.GetBackingStore().Set("orderBy", value)
    if err != nil {
        panic(err)
    }
}
// SetReportName sets the reportName property value. The reportName property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetReportName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)() {
    err := m.GetBackingStore().Set("reportName", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. The search property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetSearch(value *string)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
// SetSelect sets the select property value. The select property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetSelect(value []string)() {
    err := m.GetBackingStore().Set("selectEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetSkip sets the skip property value. The skip property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetSkip(value *int32)() {
    err := m.GetBackingStore().Set("skip", value)
    if err != nil {
        panic(err)
    }
}
// SetTop sets the top property value. The top property
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBody) SetTop(value *int32)() {
    err := m.GetBackingStore().Set("top", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyable 
type VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFilter()(*string)
    GetGroupBy()([]string)
    GetOrderBy()([]string)
    GetReportName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFilter(value *string)()
    SetGroupBy(value []string)()
    SetOrderBy(value []string)()
    SetReportName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportName)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
