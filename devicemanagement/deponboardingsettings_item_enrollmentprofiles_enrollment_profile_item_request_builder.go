package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
type DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetQueryParameters the enrollment profiles.
type DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderInternal instantiates a new DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) {
    m := &DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder instantiates a new DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property enrollmentProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExportMobileConfig provides operations to call the exportMobileConfig method.
// returns a *DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) ExportMobileConfig()(*DeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesItemExportmobileconfigExportMobileConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the enrollment profiles.
// returns a EnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable), nil
}
// Patch update the navigation property enrollmentProfiles in deviceManagement
// returns a EnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable), nil
}
// SetDefaultProfile provides operations to call the setDefaultProfile method.
// returns a *DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) SetDefaultProfile()(*DeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesItemSetdefaultprofileSetDefaultProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property enrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the enrollment profiles.
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property enrollmentProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, requestConfiguration *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateDeviceProfileAssignment provides operations to call the updateDeviceProfileAssignment method.
// returns a *DeponboardingsettingsItemEnrollmentprofilesItemUpdatedeviceprofileassignmentUpdateDeviceProfileAssignmentRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) UpdateDeviceProfileAssignment()(*DeponboardingsettingsItemEnrollmentprofilesItemUpdatedeviceprofileassignmentUpdateDeviceProfileAssignmentRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesItemUpdatedeviceprofileassignmentUpdateDeviceProfileAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder when successful
func (m *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
