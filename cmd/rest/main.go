package main

import (
	"auth-microservice/domain"
	"auth-microservice/infra"
	"auth-microservice/rest"
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cf := loadConfig()
	db := mustConnectDB(cf)

	ps, err := infra.NewPostgresStorage(db)
	if err != nil {
		log.Fatalf("Failed to create Postgres: %v\n", err)
	}

	bh := infra.BcryptHasher{}

	authSrv := domain.NewAuthService(ps, bh)
	credentialSrv := domain.NewAdminService(ps, bh, authSrv)

	mustCreateDefaultCredential(cf, ps, credentialSrv)

	credentialH := rest.NewAdminHandler(credentialSrv)
	authH := rest.NewAuthHandler(authSrv)
	authM := rest.NewAuthMiddleware(authSrv)

	e := echo.New()

	api := e.Group("/api/v1")

	api.POST("/login", authH.HandleLogin)
	api.POST("/logout", authH.HandleLogout)

	dash := api.Group("/dashboard")
	dash.Use(authM)

	dash.GET("/admins", credentialH.HandleGetAdmins)
	dash.POST("/admins", credentialH.HandleSaveAdmin)
	dash.DELETE("/admins/:id", credentialH.HandleDeleteAdmin)

	e.Logger.Fatal(e.Start(":9876"))
}

func mustConnectDB(cf *config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		`host=%s
		 user=%s
		 password=%s
		 dbname=%s
		 port=%s
		 sslmode=%s
		 TimeZone=%s`,
		cf.PostgresHost,
		cf.PostgresUser,
		cf.PostgresPassword,
		cf.PostgresDBName,
		cf.PostgresPort,
		cf.PostgresSSLMode,
		cf.PostgresTimezone,
	)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func mustCreateDefaultCredential(cf *config, s domain.Storage, as domain.AdminService) {
	ctx := context.Background()
	_, err := s.FindByName(ctx, cf.DefaultAdminEmail)
	if err == nil {
		log.Println("Skipping default core creation")
		return
	}

	if err = as.SaveAdmin(ctx, domain.SaveParams{
		FirstName: cf.DefaultAdminFirstName,
		LastName:  cf.DefaultAdminLastName,
		Email:     cf.DefaultAdminEmail,
		Password:  cf.DefaultCredentialPassword,
	}); err != nil {
		log.Fatalf("Failed to create default credential: %v", err)
	} else {
		log.Println("Default credential created")
	}
}
