package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/app/repository/commands"
	"github.com/mohnaofal/rest-go-gin/app/repository/queries"
	"github.com/mohnaofal/rest-go-gin/app/usecase"
	"github.com/mohnaofal/rest-go-gin/app/usecase/request"
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
	form := new(models.Article)
	if err := g.ShouldBind(form); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	result, err := h.articleCommandUsecase.Create(g.Request.Context(), form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success created article",
		"data":    result,
	})
}

func (h *ArticleDelivery) View(g *gin.Context) {
	form := new(request.ViewArticle)
	if err := g.ShouldBind(form); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	result, err := h.articleQueryUsecase.View(g.Request.Context(), form)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success retrieve article",
		"data":    result,
	})
}
