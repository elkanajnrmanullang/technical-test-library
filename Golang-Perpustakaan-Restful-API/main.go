package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http"

	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/user"

	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/buku"
	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/buku"

	authController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/auth"
	authService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/auth"
	authRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/auth"

	peminjamanController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/peminjaman"
	peminjamanService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/peminjaman"
	peminjamanRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/peminjaman"

	dendaController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/denda"
	dendaService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/denda"
	dendaRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/repository/mysql/denda"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
	conf "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	// "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/database"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/middleware"
)

func main() {
	db := conf.MySQLConn()
	// database.Migrate()
	// database.Seeder(db)

	// Init JWT Config
	jwtConfig := conf.GetJwtConfig()
	jwtUtility := utility.NewJWTUtility(jwtConfig)

	// User module
	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	// Buku module
	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	// Auth module
	authRepo := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, jwtUtility)
	authController := authController.NewAuthController(authService)

	// Peminjaman module
	peminjamanRepo := peminjamanRepository.NewPeminjamanRepository(db)
	peminjamanServices := peminjamanService.NewPeminjamanService(peminjamanRepo)
	peminjamanCon := peminjamanController.NewPeminjamanController(peminjamanServices)

	// Denda module
	dendaRepo := dendaRepository.NewDendaRepository(db)
	dendaServices := dendaService.NewDendaService(dendaRepo)
	dendaCon := dendaController.NewDendaController(dendaServices)

	// Middleware
	jwtAuthMiddleware := middleware.NewJWTAuthMiddleware(jwtUtility, jwtConfig)


	appConfig := conf.ServerTimeOut()
	app := fiber.New(appConfig)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	http.RegisterPath(
		app,
		userCon,
		bukuCon,
		authController,
		peminjamanCon,
		dendaCon,
		jwtAuthMiddleware,
	)

	app.Listen(":8001")

}
