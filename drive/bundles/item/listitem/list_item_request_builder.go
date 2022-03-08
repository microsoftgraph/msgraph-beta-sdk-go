package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0f7bf3dcd8a9180e95b803787fcdc127e979d88fc093bec64d6fe8a83146b799 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/versions"
    i47fea86f01c39cac46033d30e421d32c26a3d92ee0a170d3f13e5efa35c842a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i565fda6b26cb91861bc2c0fe41719ac2b0a12143ca47d6ba2f61ba23b636e4ec "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/analytics"
    i9a00f8221b70503dbdde71402e8e1592a44b402c98b33892d0435302c3cff5bb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/driveitem"
    ia25425c64f2072c81eb52d414968b6de79b1cabb599cf2cf294d69798f8993ef "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/fields"
    ia5cd7969e1fa259d40af09c1e8a22903cd3a3f0c7bc2dcd2bb0ad491fbd5742e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/createlink"
    icee52794c7989b66d152e09d8029577855b45a599efea223476b5fe6b8be95d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/activities"
    i970b73cf6ce47853a2d90e46ce94ab800487a2ad8244bdd3824c633449652ab4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/activities/item"
    ide762eb9e82329af42eea3267f0526fdb87015c3a360324d2d200ea0efaa0328 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem/versions/item"
)

// ListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*icee52794c7989b66d152e09d8029577855b45a599efea223476b5fe6b8be95d7.ActivitiesRequestBuilder) {
    return icee52794c7989b66d152e09d8029577855b45a599efea223476b5fe6b8be95d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i970b73cf6ce47853a2d90e46ce94ab800487a2ad8244bdd3824c633449652ab4.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i970b73cf6ce47853a2d90e46ce94ab800487a2ad8244bdd3824c633449652ab4.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i565fda6b26cb91861bc2c0fe41719ac2b0a12143ca47d6ba2f61ba23b636e4ec.AnalyticsRequestBuilder) {
    return i565fda6b26cb91861bc2c0fe41719ac2b0a12143ca47d6ba2f61ba23b636e4ec.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/bundles/{driveItem_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property listItem for drive
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemRequestBuilder) CreateLink()(*ia5cd7969e1fa259d40af09c1e8a22903cd3a3f0c7bc2dcd2bb0ad491fbd5742e.CreateLinkRequestBuilder) {
    return ia5cd7969e1fa259d40af09c1e8a22903cd3a3f0c7bc2dcd2bb0ad491fbd5742e.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in drive
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete delete navigation property listItem for drive
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i9a00f8221b70503dbdde71402e8e1592a44b402c98b33892d0435302c3cff5bb.DriveItemRequestBuilder) {
    return i9a00f8221b70503dbdde71402e8e1592a44b402c98b33892d0435302c3cff5bb.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*ia25425c64f2072c81eb52d414968b6de79b1cabb599cf2cf294d69798f8993ef.FieldsRequestBuilder) {
    return ia25425c64f2072c81eb52d414968b6de79b1cabb599cf2cf294d69798f8993ef.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*i47fea86f01c39cac46033d30e421d32c26a3d92ee0a170d3f13e5efa35c842a3.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i47fea86f01c39cac46033d30e421d32c26a3d92ee0a170d3f13e5efa35c842a3.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property listItem in drive
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*i0f7bf3dcd8a9180e95b803787fcdc127e979d88fc093bec64d6fe8a83146b799.VersionsRequestBuilder) {
    return i0f7bf3dcd8a9180e95b803787fcdc127e979d88fc093bec64d6fe8a83146b799.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*ide762eb9e82329af42eea3267f0526fdb87015c3a360324d2d200ea0efaa0328.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ide762eb9e82329af42eea3267f0526fdb87015c3a360324d2d200ea0efaa0328.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
