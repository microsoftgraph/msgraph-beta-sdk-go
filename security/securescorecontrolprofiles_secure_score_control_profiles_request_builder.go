package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
type SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetQueryParameters retrieves a list of secureScoreControlProfile objects for a tenant.
type SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetQueryParameters
}
// SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySecureScoreControlProfileId provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
// returns a *SecurescorecontrolprofilesSecureScoreControlProfileItemRequestBuilder when successful
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) BySecureScoreControlProfileId(secureScoreControlProfileId string)(*SecurescorecontrolprofilesSecureScoreControlProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secureScoreControlProfileId != "" {
        urlTplParams["secureScoreControlProfile%2Did"] = secureScoreControlProfileId
    }
    return NewSecurescorecontrolprofilesSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderInternal instantiates a new SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder and sets the default values.
func NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) {
    m := &SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/secureScoreControlProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder instantiates a new SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder and sets the default values.
func NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *SecurescorecontrolprofilesCountRequestBuilder when successful
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) Count()(*SecurescorecontrolprofilesCountRequestBuilder) {
    return NewSecurescorecontrolprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieves a list of secureScoreControlProfile objects for a tenant.
// returns a SecureScoreControlProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securescorecontrolprofiles-list?view=graph-rest-beta
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecureScoreControlProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileCollectionResponseable), nil
}
// Post create new navigation property to secureScoreControlProfiles for security
// returns a SecureScoreControlProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileable, requestConfiguration *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecureScoreControlProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileable), nil
}
// ToGetRequestInformation retrieves a list of secureScoreControlProfile objects for a tenant.
// returns a *RequestInformation when successful
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to secureScoreControlProfiles for security
// returns a *RequestInformation when successful
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecureScoreControlProfileable, requestConfiguration *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder when successful
func (m *SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) WithUrl(rawUrl string)(*SecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder) {
    return NewSecurescorecontrolprofilesSecureScoreControlProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
