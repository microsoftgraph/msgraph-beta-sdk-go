package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProfileCardProperty provides operations to manage the collection of organization entities.
type ProfileCardProperty struct {
    Entity
    // Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
    annotations []ProfileCardAnnotationable;
    // Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
    directoryPropertyName *string;
}
// NewProfileCardProperty instantiates a new profileCardProperty and sets the default values.
func NewProfileCardProperty()(*ProfileCardProperty) {
    m := &ProfileCardProperty{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProfileCardPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfileCardPropertyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProfileCardProperty(), nil
}
// GetAnnotations gets the annotations property value. Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
func (m *ProfileCardProperty) GetAnnotations()([]ProfileCardAnnotationable) {
    if m == nil {
        return nil
    } else {
        return m.annotations
    }
}
// GetDirectoryPropertyName gets the directoryPropertyName property value. Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
func (m *ProfileCardProperty) GetDirectoryPropertyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryPropertyName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProfileCardProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["annotations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileCardAnnotationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileCardAnnotationable, len(val))
            for i, v := range val {
                res[i] = v.(ProfileCardAnnotationable)
            }
            m.SetAnnotations(res)
        }
        return nil
    }
    res["directoryPropertyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryPropertyName(val)
        }
        return nil
    }
    return res
}
func (m *ProfileCardProperty) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProfileCardProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnnotations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAnnotations()))
        for i, v := range m.GetAnnotations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("annotations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryPropertyName", m.GetDirectoryPropertyName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnnotations sets the annotations property value. Allows an administrator to set a custom display label for the directory property and localize it for the users in their tenant.
func (m *ProfileCardProperty) SetAnnotations(value []ProfileCardAnnotationable)() {
    if m != nil {
        m.annotations = value
    }
}
// SetDirectoryPropertyName sets the directoryPropertyName property value. Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface hidden Azure Active Directory (Azure AD) properties on the Microsoft 365 profile card within their tenant. When present, the Azure AD field referenced in this field will be visible to all users in your tenant on the contact pane of the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode, StateOrProvince, Alias, CustomAttribute1,  CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5, CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11, CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
func (m *ProfileCardProperty) SetDirectoryPropertyName(value *string)() {
    if m != nil {
        m.directoryPropertyName = value
    }
}
