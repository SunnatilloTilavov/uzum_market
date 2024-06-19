package api

import (
	"net/http"

	_ "backend_course/customer_api_gateway/api/docs" //for swagger
	"backend_course/customer_api_gateway/api/handler"
	"backend_course/customer_api_gateway/config"
	"backend_course/customer_api_gateway/pkg/grpc_client"
	"backend_course/customer_api_gateway/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Bu xavfsiz emas, kerakli domenlarni qo'shishingiz kerak
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	r.POST("/create/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetCategoryByID)
	r.GET("/getlistCategory", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)
	r.GET("/getcategorywith/:id", handler.GetCategoryWithProductId)
	
	r.POST("/create/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetProductByID)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)
	r.GET("/getlistProduct", handler.GetListProduct)

	r.POST("/create/product_review", handler.CreateProductReview)
	r.GET("/product_review/:id", handler.GetProductReviewByID)
	r.PUT("/product_review/:id", handler.UpdateProductReview)
	r.DELETE("/product_review/:id", handler.DeleteProductReview)
	r.GET("/product_reviews/product/:product_id", handler.GetProductReviewsByProductID)
	r.GET("/product_reviews/customer/:customer_id", handler.GetProductReviewsByCustomerID)


	// Swagger endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
