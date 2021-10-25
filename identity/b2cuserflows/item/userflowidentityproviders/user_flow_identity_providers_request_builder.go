package userflowidentityproviders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders/ref"
    i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders/availableprovidertypes"
)

type UserFlowIdentityProvidersRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func (m *UserFlowIdentityProvidersRequestBuilder) AvailableProviderTypes()(*i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b.AvailableProviderTypesRequestBuilder) {
    return i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b.NewAvailableProviderTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    m := &UserFlowIdentityProvidersRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identity/b2cUserFlows/{b2cIdentityUserFlow_id}/userFlowIdentityProviders{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserFlowIdentityProvidersRequestBuilder) CreateGetRequestInformation(q func (value *UserFlowIdentityProvidersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UserFlowIdentityProvidersRequestBuilderGetQueryParameters)
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
func (m *UserFlowIdentityProvidersRequestBuilder) Get(q func (value *UserFlowIdentityProvidersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*UserFlowIdentityProvidersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowIdentityProvidersResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*UserFlowIdentityProvidersResponse), nil
}
func (m *UserFlowIdentityProvidersRequestBuilder) Ref()(*i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e.RefRequestBuilder) {
    return i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
