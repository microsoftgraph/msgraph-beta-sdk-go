package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluateRequest request for assignment filter evaluation for devices.
type AssignmentFilterEvaluateRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Order the devices should be sorted in. Default is ascending on device name.
    orderBy []string
    // Supported platform types.
    platform *DevicePlatformType
    // Rule definition of the Assignment Filter.
    rule *string
    // Search keyword applied to scope found devices.
    search *string
    // Number of records to skip. Default value is 0
    skip *int32
    // Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
    top *int32
}
// NewAssignmentFilterEvaluateRequest instantiates a new assignmentFilterEvaluateRequest and sets the default values.
func NewAssignmentFilterEvaluateRequest()(*AssignmentFilterEvaluateRequest) {
    m := &AssignmentFilterEvaluateRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.assignmentFilterEvaluateRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAssignmentFilterEvaluateRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluateRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluateRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluateRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluateRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluateRequest) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetOrderBy gets the orderBy property value. Order the devices should be sorted in. Default is ascending on device name.
func (m *AssignmentFilterEvaluateRequest) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *AssignmentFilterEvaluateRequest) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetRule gets the rule property value. Rule definition of the Assignment Filter.
func (m *AssignmentFilterEvaluateRequest) GetRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rule
    }
}
// GetSearch gets the search property value. Search keyword applied to scope found devices.
func (m *AssignmentFilterEvaluateRequest) GetSearch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.search
    }
}
// GetSkip gets the skip property value. Number of records to skip. Default value is 0
func (m *AssignmentFilterEvaluateRequest) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// GetTop gets the top property value. Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
func (m *AssignmentFilterEvaluateRequest) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluateRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule", m.GetRule())
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
func (m *AssignmentFilterEvaluateRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluateRequest) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetOrderBy sets the orderBy property value. Order the devices should be sorted in. Default is ascending on device name.
func (m *AssignmentFilterEvaluateRequest) SetOrderBy(value []string)() {
    if m != nil {
        m.orderBy = value
    }
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *AssignmentFilterEvaluateRequest) SetPlatform(value *DevicePlatformType)() {
    if m != nil {
        m.platform = value
    }
}
// SetRule sets the rule property value. Rule definition of the Assignment Filter.
func (m *AssignmentFilterEvaluateRequest) SetRule(value *string)() {
    if m != nil {
        m.rule = value
    }
}
// SetSearch sets the search property value. Search keyword applied to scope found devices.
func (m *AssignmentFilterEvaluateRequest) SetSearch(value *string)() {
    if m != nil {
        m.search = value
    }
}
// SetSkip sets the skip property value. Number of records to skip. Default value is 0
func (m *AssignmentFilterEvaluateRequest) SetSkip(value *int32)() {
    if m != nil {
        m.skip = value
    }
}
// SetTop sets the top property value. Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
func (m *AssignmentFilterEvaluateRequest) SetTop(value *int32)() {
    if m != nil {
        m.top = value
    }
}
