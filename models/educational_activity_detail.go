package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationalActivityDetail 
type EducationalActivityDetail struct {
    // Shortened name of the degree or program (example: PhD, MBA)
    abbreviation *string
    // Extracurricular activities undertaken alongside the program.
    activities []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Any awards or honors associated with the program.
    awards []string
    // Short description of the program provided by the user.
    description *string
    // Long-form name of the program that the user has provided.
    displayName *string
    // Majors and minors associated with the program. (if applicable)
    fieldsOfStudy []string
    // The final grade, class, GPA or score.
    grade *string
    // Additional notes the user has provided.
    notes *string
    // The OdataType property
    odataType *string
    // Link to the degree or program page.
    webUrl *string
}
// NewEducationalActivityDetail instantiates a new educationalActivityDetail and sets the default values.
func NewEducationalActivityDetail()(*EducationalActivityDetail) {
    m := &EducationalActivityDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.educationalActivityDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationalActivityDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationalActivityDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationalActivityDetail(), nil
}
// GetAbbreviation gets the abbreviation property value. Shortened name of the degree or program (example: PhD, MBA)
func (m *EducationalActivityDetail) GetAbbreviation()(*string) {
    return m.abbreviation
}
// GetActivities gets the activities property value. Extracurricular activities undertaken alongside the program.
func (m *EducationalActivityDetail) GetActivities()([]string) {
    return m.activities
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationalActivityDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAwards gets the awards property value. Any awards or honors associated with the program.
func (m *EducationalActivityDetail) GetAwards()([]string) {
    return m.awards
}
// GetDescription gets the description property value. Short description of the program provided by the user.
func (m *EducationalActivityDetail) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Long-form name of the program that the user has provided.
func (m *EducationalActivityDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivityDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["abbreviation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAbbreviation)
    res["activities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetActivities)
    res["awards"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAwards)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["fieldsOfStudy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetFieldsOfStudy)
    res["grade"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGrade)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["webUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebUrl)
    return res
}
// GetFieldsOfStudy gets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
func (m *EducationalActivityDetail) GetFieldsOfStudy()([]string) {
    return m.fieldsOfStudy
}
// GetGrade gets the grade property value. The final grade, class, GPA or score.
func (m *EducationalActivityDetail) GetGrade()(*string) {
    return m.grade
}
// GetNotes gets the notes property value. Additional notes the user has provided.
func (m *EducationalActivityDetail) GetNotes()(*string) {
    return m.notes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationalActivityDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetWebUrl gets the webUrl property value. Link to the degree or program page.
func (m *EducationalActivityDetail) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *EducationalActivityDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.abbreviation = value
}
// SetActivities sets the activities property value. Extracurricular activities undertaken alongside the program.
func (m *EducationalActivityDetail) SetActivities(value []string)() {
    m.activities = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationalActivityDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAwards sets the awards property value. Any awards or honors associated with the program.
func (m *EducationalActivityDetail) SetAwards(value []string)() {
    m.awards = value
}
// SetDescription sets the description property value. Short description of the program provided by the user.
func (m *EducationalActivityDetail) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Long-form name of the program that the user has provided.
func (m *EducationalActivityDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFieldsOfStudy sets the fieldsOfStudy property value. Majors and minors associated with the program. (if applicable)
func (m *EducationalActivityDetail) SetFieldsOfStudy(value []string)() {
    m.fieldsOfStudy = value
}
// SetGrade sets the grade property value. The final grade, class, GPA or score.
func (m *EducationalActivityDetail) SetGrade(value *string)() {
    m.grade = value
}
// SetNotes sets the notes property value. Additional notes the user has provided.
func (m *EducationalActivityDetail) SetNotes(value *string)() {
    m.notes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationalActivityDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetWebUrl sets the webUrl property value. Link to the degree or program page.
func (m *EducationalActivityDetail) SetWebUrl(value *string)() {
    m.webUrl = value
}
