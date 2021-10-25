package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerCategoryDescriptions struct {
    additionalData map[string]interface{};
    category1 *string;
    category10 *string;
    category11 *string;
    category12 *string;
    category13 *string;
    category14 *string;
    category15 *string;
    category16 *string;
    category17 *string;
    category18 *string;
    category19 *string;
    category2 *string;
    category20 *string;
    category21 *string;
    category22 *string;
    category23 *string;
    category24 *string;
    category25 *string;
    category3 *string;
    category4 *string;
    category5 *string;
    category6 *string;
    category7 *string;
    category8 *string;
    category9 *string;
}
func NewPlannerCategoryDescriptions()(*PlannerCategoryDescriptions) {
    m := &PlannerCategoryDescriptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PlannerCategoryDescriptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PlannerCategoryDescriptions) GetCategory1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category1
    }
}
func (m *PlannerCategoryDescriptions) GetCategory10()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category10
    }
}
func (m *PlannerCategoryDescriptions) GetCategory11()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category11
    }
}
func (m *PlannerCategoryDescriptions) GetCategory12()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category12
    }
}
func (m *PlannerCategoryDescriptions) GetCategory13()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category13
    }
}
func (m *PlannerCategoryDescriptions) GetCategory14()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category14
    }
}
func (m *PlannerCategoryDescriptions) GetCategory15()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category15
    }
}
func (m *PlannerCategoryDescriptions) GetCategory16()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category16
    }
}
func (m *PlannerCategoryDescriptions) GetCategory17()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category17
    }
}
func (m *PlannerCategoryDescriptions) GetCategory18()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category18
    }
}
func (m *PlannerCategoryDescriptions) GetCategory19()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category19
    }
}
func (m *PlannerCategoryDescriptions) GetCategory2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category2
    }
}
func (m *PlannerCategoryDescriptions) GetCategory20()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category20
    }
}
func (m *PlannerCategoryDescriptions) GetCategory21()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category21
    }
}
func (m *PlannerCategoryDescriptions) GetCategory22()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category22
    }
}
func (m *PlannerCategoryDescriptions) GetCategory23()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category23
    }
}
func (m *PlannerCategoryDescriptions) GetCategory24()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category24
    }
}
func (m *PlannerCategoryDescriptions) GetCategory25()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category25
    }
}
func (m *PlannerCategoryDescriptions) GetCategory3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category3
    }
}
func (m *PlannerCategoryDescriptions) GetCategory4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category4
    }
}
func (m *PlannerCategoryDescriptions) GetCategory5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category5
    }
}
func (m *PlannerCategoryDescriptions) GetCategory6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category6
    }
}
func (m *PlannerCategoryDescriptions) GetCategory7()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category7
    }
}
func (m *PlannerCategoryDescriptions) GetCategory8()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category8
    }
}
func (m *PlannerCategoryDescriptions) GetCategory9()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category9
    }
}
func (m *PlannerCategoryDescriptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory1(val)
        return nil
    }
    res["category10"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory10(val)
        return nil
    }
    res["category11"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory11(val)
        return nil
    }
    res["category12"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory12(val)
        return nil
    }
    res["category13"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory13(val)
        return nil
    }
    res["category14"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory14(val)
        return nil
    }
    res["category15"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory15(val)
        return nil
    }
    res["category16"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory16(val)
        return nil
    }
    res["category17"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory17(val)
        return nil
    }
    res["category18"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory18(val)
        return nil
    }
    res["category19"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory19(val)
        return nil
    }
    res["category2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory2(val)
        return nil
    }
    res["category20"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory20(val)
        return nil
    }
    res["category21"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory21(val)
        return nil
    }
    res["category22"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory22(val)
        return nil
    }
    res["category23"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory23(val)
        return nil
    }
    res["category24"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory24(val)
        return nil
    }
    res["category25"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory25(val)
        return nil
    }
    res["category3"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory3(val)
        return nil
    }
    res["category4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory4(val)
        return nil
    }
    res["category5"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory5(val)
        return nil
    }
    res["category6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory6(val)
        return nil
    }
    res["category7"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory7(val)
        return nil
    }
    res["category8"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory8(val)
        return nil
    }
    res["category9"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory9(val)
        return nil
    }
    return res
}
func (m *PlannerCategoryDescriptions) IsNil()(bool) {
    return m == nil
}
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
func (m *PlannerCategoryDescriptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PlannerCategoryDescriptions) SetCategory1(value *string)() {
    m.category1 = value
}
func (m *PlannerCategoryDescriptions) SetCategory10(value *string)() {
    m.category10 = value
}
func (m *PlannerCategoryDescriptions) SetCategory11(value *string)() {
    m.category11 = value
}
func (m *PlannerCategoryDescriptions) SetCategory12(value *string)() {
    m.category12 = value
}
func (m *PlannerCategoryDescriptions) SetCategory13(value *string)() {
    m.category13 = value
}
func (m *PlannerCategoryDescriptions) SetCategory14(value *string)() {
    m.category14 = value
}
func (m *PlannerCategoryDescriptions) SetCategory15(value *string)() {
    m.category15 = value
}
func (m *PlannerCategoryDescriptions) SetCategory16(value *string)() {
    m.category16 = value
}
func (m *PlannerCategoryDescriptions) SetCategory17(value *string)() {
    m.category17 = value
}
func (m *PlannerCategoryDescriptions) SetCategory18(value *string)() {
    m.category18 = value
}
func (m *PlannerCategoryDescriptions) SetCategory19(value *string)() {
    m.category19 = value
}
func (m *PlannerCategoryDescriptions) SetCategory2(value *string)() {
    m.category2 = value
}
func (m *PlannerCategoryDescriptions) SetCategory20(value *string)() {
    m.category20 = value
}
func (m *PlannerCategoryDescriptions) SetCategory21(value *string)() {
    m.category21 = value
}
func (m *PlannerCategoryDescriptions) SetCategory22(value *string)() {
    m.category22 = value
}
func (m *PlannerCategoryDescriptions) SetCategory23(value *string)() {
    m.category23 = value
}
func (m *PlannerCategoryDescriptions) SetCategory24(value *string)() {
    m.category24 = value
}
func (m *PlannerCategoryDescriptions) SetCategory25(value *string)() {
    m.category25 = value
}
func (m *PlannerCategoryDescriptions) SetCategory3(value *string)() {
    m.category3 = value
}
func (m *PlannerCategoryDescriptions) SetCategory4(value *string)() {
    m.category4 = value
}
func (m *PlannerCategoryDescriptions) SetCategory5(value *string)() {
    m.category5 = value
}
func (m *PlannerCategoryDescriptions) SetCategory6(value *string)() {
    m.category6 = value
}
func (m *PlannerCategoryDescriptions) SetCategory7(value *string)() {
    m.category7 = value
}
func (m *PlannerCategoryDescriptions) SetCategory8(value *string)() {
    m.category8 = value
}
func (m *PlannerCategoryDescriptions) SetCategory9(value *string)() {
    m.category9 = value
}
