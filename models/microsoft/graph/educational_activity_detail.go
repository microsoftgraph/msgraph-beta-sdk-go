package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationalActivityDetail struct {
    abbreviation *string;
    activities []string;
    additionalData map[string]interface{};
    awards []string;
    description *string;
    displayName *string;
    fieldsOfStudy []string;
    grade *string;
    notes *string;
    webUrl *string;
}
func NewEducationalActivityDetail()(*EducationalActivityDetail) {
    m := &EducationalActivityDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EducationalActivityDetail) GetAbbreviation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.abbreviation
    }
}
func (m *EducationalActivityDetail) GetActivities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
func (m *EducationalActivityDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EducationalActivityDetail) GetAwards()([]string) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
func (m *EducationalActivityDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *EducationalActivityDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *EducationalActivityDetail) GetFieldsOfStudy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fieldsOfStudy
    }
}
func (m *EducationalActivityDetail) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
func (m *EducationalActivityDetail) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *EducationalActivityDetail) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *EducationalActivityDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["abbreviation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAbbreviation(val)
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetActivities(res)
        return nil
    }
    res["awards"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAwards(res)
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
    res["fieldsOfStudy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetFieldsOfStudy(res)
        return nil
    }
    res["grade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGrade(val)
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotes(val)
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
func (m *EducationalActivityDetail) IsNil()(bool) {
    return m == nil
}
func (m *EducationalActivityDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("abbreviation", m.GetAbbreviation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("activities", m.GetActivities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("awards", m.GetAwards())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("fieldsOfStudy", m.GetFieldsOfStudy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *EducationalActivityDetail) SetAbbreviation(value *string)() {
    m.abbreviation = value
}
func (m *EducationalActivityDetail) SetActivities(value []string)() {
    m.activities = value
}
func (m *EducationalActivityDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EducationalActivityDetail) SetAwards(value []string)() {
    m.awards = value
}
func (m *EducationalActivityDetail) SetDescription(value *string)() {
    m.description = value
}
func (m *EducationalActivityDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *EducationalActivityDetail) SetFieldsOfStudy(value []string)() {
    m.fieldsOfStudy = value
}
func (m *EducationalActivityDetail) SetGrade(value *string)() {
    m.grade = value
}
func (m *EducationalActivityDetail) SetNotes(value *string)() {
    m.notes = value
}
func (m *EducationalActivityDetail) SetWebUrl(value *string)() {
    m.webUrl = value
}
