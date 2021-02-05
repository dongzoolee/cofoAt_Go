module cofoAt.go

go 1.15

replace cf/getData => ../getData

replace cf/updateUser => ../functional

require (
	cf/getData v0.0.0-00010101000000-000000000000
	cf/updateUser v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
)
