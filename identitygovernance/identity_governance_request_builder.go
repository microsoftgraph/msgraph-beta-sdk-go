package identitygovernance

import (
    i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews"
    i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement"
    ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/appconsent"
    ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type IdentityGovernanceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.AccessReviewsRequestBuilder) {
    return i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.AppConsentRequestBuilder) {
    return ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IdentityGovernanceRequestBuilder) CreateGetRequestInformation(q func (value *IdentityGovernanceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(IdentityGovernanceRequestBuilderGetQueryParameters)
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
func (m *IdentityGovernanceRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.EntitlementManagementRequestBuilder) {
    return i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IdentityGovernanceRequestBuilder) Get(q func (value *IdentityGovernanceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIdentityGovernance() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance), nil
}
func (m *IdentityGovernanceRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.TermsOfUseRequestBuilder) {
    return ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.NewTermsOfUseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
