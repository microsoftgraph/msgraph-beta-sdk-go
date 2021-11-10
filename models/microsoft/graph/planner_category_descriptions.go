package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerCategoryDescriptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The label associated with Category 1
    category1 *string;
    // The label associated with Category 10
    category10 *string;
    // The label associated with Category 11
    category11 *string;
    // The label associated with Category 12
    category12 *string;
    // The label associated with Category 13
    category13 *string;
    // The label associated with Category 14
    category14 *string;
    // The label associated with Category 15
    category15 *string;
    // The label associated with Category 16
    category16 *string;
    // The label associated with Category 17
    category17 *string;
    // The label associated with Category 18
    category18 *string;
    // The label associated with Category 19
    category19 *string;
    // The label associated with Category 2
    category2 *string;
    // The label associated with Category 20
    category20 *string;
    // The label associated with Category 21
    category21 *string;
    // The label associated with Category 22
    category22 *string;
    // The label associated with Category 23
    category23 *string;
    // The label associated with Category 24
    category24 *string;
    // The label associated with Category 25
    category25 *string;
    // The label associated with Category 3
    category3 *string;
    // The label associated with Category 4
    category4 *string;
    // The label associated with Category 5
    category5 *string;
    // The label associated with Category 6
    category6 *string;
    // The label associated with Category 7
    category7 *string;
    // The label associated with Category 8
    category8 *string;
    // The label associated with Category 9
    category9 *string;
}
// Instantiates a new plannerCategoryDescriptions and sets the default values.
func NewPlannerCategoryDescriptions()(*PlannerCategoryDescriptions) {
    m := &PlannerCategoryDescriptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerCategoryDescriptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the category1 property value. The label associated with Category 1
func (m *PlannerCategoryDescriptions) GetCategory1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category1
    }
}
// Gets the category10 property value. The label associated with Category 10
func (m *PlannerCategoryDescriptions) GetCategory10()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category10
    }
}
// Gets the category11 property value. The label associated with Category 11
func (m *PlannerCategoryDescriptions) GetCategory11()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category11
    }
}
// Gets the category12 property value. The label associated with Category 12
func (m *PlannerCategoryDescriptions) GetCategory12()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category12
    }
}
// Gets the category13 property value. The label associated with Category 13
func (m *PlannerCategoryDescriptions) GetCategory13()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category13
    }
}
// Gets the category14 property value. The label associated with Category 14
func (m *PlannerCategoryDescriptions) GetCategory14()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category14
    }
}
// Gets the category15 property value. The label associated with Category 15
func (m *PlannerCategoryDescriptions) GetCategory15()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category15
    }
}
// Gets the category16 property value. The label associated with Category 16
func (m *PlannerCategoryDescriptions) GetCategory16()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category16
    }
}
// Gets the category17 property value. The label associated with Category 17
func (m *PlannerCategoryDescriptions) GetCategory17()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category17
    }
}
// Gets the category18 property value. The label associated with Category 18
func (m *PlannerCategoryDescriptions) GetCategory18()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category18
    }
}
// Gets the category19 property value. The label associated with Category 19
func (m *PlannerCategoryDescriptions) GetCategory19()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category19
    }
}
// Gets the category2 property value. The label associated with Category 2
func (m *PlannerCategoryDescriptions) GetCategory2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category2
    }
}
// Gets the category20 property value. The label associated with Category 20
func (m *PlannerCategoryDescriptions) GetCategory20()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category20
    }
}
// Gets the category21 property value. The label associated with Category 21
func (m *PlannerCategoryDescriptions) GetCategory21()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category21
    }
}
// Gets the category22 property value. The label associated with Category 22
func (m *PlannerCategoryDescriptions) GetCategory22()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category22
    }
}
// Gets the category23 property value. The label associated with Category 23
func (m *PlannerCategoryDescriptions) GetCategory23()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category23
    }
}
// Gets the category24 property value. The label associated with Category 24
func (m *PlannerCategoryDescriptions) GetCategory24()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category24
    }
}
// Gets the category25 property value. The label associated with Category 25
func (m *PlannerCategoryDescriptions) GetCategory25()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category25
    }
}
// Gets the category3 property value. The label associated with Category 3
func (m *PlannerCategoryDescriptions) GetCategory3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category3
    }
}
// Gets the category4 property value. The label associated with Category 4
func (m *PlannerCategoryDescriptions) GetCategory4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category4
    }
}
// Gets the category5 property value. The label associated with Category 5
func (m *PlannerCategoryDescriptions) GetCategory5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category5
    }
}
// Gets the category6 property value. The label associated with Category 6
func (m *PlannerCategoryDescriptions) GetCategory6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category6
    }
}
// Gets the category7 property value. The label associated with Category 7
func (m *PlannerCategoryDescriptions) GetCategory7()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category7
    }
}
// Gets the category8 property value. The label associated with Category 8
func (m *PlannerCategoryDescriptions) GetCategory8()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category8
    }
}
// Gets the category9 property value. The label associated with Category 9
func (m *PlannerCategoryDescriptions) GetCategory9()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category9
    }
}
// The deserialization information for the current model
func (m *PlannerCategoryDescriptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory1(val)
        }
        return nil
    }
    res["category10"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory10(val)
        }
        return nil
    }
    res["category11"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory11(val)
        }
        return nil
    }
    res["category12"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory12(val)
        }
        return nil
    }
    res["category13"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory13(val)
        }
        return nil
    }
    res["category14"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory14(val)
        }
        return nil
    }
    res["category15"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory15(val)
        }
        return nil
    }
    res["category16"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory16(val)
        }
        return nil
    }
    res["category17"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory17(val)
        }
        return nil
    }
    res["category18"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory18(val)
        }
        return nil
    }
    res["category19"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory19(val)
        }
        return nil
    }
    res["category2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory2(val)
        }
        return nil
    }
    res["category20"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory20(val)
        }
        return nil
    }
    res["category21"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory21(val)
        }
        return nil
    }
    res["category22"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory22(val)
        }
        return nil
    }
    res["category23"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory23(val)
        }
        return nil
    }
    res["category24"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory24(val)
        }
        return nil
    }
    res["category25"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory25(val)
        }
        return nil
    }
    res["category3"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory3(val)
        }
        return nil
    }
    res["category4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory4(val)
        }
        return nil
    }
    res["category5"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory5(val)
        }
        return nil
    }
    res["category6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory6(val)
        }
        return nil
    }
    res["category7"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory7(val)
        }
        return nil
    }
    res["category8"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory8(val)
        }
        return nil
    }
    res["category9"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory9(val)
        }
        return nil
    }
    return res
}
func (m *PlannerCategoryDescriptions) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PlannerCategoryDescriptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category1", m.GetCategory1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category10", m.GetCategory10())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category11", m.GetCategory11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category12", m.GetCategory12())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category13", m.GetCategory13())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category14", m.GetCategory14())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category15", m.GetCategory15())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category16", m.GetCategory16())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category17", m.GetCategory17())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category18", m.GetCategory18())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category19", m.GetCategory19())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category2", m.GetCategory2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category20", m.GetCategory20())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category21", m.GetCategory21())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category22", m.GetCategory22())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category23", m.GetCategory23())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category24", m.GetCategory24())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category25", m.GetCategory25())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category3", m.GetCategory3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category4", m.GetCategory4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category5", m.GetCategory5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category6", m.GetCategory6())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category7", m.GetCategory7())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category8", m.GetCategory8())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category9", m.GetCategory9())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PlannerCategoryDescriptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the category1 property value. The label associated with Category 1
// Parameters:
//  - value : Value to set for the category1 property.
func (m *PlannerCategoryDescriptions) SetCategory1(value *string)() {
    m.category1 = value
}
// Sets the category10 property value. The label associated with Category 10
// Parameters:
//  - value : Value to set for the category10 property.
func (m *PlannerCategoryDescriptions) SetCategory10(value *string)() {
    m.category10 = value
}
// Sets the category11 property value. The label associated with Category 11
// Parameters:
//  - value : Value to set for the category11 property.
func (m *PlannerCategoryDescriptions) SetCategory11(value *string)() {
    m.category11 = value
}
// Sets the category12 property value. The label associated with Category 12
// Parameters:
//  - value : Value to set for the category12 property.
func (m *PlannerCategoryDescriptions) SetCategory12(value *string)() {
    m.category12 = value
}
// Sets the category13 property value. The label associated with Category 13
// Parameters:
//  - value : Value to set for the category13 property.
func (m *PlannerCategoryDescriptions) SetCategory13(value *string)() {
    m.category13 = value
}
// Sets the category14 property value. The label associated with Category 14
// Parameters:
//  - value : Value to set for the category14 property.
func (m *PlannerCategoryDescriptions) SetCategory14(value *string)() {
    m.category14 = value
}
// Sets the category15 property value. The label associated with Category 15
// Parameters:
//  - value : Value to set for the category15 property.
func (m *PlannerCategoryDescriptions) SetCategory15(value *string)() {
    m.category15 = value
}
// Sets the category16 property value. The label associated with Category 16
// Parameters:
//  - value : Value to set for the category16 property.
func (m *PlannerCategoryDescriptions) SetCategory16(value *string)() {
    m.category16 = value
}
// Sets the category17 property value. The label associated with Category 17
// Parameters:
//  - value : Value to set for the category17 property.
func (m *PlannerCategoryDescriptions) SetCategory17(value *string)() {
    m.category17 = value
}
// Sets the category18 property value. The label associated with Category 18
// Parameters:
//  - value : Value to set for the category18 property.
func (m *PlannerCategoryDescriptions) SetCategory18(value *string)() {
    m.category18 = value
}
// Sets the category19 property value. The label associated with Category 19
// Parameters:
//  - value : Value to set for the category19 property.
func (m *PlannerCategoryDescriptions) SetCategory19(value *string)() {
    m.category19 = value
}
// Sets the category2 property value. The label associated with Category 2
// Parameters:
//  - value : Value to set for the category2 property.
func (m *PlannerCategoryDescriptions) SetCategory2(value *string)() {
    m.category2 = value
}
// Sets the category20 property value. The label associated with Category 20
// Parameters:
//  - value : Value to set for the category20 property.
func (m *PlannerCategoryDescriptions) SetCategory20(value *string)() {
    m.category20 = value
}
// Sets the category21 property value. The label associated with Category 21
// Parameters:
//  - value : Value to set for the category21 property.
func (m *PlannerCategoryDescriptions) SetCategory21(value *string)() {
    m.category21 = value
}
// Sets the category22 property value. The label associated with Category 22
// Parameters:
//  - value : Value to set for the category22 property.
func (m *PlannerCategoryDescriptions) SetCategory22(value *string)() {
    m.category22 = value
}
// Sets the category23 property value. The label associated with Category 23
// Parameters:
//  - value : Value to set for the category23 property.
func (m *PlannerCategoryDescriptions) SetCategory23(value *string)() {
    m.category23 = value
}
// Sets the category24 property value. The label associated with Category 24
// Parameters:
//  - value : Value to set for the category24 property.
func (m *PlannerCategoryDescriptions) SetCategory24(value *string)() {
    m.category24 = value
}
// Sets the category25 property value. The label associated with Category 25
// Parameters:
//  - value : Value to set for the category25 property.
func (m *PlannerCategoryDescriptions) SetCategory25(value *string)() {
    m.category25 = value
}
// Sets the category3 property value. The label associated with Category 3
// Parameters:
//  - value : Value to set for the category3 property.
func (m *PlannerCategoryDescriptions) SetCategory3(value *string)() {
    m.category3 = value
}
// Sets the category4 property value. The label associated with Category 4
// Parameters:
//  - value : Value to set for the category4 property.
func (m *PlannerCategoryDescriptions) SetCategory4(value *string)() {
    m.category4 = value
}
// Sets the category5 property value. The label associated with Category 5
// Parameters:
//  - value : Value to set for the category5 property.
func (m *PlannerCategoryDescriptions) SetCategory5(value *string)() {
    m.category5 = value
}
// Sets the category6 property value. The label associated with Category 6
// Parameters:
//  - value : Value to set for the category6 property.
func (m *PlannerCategoryDescriptions) SetCategory6(value *string)() {
    m.category6 = value
}
// Sets the category7 property value. The label associated with Category 7
// Parameters:
//  - value : Value to set for the category7 property.
func (m *PlannerCategoryDescriptions) SetCategory7(value *string)() {
    m.category7 = value
}
// Sets the category8 property value. The label associated with Category 8
// Parameters:
//  - value : Value to set for the category8 property.
func (m *PlannerCategoryDescriptions) SetCategory8(value *string)() {
    m.category8 = value
}
// Sets the category9 property value. The label associated with Category 9
// Parameters:
//  - value : Value to set for the category9 property.
func (m *PlannerCategoryDescriptions) SetCategory9(value *string)() {
    m.category9 = value
}
