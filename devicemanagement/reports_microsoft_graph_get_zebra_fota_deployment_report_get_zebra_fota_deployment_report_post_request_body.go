package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody 
type ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The filter property
    filter *string
    // The groupBy property
    groupBy []string
    // The name property
    name *string
    // The orderBy property
    orderBy []string
    // The search property
    search *string
    // The select property
    select_escaped []string
    // The sessionId property
    sessionId *string
    // The skip property
    skip *int32
    // The top property
    top *int32
}
// NewReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody instantiates a new ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody and sets the default values.
func NewReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody()(*ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) {
    m := &ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
                res[i] = *(v.(*string))
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
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
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
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetFilter()(*string) {
    return m.filter
}
// GetGroupBy gets the groupBy property value. The groupBy property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetGroupBy()([]string) {
    return m.groupBy
}
// GetName gets the name property value. The name property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetName()(*string) {
    return m.name
}
// GetOrderBy gets the orderBy property value. The orderBy property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetOrderBy()([]string) {
    return m.orderBy
}
// GetSearch gets the search property value. The search property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetSearch()(*string) {
    return m.search
}
// GetSelect gets the select property value. The select property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetSelect()([]string) {
    return m.select_escaped
}
// GetSessionId gets the sessionId property value. The sessionId property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetSessionId()(*string) {
    return m.sessionId
}
// GetSkip gets the skip property value. The skip property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetSkip()(*int32) {
    return m.skip
}
// GetTop gets the top property value. The top property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) GetTop()(*int32) {
    return m.top
}
// Serialize serializes information the current object
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetSelect() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelect())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFilter sets the filter property value. The filter property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetFilter(value *string)() {
    m.filter = value
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetGroupBy(value []string)() {
    m.groupBy = value
}
// SetName sets the name property value. The name property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetOrderBy(value []string)() {
    m.orderBy = value
}
// SetSearch sets the search property value. The search property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetSearch(value *string)() {
    m.search = value
}
// SetSelect sets the select property value. The select property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetSelect(value []string)() {
    m.select_escaped = value
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
// SetSkip sets the skip property value. The skip property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// SetTop sets the top property value. The top property
func (m *ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportPostRequestBody) SetTop(value *int32)() {
    m.top = value
}
