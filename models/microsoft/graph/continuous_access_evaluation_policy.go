package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContinuousAccessEvaluationPolicy 
type ContinuousAccessEvaluationPolicy struct {
    Entity
    // Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
    description *string;
    // The value is always Continuous Access Evaluation. Read-only.
    displayName *string;
    // The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
    groups []string;
    // true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
    isEnabled *bool;
    // true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
    migrate *bool;
    // The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
    users []string;
}
// NewContinuousAccessEvaluationPolicy instantiates a new continuousAccessEvaluationPolicy and sets the default values.
func NewContinuousAccessEvaluationPolicy()(*ContinuousAccessEvaluationPolicy) {
    m := &ContinuousAccessEvaluationPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGroups gets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// GetIsEnabled gets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetMigrate gets the migrate property value. true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
func (m *ContinuousAccessEvaluationPolicy) GetMigrate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.migrate
    }
}
// GetUsers gets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContinuousAccessEvaluationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["migrate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMigrate(val)
        }
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
func (m *ContinuousAccessEvaluationPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetGroups() != nil {
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
    if m.GetUsers() != nil {
        err = writer.WriteCollectionOfStringValues("users", m.GetUsers())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGroups sets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetGroups(value []string)() {
    if m != nil {
        m.groups = value
    }
}
// SetIsEnabled sets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetMigrate sets the migrate property value. true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
func (m *ContinuousAccessEvaluationPolicy) SetMigrate(value *bool)() {
    if m != nil {
        m.migrate = value
    }
}
// SetUsers sets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetUsers(value []string)() {
    if m != nil {
        m.users = value
    }
}
