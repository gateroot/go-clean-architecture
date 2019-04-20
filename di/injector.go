package di

import (
	"awesomeProject1/trace"
	"awesomeProject1/user"
	"awesomeProject1/user/delivery"
	"awesomeProject1/user/delivery/echo"
	"awesomeProject1/user/repository"
	"awesomeProject1/user/repository/sqlite3"
	"database/sql"
	"os"
)

func InjectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	return db
}

func InjectTracer() trace.Tracer {
	return trace.New(os.Stdout)
}

func InjectUserRepository() repository.UserRepository {
	return sqlite3.NewUserRepository(InjectDB())
}

func InjectUserUsecase() user.UserUsecase {
	return user.NewUserUsecase(InjectUserRepository(), InjectTracer())
}

func InjectUserDelivery() delivery.UserDelivery {
	return echo.NewUserDelivery(InjectUserUsecase())
}
