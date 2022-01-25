package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonCertification 
type PersonCertification struct {
    ItemFacet
    // The referenceable identifier for the certification.
    certificationId *string;
    // Description of the certification.
    description *string;
    // Title of the certification.
    displayName *string;
    // The date that the certification expires.
    endDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The date that the certification was issued.
    issuedDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Authority which granted the certification.
    issuingAuthority *string;
    // Company which granted the certification.
    issuingCompany *string;
    // The date that the certification became valid.
    startDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // URL referencing a thumbnail of the certification.
    thumbnailUrl *string;
    // URL referencing the certification.
    webUrl *string;
}
// NewPersonCertification instantiates a new personCertification and sets the default values.
func NewPersonCertification()(*PersonCertification) {
    m := &PersonCertification{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetCertificationId gets the certificationId property value. The referenceable identifier for the certification.
func (m *PersonCertification) GetCertificationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationId
    }
}
// GetDescription gets the description property value. Description of the certification.
func (m *PersonCertification) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Title of the certification.
func (m *PersonCertification) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndDate gets the endDate property value. The date that the certification expires.
func (m *PersonCertification) GetEndDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// GetIssuedDate gets the issuedDate property value. The date that the certification was issued.
func (m *PersonCertification) GetIssuedDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the certification.
func (m *PersonCertification) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// GetIssuingCompany gets the issuingCompany property value. Company which granted the certification.
func (m *PersonCertification) GetIssuingCompany()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingCompany
    }
}
// GetStartDate gets the startDate property value. The date that the certification became valid.
func (m *PersonCertification) GetStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
func (m *PersonCertification) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetWebUrl gets the webUrl property value. URL referencing the certification.
func (m *PersonCertification) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonCertification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["certificationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationId(val)
        }
        return nil
    }
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
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["issuedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuedDate(val)
        }
        return nil
    }
    res["issuingAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuingAuthority(val)
        }
        return nil
    }
    res["issuingCompany"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuingCompany(val)
        }
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *PersonCertification) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("issuedDate", m.GetIssuedDate())
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
        err = writer.WriteDateOnlyValue("startDate", m.GetStartDate())
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
// SetCertificationId sets the certificationId property value. The referenceable identifier for the certification.
func (m *PersonCertification) SetCertificationId(value *string)() {
    if m != nil {
        m.certificationId = value
    }
}
// SetDescription sets the description property value. Description of the certification.
func (m *PersonCertification) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Title of the certification.
func (m *PersonCertification) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndDate sets the endDate property value. The date that the certification expires.
func (m *PersonCertification) SetEndDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endDate = value
    }
}
// SetIssuedDate sets the issuedDate property value. The date that the certification was issued.
func (m *PersonCertification) SetIssuedDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.issuedDate = value
    }
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the certification.
func (m *PersonCertification) SetIssuingAuthority(value *string)() {
    if m != nil {
        m.issuingAuthority = value
    }
}
// SetIssuingCompany sets the issuingCompany property value. Company which granted the certification.
func (m *PersonCertification) SetIssuingCompany(value *string)() {
    if m != nil {
        m.issuingCompany = value
    }
}
// SetStartDate sets the startDate property value. The date that the certification became valid.
func (m *PersonCertification) SetStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startDate = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
func (m *PersonCertification) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
// SetWebUrl sets the webUrl property value. URL referencing the certification.
func (m *PersonCertification) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
