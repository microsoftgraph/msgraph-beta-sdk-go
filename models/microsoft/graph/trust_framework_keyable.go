package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFrameworkKeyable 
type TrustFrameworkKeyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetD()(*string)
    GetDp()(*string)
    GetDq()(*string)
    GetE()(*string)
    GetExp()(*int64)
    GetK()(*string)
    GetKid()(*string)
    GetKty()(*string)
    GetN()(*string)
    GetNbf()(*int64)
    GetP()(*string)
    GetQ()(*string)
    GetQi()(*string)
    GetUse()(*string)
    GetX5c()([]string)
    GetX5t()(*string)
    SetD(value *string)()
    SetDp(value *string)()
    SetDq(value *string)()
    SetE(value *string)()
    SetExp(value *int64)()
    SetK(value *string)()
    SetKid(value *string)()
    SetKty(value *string)()
    SetN(value *string)()
    SetNbf(value *int64)()
    SetP(value *string)()
    SetQ(value *string)()
    SetQi(value *string)()
    SetUse(value *string)()
    SetX5c(value []string)()
    SetX5t(value *string)()
}
