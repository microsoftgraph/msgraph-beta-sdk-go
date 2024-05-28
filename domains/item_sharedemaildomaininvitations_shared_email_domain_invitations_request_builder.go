package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder provides operations to manage the sharedEmailDomainInvitations property of the microsoft.graph.domain entity.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetQueryParameters get sharedEmailDomainInvitations from domains
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetQueryParameters struct {
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
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetQueryParameters
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySharedEmailDomainInvitationId provides operations to manage the sharedEmailDomainInvitations property of the microsoft.graph.domain entity.
// returns a *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) BySharedEmailDomainInvitationId(sharedEmailDomainInvitationId string)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sharedEmailDomainInvitationId != "" {
        urlTplParams["sharedEmailDomainInvitation%2Did"] = sharedEmailDomainInvitationId
    }
    return NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderInternal instantiates a new ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder and sets the default values.
func NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) {
    m := &ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/sharedEmailDomainInvitations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder instantiates a new ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder and sets the default values.
func NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSharedemaildomaininvitationsCountRequestBuilder when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) Count()(*ItemSharedemaildomaininvitationsCountRequestBuilder) {
    return NewItemSharedemaildomaininvitationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get sharedEmailDomainInvitations from domains
// returns a SharedEmailDomainInvitationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharedEmailDomainInvitationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationCollectionResponseable), nil
}
// Post create new navigation property to sharedEmailDomainInvitations for domains
// returns a SharedEmailDomainInvitationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharedEmailDomainInvitationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable), nil
}
// ToGetRequestInformation get sharedEmailDomainInvitations from domains
// returns a *RequestInformation when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to sharedEmailDomainInvitations for domains
// returns a *RequestInformation when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) WithUrl(rawUrl string)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder) {
    return NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
