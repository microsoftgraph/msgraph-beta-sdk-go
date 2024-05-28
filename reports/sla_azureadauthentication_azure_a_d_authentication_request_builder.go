package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SlaAzureadauthenticationAzureADAuthenticationRequestBuilder provides operations to manage the azureADAuthentication property of the microsoft.graph.serviceLevelAgreementRoot entity.
type SlaAzureadauthenticationAzureADAuthenticationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SlaAzureadauthenticationAzureADAuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SlaAzureadauthenticationAzureADAuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetQueryParameters read the properties and relationships of an azureADAuthentication object to find the level of Microsoft Entra authentication availability for your tenant. The Microsoft Entra service Level Agreement (SLA) commits to at least 99.99% authentication availability, as described in Microsoft Entra SLA performance. This object provides you with your tenant's actual performance against this commitment.
type SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetQueryParameters
}
// SlaAzureadauthenticationAzureADAuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SlaAzureadauthenticationAzureADAuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilderInternal instantiates a new SlaAzureadauthenticationAzureADAuthenticationRequestBuilder and sets the default values.
func NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) {
    m := &SlaAzureadauthenticationAzureADAuthenticationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/sla/azureADAuthentication{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilder instantiates a new SlaAzureadauthenticationAzureADAuthenticationRequestBuilder and sets the default values.
func NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property azureADAuthentication for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an azureADAuthentication object to find the level of Microsoft Entra authentication availability for your tenant. The Microsoft Entra service Level Agreement (SLA) commits to at least 99.99% authentication availability, as described in Microsoft Entra SLA performance. This object provides you with your tenant's actual performance against this commitment.
// returns a AzureADAuthenticationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/azureadauthentication-get?view=graph-rest-beta
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAzureADAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable), nil
}
// Patch update the navigation property azureADAuthentication in reports
// returns a AzureADAuthenticationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAzureADAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable), nil
}
// ToDeleteRequestInformation delete navigation property azureADAuthentication for reports
// returns a *RequestInformation when successful
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an azureADAuthentication object to find the level of Microsoft Entra authentication availability for your tenant. The Microsoft Entra service Level Agreement (SLA) commits to at least 99.99% authentication availability, as described in Microsoft Entra SLA performance. This object provides you with your tenant's actual performance against this commitment.
// returns a *RequestInformation when successful
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property azureADAuthentication in reports
// returns a *RequestInformation when successful
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AzureADAuthenticationable, requestConfiguration *SlaAzureadauthenticationAzureADAuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder when successful
func (m *SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) WithUrl(rawUrl string)(*SlaAzureadauthenticationAzureADAuthenticationRequestBuilder) {
    return NewSlaAzureadauthenticationAzureADAuthenticationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
