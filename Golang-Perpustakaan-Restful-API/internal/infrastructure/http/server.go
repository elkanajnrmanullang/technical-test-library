package http

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/auth"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/peminjaman"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/denda"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(f *fiber.App,
	userCon *user.Controller,
	bukuCon *buku.Controller,
	authCon *auth.Controller,
	peminjamanCon *peminjaman.Controller,
	dendaCon *denda.Controller,
	jwtAuth *middleware.JWTAuthMiddleware) {

	route := f.Group("/api")
	v1 := route.Group("/v1")

	// Authentication route
	v1.Post("/login", authCon.Login)

	// Public book routes
	publicBukuRoutes := v1.Group("/buku")
	publicBukuRoutes.Get("/", bukuCon.GetAllBuku)
	publicBukuRoutes.Get("/:id", bukuCon.GetBukuByID)

	// Protected routes (example)
	adminRoutes := v1.Group("/admin")
	adminRoutes.Use(jwtAuth.Protected()) // Apply JWT protection to admin routes

	// pegawai route
	pegawaiRoutes := adminRoutes.Group("/pegawai")
	pegawaiRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "sukses",
		})
	})

	//end pegawai area

	//buku area
	bukuRoute := adminRoutes.Group("/buku") // Apply JWT protection to buku routes

	//jenis buku route
	jenBuk := bukuRoute.Group("/jenbuk")
	jenBuk.Get("/", bukuCon.GetAllJenisBuku)
	jenBuk.Get("/:id", bukuCon.GetJenisBukuById)
	jenBuk.Post("create", bukuCon.CreateJenisBuku)
	jenBuk.Put("update", bukuCon.UpdateJenisBuku)
	jenBuk.Delete("delete", bukuCon.DeleteJenisBuku)
	// end jenis buku

	// penerbit buku area
	penbuk := bukuRoute.Group("/penbuk")
	penbuk.Get("/", bukuCon.GetAllPenerbitBuku)
	penbuk.Get("/:id", bukuCon.GetPenerbitBukuById)
	penbuk.Post("create", bukuCon.CreatePenerbitBuku)
	penbuk.Put("update", bukuCon.UpdatePenerbitBuku)
	penbuk.Delete("delete", bukuCon.DeletePenerbitBuku)
	// end penerbit buku

	// penulis buku area
	author := bukuRoute.Group("/author")
	author.Get("/", bukuCon.GetAllPenulisBuku)
	author.Get("/:id", bukuCon.GetPenulisBukuById)
	author.Post("create", bukuCon.CreatePenulisBuku)
	author.Put("update", bukuCon.UpdatePenulisBuku)
	author.Delete("delete", bukuCon.DeletePenulisBuku)
	// end penulis buku area
	//end buku area

	// peminjaman area
	peminjamanRoutes := adminRoutes.Group("/peminjaman")
	peminjamanRoutes.Get("/", peminjamanCon.GetAllPeminjaman)
	peminjamanRoutes.Get("/:id", peminjamanCon.GetPeminjamanByID)
	peminjamanRoutes.Get("/detail/:id", peminjamanCon.GetDetailPeminjaman)
	peminjamanRoutes.Post("create", peminjamanCon.CreatePeminjaman)
	peminjamanRoutes.Put("update", peminjamanCon.UpdatePeminjaman)
	peminjamanRoutes.Delete("delete", peminjamanCon.DeletePeminjaman)
	// end peminjaman area

	// denda area
	dendaRoutes := adminRoutes.Group("/denda")
	dendaRoutes.Get("/", dendaCon.GetAllDenda)
	dendaRoutes.Get("/:id", dendaCon.GetDendaByID)
	dendaRoutes.Post("create", dendaCon.CreateDenda)
	dendaRoutes.Put("update", dendaCon.UpdateDenda)
	dendaRoutes.Delete("delete", dendaCon.DeleteDenda)
	// end denda area
}
