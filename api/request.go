package main

type UserCreateReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserUpdateReq struct {
	Name *string `json:"name,omitempty"`
	Age  *int    `json:"age,omitempty"`
}

type ListRequest struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}

type ListUserResponse struct {
	Users         []*User `json:"users"`
	NextPageToken string  `json:"next_page_token"`
}

type ListResponse[T any] struct {
	Items         []T    `json:"items"`
	NextPageToken string `json:"next_page_token"`
}
