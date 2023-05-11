/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"github.com/permitio/permit-golang/pkg/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// TenantsApiService TenantsApi service
type TenantsApiService service

type ApiCreateTenantRequest struct {
	ctx          context.Context
	ApiService   *TenantsApiService
	projId       string
	envId        string
	tenantCreate *models.TenantCreate
}

func (r ApiCreateTenantRequest) TenantCreate(tenantCreate models.TenantCreate) ApiCreateTenantRequest {
	r.tenantCreate = &tenantCreate
	return r
}

func (r ApiCreateTenantRequest) Execute() (*models.TenantRead, *http.Response, error) {
	return r.ApiService.CreateTenantExecute(r)
}

/*
CreateTenant Create Tenant

Creates a new tenant inside the Permit.io system.

If the tenant is already created: will return 200 instead of 201,
and will return the existing tenant object in the response body.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@return ApiCreateTenantRequest
*/
func (a *TenantsApiService) CreateTenant(ctx context.Context, projId string, envId string) ApiCreateTenantRequest {
	return ApiCreateTenantRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//
//	@return TenantRead
func (a *TenantsApiService) CreateTenantExecute(r ApiCreateTenantRequest) (*models.TenantRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TenantRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.CreateTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantCreate == nil {
		return localVarReturnValue, nil, reportError("tenantCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tenantCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteTenantRequest struct {
	ctx        context.Context
	ApiService *TenantsApiService
	projId     string
	envId      string
	tenantId   string
}

func (r ApiDeleteTenantRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTenantExecute(r)
}

/*
DeleteTenant Delete Tenant

Deletes the tenant and all its related data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@param tenantId Either the unique id of the tenant, or the URL-friendly key of the tenant (i.e: the \"slug\").
	@return ApiDeleteTenantRequest
*/
func (a *TenantsApiService) DeleteTenant(ctx context.Context, projId string, envId string, tenantId string) ApiDeleteTenantRequest {
	return ApiDeleteTenantRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
		tenantId:   tenantId,
	}
}

// Execute executes the request
func (a *TenantsApiService) DeleteTenantExecute(r ApiDeleteTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.DeleteTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants/{tenant_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteTenantUserRequest struct {
	ctx        context.Context
	ApiService *TenantsApiService
	projId     string
	envId      string
	tenantId   string
	userId     string
}

func (r ApiDeleteTenantUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTenantUserExecute(r)
}

/*
DeleteTenantUser Delete Tenant User

Deletes a user under a tenant.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@param tenantId Either the unique id of the tenant, or the URL-friendly key of the tenant (i.e: the \"slug\").
	@param userId Either the unique id of the user, or the URL-friendly key of the user (i.e: the \"slug\").
	@return ApiDeleteTenantUserRequest
*/
func (a *TenantsApiService) DeleteTenantUser(ctx context.Context, projId string, envId string, tenantId string, userId string) ApiDeleteTenantUserRequest {
	return ApiDeleteTenantUserRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
		tenantId:   tenantId,
		userId:     userId,
	}
}

// Execute executes the request
func (a *TenantsApiService) DeleteTenantUserExecute(r ApiDeleteTenantUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.DeleteTenantUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants/{tenant_id}/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTenantRequest struct {
	ctx        context.Context
	ApiService *TenantsApiService
	projId     string
	envId      string
	tenantId   string
}

func (r ApiGetTenantRequest) Execute() (*models.TenantRead, *http.Response, error) {
	return r.ApiService.GetTenantExecute(r)
}

/*
GetTenant Get Tenant

Gets a tenant, if such tenant exists. Otherwise returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@param tenantId Either the unique id of the tenant, or the URL-friendly key of the tenant (i.e: the \"slug\").
	@return ApiGetTenantRequest
*/
func (a *TenantsApiService) GetTenant(ctx context.Context, projId string, envId string, tenantId string) ApiGetTenantRequest {
	return ApiGetTenantRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
		tenantId:   tenantId,
	}
}

// Execute executes the request
//
//	@return TenantRead
func (a *TenantsApiService) GetTenantExecute(r ApiGetTenantRequest) (*models.TenantRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TenantRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.GetTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants/{tenant_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListTenantUsersRequest struct {
	ctx        context.Context
	ApiService *TenantsApiService
	projId     string
	tenantId   string
	envId      string
	search     *string
	page       *int32
	perPage    *int32
}

// Text search for the email field
func (r ApiListTenantUsersRequest) Search(search string) ApiListTenantUsersRequest {
	r.search = &search
	return r
}

// Page number of the results to fetch, starting at 1.
func (r ApiListTenantUsersRequest) Page(page int32) ApiListTenantUsersRequest {
	r.page = &page
	return r
}

// The number of results per page (max 100).
func (r ApiListTenantUsersRequest) PerPage(perPage int32) ApiListTenantUsersRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListTenantUsersRequest) Execute() (*models.PaginatedResultUserRead, *http.Response, error) {
	return r.ApiService.ListTenantUsersExecute(r)
}

/*
ListTenantUsers List Tenant Users

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param tenantId Either the unique id of the tenant, or the URL-friendly key of the tenant (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@return ApiListTenantUsersRequest
*/
func (a *TenantsApiService) ListTenantUsers(ctx context.Context, projId string, tenantId string, envId string) ApiListTenantUsersRequest {
	return ApiListTenantUsersRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		tenantId:   tenantId,
		envId:      envId,
	}
}

// Execute executes the request
//
//	@return PaginatedResultUserRead
func (a *TenantsApiService) ListTenantUsersExecute(r ApiListTenantUsersRequest) (*models.PaginatedResultUserRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PaginatedResultUserRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.ListTenantUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants/{tenant_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListTenantsRequest struct {
	ctx              context.Context
	ApiService       *TenantsApiService
	projId           string
	envId            string
	search           *string
	attributesFilter map[string]interface{}
	page             *int32
	perPage          *int32
}

// Text search for the tenant name or key
func (r ApiListTenantsRequest) Search(search string) ApiListTenantsRequest {
	r.search = &search
	return r
}

// Filter for tenant with specific attributes
func (r ApiListTenantsRequest) AttributeFilter(attributesFilter map[string]interface{}) ApiListTenantsRequest {
	r.attributesFilter = attributesFilter
	return r
}

// Page number of the results to fetch, starting at 1.
func (r ApiListTenantsRequest) Page(page int32) ApiListTenantsRequest {
	r.page = &page
	return r
}

// The number of results per page (max 100).
func (r ApiListTenantsRequest) PerPage(perPage int32) ApiListTenantsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListTenantsRequest) Execute() ([]models.TenantRead, *http.Response, error) {
	return r.ApiService.ListTenantsExecute(r)
}

/*
ListTenants List Tenants

Lists all the tenants defined within an environment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@return ApiListTenantsRequest
*/
func (a *TenantsApiService) ListTenants(ctx context.Context, projId string, envId string) ApiListTenantsRequest {
	return ApiListTenantsRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//
//	@return []TenantRead
func (a *TenantsApiService) ListTenantsExecute(r ApiListTenantsRequest) ([]models.TenantRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.TenantRead
	)
	const attributeFilterPrefix = "attr_"

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.ListTenants")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.attributesFilter != nil {
		for k, v := range r.attributesFilter {
			localVarQueryParams.Add(attributeFilterPrefix+k, parameterToString(v, ""))
		}
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateTenantRequest struct {
	ctx          context.Context
	ApiService   *TenantsApiService
	projId       string
	envId        string
	tenantId     string
	tenantUpdate *models.TenantUpdate
}

func (r ApiUpdateTenantRequest) TenantUpdate(tenantUpdate models.TenantUpdate) ApiUpdateTenantRequest {
	r.tenantUpdate = &tenantUpdate
	return r
}

func (r ApiUpdateTenantRequest) Execute() (*models.TenantRead, *http.Response, error) {
	return r.ApiService.UpdateTenantExecute(r)
}

/*
UpdateTenant Update Tenant

Partially updates the tenant definition.
Fields that will be provided will be completely overwritten.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
	@param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
	@param tenantId Either the unique id of the tenant, or the URL-friendly key of the tenant (i.e: the \"slug\").
	@return ApiUpdateTenantRequest
*/
func (a *TenantsApiService) UpdateTenant(ctx context.Context, projId string, envId string, tenantId string) ApiUpdateTenantRequest {
	return ApiUpdateTenantRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
		tenantId:   tenantId,
	}
}

// Execute executes the request
//
//	@return TenantRead
func (a *TenantsApiService) UpdateTenantExecute(r ApiUpdateTenantRequest) (*models.TenantRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.TenantRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantsApiService.UpdateTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/facts/{proj_id}/{env_id}/tenants/{tenant_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantUpdate == nil {
		return localVarReturnValue, nil, reportError("tenantUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tenantUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
