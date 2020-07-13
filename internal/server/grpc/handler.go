package grpc

import (
	"context"

	grapi "github.com/jerosa/grapi/internal"
	"github.com/jerosa/grapi/internal/adding"
	"github.com/jerosa/grapi/internal/creating"
	"github.com/jerosa/grapi/internal/listing"
	pb "github.com/jerosa/grapi/proto"
	readlistgrpc "github.com/jerosa/grapi/proto"
)

type readListHandler struct {
	creatingService creating.Service
	addingService   adding.Service
	listingService  listing.Service
}

// NewReadListServer provides WishList gRPC operations
func NewReadListServer(
	cS creating.Service,
	aS adding.Service,
	lS listing.Service,
) pb.ReadListServiceServer {
	return &readListHandler{creatingService: cS, addingService: aS, listingService: lS}
}

func (s readListHandler) Create(ctx context.Context, req *pb.CreateReadListReq) (*pb.CreateReadListResp, error) {
	id, err := s.creatingService.Create(req.ReadList.Name, grapi.Status(req.ReadList.Status))
	if err != nil {
		return nil, err
	}
	return &readlistgrpc.CreateReadListResp{ReadListId: id}, nil
}

func (s readListHandler) Add(ctx context.Context, req *pb.AddArticleReq) (*pb.AddArticleResp, error) {
	id, err := s.addingService.AddArticle(
		req.Article.ReadListId,
		req.Article.Name,
		req.Article.Link,
	)
	if err != nil {
		return nil, err
	}
	return &readlistgrpc.AddArticleResp{ArticleId: id}, nil
}

func (s readListHandler) List(ctx context.Context, req *pb.ListReadListReq) (*pb.ListReadListResp, error) {
	articles, err := s.listingService.ListArticles(req.ReadListId)
	if err != nil {
		return nil, err
	}
	return &readlistgrpc.ListReadListResp{Articles: mapSliceOfItems(articles)}, nil
}

func mapSliceOfItems(domainItems []grapi.Article) (grpcItems []*readlistgrpc.Article) {
	for _, i := range domainItems {
		grpcItems = append(grpcItems, &readlistgrpc.Article{
			Id:         i.ID,
			ReadListId: i.ReadListID,
			Name:       i.Name,
			Link:       i.Link,
		})
	}
	return
}
