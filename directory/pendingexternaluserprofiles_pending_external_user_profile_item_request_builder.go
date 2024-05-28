package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder provides operations to manage the pendingExternalUserProfiles property of the microsoft.graph.directory entity.
type PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetQueryParameters retrieve the properties of a specific pendingExternalUserProfile.
type PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetQueryParameters
}
// PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderInternal instantiates a new PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder and sets the default values.
func NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) {
    m := &PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/pendingExternalUserProfiles/{pendingExternalUserProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder instantiates a new PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder and sets the default values.
func NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a pendingExternalUserProfile object. Note: To permanently delete the pendingExternalUserProfile, follow permanently delete an item. To restore a pendingExternalUserProfile, follow restore a deleted item.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-delete-pendingexternaluserprofiles?view=graph-rest-beta
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties of a specific pendingExternalUserProfile.
// returns a PendingExternalUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/pendingexternaluserprofile-get?view=graph-rest-beta
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePendingExternalUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable), nil
}
// Patch update the properties of a pendingExternalUserProfile object.
// returns a PendingExternalUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/pendingexternaluserprofile-update?view=graph-rest-beta
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePendingExternalUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable), nil
}
// ToDeleteRequestInformation delete a pendingExternalUserProfile object. Note: To permanently delete the pendingExternalUserProfile, follow permanently delete an item. To restore a pendingExternalUserProfile, follow restore a deleted item.
// returns a *RequestInformation when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties of a specific pendingExternalUserProfile.
// returns a *RequestInformation when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a pendingExternalUserProfile object.
// returns a *RequestInformation when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PendingExternalUserProfileable, requestConfiguration *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder when successful
func (m *PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) WithUrl(rawUrl string)(*PendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder) {
    return NewPendingexternaluserprofilesPendingExternalUserProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
