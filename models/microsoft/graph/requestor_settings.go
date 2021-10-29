package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RequestorSettings struct {
    // Indicates whether new requests are accepted on this policy.
    acceptRequests *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
    allowedRequestors []UserSet;
    // Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
    scopeType *string;
}
// Instantiates a new requestorSettings and sets the default values.
func NewRequestorSettings()(*RequestorSettings) {
    m := &RequestorSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the acceptRequests property value. Indicates whether new requests are accepted on this policy.
func (m *RequestorSettings) GetAcceptRequests()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptRequests
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestorSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedRequestors property value. The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
func (m *RequestorSettings) GetAllowedRequestors()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.allowedRequestors
    }
}
// Gets the scopeType property value. Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
func (m *RequestorSettings) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
// The deserialization information for the current model
func (m *RequestorSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acceptRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAcceptRequests(val)
        return nil
    }
    res["allowedRequestors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetAllowedRequestors(res)
        return nil
    }
    res["scopeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScopeType(val)
        return nil
    }
    return res
}
func (m *RequestorSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RequestorSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptRequests", m.GetAcceptRequests())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedRequestors()))
        for i, v := range m.GetAllowedRequestors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("allowedRequestors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scopeType", m.GetScopeType())
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
// Sets the acceptRequests property value. Indicates whether new requests are accepted on this policy.
// Parameters:
//  - value : Value to set for the acceptRequests property.
func (m *RequestorSettings) SetAcceptRequests(value *bool)() {
    m.acceptRequests = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RequestorSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedRequestors property value. The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
// Parameters:
//  - value : Value to set for the allowedRequestors property.
func (m *RequestorSettings) SetAllowedRequestors(value []UserSet)() {
    m.allowedRequestors = value
}
// Sets the scopeType property value. Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
// Parameters:
//  - value : Value to set for the scopeType property.
func (m *RequestorSettings) SetScopeType(value *string)() {
    m.scopeType = value
}
