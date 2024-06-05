package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody instantiates a new ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody and sets the default values.
func NewReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody()(*ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) {
    m := &ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["groupby"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupby(val)
        }
        return nil
    }
    res["orderby"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderby(val)
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectEscaped(val)
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
// returns a *string when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetFilter()(*string) {
    val, err := m.GetBackingStore().Get("filter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetGroupby gets the groupby property value. The groupby property
// returns a *string when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetGroupby()(*string) {
    val, err := m.GetBackingStore().Get("groupby")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrderby gets the orderby property value. The orderby property
// returns a *string when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetOrderby()(*string) {
    val, err := m.GetBackingStore().Get("orderby")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSearch gets the search property value. The search property
// returns a *string when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetSearch()(*string) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSelectEscaped gets the select property value. The select property
// returns a *string when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetSelectEscaped()(*string) {
    val, err := m.GetBackingStore().Get("selectEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkip gets the skip property value. The skip property
// returns a *int32 when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetSkip()(*int32) {
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
// returns a *int32 when successful
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) GetTop()(*int32) {
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
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupby", m.GetGroupby())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("orderby", m.GetOrderby())
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
    {
        err := writer.WriteStringValue("select", m.GetSelectEscaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFilter sets the filter property value. The filter property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetFilter(value *string)() {
    err := m.GetBackingStore().Set("filter", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupby sets the groupby property value. The groupby property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetGroupby(value *string)() {
    err := m.GetBackingStore().Set("groupby", value)
    if err != nil {
        panic(err)
    }
}
// SetOrderby sets the orderby property value. The orderby property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetOrderby(value *string)() {
    err := m.GetBackingStore().Set("orderby", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. The search property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetSearch(value *string)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
// SetSelectEscaped sets the select property value. The select property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetSelectEscaped(value *string)() {
    err := m.GetBackingStore().Set("selectEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetSkip sets the skip property value. The skip property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetSkip(value *int32)() {
    err := m.GetBackingStore().Set("skip", value)
    if err != nil {
        panic(err)
    }
}
// SetTop sets the top property value. The top property
func (m *ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBody) SetTop(value *int32)() {
    err := m.GetBackingStore().Set("top", value)
    if err != nil {
        panic(err)
    }
}
type ReportsRetrieveassignedapplicationsreportRetrieveAssignedApplicationsReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFilter()(*string)
    GetGroupby()(*string)
    GetOrderby()(*string)
    GetSearch()(*string)
    GetSelectEscaped()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFilter(value *string)()
    SetGroupby(value *string)()
    SetOrderby(value *string)()
    SetSearch(value *string)()
    SetSelectEscaped(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
