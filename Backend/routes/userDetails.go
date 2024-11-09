package routes

import (
	"github.com/adityjoshi/avinya/Backend/controllers"
	"github.com/adityjoshi/avinya/Backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserInfoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/updatePatientInfo/:id", middleware.AuthRequired("Patient", ""), middleware.OtpAuthRequireed, controllers.AddPatientDetails)
	// incomingRoutes.GET("/getPatientId/:id", middleware.AuthUser(), controllers.GetPatientDetails)

}
