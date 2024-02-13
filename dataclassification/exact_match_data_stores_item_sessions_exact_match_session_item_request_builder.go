package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.exactMatchDataStore entity.
type ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters get sessions from dataClassification
type ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetQueryParameters
}
// ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *ExactMatchDataStoresItemSessionsItemCancelRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Cancel()(*ExactMatchDataStoresItemSessionsItemCancelRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commit provides operations to call the commit method.
// returns a *ExactMatchDataStoresItemSessionsItemCommitRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Commit()(*ExactMatchDataStoresItemSessionsItemCommitRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemCommitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal instantiates a new ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    m := &ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder instantiates a new ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder and sets the default values.
func NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sessions for dataClassification
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get sessions from dataClassification
// returns a ExactMatchSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable), nil
}
// Patch update the navigation property sessions in dataClassification
// returns a ExactMatchSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable), nil
}
// Renew provides operations to call the renew method.
// returns a *ExactMatchDataStoresItemSessionsItemRenewRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) Renew()(*ExactMatchDataStoresItemSessionsItemRenewRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemRenewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sessions for dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get sessions from dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sessions in dataClassification
// returns a *RequestInformation when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchSessionable, requestConfiguration *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore%2Did}/sessions/{exactMatchSession%2Did}", m.BaseRequestBuilder.PathParameters)
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
// UploadAgent provides operations to manage the uploadAgent property of the microsoft.graph.exactMatchSession entity.
// returns a *ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) UploadAgent()(*ExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsItemUploadAgentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder when successful
func (m *ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) WithUrl(rawUrl string)(*ExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder) {
    return NewExactMatchDataStoresItemSessionsExactMatchSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
