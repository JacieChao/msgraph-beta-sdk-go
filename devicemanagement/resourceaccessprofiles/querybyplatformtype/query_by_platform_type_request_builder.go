package querybyplatformtype

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// queryByPlatformTypeRequestBuilder builds and executes requests for operations under \deviceManagement\resourceAccessProfiles\microsoft.graph.queryByPlatformType
type QueryByPlatformTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// QueryByPlatformTypeRequestBuilderPostOptions options for Post
type QueryByPlatformTypeRequestBuilderPostOptions struct {
    // 
    Body *QueryByPlatformTypeRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewQueryByPlatformTypeRequestBuilderInternal instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewQueryByPlatformTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*QueryByPlatformTypeRequestBuilder) {
    m := &QueryByPlatformTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/microsoft.graph.queryByPlatformType";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewQueryByPlatformTypeRequestBuilder instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewQueryByPlatformTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*QueryByPlatformTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewQueryByPlatformTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) CreatePostRequestInformation(options *QueryByPlatformTypeRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) Post(options *QueryByPlatformTypeRequestBuilderPostOptions)([]QueryByPlatformType, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewQueryByPlatformType() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]QueryByPlatformType, len(res))
    for i, v := range res {
        val[i] = *(v.(*QueryByPlatformType))
    }
    return val, nil
}
