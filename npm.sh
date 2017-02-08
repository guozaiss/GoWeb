echo "正在下载go lib"

echo "github.com/astaxie/beego"
go get github.com/astaxie/beego
cd $GOPATH
cd src/github.com/astaxie/beego
git checkout v1.7.2

echo "github.com/gorilla/mux"
go get github.com/gorilla/mux
cd $GOPATH
cd src/github.com/gorilla/mux
git checkout v1.3.0

echo "github.com/go-sql-driver/mysql"
go get github.com/go-sql-driver/mysql
cd $GOPATH
cd src/github.com/go-sql-driver/mysql
git checkout v1.3

echo "github.com/go-xorm/xorm"
go get github.com/go-xorm/xorm
cd $GOPATH
cd src/github.com/go-xorm/xorm
git checkout v0.5.8

echo "下载完毕"