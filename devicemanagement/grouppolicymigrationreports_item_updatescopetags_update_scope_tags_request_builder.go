package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder provides operations to call the updateScopeTags method.
type GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderInternal instantiates a new GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) {
    m := &GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/updateScopeTags", pathParameters),
    }
    return m
}
// NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder instantiates a new GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateScopeTags
// Deprecated: This method is obsolete. Use PostAsUpdateScopeTagsPostResponse instead.
// returns a GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) Post(ctx context.Context, body GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsResponseable), nil
}
// PostAsUpdateScopeTagsPostResponse invoke action updateScopeTags
// returns a GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) PostAsUpdateScopeTagsPostResponse(ctx context.Context, body GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostResponseable), nil
}
// ToPostRequestInformation invoke action updateScopeTags
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder when successful
func (m *GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) WithUrl(rawUrl string)(*GrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder) {
    return NewGrouppolicymigrationreportsItemUpdatescopetagsUpdateScopeTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
