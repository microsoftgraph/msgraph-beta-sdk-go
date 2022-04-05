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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteOptions options for Delete
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetOptions options for Get
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchOptions options for Patch
type AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
        urlTplParams["appleEnrollmentProfileAssignment_id"] = id
    }
    return i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56.NewAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) Delete(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) Get(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// Patch update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) Patch(options *AppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SetPriority the setPriority property
func (m *AppleUserInitiatedEnrollmentProfileItemRequestBuilder) SetPriority()(*i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.SetPriorityRequestBuilder) {
    return i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
