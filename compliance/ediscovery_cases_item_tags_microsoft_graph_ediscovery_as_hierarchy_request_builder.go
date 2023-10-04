package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder provides operations to call the asHierarchy method.
type EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters invoke function asHierarchy
type EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryAsHierarchyRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    m := &EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/tags/microsoft.graph.ediscovery.asHierarchy(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder instantiates a new MicrosoftGraphEdiscoveryAsHierarchyRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function asHierarchy
// Deprecated: This method is obsolete. Use GetAsAsHierarchyGetResponse instead.
func (m *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyResponseable), nil
}
// GetAsAsHierarchyGetResponse invoke function asHierarchy
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) GetAsAsHierarchyGetResponse(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyAsHierarchyGetResponseable), nil
}
// ToGetRequestInformation invoke function asHierarchy
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    return NewEdiscoveryCasesItemTagsMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
