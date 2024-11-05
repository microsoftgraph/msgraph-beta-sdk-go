package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder provides operations to manage the reflectCheckInResponses property of the microsoft.graph.reportsRoot entity.
type ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetQueryParameters get reflectCheckInResponses from education
type ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetQueryParameters
}
// ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderInternal instantiates a new ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder and sets the default values.
func NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) {
    m := &ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/reports/reflectCheckInResponses/{reflectCheckInResponse%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder instantiates a new ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder and sets the default values.
func NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reflectCheckInResponses for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get reflectCheckInResponses from education
// returns a ReflectCheckInResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReflectCheckInResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable), nil
}
// Patch update the navigation property reflectCheckInResponses in education
// returns a ReflectCheckInResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReflectCheckInResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable), nil
}
// ToDeleteRequestInformation delete navigation property reflectCheckInResponses for education
// returns a *RequestInformation when successful
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get reflectCheckInResponses from education
// returns a *RequestInformation when successful
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property reflectCheckInResponses in education
// returns a *RequestInformation when successful
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReflectCheckInResponseable, requestConfiguration *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder when successful
func (m *ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) WithUrl(rawUrl string)(*ReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder) {
    return NewReportsReflectCheckInResponsesReflectCheckInResponseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
