package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonCertification struct {
    ItemFacet
    // The referenceable identifier for the certification.
    certificationId *string;
    // Description of the certification.
    description *string;
    // Title of the certification.
    displayName *string;
    // The date that the certification expires.
    endDate *string;
    // The date that the certification was issued.
    issuedDate *string;
    // Authority which granted the certification.
    issuingAuthority *string;
    // Company which granted the certification.
    issuingCompany *string;
    // The date that the certification became valid.
    startDate *string;
    // URL referencing a thumbnail of the certification.
    thumbnailUrl *string;
    // URL referencing the certification.
    webUrl *string;
}
// Instantiates a new personCertification and sets the default values.
func NewPersonCertification()(*PersonCertification) {
    m := &PersonCertification{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the certificationId property value. The referenceable identifier for the certification.
func (m *PersonCertification) GetCertificationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationId
    }
}
// Gets the description property value. Description of the certification.
func (m *PersonCertification) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Title of the certification.
func (m *PersonCertification) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDate property value. The date that the certification expires.
func (m *PersonCertification) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// Gets the issuedDate property value. The date that the certification was issued.
func (m *PersonCertification) GetIssuedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// Gets the issuingAuthority property value. Authority which granted the certification.
func (m *PersonCertification) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// Gets the issuingCompany property value. Company which granted the certification.
func (m *PersonCertification) GetIssuingCompany()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingCompany
    }
}
// Gets the startDate property value. The date that the certification became valid.
func (m *PersonCertification) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// Gets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
func (m *PersonCertification) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the webUrl property value. URL referencing the certification.
func (m *PersonCertification) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *PersonCertification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["certificationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCertificationId(val)
        return nil
    }
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
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndDate(val)
        return nil
    }
    res["issuedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuedDate(val)
        return nil
    }
    res["issuingAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuingAuthority(val)
        return nil
    }
    res["issuingCompany"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuingCompany(val)
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartDate(val)
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbnailUrl(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *PersonCertification) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonCertification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificationId", m.GetCertificationId())
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
        err = writer.WriteStringValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuedDate", m.GetIssuedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingAuthority", m.GetIssuingAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingCompany", m.GetIssuingCompany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the certificationId property value. The referenceable identifier for the certification.
// Parameters:
//  - value : Value to set for the certificationId property.
func (m *PersonCertification) SetCertificationId(value *string)() {
    m.certificationId = value
}
// Sets the description property value. Description of the certification.
// Parameters:
//  - value : Value to set for the description property.
func (m *PersonCertification) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Title of the certification.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonCertification) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDate property value. The date that the certification expires.
// Parameters:
//  - value : Value to set for the endDate property.
func (m *PersonCertification) SetEndDate(value *string)() {
    m.endDate = value
}
// Sets the issuedDate property value. The date that the certification was issued.
// Parameters:
//  - value : Value to set for the issuedDate property.
func (m *PersonCertification) SetIssuedDate(value *string)() {
    m.issuedDate = value
}
// Sets the issuingAuthority property value. Authority which granted the certification.
// Parameters:
//  - value : Value to set for the issuingAuthority property.
func (m *PersonCertification) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// Sets the issuingCompany property value. Company which granted the certification.
// Parameters:
//  - value : Value to set for the issuingCompany property.
func (m *PersonCertification) SetIssuingCompany(value *string)() {
    m.issuingCompany = value
}
// Sets the startDate property value. The date that the certification became valid.
// Parameters:
//  - value : Value to set for the startDate property.
func (m *PersonCertification) SetStartDate(value *string)() {
    m.startDate = value
}
// Sets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *PersonCertification) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the webUrl property value. URL referencing the certification.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *PersonCertification) SetWebUrl(value *string)() {
    m.webUrl = value
}
