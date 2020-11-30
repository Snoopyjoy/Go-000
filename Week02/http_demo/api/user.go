package api

import (
	"context"
	"errors"
	"fmt"
	"http_demo/dao"
	"http_demo/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler() Handler {
	return &UserHandler{}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path[:6] == "/user/" {
		idstr := req.URL.Path[6:]
		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			// 接口层处理err
			fmt.Printf("invalid user id %s\n", idstr)
			ClientError(w, []byte("invalid user id"))
			return
		}
		ctx := context.Background()
		userData, err := h.service.GetUserByID(ctx, id)
		if err != nil {
			// 接口层处理err
			// 日志打印
			fmt.Printf("find user failed %+v\n", err)

			if errors.Is(err, dao.ErrNotFound) {
				// 判断是否为记录不存在
				// 根据dao定义的sentinel err判断
				NotFound(w, []byte("user not exists"))
				return
			} else {
				ServeError(w, []byte("service internal error"))
				return
			}
		}

		OK(w, []byte(fmt.Sprintf("hello %+v", userData)))
		return
	}
	PageNotFound(w)
}
