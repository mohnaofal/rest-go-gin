package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohnaofal/rest-go-gin/app/repository/commands"
	"github.com/mohnaofal/rest-go-gin/app/repository/queries"
	"github.com/mohnaofal/rest-go-gin/app/usecase"
	"github.com/mohnaofal/rest-go-gin/config"
)

type ArticleDelivery struct {
	articleCommandUsecase usecase.ArticleCommandUsecase
	articleQueryUsecase   usecase.ArticleQueryUsecase
}

func NewArticleDelivery(cfg *config.Config) ArticleDelivery {
	articleCommand := commands.NewArticleCommand(cfg)
	articleQuery := queries.NewArticleQuery(cfg)
	articleCommandUsecase := usecase.NewArticleCommandUsecase(articleCommand, articleQuery)
	articleQueryUsecase := usecase.NewArticleQueryUsecase(articleQuery)

	return ArticleDelivery{
		articleCommandUsecase: articleCommandUsecase,
		articleQueryUsecase:   articleQueryUsecase,
	}
}

func (h *ArticleDelivery) Apply(c *gin.RouterGroup) {
	c.POST("", h.Create)
	c.GET("", h.View)
}

func (h *ArticleDelivery) Create(g *gin.Context) {
	g.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success created article",
	})
}

func (h *ArticleDelivery) View(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success retrieve article",
	})
}
