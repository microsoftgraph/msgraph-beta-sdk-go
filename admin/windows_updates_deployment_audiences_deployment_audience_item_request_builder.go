package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder provides operations to manage the deploymentAudiences property of the microsoft.graph.adminWindowsUpdates entity.
type WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetQueryParameters read the properties and relationships of a deploymentAudience object.
type WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderInternal instantiates a new DeploymentAudienceItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder instantiates a new DeploymentAudienceItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a deploymentAudience object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-delete?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Exclusions provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) Exclusions()(*WindowsUpdatesDeploymentAudiencesItemExclusionsRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a deploymentAudience object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-get?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) Members()(*WindowsUpdatesDeploymentAudiencesItemMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUpdateAudience provides operations to call the updateAudience method.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) MicrosoftGraphWindowsUpdatesUpdateAudience()(*WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphWindowsUpdatesUpdateAudienceById provides operations to call the updateAudienceById method.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) MicrosoftGraphWindowsUpdatesUpdateAudienceById()(*WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deploymentAudiences in admin
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// ToDeleteRequestInformation delete a deploymentAudience object.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a deploymentAudience object.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property deploymentAudiences in admin
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesDeploymentAudienceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
