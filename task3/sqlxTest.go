package task3

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 或其他数据库驱动（如 postgres, sqlite 等）
	"github.com/jmoiron/sqlx"
	"time"
)

// 查询指定部门的所有员工
func GetEmployeesByDepartment(db *sqlx.DB, params string) ([]EmployeeInfo, error) {
	var employeeList []EmployeeInfo
	query := `SELECT id, name, department, salary,created_at,updated_at FROM employees WHERE department = ?`
	err := db.Select(&employeeList, query, params)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return employeeList, nil
}

// 查询薪资最高的员工
func GetHighestPaidEmployee(db *sqlx.DB) (EmployeeInfo, error) {
	var employee EmployeeInfo
	query := `SELECT id, name, department, salary ,created_at,updated_at FROM employees ORDER BY salary DESC LIMIT 1`
	err := db.Get(&employee, query)
	if err != nil {
		return EmployeeInfo{}, fmt.Errorf("查询失败: %v", err)
	}
	return employee, nil
}

type EmployeeInfo struct {
	ID         int        `db:"id"`
	Name       string     `db:"name"`
	Department string     `db:"department"`
	Salary     float64    `db:"salary"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

func QueryExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book

	// 使用参数化查询防止SQL注入
	query := `SELECT id, title, author, price 
	          FROM books 
	          WHERE price > ? 
	          ORDER BY price DESC`

	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return books, nil
}

type Book struct {
	ID     int     `db:"id"`     // 对应books表的id字段
	Title  string  `db:"title"`  // 对应books表的title字段
	Author string  `db:"author"` // 对应books表的author字段
	Price  float64 `db:"price"`  // 对应books表的price字段
}
