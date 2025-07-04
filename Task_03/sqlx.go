package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func func_sqlx01() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	var employees []Employee
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", employees)
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	var maxSalaryEmployee Employee
	err = db.Get(&maxSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", maxSalaryEmployee)
	fmt.Printf("Highest Salary Employee: ID: %d, Name: %s, Department: %s, Salary: %.2f\n",
		maxSalaryEmployee.ID, maxSalaryEmployee.Name, maxSalaryEmployee.Department, maxSalaryEmployee.Salary)
}

func func_gorm01() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	// 获取底层 sql.DB 连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database: %v", err)
	}
	fmt.Println("连接数据库成功")
	defer sqlDB.Close()

	var employees []Employee
	res := db.Where("department = ?", "技术部").Find(&employees)
	if res.Error != nil {
		// 处理错误
		return
	}

	fmt.Println("找到记录数:", res.RowsAffected)

	log.Printf("%+v\n", employees)
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	var maxSalaryEmployee Employee
	res = db.Order("salary DESC").First(&maxSalaryEmployee)
	if res.Error != nil {
		// 处理错误
		return
	}

	fmt.Println("找到记录数:", res.RowsAffected)

	log.Printf("%+v\n", maxSalaryEmployee)
	fmt.Printf("Highest Salary Employee: ID: %d, Name: %s, Department: %s, Salary: %.2f\n",
		maxSalaryEmployee.ID, maxSalaryEmployee.Name, maxSalaryEmployee.Department, maxSalaryEmployee.Salary)
}

type Book struct {
	Id     uint64  `db:"id"`
	Title  string  `db:"Title"`
	Author string  `db:"Author"`
	Price  float64 `db:"Price"`
}

func func_sqlx02() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	var books []Book
	err = db.Select(&books, "select * from books where price > 50")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", books)
	for _, emp := range books {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", emp.Id, emp.Title, emp.Author, emp.Price)
	}

}

func func_gorm02() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	// 获取底层 sql.DB 连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database: %v", err)
	}
	fmt.Println("连接数据库成功")
	defer sqlDB.Close()

	var books []Book
	//res := db.Where("price > ?", 50).Find(&books)
	res := db.Where("price > 50").Find(&books)
	if res.Error != nil {
		return
	}

	fmt.Println("找到记录数:", res.RowsAffected)

	log.Printf("%+v\n", books)
	for _, emp := range books {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", emp.Id, emp.Title, emp.Author, emp.Price)
	}

}

func main() {
	func_sqlx01()
	func_gorm01()
	func_sqlx02()
	func_gorm02()
}
