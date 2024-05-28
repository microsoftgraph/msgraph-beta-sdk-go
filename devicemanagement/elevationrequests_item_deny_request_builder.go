package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ElevationrequestsItemDenyRequestBuilder provides operations to call the deny method.
type ElevationrequestsItemDenyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ElevationrequestsItemDenyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ElevationrequestsItemDenyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewElevationrequestsItemDenyRequestBuilderInternal instantiates a new ElevationrequestsItemDenyRequestBuilder and sets the default values.
func NewElevationrequestsItemDenyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsItemDenyRequestBuilder) {
    m := &ElevationrequestsItemDenyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/elevationRequests/{privilegeManagementElevationRequest%2Did}/deny", pathParameters),
    }
    return m
}
// NewElevationrequestsItemDenyRequestBuilder instantiates a new ElevationrequestsItemDenyRequestBuilder and sets the default values.
func NewElevationrequestsItemDenyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ElevationrequestsItemDenyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewElevationrequestsItemDenyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action deny
// returns a PrivilegeManagementElevationRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ElevationrequestsItemDenyRequestBuilder) Post(ctx context.Context, body ElevationrequestsItemDenyPostRequestBodyable, requestConfiguration *ElevationrequestsItemDenyRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegeManagementElevationRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationRequestable), nil
}
// ToPostRequestInformation invoke action deny
// returns a *RequestInformation when successful
func (m *ElevationrequestsItemDenyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ElevationrequestsItemDenyPostRequestBodyable, requestConfiguration *ElevationrequestsItemDenyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ElevationrequestsItemDenyRequestBuilder when successful
func (m *ElevationrequestsItemDenyRequestBuilder) WithUrl(rawUrl string)(*ElevationrequestsItemDenyRequestBuilder) {
    return NewElevationrequestsItemDenyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
