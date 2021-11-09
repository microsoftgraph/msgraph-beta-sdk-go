package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new educationalActivityDetail and sets the default values.
func NewEducationalActivityDetail()(*EducationalActivityDetail) {
    m := &EducationalActivityDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the abbreviation property value. Shortened name of the degree or program (example: PhD, MBA)
func (m *EducationalActivityDetail) GetAbbreviation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.abbreviation
    }
}
// Gets the activities property value. Extracurricular activities undertaken alongside the program.
func (m *EducationalActivityDetail) GetActivities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationalActivityDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the awards property value. Any awards or honors associated with the program.
func (m *EducationalActivityDetail) GetAwards()([]string) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
// Gets the description property value. Short description of the program provided by the user.
func (m *EducationalActivityDetail) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Long-form name of the program that the user has provided.
func (m *EducationalActivityDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
func (m *EducationalActivityDetail) GetFieldsOfStudy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fieldsOfStudy
    }
}
// Gets the grade property value. The final grade, class, GPA or score.
func (m *EducationalActivityDetail) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
// Gets the notes property value. Additional notes the user has provided.
func (m *EducationalActivityDetail) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the webUrl property value. Link to the degree or program page.
func (m *EducationalActivityDetail) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the abbreviation property value. Shortened name of the degree or program (example: PhD, MBA)
// Parameters:
//  - value : Value to set for the abbreviation property.
func (m *EducationalActivityDetail) SetAbbreviation(value *string)() {
    m.abbreviation = value
}
// Sets the activities property value. Extracurricular activities undertaken alongside the program.
// Parameters:
//  - value : Value to set for the activities property.
func (m *EducationalActivityDetail) SetActivities(value []string)() {
    m.activities = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EducationalActivityDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the awards property value. Any awards or honors associated with the program.
// Parameters:
//  - value : Value to set for the awards property.
func (m *EducationalActivityDetail) SetAwards(value []string)() {
    m.awards = value
}
// Sets the description property value. Short description of the program provided by the user.
// Parameters:
//  - value : Value to set for the description property.
func (m *EducationalActivityDetail) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Long-form name of the program that the user has provided.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationalActivityDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
// Parameters:
//  - value : Value to set for the fieldsOfStudy property.
func (m *EducationalActivityDetail) SetFieldsOfStudy(value []string)() {
    m.fieldsOfStudy = value
}
// Sets the grade property value. The final grade, class, GPA or score.
// Parameters:
//  - value : Value to set for the grade property.
func (m *EducationalActivityDetail) SetGrade(value *string)() {
    m.grade = value
}
// Sets the notes property value. Additional notes the user has provided.
// Parameters:
//  - value : Value to set for the notes property.
func (m *EducationalActivityDetail) SetNotes(value *string)() {
    m.notes = value
}
// Sets the webUrl property value. Link to the degree or program page.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *EducationalActivityDetail) SetWebUrl(value *string)() {
    m.webUrl = value
}
