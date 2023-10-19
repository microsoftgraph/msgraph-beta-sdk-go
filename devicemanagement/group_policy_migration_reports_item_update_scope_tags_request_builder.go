package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder provides operations to call the updateScopeTags method.
type GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/updateScopeTags", pathParameters),
    }
    return m
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateScopeTags
// Deprecated: This method is obsolete. Use PostAsUpdateScopeTagsPostResponse instead.
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) Post(ctx context.Context, body GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsItemUpdateScopeTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable), nil
}
// PostAsUpdateScopeTagsPostResponse invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) PostAsUpdateScopeTagsPostResponse(ctx context.Context, body GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsItemUpdateScopeTagsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsItemUpdateScopeTagsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsItemUpdateScopeTagsPostResponseable), nil
}
// ToPostRequestInformation invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) {
    return NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
