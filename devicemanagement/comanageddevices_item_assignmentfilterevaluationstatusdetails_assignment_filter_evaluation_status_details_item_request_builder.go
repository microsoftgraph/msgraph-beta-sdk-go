package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
type ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters managed device mobile app configuration states for this device.
type ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters
}
// ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal instantiates a new ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder and sets the default values.
func NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    m := &ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/assignmentFilterEvaluationStatusDetails/{assignmentFilterEvaluationStatusDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder instantiates a new ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder and sets the default values.
func NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentFilterEvaluationStatusDetails for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get managed device mobile app configuration states for this device.
// returns a AssignmentFilterEvaluationStatusDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable), nil
}
// Patch update the navigation property assignmentFilterEvaluationStatusDetails in deviceManagement
// returns a AssignmentFilterEvaluationStatusDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable), nil
}
// ToDeleteRequestInformation delete navigation property assignmentFilterEvaluationStatusDetails for deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation managed device mobile app configuration states for this device.
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentFilterEvaluationStatusDetails in deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, requestConfiguration *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder when successful
func (m *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    return NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
