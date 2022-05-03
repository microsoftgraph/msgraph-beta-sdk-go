package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/setpriority"
    i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/assignments"
    i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/assignments/item"
)

// AppleUserInitiatedEnrollmentProfileItemRequestBuilder provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AppleUserInitiatedEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments the assignments property
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) Assignments()(*i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031.AssignmentsRequestBuilder) {
    return i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.appleUserInitiatedEnrollmentProfiles.item.assignments.item collection
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) AssignmentsById(id string)(*i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56.AppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleEnrollmentProfileAssignment%2Did"] = id
    }
    return i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56.NewAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppleUserInitiatedEnrollmentProfileItemRequestBuilder instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// PatchWithResponseHandler update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SetPriority the setPriority property
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) SetPriority()(*i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.SetPriorityRequestBuilder) {
    return i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
