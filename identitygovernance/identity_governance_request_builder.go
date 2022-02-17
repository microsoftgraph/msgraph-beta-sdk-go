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

// IdentityGovernanceRequestBuilder builds and executes requests for operations under \identityGovernance
type IdentityGovernanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IdentityGovernanceRequestBuilderGetOptions options for Get
type IdentityGovernanceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IdentityGovernanceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IdentityGovernanceRequestBuilderGetQueryParameters get identityGovernance
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IdentityGovernanceRequestBuilderPatchOptions options for Patch
type IdentityGovernanceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.AccessReviewsRequestBuilder) {
    return i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.AppConsentRequestBuilder) {
    return ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceRequestBuilderInternal instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceRequestBuilder instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identityGovernance
func (m *IdentityGovernanceRequestBuilder) CreateGetRequestInformation(options *IdentityGovernanceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update identityGovernance
func (m *IdentityGovernanceRequestBuilder) CreatePatchRequestInformation(options *IdentityGovernanceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.EntitlementManagementRequestBuilder) {
    return i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get identityGovernance
func (m *IdentityGovernanceRequestBuilder) Get(options *IdentityGovernanceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIdentityGovernance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityGovernance), nil
}
// Patch update identityGovernance
func (m *IdentityGovernanceRequestBuilder) Patch(options *IdentityGovernanceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.TermsOfUseRequestBuilder) {
    return ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.NewTermsOfUseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
