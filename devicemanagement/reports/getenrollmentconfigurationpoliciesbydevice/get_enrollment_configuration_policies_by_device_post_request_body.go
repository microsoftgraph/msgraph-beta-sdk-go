package getenrollmentconfigurationpoliciesbydevice

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetEnrollmentConfigurationPoliciesByDevicePostRequestBody provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
type GetEnrollmentConfigurationPoliciesByDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    // The skip property
    skip *int32
    // The top property
    top *int32
}
// NewGetEnrollmentConfigurationPoliciesByDevicePostRequestBody instantiates a new getEnrollmentConfigurationPoliciesByDevicePostRequestBody and sets the default values.
func NewGetEnrollmentConfigurationPoliciesByDevicePostRequestBody()(*GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) {
    m := &GetEnrollmentConfigurationPoliciesByDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetEnrollmentConfigurationPoliciesByDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetEnrollmentConfigurationPoliciesByDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetEnrollmentConfigurationPoliciesByDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetGroupBy gets the groupBy property value. The groupBy property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetGroupBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupBy
    }
}
// GetName gets the name property value. The name property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOrderBy gets the orderBy property value. The orderBy property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// GetSearch gets the search property value. The search property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetSearch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// GetSelect gets the select property value. The select property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetSelect()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// GetSkip gets the skip property value. The skip property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// GetTop gets the top property value. The top property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// Serialize serializes information the current object
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFilter sets the filter property value. The filter property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetFilter(value *string)() {
    if m != nil {
        m.filter = value
    }
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetGroupBy(value []string)() {
    if m != nil {
        m.groupBy = value
    }
}
// SetName sets the name property value. The name property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetOrderBy(value []string)() {
    if m != nil {
        m.orderBy = value
    }
}
// SetSearch sets the search property value. The search property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetSearch(value *string)() {
    if m != nil {
        m.search = value
    }
}
// SetSelect sets the select property value. The select property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetSelect(value []string)() {
    if m != nil {
        m.select_escaped = value
    }
}
// SetSkip sets the skip property value. The skip property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetSkip(value *int32)() {
    if m != nil {
        m.skip = value
    }
}
// SetTop sets the top property value. The top property
func (m *GetEnrollmentConfigurationPoliciesByDevicePostRequestBody) SetTop(value *int32)() {
    if m != nil {
        m.top = value
    }
}
