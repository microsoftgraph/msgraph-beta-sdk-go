package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// UnifiedGroupSource provides operations to manage the compliance singleton.
type UnifiedGroupSource struct {
    DataSource
    // 
    group Groupable;
    // Specifies which sources are included in this group. Possible values are: mailbox, site.
    includedSources *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.SourceType;
}
// NewUnifiedGroupSource instantiates a new unifiedGroupSource and sets the default values.
func NewUnifiedGroupSource()(*UnifiedGroupSource) {
    m := &UnifiedGroupSource{
        DataSource: *NewDataSource(),
    }
    return m
}
// CreateUnifiedGroupSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedGroupSourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnifiedGroupSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedGroupSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
    res["includedSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedSources(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.SourceType))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. 
func (m *UnifiedGroupSource) GetGroup()(Groupable) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// GetIncludedSources gets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UnifiedGroupSource) GetIncludedSources()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.SourceType) {
    if m == nil {
        return nil
    } else {
        return m.includedSources
    }
}
func (m *UnifiedGroupSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UnifiedGroupSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedSources() != nil {
        cast := (*m.GetIncludedSources()).String()
        err = writer.WriteStringValue("includedSources", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. 
func (m *UnifiedGroupSource) SetGroup(value Groupable)() {
    if m != nil {
        m.group = value
    }
}
// SetIncludedSources sets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UnifiedGroupSource) SetIncludedSources(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.SourceType)() {
    if m != nil {
        m.includedSources = value
    }
}
