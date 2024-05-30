package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody instantiates a new ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody and sets the default values.
func NewReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody()(*ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) {
    m := &ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetGroupBy(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOrderBy(res)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSelectEscaped(res)
        }
        return nil
    }
    res["sessionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetFilter()(*string) {
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
// returns a []string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetGroupBy()([]string) {
    val, err := m.GetBackingStore().Get("groupBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrderBy gets the orderBy property value. The orderBy property
// returns a []string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetOrderBy()([]string) {
    val, err := m.GetBackingStore().Get("orderBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSearch gets the search property value. The search property
// returns a *string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetSearch()(*string) {
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
// returns a []string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetSelectEscaped()([]string) {
    val, err := m.GetBackingStore().Get("selectEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSessionId gets the sessionId property value. The sessionId property
// returns a *string when successful
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetSessionId()(*string) {
    val, err := m.GetBackingStore().Get("sessionId")
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetSkip()(*int32) {
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) GetTop()(*int32) {
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    if m.GetSelectEscaped() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelectEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
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
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFilter sets the filter property value. The filter property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetFilter(value *string)() {
    err := m.GetBackingStore().Set("filter", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetGroupBy(value []string)() {
    err := m.GetBackingStore().Set("groupBy", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetOrderBy(value []string)() {
    err := m.GetBackingStore().Set("orderBy", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. The search property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetSearch(value *string)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
// SetSelectEscaped sets the select property value. The select property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetSelectEscaped(value []string)() {
    err := m.GetBackingStore().Set("selectEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetSessionId(value *string)() {
    err := m.GetBackingStore().Set("sessionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSkip sets the skip property value. The skip property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetSkip(value *int32)() {
    err := m.GetBackingStore().Set("skip", value)
    if err != nil {
        panic(err)
    }
}
// SetTop sets the top property value. The top property
func (m *ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBody) SetTop(value *int32)() {
    err := m.GetBackingStore().Set("top", value)
    if err != nil {
        panic(err)
    }
}
type ReportsGetmobileapplicationmanagementappregistrationsummaryreportGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFilter()(*string)
    GetGroupBy()([]string)
    GetName()(*string)
    GetOrderBy()([]string)
    GetSearch()(*string)
    GetSelectEscaped()([]string)
    GetSessionId()(*string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFilter(value *string)()
    SetGroupBy(value []string)()
    SetName(value *string)()
    SetOrderBy(value []string)()
    SetSearch(value *string)()
    SetSelectEscaped(value []string)()
    SetSessionId(value *string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
