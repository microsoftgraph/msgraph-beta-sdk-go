package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement"
    ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/appconsent"
    ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse"
)

// IdentityGovernanceRequestBuilder provides operations to manage the identityGovernance singleton.
type IdentityGovernanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceRequestBuilderGetOptions options for Get
type IdentityGovernanceRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// IdentityGovernanceRequestBuilderGetQueryParameters get identityGovernance
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceRequestBuilderPatchOptions options for Patch
type IdentityGovernanceRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessReviews the accessReviews property
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.AccessReviewsRequestBuilder) {
    return i1634d124e13ffd0cab6d792b39af6165594c5183595443f6dbd586fb2dacb598.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsent the appConsent property
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.AppConsentRequestBuilder) {
    return ia84e7b6e98da4aca360dbefe526b5e9895268f91ebdbd56a0d404de509c7705b.NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceRequestBuilderInternal instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceRequestBuilder instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identityGovernance
func (m *IdentityGovernanceRequestBuilder) CreateGetRequestInformation(options *IdentityGovernanceRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update identityGovernance
func (m *IdentityGovernanceRequestBuilder) CreatePatchRequestInformation(options *IdentityGovernanceRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// EntitlementManagement the entitlementManagement property
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.EntitlementManagementRequestBuilder) {
    return i980a16b671cb1d9b9251e92f61e0a38b85d73c7e3da0d7d33055d60e7e5583df.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get identityGovernance
func (m *IdentityGovernanceRequestBuilder) Get(options *IdentityGovernanceRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityGovernanceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable), nil
}
// Patch update identityGovernance
func (m *IdentityGovernanceRequestBuilder) Patch(options *IdentityGovernanceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// TermsOfUse the termsOfUse property
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.TermsOfUseRequestBuilder) {
    return ic8d738e0ef5dd83881ad42286a210c4014b53022a6b1ca33fb97933713ff69c6.NewTermsOfUseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
