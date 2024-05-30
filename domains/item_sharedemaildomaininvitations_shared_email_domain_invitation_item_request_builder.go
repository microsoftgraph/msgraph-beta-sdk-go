package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder provides operations to manage the sharedEmailDomainInvitations property of the microsoft.graph.domain entity.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetQueryParameters get sharedEmailDomainInvitations from domains
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetQueryParameters
}
// ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderInternal instantiates a new ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder and sets the default values.
func NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) {
    m := &ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/sharedEmailDomainInvitations/{sharedEmailDomainInvitation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder instantiates a new ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder and sets the default values.
func NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sharedEmailDomainInvitations for domains
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get sharedEmailDomainInvitations from domains
// returns a SharedEmailDomainInvitationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property sharedEmailDomainInvitations in domains
// returns a SharedEmailDomainInvitationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property sharedEmailDomainInvitations for domains
// returns a *RequestInformation when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get sharedEmailDomainInvitations from domains
// returns a *RequestInformation when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sharedEmailDomainInvitations in domains
// returns a *RequestInformation when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedEmailDomainInvitationable, requestConfiguration *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder when successful
func (m *ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) WithUrl(rawUrl string)(*ItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder) {
    return NewItemSharedemaildomaininvitationsSharedEmailDomainInvitationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
