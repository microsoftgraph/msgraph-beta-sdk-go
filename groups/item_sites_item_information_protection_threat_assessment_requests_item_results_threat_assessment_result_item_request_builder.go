package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    m := &ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results/{threatAssessmentResult%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder instantiates a new ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property results for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
// returns a ThreatAssessmentResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable), nil
}
// Patch update the navigation property results in groups
// returns a ThreatAssessmentResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable), nil
}
// ToDeleteRequestInformation delete navigation property results for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results/{threatAssessmentResult%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property results in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results/{threatAssessmentResult%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
