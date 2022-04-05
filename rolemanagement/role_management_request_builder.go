package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement"
    i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory"
    ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc"
)

// RoleManagementRequestBuilder provides operations to manage the roleManagement singleton.
type RoleManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleManagementRequestBuilderGetOptions options for Get
type RoleManagementRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *RoleManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// RoleManagementRequestBuilderGetQueryParameters get roleManagement
type RoleManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RoleManagementRequestBuilderPatchOptions options for Patch
type RoleManagementRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// CloudPC the cloudPC property
func (m *RoleManagementRequestBuilder) CloudPC()(*ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.CloudPCRequestBuilder) {
    return ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.NewCloudPCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementRequestBuilderInternal instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementRequestBuilder) {
    m := &RoleManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementRequestBuilder instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleManagement
func (m *RoleManagementRequestBuilder) CreateGetRequestInformation(options *RoleManagementRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update roleManagement
func (m *RoleManagementRequestBuilder) CreatePatchRequestInformation(options *RoleManagementRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeviceManagement the deviceManagement property
func (m *RoleManagementRequestBuilder) DeviceManagement()(*i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.DeviceManagementRequestBuilder) {
    return i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Directory the directory property
func (m *RoleManagementRequestBuilder) Directory()(*i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.DirectoryRequestBuilder) {
    return i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntitlementManagement the entitlementManagement property
func (m *RoleManagementRequestBuilder) EntitlementManagement()(*i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.EntitlementManagementRequestBuilder) {
    return i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleManagement
func (m *RoleManagementRequestBuilder) Get(options *RoleManagementRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleManagementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementable), nil
}
// Patch update roleManagement
func (m *RoleManagementRequestBuilder) Patch(options *RoleManagementRequestBuilderPatchOptions)(error) {
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
