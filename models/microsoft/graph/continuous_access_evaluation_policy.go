package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ContinuousAccessEvaluationPolicy struct {
    Entity
    // Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
    description *string;
    // The value is always Continuous Access Evaluation. Read-only.
    displayName *string;
    // The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty.
    groups []string;
    // true to indicate whether continuous access evaluation should be performed; otherwise false.
    isEnabled *bool;
    // 
    migrate *bool;
    // The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty.
    users []string;
}
// Instantiates a new continuousAccessEvaluationPolicy and sets the default values.
func NewContinuousAccessEvaluationPolicy()(*ContinuousAccessEvaluationPolicy) {
    m := &ContinuousAccessEvaluationPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty.
func (m *ContinuousAccessEvaluationPolicy) GetGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// Gets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false.
func (m *ContinuousAccessEvaluationPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the migrate property value. 
func (m *ContinuousAccessEvaluationPolicy) GetMigrate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.migrate
    }
}
// Gets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty.
func (m *ContinuousAccessEvaluationPolicy) GetUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// The deserialization information for the current model
func (m *ContinuousAccessEvaluationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetGroups(res)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["migrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMigrate(val)
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetUsers(res)
        return nil
    }
    return res
}
func (m *ContinuousAccessEvaluationPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ContinuousAccessEvaluationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("migrate", m.GetMigrate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("users", m.GetUsers())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *ContinuousAccessEvaluationPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ContinuousAccessEvaluationPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty.
// Parameters:
//  - value : Value to set for the groups property.
func (m *ContinuousAccessEvaluationPolicy) SetGroups(value []string)() {
    m.groups = value
}
// Sets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *ContinuousAccessEvaluationPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the migrate property value. 
// Parameters:
//  - value : Value to set for the migrate property.
func (m *ContinuousAccessEvaluationPolicy) SetMigrate(value *bool)() {
    m.migrate = value
}
// Sets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty.
// Parameters:
//  - value : Value to set for the users property.
func (m *ContinuousAccessEvaluationPolicy) SetUsers(value []string)() {
    m.users = value
}
