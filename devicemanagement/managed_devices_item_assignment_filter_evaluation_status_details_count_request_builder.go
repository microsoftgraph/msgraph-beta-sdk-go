package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder provides operations to count the resources in the collection.
type ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetQueryParameters get the number of the resource
type ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetQueryParameters
}
// NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) {
    m := &ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/assignmentFilterEvaluationStatusDetails/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder) {
    return NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
