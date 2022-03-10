package members

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
    i079094e832db2fdc3708f4c93306944a98d15be1db5ce29dbf96023069c7e745 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/enrollassets"
    i0e3ec33249cfa006a9d1754f0677f112f76990d7829004b50cb28caa94546d94 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/unenrollassetsbyid"
    i10ab5f35138e0963fa62739ea774a4cfd4a425f8ad2e62dde35ac49854f903c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/count"
    i4c88dc2f576d570f1868d24fb9515df3d9593c50de671743df58e4fb17d47336 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/enrollassetsbyid"
    ib620cf4d2dcbe5b56d595a4921e6f8f470341999bb89c2a3526a70d51d228519 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/unenrollassets"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MembersRequestBuilderGetOptions options for Get
type MembersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MembersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MembersRequestBuilderGetQueryParameters specifies the assets to include in the audience.
type MembersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// MembersRequestBuilderPostOptions options for Post
type MembersRequestBuilderPostOptions struct {
    // 
    Body ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience/members{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MembersRequestBuilder) Count()(*i10ab5f35138e0963fa62739ea774a4cfd4a425f8ad2e62dde35ac49854f903c2.CountRequestBuilder) {
    return i10ab5f35138e0963fa62739ea774a4cfd4a425f8ad2e62dde35ac49854f903c2.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation specifies the assets to include in the audience.
func (m *MembersRequestBuilder) CreateGetRequestInformation(options *MembersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to members for admin
func (m *MembersRequestBuilder) CreatePostRequestInformation(options *MembersRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MembersRequestBuilder) EnrollAssets()(*i079094e832db2fdc3708f4c93306944a98d15be1db5ce29dbf96023069c7e745.EnrollAssetsRequestBuilder) {
    return i079094e832db2fdc3708f4c93306944a98d15be1db5ce29dbf96023069c7e745.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MembersRequestBuilder) EnrollAssetsById()(*i4c88dc2f576d570f1868d24fb9515df3d9593c50de671743df58e4fb17d47336.EnrollAssetsByIdRequestBuilder) {
    return i4c88dc2f576d570f1868d24fb9515df3d9593c50de671743df58e4fb17d47336.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get specifies the assets to include in the audience.
func (m *MembersRequestBuilder) Get(options *MembersRequestBuilderGetOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateUpdatableAssetCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetCollectionResponseable), nil
}
// Post create new navigation property to members for admin
func (m *MembersRequestBuilder) Post(options *MembersRequestBuilderPostOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateUpdatableAssetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.UpdatableAssetable), nil
}
func (m *MembersRequestBuilder) UnenrollAssets()(*ib620cf4d2dcbe5b56d595a4921e6f8f470341999bb89c2a3526a70d51d228519.UnenrollAssetsRequestBuilder) {
    return ib620cf4d2dcbe5b56d595a4921e6f8f470341999bb89c2a3526a70d51d228519.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MembersRequestBuilder) UnenrollAssetsById()(*i0e3ec33249cfa006a9d1754f0677f112f76990d7829004b50cb28caa94546d94.UnenrollAssetsByIdRequestBuilder) {
    return i0e3ec33249cfa006a9d1754f0677f112f76990d7829004b50cb28caa94546d94.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
