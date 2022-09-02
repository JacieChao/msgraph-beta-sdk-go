package getrolescopetagsbyresourcewithresource

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetRoleScopeTagsByResourceWithResourceRequestBuilder provides operations to call the getRoleScopeTagsByResource method.
type GetRoleScopeTagsByResourceWithResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, resource *string)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    m := &GetRoleScopeTagsByResourceWithResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getRoleScopeTagsByResource(resource='{resource}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if resource != nil {
        urlTplParams["resource"] = *resource
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRoleScopeTagsByResourceWithResourceRequestBuilder instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByResourceWithResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getRoleScopeTagsByResource
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getRoleScopeTagsByResource
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getRoleScopeTagsByResource
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *GetRoleScopeTagsByResourceWithResourceRequestBuilderGetRequestConfiguration)(GetRoleScopeTagsByResourceWithResourceResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetRoleScopeTagsByResourceWithResourceResponseable), nil
}
