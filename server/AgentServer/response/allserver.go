package response

import (
	"KServer/manage"
	"KServer/proto"
	"fmt"
)

type AllServerResponse struct {
	IManage manage.IManage
}

func NewAllServerResponse(m manage.IManage) *AllServerResponse {
	return &AllServerResponse{IManage: m}
}

func (s *AllServerResponse) ResponseAllServer(data proto.IDataPack) {
	//s.IManage.Message().DataPack().UnPack(req.GetData().Bytes())

	fmt.Println("服务器全体收到消息", s.IManage.Message().DataPack().GetData().String())

	//switch s.IManage.DataPack().GetMsgId() {

}

// 移除客户端
func (s *AllServerResponse) RemoveClient(data proto.IDataPack) {
	c := s.IManage.Socket().Client().GetClient(data.GetClientId())
	if c == nil {
		fmt.Println("客户端不在此服务器")
		return
	}

	err := c.Send(data.GetRawData())
	if err != nil {
		fmt.Println("客户端回调消息失败")
	}
	c.Stop()

}
