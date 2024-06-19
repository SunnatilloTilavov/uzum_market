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


	// orders
	r.POST("/api/v1/order", handler.CreateOrder)
	r.GET("/api/v1/order/:id", handler.GetByIdOrder)
	r.PUT("/api/v1/order", handler.UpdateOrder)
	r.DELETE("/api/v1/order/:id", handler.DeleteOrder)
	r.GET("/api/v1/orders", handler.GetAllOrder)

	// order-products
	r.POST("/api/v1/orderProduct", handler.CreateOrderProduct)
	r.GET("/api/v1/orderProduct/:id", handler.GetByIdOrderProduct)
	r.PUT("/api/v1/orderProduct", handler.UpdateOrderProduct)
	r.DELETE("/api/v1/orderProduct/:id", handler.DeleteOrderProduct)
	r.GET("/api/v1/orderProducts", handler.GetAllOrderProducts)

	// order-notes
	r.POST("/api/v1/orderNote", handler.CreateOrderNotes)
	r.GET("/api/v1/orderNote/:id", handler.GetByIdOrderNotes)
	r.PUT("/api/v1/orderNote", handler.UpdateOrderNotes)
	r.DELETE("/api/v1/orderNote/:id", handler.DeleteOrderNotes)
	r.GET("/api/v1/orderNotes", handler.GetAllOrderNotes)

	r.POST("/v1/branch/create", handler.CreateBranch)
	r.POST("/v1/branch/getbyid", handler.GetByID)
	r.POST("/v1/branch/getlist", handler.GetList)
	r.PUT("/v1/branch/updatebranch", handler.Update)
	r.DELETE("/v1/branch/delete", handler.Delete)

	r.POST("/v1/customer/create", handler.CreateCustomer)
	r.POST("/v1/customer/getbyid", handler.GetByIdCustomer)
	r.POST("/v1/customer/getlist", handler.GetListCustomer)
	r.PUT("/v1/customer/update", handler.UpdateCustomer)
	r.DELETE("/v1/customer/delete", handler.DeleteCustomer)

	r.POST("/v1/shop/create", handler.CreateShop)
	r.POST("/v1/shop/getbyid", handler.GetByIdShop)
	r.POST("/v1/shop/getlist", handler.GetListShop)
	r.PUT("/v1/shop/update", handler.UpdateShop)
	r.DELETE("/v1/shop/delete", handler.DeleteShop)

	r.POST("/v1/seller/create", handler.CreateSeller)
	r.POST("/v1/seller/getbyid", handler.GetByIdSeller)
	r.POST("/v1/seller/getlist", handler.GetListSeller)
	r.PUT("/v1/seller/update", handler.UpdateSeller)
	r.DELETE("/v1/seller/delete", handler.DeleteSeller)

	r.POST("/v1/system-user/create", handler.CreateSystemUser)
	r.POST("/v1/system-user/getbyid", handler.GetByIdSystemUser)
	r.POST("/v1/system-user/getlist", handler.GetListSystemUser)
	r.PUT("/v1/system-user/update", handler.UpdateSystemUser)	
	r.DELETE("/v1/system-user/delete", handler.DeleteSystemUser)


	r.POST("/v1/customer/register", handler.RegisterCustomer)
	r.POST("/v1/customer/registercomfirm",handler.RegisterConfirmCustomer)
	r.POST("/v1/customer/login",handler.LoginCustomer)
	r.POST("/v1/customer/loginconfirm",handler.LoginConfirm)
	r.PUT("/v1/customer/resetpassword",handler.ResetPasswordCustomer)
	r.PUT("/v1/customer/resetconfirm",handler.ResetPasswordConfirm)

	r.POST("/v1/seller/register", handler.RegisterSeller)
	r.POST("/v1/seller/registercomfirm",handler.RegisterConfirmSeller)
	r.POST("/v1/seller/login",handler.LoginSeller)
	r.POST("/v1/seller/loginconfirm",handler.SellerLoginConfirm)
	r.PUT("/v1/seller/resetpassword",handler.ResetPasswordSeller)
	r.PUT("/v1/seller/resetconfirm",handler.SellerResetPasswordConfirm)

	r.POST("/v1/systemuser/register", handler.RegisterSystemUser)
	r.POST("/v1/systemuser/registercomfirm",handler.RegisterConfirmSystemUser)
	r.POST("/v1/systemuser/login",handler.LoginSystemUser)
	r.POST("/v1/systemuser/loginconfirm",handler.SystemUserLoginConfirm)
	r.PUT("/v1/systemuser/resetpassword",handler.ResetPasswordSystemUser)
	r.PUT("/v1/systemuser/resetconfirm",handler.SystemUserResetPasswordConfirm)



	// Swagger endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
