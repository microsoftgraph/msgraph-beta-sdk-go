package authenticationstrengths

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1ae18e0e9da292b2ad347484d67b5f5431c31339e39ab70ae2414154fe008621 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationstrengths/authenticationmethodmodes"
    i7fac19ff62f87e78800afbe47d58ad82ca2edd4d1022323c5a291949f49c64d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationstrengths/policies"
    i62af1e0b68f40453793fa45c578271f673697b4aaed26e0642d99442b50d4656 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationstrengths/authenticationmethodmodes/item"
    ie490c37e832691221d12185a86be2a2b6c2b300f1f51ea20d762a0cf30744096 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationstrengths/policies/item"
)

// AuthenticationStrengthsRequestBuilder provides operations to manage the authenticationStrengths property of the microsoft.graph.conditionalAccessRoot entity.
type AuthenticationStrengthsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationStrengthsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationStrengthsRequestBuilderGetQueryParameters get authenticationStrengths from identity
type AuthenticationStrengthsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationStrengthsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationStrengthsRequestBuilderGetQueryParameters
}
// AuthenticationStrengthsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationMethodModes the authenticationMethodModes property
func (m *AuthenticationStrengthsRequestBuilder) AuthenticationMethodModes()(*i1ae18e0e9da292b2ad347484d67b5f5431c31339e39ab70ae2414154fe008621.AuthenticationMethodModesRequestBuilder) {
    return i1ae18e0e9da292b2ad347484d67b5f5431c31339e39ab70ae2414154fe008621.NewAuthenticationMethodModesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodModesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.authenticationStrengths.authenticationMethodModes.item collection
func (m *AuthenticationStrengthsRequestBuilder) AuthenticationMethodModesById(id string)(*i62af1e0b68f40453793fa45c578271f673697b4aaed26e0642d99442b50d4656.AuthenticationMethodModeDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethodModeDetail%2Did"] = id
    }
    return i62af1e0b68f40453793fa45c578271f673697b4aaed26e0642d99442b50d4656.NewAuthenticationMethodModeDetailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAuthenticationStrengthsRequestBuilderInternal instantiates a new AuthenticationStrengthsRequestBuilder and sets the default values.
func NewAuthenticationStrengthsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthsRequestBuilder) {
    m := &AuthenticationStrengthsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess/authenticationStrengths{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationStrengthsRequestBuilder instantiates a new AuthenticationStrengthsRequestBuilder and sets the default values.
func NewAuthenticationStrengthsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationStrengthsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authenticationStrengths for identity
func (m *AuthenticationStrengthsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get authenticationStrengths from identity
func (m *AuthenticationStrengthsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property authenticationStrengths in identity
func (m *AuthenticationStrengthsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, requestConfiguration *AuthenticationStrengthsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property authenticationStrengths for identity
func (m *AuthenticationStrengthsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationStrengthsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get authenticationStrengths from identity
func (m *AuthenticationStrengthsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationStrengthsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable), nil
}
// Patch update the navigation property authenticationStrengths in identity
func (m *AuthenticationStrengthsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, requestConfiguration *AuthenticationStrengthsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationStrengthRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationStrengthRootable), nil
}
// Policies the policies property
func (m *AuthenticationStrengthsRequestBuilder) Policies()(*i7fac19ff62f87e78800afbe47d58ad82ca2edd4d1022323c5a291949f49c64d4.PoliciesRequestBuilder) {
    return i7fac19ff62f87e78800afbe47d58ad82ca2edd4d1022323c5a291949f49c64d4.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.authenticationStrengths.policies.item collection
func (m *AuthenticationStrengthsRequestBuilder) PoliciesById(id string)(*ie490c37e832691221d12185a86be2a2b6c2b300f1f51ea20d762a0cf30744096.AuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = id
    }
    return ie490c37e832691221d12185a86be2a2b6c2b300f1f51ea20d762a0cf30744096.NewAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
