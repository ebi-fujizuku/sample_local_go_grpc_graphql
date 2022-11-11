package client

import (
	"context"
	"fmt"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/common"
	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/pb"
	"google.golang.org/grpc"
)

type Client struct{
	Conn   *grpc.ClientConn
	Client pb.ArticleServiceClient
}

func (c *Client)Create() {
	common.PrintStart("")

	name := "fujito"

	req := &pb.CreateArticleRequest{
		ArticleInput: name,
	}
	res, err := c.Client.CreateArticle(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetArticle())
	}

	common.PrintEnd("")
}