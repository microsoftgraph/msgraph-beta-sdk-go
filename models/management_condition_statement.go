package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementConditionStatement 
type ManagementConditionStatement struct {
    Entity
    // The applicable platforms for this management condition statement.
    applicablePlatforms []DevicePlatformType
    // The time the management condition statement was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The admin defined description of the management condition statement.
    description *string
    // The admin defined name of the management condition statement.
    displayName *string
    // ETag of the management condition statement. Updated service side.
    eTag *string
    // The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
    expression ManagementConditionExpressionable
    // The management conditions associated to the management condition statement.
    managementConditions []ManagementConditionable
    // The time the management condition statement was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewManagementConditionStatement instantiates a new managementConditionStatement and sets the default values.
func NewManagementConditionStatement()(*ManagementConditionStatement) {
    m := &ManagementConditionStatement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagementConditionStatementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementConditionStatementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementConditionStatement(), nil
}
// GetApplicablePlatforms gets the applicablePlatforms property value. The applicable platforms for this management condition statement.
func (m *ManagementConditionStatement) GetApplicablePlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatforms
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The time the management condition statement was created. Generated service side.
func (m *ManagementConditionStatement) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The admin defined description of the management condition statement.
func (m *ManagementConditionStatement) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The admin defined name of the management condition statement.
func (m *ManagementConditionStatement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetETag gets the eTag property value. ETag of the management condition statement. Updated service side.
func (m *ManagementConditionStatement) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// GetExpression gets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
func (m *ManagementConditionStatement) GetExpression()(ManagementConditionExpressionable) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementConditionStatement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicablePlatforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DevicePlatformType, len(val))
            for i, v := range val {
                res[i] = *(v.(*DevicePlatformType))
            }
            m.SetApplicablePlatforms(res)
        }
        return nil
    }
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["eTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetETag(val)
        }
        return nil
    }
    res["expression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementConditionExpressionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val.(ManagementConditionExpressionable))
        }
        return nil
    }
    res["managementConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementConditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementConditionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementConditionable)
            }
            m.SetManagementConditions(res)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetManagementConditions gets the managementConditions property value. The management conditions associated to the management condition statement.
func (m *ManagementConditionStatement) GetManagementConditions()([]ManagementConditionable) {
    if m == nil {
        return nil
    } else {
        return m.managementConditions
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
func (m *ManagementConditionStatement) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Serialize serializes information the current object
func (m *ManagementConditionStatement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicablePlatforms() != nil {
        err = writer.WriteCollectionOfStringValues("applicablePlatforms", SerializeDevicePlatformType(m.GetApplicablePlatforms()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    if m.GetManagementConditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementConditions()))
        for i, v := range m.GetManagementConditions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementConditions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicablePlatforms sets the applicablePlatforms property value. The applicable platforms for this management condition statement.
func (m *ManagementConditionStatement) SetApplicablePlatforms(value []DevicePlatformType)() {
    if m != nil {
        m.applicablePlatforms = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the management condition statement was created. Generated service side.
func (m *ManagementConditionStatement) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The admin defined description of the management condition statement.
func (m *ManagementConditionStatement) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The admin defined name of the management condition statement.
func (m *ManagementConditionStatement) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetETag sets the eTag property value. ETag of the management condition statement. Updated service side.
func (m *ManagementConditionStatement) SetETag(value *string)() {
    if m != nil {
        m.eTag = value
    }
}
// SetExpression sets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
func (m *ManagementConditionStatement) SetExpression(value ManagementConditionExpressionable)() {
    if m != nil {
        m.expression = value
    }
}
// SetManagementConditions sets the managementConditions property value. The management conditions associated to the management condition statement.
func (m *ManagementConditionStatement) SetManagementConditions(value []ManagementConditionable)() {
    if m != nil {
        m.managementConditions = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
func (m *ManagementConditionStatement) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
