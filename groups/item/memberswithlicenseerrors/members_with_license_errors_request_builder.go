package memberswithlicenseerrors

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i266c609a4a0b2ea717fcb9f2b5bd204123ad620a4c9d6ac6d46ac21cb41e5d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/count"
    i4fe7da1c99253943ddac41ce7c2dc39031462186f7e483abbd38ffe0ac35bbe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/serviceprincipal"
    i9c2e1a0bee7b2938a350f62e4e9e6d9b5c95bc5c7b1a22f1513ebd7b64441fe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/user"
    ia86add01b2f0bd76252b1b785e6434aad18c4d0c88e021524ab092203ca55552 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/orgcontact"
    ic2076f0c68781bab805ae88e1fb6575d94374bc5b6d3360f211d204169415808 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/application"
    id463127ae515fce34783aeeb345c7cefa9fef6e27311c3d35fc0e7f89cbff7fb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/group"
    if6158cfb9c5935af0795851e895bbba23c924418ac76059a4b88e0b52d9e2831 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/device"
)

// MembersWithLicenseErrorsRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type MembersWithLicenseErrorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersWithLicenseErrorsRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type MembersWithLicenseErrorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MembersWithLicenseErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembersWithLicenseErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MembersWithLicenseErrorsRequestBuilderGetQueryParameters
}
// Application the application property
func (m *MembersWithLicenseErrorsRequestBuilder) Application()(*ic2076f0c68781bab805ae88e1fb6575d94374bc5b6d3360f211d204169415808.ApplicationRequestBuilder) {
    return ic2076f0c68781bab805ae88e1fb6575d94374bc5b6d3360f211d204169415808.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMembersWithLicenseErrorsRequestBuilderInternal instantiates a new MembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewMembersWithLicenseErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersWithLicenseErrorsRequestBuilder) {
    m := &MembersWithLicenseErrorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMembersWithLicenseErrorsRequestBuilder instantiates a new MembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewMembersWithLicenseErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersWithLicenseErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembersWithLicenseErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *MembersWithLicenseErrorsRequestBuilder) Count()(*i266c609a4a0b2ea717fcb9f2b5bd204123ad620a4c9d6ac6d46ac21cb41e5d12.CountRequestBuilder) {
    return i266c609a4a0b2ea717fcb9f2b5bd204123ad620a4c9d6ac6d46ac21cb41e5d12.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
func (m *MembersWithLicenseErrorsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device the device property
func (m *MembersWithLicenseErrorsRequestBuilder) Device()(*if6158cfb9c5935af0795851e895bbba23c924418ac76059a4b88e0b52d9e2831.DeviceRequestBuilder) {
    return if6158cfb9c5935af0795851e895bbba23c924418ac76059a4b88e0b52d9e2831.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
func (m *MembersWithLicenseErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *MembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *MembersWithLicenseErrorsRequestBuilder) Group()(*id463127ae515fce34783aeeb345c7cefa9fef6e27311c3d35fc0e7f89cbff7fb.GroupRequestBuilder) {
    return id463127ae515fce34783aeeb345c7cefa9fef6e27311c3d35fc0e7f89cbff7fb.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MembersWithLicenseErrorsRequestBuilder) OrgContact()(*ia86add01b2f0bd76252b1b785e6434aad18c4d0c88e021524ab092203ca55552.OrgContactRequestBuilder) {
    return ia86add01b2f0bd76252b1b785e6434aad18c4d0c88e021524ab092203ca55552.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MembersWithLicenseErrorsRequestBuilder) ServicePrincipal()(*i4fe7da1c99253943ddac41ce7c2dc39031462186f7e483abbd38ffe0ac35bbe8.ServicePrincipalRequestBuilder) {
    return i4fe7da1c99253943ddac41ce7c2dc39031462186f7e483abbd38ffe0ac35bbe8.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MembersWithLicenseErrorsRequestBuilder) User()(*i9c2e1a0bee7b2938a350f62e4e9e6d9b5c95bc5c7b1a22f1513ebd7b64441fe7.UserRequestBuilder) {
    return i9c2e1a0bee7b2938a350f62e4e9e6d9b5c95bc5c7b1a22f1513ebd7b64441fe7.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
