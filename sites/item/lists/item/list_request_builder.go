package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities"
    i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes"
    i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns"
    i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/drive"
    ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items"
    iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions"
    i310254797e5d342cdb1f2d5ff07a899b7c9adeb765730d80690cdbb88a73e453 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item"
    i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions/item"
    i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns/item"
    i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item"
    i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item"
)

type ListRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ListRequestBuilder) Activities()(*i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.ActivitiesRequestBuilder) {
    return i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ActivitiesById(id string)(*i310254797e5d342cdb1f2d5ff07a899b7c9adeb765730d80690cdbb88a73e453.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i310254797e5d342cdb1f2d5ff07a899b7c9adeb765730d80690cdbb88a73e453.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.ColumnsRequestBuilder) {
    return i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ColumnsById(id string)(*i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/lists/{list_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.ContentTypesRequestBuilder) {
    return i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ContentTypesById(id string)(*i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListRequestBuilder) CreateGetRequestInformation(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Drive()(*i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.DriveRequestBuilder) {
    return i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) Get(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewList() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List), nil
}
func (m *ListRequestBuilder) Items()(*ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.ItemsRequestBuilder) {
    return ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ItemsById(id string)(*i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Subscriptions()(*iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.SubscriptionsRequestBuilder) {
    return iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
