package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationalActivityDetail 
type EducationalActivityDetail struct {
    // Shortened name of the degree or program (example: PhD, MBA)
    abbreviation *string;
    // Extracurricular activities undertaken alongside the program.
    activities []string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Any awards or honors associated with the program.
    awards []string;
    // Short description of the program provided by the user.
    description *string;
    // Long-form name of the program that the user has provided.
    displayName *string;
    // Majors and minors associated with the program. (if applicable)
    fieldsOfStudy []string;
    // The final grade, class, GPA or score.
    grade *string;
    // Additional notes the user has provided.
    notes *string;
    // Link to the degree or program page.
    webUrl *string;
}
// NewEducationalActivityDetail instantiates a new educationalActivityDetail and sets the default values.
func NewEducationalActivityDetail()(*EducationalActivityDetail) {
    m := &EducationalActivityDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAbbreviation gets the abbreviation property value. Shortened name of the degree or program (example: PhD, MBA)
func (m *EducationalActivityDetail) GetAbbreviation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.abbreviation
    }
}
// GetActivities gets the activities property value. Extracurricular activities undertaken alongside the program.
func (m *EducationalActivityDetail) GetActivities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationalActivityDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAwards gets the awards property value. Any awards or honors associated with the program.
func (m *EducationalActivityDetail) GetAwards()([]string) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
// GetDescription gets the description property value. Short description of the program provided by the user.
func (m *EducationalActivityDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Long-form name of the program that the user has provided.
func (m *EducationalActivityDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldsOfStudy gets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
func (m *EducationalActivityDetail) GetFieldsOfStudy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fieldsOfStudy
    }
}
// GetGrade gets the grade property value. The final grade, class, GPA or score.
func (m *EducationalActivityDetail) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
// GetNotes gets the notes property value. Additional notes the user has provided.
func (m *EducationalActivityDetail) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetWebUrl gets the webUrl property value. Link to the degree or program page.
func (m *EducationalActivityDetail) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivityDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["abbreviation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAbbreviation(val)
        }
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["awards"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAwards(res)
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
    res["fieldsOfStudy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFieldsOfStudy(res)
        }
        return nil
    }
    res["grade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrade(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
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
func (m *EducationalActivityDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationalActivityDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("abbreviation", m.GetAbbreviation())
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        err := writer.WriteCollectionOfStringValues("activities", m.GetActivities())
        if err != nil {
            return err
        }
    }
    if m.GetAwards() != nil {
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
    if m.GetFieldsOfStudy() != nil {
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
// SetAbbreviation sets the abbreviation property value. Shortened name of the degree or program (example: PhD, MBA)
func (m *EducationalActivityDetail) SetAbbreviation(value *string)() {
    if m != nil {
        m.abbreviation = value
    }
}
// SetActivities sets the activities property value. Extracurricular activities undertaken alongside the program.
func (m *EducationalActivityDetail) SetActivities(value []string)() {
    if m != nil {
        m.activities = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationalActivityDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAwards sets the awards property value. Any awards or honors associated with the program.
func (m *EducationalActivityDetail) SetAwards(value []string)() {
    if m != nil {
        m.awards = value
    }
}
// SetDescription sets the description property value. Short description of the program provided by the user.
func (m *EducationalActivityDetail) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Long-form name of the program that the user has provided.
func (m *EducationalActivityDetail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFieldsOfStudy sets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
func (m *EducationalActivityDetail) SetFieldsOfStudy(value []string)() {
    if m != nil {
        m.fieldsOfStudy = value
    }
}
// SetGrade sets the grade property value. The final grade, class, GPA or score.
func (m *EducationalActivityDetail) SetGrade(value *string)() {
    if m != nil {
        m.grade = value
    }
}
// SetNotes sets the notes property value. Additional notes the user has provided.
func (m *EducationalActivityDetail) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetWebUrl sets the webUrl property value. Link to the degree or program page.
func (m *EducationalActivityDetail) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
