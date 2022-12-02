package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetQueryParameters the enrollment profiles.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderInternal instantiates a new EnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) {
    m := &DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder instantiates a new EnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property enrollmentProfiles for deviceManagement
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the enrollment profiles.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property enrollmentProfiles in deviceManagement
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property enrollmentProfiles for deviceManagement
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportMobileConfig provides operations to call the exportMobileConfig method.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) ExportMobileConfig()(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the enrollment profiles.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable), nil
}
// Patch update the navigation property enrollmentProfiles in deviceManagement
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentProfileable), nil
}
// SetDefaultProfile provides operations to call the setDefaultProfile method.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) SetDefaultProfile()(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateDeviceProfileAssignment provides operations to call the updateDeviceProfileAssignment method.
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) UpdateDeviceProfileAssignment()(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemUpdateDeviceProfileAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
