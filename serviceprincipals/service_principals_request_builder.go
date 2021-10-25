package serviceprincipals

import (
    i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/delta"
    i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/getbyids"
    i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/validateproperties"
    ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/getuserownedobjects"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type ServicePrincipalsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ServicePrincipalsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func NewServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    m := &ServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/servicePrincipals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewServicePrincipalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ServicePrincipalsRequestBuilder) CreateGetRequestInformation(q func (value *ServicePrincipalsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ServicePrincipalsRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalsRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalsRequestBuilder) Delta()(*i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744.DeltaRequestBuilder) {
    return i07ea82e8cdb01938ad025e6648ae29c3c96f2adbafa7de6af15c650980651744.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalsRequestBuilder) Get(q func (value *ServicePrincipalsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ServicePrincipalsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipalsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ServicePrincipalsResponse), nil
}
func (m *ServicePrincipalsRequestBuilder) GetByIds()(*i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62.GetByIdsRequestBuilder) {
    return i13ad01682da78442a9cf027f302e407c5d377fb9738dde4e75f036b89203db62.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalsRequestBuilder) GetUserOwnedObjects()(*ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a.GetUserOwnedObjectsRequestBuilder) {
    return ic1b17114e49e92b3da03ac7e6328ba8f143e419e9a0cb2b80fe833a35584283a.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalsRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServicePrincipal() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal), nil
}
func (m *ServicePrincipalsRequestBuilder) ValidateProperties()(*i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7.ValidatePropertiesRequestBuilder) {
    return i59f9b4fdf3f63fe1a7bfbabda2130077511ff77c1f090e63ad3afc1690b2a0b7.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
