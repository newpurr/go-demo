package inject

import (
	"fmt"
	"github.com/facebookgo/inject"
	"testing"
)

type Conf struct {
}

func loadConf() *Conf {
	return &Conf{}
}

type DB struct {
}

func connectDB() *DB {
	return &DB{}
}

type UserController struct {
	UserService *UserService `inject:""`
	Conf        *Conf        `inject:""`
}

type PostController struct {
	UserService *UserService `inject:""`
	PostService *PostService `inject:""`
	Conf        *Conf        `inject:""`
}

type UserService struct {
	Db   *DB   `inject:""`
	Conf *Conf `inject:""`
}

type PostService struct {
	Db  *DB `inject:""`
	Db2 *DB `inject:"db_named"`
}

type Server struct {
	UserApi *UserController `inject:""`
	PostApi *PostController `inject:""`
}

func (s Server) Run() {
	fmt.Println("run ...")
}

// https://blog.csdn.net/weixin_39603217/article/details/110805219
func TestRun(t *testing.T) {
	conf := loadConf() // *Conf
	db := connectDB()  // *DB
	db2 := connectDB() // *DB

	server := Server{}

	graph := inject.Graph{}

	if err := graph.Provide(
		&inject.Object{
			Value: &server,
		},
		&inject.Object{
			Value: conf,
		},
		&inject.Object{
			Value: db,
		},
		&inject.Object{
			Value: db2,
			Name:  "db_named",
		},
	); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}

	server.Run()
}
